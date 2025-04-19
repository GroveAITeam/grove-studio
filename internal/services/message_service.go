package services

import (
    "context"
    "errors"
    "fmt"
    "grove-studio/internal/database"
    "grove-studio/internal/models"
    "grove-studio/internal/utils"

    "github.com/openai/openai-go"
    "github.com/openai/openai-go/option"
    "github.com/openai/openai-go/packages/param"
    "github.com/wailsapp/wails/v2/pkg/runtime"
)

type MessageService struct {
    ctx    context.Context
    logger *utils.Logger
}

func NewMessageService(ctx context.Context) *MessageService {
    return &MessageService{
        ctx:    ctx,
        logger: utils.NewLogger(ctx),
    }
}

type MessageRequestParams struct {
    CloudLLMId          int     `json:"cloud_llm_id"`
    ConversationId      int     `json:"conversation_id"`
    Question            string  `json:"question"`
    ModelName           string  `json:"model_name"`
    Temperature         float64 `json:"temperature"`
    MaxCompletionTokens uint32  `json:"max_completion_tokens"`
    HistoryLength       uint32  `json:"history_length"`
}

type MessagePageResult struct {
    Items []models.Message `json:"items"`
}

// GetList 获取消息列表
func (n *MessageService) GetList(conversationId, minId, size int) (*MessagePageResult, error) {
    minId = max(minId, 0)
    size = max(size, 20)

    if conversationId < 0 {
        return nil, errors.New("会话ID不能为空")
    }

    var items []models.Message
    db := database.DB.
        Model(&models.Message{}).
        Where("conversation_id = ?", conversationId)
    if minId > 0 {
        db = db.Where("min_id < ?", minId)
    }
    if err := db.Limit(size).Find(&items).Error; err != nil {
        n.logger.Error("消息列表获取失败: %v", err)
        return nil, err
    }

    return &MessagePageResult{Items: items}, nil
}

// Request 给大模型发消息
func (n *MessageService) Request(params MessageRequestParams) (int64, error) {
    if params.CloudLLMId <= 0 {
        return 0, errors.New("模型ID不能为空")
    }
    if params.Question == "" {
        return 0, errors.New("问题不能为空")
    }
    if params.ModelName == "" {
        return 0, errors.New("模型名称不能为空")
    }
    if params.Temperature <= 0 {
        return 0, errors.New("温蒂不能为空")
    }

    var cloudLLM models.CloudLLMModel
    var conversation models.Conversation
    var historyMessages []models.Message
    var message models.Message

    if err := database.DB.Where("id = ?", params.CloudLLMId).First(&cloudLLM).Error; err != nil {
        n.logger.Error("查找模型失败: %v", err)
        return 0, err
    }
    if cloudLLM.ID == 0 {
        msg := fmt.Sprintf("ID=%d的模型不存在", params.CloudLLMId)
        n.logger.Error(msg)
        return 0, errors.New(msg)
    }

    if params.ConversationId > 0 {
        if err := database.DB.Where("id = ?", params.ConversationId).First(&conversation).Error; err != nil {
            n.logger.Error("查询会话失败: %v", err)
            return 0, err
        }
        if conversation.ID == 0 {
            msg := fmt.Sprintf("ID=%d的会话不存在", params.ConversationId)
            n.logger.Error(msg)
            return 0, errors.New(msg)
        }
        if err := database.DB.Where("conversation_id = ?", params.ConversationId).Order("id desc").Limit(int(params.HistoryLength) * 2).Find(&historyMessages).Error; err != nil {
            return 0, err
        }
    } else {
        if err := database.DB.Create(&conversation).Error; err != nil {
            return 0, err
        }
    }

    var messages []openai.ChatCompletionMessageParamUnion
    for i := len(historyMessages) - 1; i >= 0; i-- {
        msg := historyMessages[i]
        if msg.Role == "user" {
            messages = append(messages, openai.UserMessage(msg.Content))
        }
        if msg.Role == "assistant" {
            messages = append(messages, openai.AssistantMessage(msg.Content))
        }
    }
    messages = append(messages, openai.UserMessage(params.Question))

    openaiParams := openai.ChatCompletionNewParams{
        Messages: messages,
        Model:    params.ModelName,
        StreamOptions: openai.ChatCompletionStreamOptionsParam{
            IncludeUsage: param.NewOpt(true),
        },
        Temperature: param.NewOpt(params.Temperature),
    }
    if params.MaxCompletionTokens > 0 {
        openaiParams.MaxCompletionTokens = param.NewOpt(int64(params.MaxCompletionTokens))
    }

    client := openai.NewClient(option.WithAPIKey(cloudLLM.ApiKey))
    stream := client.Chat.Completions.NewStreaming(n.ctx, openaiParams)
    acc := openai.ChatCompletionAccumulator{}

    chunkIndex := 0
    for stream.Next() {
        chunk := stream.Current()
        acc.AddChunk(chunk)
        if _, ok := acc.JustFinishedContent(); ok {
            runtime.EventsEmit(n.ctx, "stream-request-message", map[string]any{
                "content": "",
                "done":    true,
                "index":   chunkIndex,
            })
        }

        if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
            runtime.EventsEmit(n.ctx, "stream-request-message", map[string]any{
                "content": chunk.Choices[0].Delta.Content,
                "done":    false,
                "index":   chunkIndex,
            })
            chunkIndex++
        }
    }

    if err := stream.Err(); err != nil {
        n.logger.Error("流式输出请求失败: %v", err)
        return 0, err
    }

    // 保存用户消息
    message = models.Message{
        ConversationID: conversation.ID,
        Role:           "user",
        Content:        params.Question,
    }
    if err := database.DB.Create(&message).Error; err != nil {
        return 0, err
    }

    // 保存AI消息
    message = models.Message{
        ConversationID: conversation.ID,
        Role:           "assistant",
        Content:        acc.Choices[0].Message.Content,
    }
    if err := database.DB.Create(&message).Error; err != nil {
        return 0, err
    }

    return acc.Usage.TotalTokens, nil
}

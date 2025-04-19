package services

import (
	"context"
	"errors"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"grove-studio/internal/database"
	"grove-studio/internal/models"
	"grove-studio/internal/utils"
	"time"
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
	CloudLLMId   int    `json:"cloud_llm_id"`
	Conversation int    `json:"conversation"`
	Question     string `json:"question"`
	ModelName    string `json:"model_name"`
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
func (n *MessageService) Request(params MessageRequestParams) error {
	// 模拟的AI响应
	mockResponse := "你好，这是 AI 的回复消息"
	var words []string
	for _, char := range mockResponse {
		words = append(words, string(char))
	}

	for _, word := range words {
		runtime.EventsEmit(n.ctx, "stream-request-message", map[string]any{
			"content": word,
			"done":    false,
		})
		time.Sleep(100 * time.Millisecond)
	}

	runtime.EventsEmit(n.ctx, "stream-request-message", map[string]any{
		"content": "",
		"done":    true,
	})

	return nil
}

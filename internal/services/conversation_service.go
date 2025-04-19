package services

import (
    "context"
    "grove-studio/internal/database"
    "grove-studio/internal/models"
    "grove-studio/internal/utils"
)

type ConversationService struct {
    ctx    context.Context
    logger *utils.Logger
}

type ConversationPageResult struct {
    Total int64                 `json:"total"`
    Items []models.Conversation `json:"items"`
}

func NewConversationService(ctx context.Context) *ConversationService {
    return &ConversationService{
        ctx:    ctx,
        logger: utils.NewLogger(ctx),
    }
}

// GetList 分页获取会话列表
func (n *ConversationService) GetList(page, size int) (*ConversationPageResult, error) {
    page = max(page, 1)
    size = max(size, 50)

    var total int64
    var items []models.Conversation

    // 查询总数
    if err := database.DB.Model(&models.Conversation{}).Count(&total).Error; err != nil {
        n.logger.Error("获取会话总数失败: %v", err)
        return nil, err
    }

    // 分页查询
    if err := database.DB.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
        n.logger.Error("分页查询会话列表失败: %v", err)
    }

    return &ConversationPageResult{
        Total: total,
        Items: items,
    }, nil
}

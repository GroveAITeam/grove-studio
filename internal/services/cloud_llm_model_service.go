package services

import (
    "context"
    "grove-studio/internal/utils"
)

type CloudLLMModelService struct {
    logger *utils.Logger
}

func NewCloudLLMModelService(ctx context.Context) *CloudLLMModelService {
    return &CloudLLMModelService{
        logger: utils.NewLogger(ctx),
    }
}

// GetList 获取模型列表
func (c *CloudLLMModelService) GetList(page, size int, enabled bool) ([]CloudLLMModelService, error) {

    return nil, nil
}

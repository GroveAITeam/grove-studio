package services

import (
	"context"
	"errors"
	"grove-studio/internal/database"
	"grove-studio/internal/models"
	"grove-studio/internal/utils"
)

// CloudLLMModelService 云端LLM模型服务
type CloudLLMModelService struct {
	ctx    context.Context
	logger *utils.Logger
}

// CloudLLMModelPageResult 分页查询结果
type CloudLLMModelPageResult struct {
	Total int64                  `json:"total"`
	Items []models.CloudLLMModel `json:"items"`
}

// NewCloudLLMModelService 创建云端模型服务
func NewCloudLLMModelService(ctx context.Context) *CloudLLMModelService {
	return &CloudLLMModelService{
		ctx:    ctx,
		logger: utils.NewLogger(ctx),
	}
}

// GetList 获取模型列表，支持分页
func (c *CloudLLMModelService) GetList(page, size int) (*CloudLLMModelPageResult, error) {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = 10
	}

	var total int64
	var items []models.CloudLLMModel

	// 查询总数
	if err := database.DB.Model(&models.CloudLLMModel{}).Count(&total).Error; err != nil {
		c.logger.Error("获取云端模型总数失败: %v", err)
		return nil, err
	}

	// 分页查询
	if err := database.DB.Offset((page - 1) * size).Limit(size).Find(&items).Error; err != nil {
		c.logger.Error("分页查询云端模型失败: %v", err)
		return nil, err
	}

	return &CloudLLMModelPageResult{
		Total: total,
		Items: items,
	}, nil
}

// GetByID 根据ID获取云端模型
func (c *CloudLLMModelService) GetByID(id uint) (*models.CloudLLMModel, error) {
	var model models.CloudLLMModel
	if err := database.DB.First(&model, id).Error; err != nil {
		c.logger.Error("获取云端模型详情失败: %v", err)
		return nil, err
	}
	return &model, nil
}

// Create 创建云端模型
func (c *CloudLLMModelService) Create(model *models.CloudLLMModel) error {
	if model.Name == "" || model.Provider == "" || model.ApiKey == "" {
		return errors.New("模型名称、提供商和API密钥不能为空")
	}

	if err := database.DB.Create(model).Error; err != nil {
		c.logger.Error("创建云端模型失败: %v", err)
		return err
	}
	return nil
}

// Update 更新云端模型
func (c *CloudLLMModelService) Update(model *models.CloudLLMModel) error {
	if model.ID == 0 {
		return errors.New("模型ID不能为空")
	}

	// 先检查是否存在
	var count int64
	database.DB.Model(&models.CloudLLMModel{}).Where("id = ?", model.ID).Count(&count)
	if count == 0 {
		return errors.New("模型不存在")
	}

	// 更新模型
	if err := database.DB.Save(model).Error; err != nil {
		c.logger.Error("更新云端模型失败: %v", err)
		return err
	}
	return nil
}

// Delete 删除云端模型
func (c *CloudLLMModelService) Delete(id uint) error {
	if id == 0 {
		return errors.New("模型ID不能为空")
	}

	// 执行删除
	if err := database.DB.Delete(&models.CloudLLMModel{}, id).Error; err != nil {
		c.logger.Error("删除云端模型失败: %v", err)
		return err
	}
	return nil
}

// ToggleEnabled 切换启用状态
func (c *CloudLLMModelService) ToggleEnabled(id uint, enabled bool) error {
	if id == 0 {
		return errors.New("模型ID不能为空")
	}

	if err := database.DB.Model(&models.CloudLLMModel{}).Where("id = ?", id).Update("enabled", enabled).Error; err != nil {
		c.logger.Error("切换云端模型状态失败: %v", err)
		return err
	}
	return nil
}

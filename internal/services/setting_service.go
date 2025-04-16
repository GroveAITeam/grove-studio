package services

import (
	"grove-studio/internal/database"
	"grove-studio/internal/models"
)

// SettingService 设置服务
type SettingService struct{}

// NewSettingService 创建新的设置服务
func NewSettingService() *SettingService {
	return &SettingService{}
}

// GetSetting 获取设置值
func (s *SettingService) GetSetting(key string) (string, error) {
	var setting models.Setting
	result := database.DB.Where("key = ?", key).First(&setting)
	if result.Error != nil {
		return "", result.Error
	}
	return setting.Value, nil
}

// SetSetting 设置值
func (s *SettingService) SetSetting(key, value string) error {
	var setting models.Setting
	result := database.DB.Where("key = ?", key).First(&setting)

	if result.Error == nil {
		// 更新已存在的设置
		setting.Value = value
		return database.DB.Save(&setting).Error
	} else {
		// 创建新设置
		setting = models.Setting{
			Key:   key,
			Value: value,
		}
		return database.DB.Create(&setting).Error
	}
}

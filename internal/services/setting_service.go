package services

import (
    "context"
    "grove-studio/internal/database"
    "grove-studio/internal/models"
    "grove-studio/internal/utils"
)

// SettingService 设置服务
type SettingService struct {
    logger *utils.Logger
}

// NewSettingService 创建新的设置服务
func NewSettingService(ctx context.Context) *SettingService {
    return &SettingService{
        logger: utils.NewLogger(ctx),
    }
}

// GetSetting 获取设置值
func (s *SettingService) GetSetting(key string) (string, error) {
    var setting models.Setting
    result := database.DB.Where("key = ?", key).First(&setting)
    if result.Error != nil {
        s.logger.Error("获取设置失败: %v", result.Error)
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
        if err := database.DB.Save(&setting).Error; err != nil {
            s.logger.Error("更新设置失败: %v", err)
            return err
        }
    } else {
        // 创建新设置
        setting = models.Setting{
            Key:   key,
            Value: value,
        }
        if err := database.DB.Create(&setting).Error; err != nil {
            s.logger.Error("创建设置失败: %v", err)
            return err
        }
    }

    return nil
}

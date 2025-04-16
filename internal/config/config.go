package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"
)

// AppConfig 应用程序配置
type AppConfig struct {
	// 应用程序版本
	Version string `json:"version"`
	// 数据目录路径
	DataPath string `json:"data_path"`
	// 是否开启调试模式
	DebugMode bool `json:"debug_mode"`

	// 更多配置项
	// ...
}

var (
	// 默认配置
	defaultConfig = AppConfig{
		Version:   "1.0.0",
		DebugMode: false,
	}

	// 全局配置实例
	instance     *AppConfig
	instanceOnce sync.Once
)

// GetConfig 获取配置实例（单例模式）
func GetConfig() *AppConfig {
	instanceOnce.Do(func() {
		instance = &defaultConfig
	})
	return instance
}

// LoadConfig 从文件加载配置
func LoadConfig(configPath string) error {
	// 确保配置目录存在
	configDir := filepath.Dir(configPath)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return err
	}

	// 检查配置文件是否存在
	_, err := os.Stat(configPath)
	if os.IsNotExist(err) {
		// 文件不存在，创建默认配置
		return SaveConfig(configPath)
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// 解析JSON到配置结构
	config := GetConfig()
	if err := json.Unmarshal(data, config); err != nil {
		return err
	}

	return nil
}

// SaveConfig 保存配置到文件
func SaveConfig(configPath string) error {
	// 序列化配置为JSON
	data, err := json.MarshalIndent(GetConfig(), "", "  ")
	if err != nil {
		return err
	}

	// 写入配置文件
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		return err
	}

	return nil
}

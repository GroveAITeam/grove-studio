package utils

import (
	"os"
	"path/filepath"
)

// GetAppDataPath 获取应用程序数据路径
func GetAppDataPath() (string, error) {
	// 获取用户主目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// 创建应用程序数据目录
	appDataPath := filepath.Join(homeDir, ".grove-studio")
	if err := os.MkdirAll(appDataPath, 0755); err != nil {
		return "", err
	}

	return appDataPath, nil
}

package database

import (
	"fmt"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(appDataPath string, debugMode bool) error {
	// 确保数据目录存在
	dbDir := filepath.Join(appDataPath, "data")
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		return fmt.Errorf("创建数据目录失败: %w", err)
	}

	dbPath := filepath.Join(dbDir, "grove.db")

	// 根据调试模式设置日志级别
	logLevel := logger.Silent
	if debugMode {
		logLevel = logger.Info // 调试模式使用Info级别
	}

	// 配置GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	// 连接数据库
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), config)
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 获取底层SQL连接并设置连接池参数
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("获取DB实例失败: %w", err)
	}

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)

	return nil
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

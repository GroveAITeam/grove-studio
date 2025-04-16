package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	// InfoLogger 信息日志
	InfoLogger *log.Logger
	// ErrorLogger 错误日志
	ErrorLogger *log.Logger
)

// InitLogger 初始化日志系统
func InitLogger(appDataPath string) error {
	// 创建日志目录
	logDir := filepath.Join(appDataPath, "logs")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}

	// 打开日志文件
	logFile, err := os.OpenFile(
		filepath.Join(logDir, "app.log"),
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return err
	}

	// 设置日志输出到控制台和文件
	multiWriter := io.MultiWriter(os.Stdout, logFile)

	// 创建日志记录器
	InfoLogger = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(multiWriter, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

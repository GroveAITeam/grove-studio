package utils

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Logger 日志工具结构体
type Logger struct {
	ctx context.Context
}

// NewLogger 创建新的日志工具实例
func NewLogger(ctx context.Context) *Logger {
	return &Logger{ctx: ctx}
}

// Info 记录信息日志
func (l *Logger) Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	runtime.LogInfo(l.ctx, msg)
}

// Error 记录错误日志
func (l *Logger) Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	runtime.LogError(l.ctx, msg)
}

// Debug 记录调试日志
func (l *Logger) Debug(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	runtime.LogDebug(l.ctx, msg)
}

// Fatal 记录致命错误日志
func (l *Logger) Fatal(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	runtime.LogFatal(l.ctx, msg)
}

// Warning 记录警告日志
func (l *Logger) Warning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	runtime.LogWarning(l.ctx, msg)
}

package services

import (
	"context"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type AIService struct {
	ctx context.Context // 存储 Wails 上下文
}

func NewAIService(ctx context.Context) *AIService {
	return &AIService{
		ctx: ctx,
	}
}

// SendMessageStream 模拟流式消息
func (s *AIService) SendMessageStream(message string) error {
	// 模拟的 AI 响应
	mockResponse := "Hello! This is a simulated AI response to your message: " + message
	var words []string // 按字符或词分割
	for _, char := range mockResponse {
		words = append(words, string(char))
	}

	// 模拟流式输出
	for _, word := range words {
		// 发送数据块到前端
		runtime.EventsEmit(s.ctx, "stream-message", map[string]string{
			"content": word,
			"done":    "false",
		})

		// 模拟网络延迟
		time.Sleep(100 * time.Millisecond)
	}

	// 发送流结束事件
	runtime.EventsEmit(s.ctx, "stream-message", map[string]string{
		"content": "",
		"done":    "true",
	})

	return nil
}

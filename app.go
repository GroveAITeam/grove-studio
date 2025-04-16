package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"grove-studio/internal/config"
	"grove-studio/internal/database"
	"grove-studio/internal/models"
	"grove-studio/internal/services"
	"grove-studio/internal/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx            context.Context
	settingService *services.SettingService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		settingService: services.NewSettingService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	// 获取应用数据路径
	appDataPath, err := utils.GetAppDataPath()
	if err != nil {
		log.Fatalf("获取应用数据路径失败: %v", err)
	}

	// 初始化日志
	if err := utils.InitLogger(appDataPath); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}

	// 加载配置
	configPath := filepath.Join(appDataPath, "config.json")
	if err := config.LoadConfig(configPath); err != nil {
		utils.ErrorLogger.Fatalf("加载配置失败: %v", err)
	}

	// 保存数据路径到配置
	config.GetConfig().DataPath = appDataPath
	if err := config.SaveConfig(configPath); err != nil {
		utils.ErrorLogger.Printf("保存配置失败: %v", err)
	}

	// 初始化数据库
	if err := database.InitDB(appDataPath); err != nil {
		utils.ErrorLogger.Fatalf("初始化数据库失败: %v", err)
	}

	// 自动迁移数据库表结构
	if err := database.DB.AutoMigrate(&models.Setting{}); err != nil {
		utils.ErrorLogger.Fatalf("数据库迁移失败: %v", err)
	}

	utils.InfoLogger.Println("应用初始化完成")
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
	// 前端加载完成后通知
	runtime.EventsEmit(ctx, "app:ready", true)
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// 关闭数据库连接
	if err := database.CloseDB(); err != nil {
		utils.ErrorLogger.Printf("关闭数据库连接失败: %v", err)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("你好 %s, 欢迎体验智擎平台!", name)
}

// GetSetting 获取设置值 (供前端调用)
func (a *App) GetSetting(key string) (string, error) {
	return a.settingService.GetSetting(key)
}

// SetSetting 设置值 (供前端调用)
func (a *App) SetSetting(key, value string) error {
	return a.settingService.SetSetting(key, value)
}

// GetAppConfig 获取应用配置 (供前端调用)
func (a *App) GetAppConfig() (*config.AppConfig, error) {
	return config.GetConfig(), nil
}

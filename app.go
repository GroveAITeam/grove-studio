package main

import (
    "context"
    "fmt"
    "path/filepath"

    "github.com/wailsapp/wails/v2/pkg/logger"

    "grove-studio/internal/config"
    "grove-studio/internal/database"
    "grove-studio/internal/models"
    "grove-studio/internal/services"
    "grove-studio/internal/utils"

    "github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
    ctx                  context.Context
    settingService       *services.SettingService
    cloudLLMModelService *services.CloudLLMModelService
    logger               *utils.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
    return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
    a.ctx = ctx
    a.logger = utils.NewLogger(ctx)
    a.settingService = services.NewSettingService(ctx)
    a.cloudLLMModelService = services.NewCloudLLMModelService(ctx)

    // 获取应用数据路径
    appDataPath, err := utils.GetAppDataPath()
    if err != nil {
        a.logger.Fatal("获取应用数据路径失败: %v", err)
        return
    }

    // 获取环境信息，检查是否为开发模式
    envInfo := runtime.Environment(ctx)
    isDevMode := envInfo.BuildType == "dev"
    if isDevMode {
        runtime.LogSetLogLevel(ctx, logger.DEBUG)
    }

    // 使用日志记录运行模式
    a.logger.Info("当前运行模式: %s", envInfo.BuildType)

    // 加载配置
    configPath := filepath.Join(appDataPath, "config.json")
    if err := config.LoadConfig(configPath); err != nil {
        runtime.LogError(ctx, fmt.Sprintf("加载配置失败: %v", err))
        return
    }

    // 根据运行模式设置调试模式
    config.GetConfig().DebugMode = isDevMode

    // 保存数据路径到配置
    config.GetConfig().DataPath = appDataPath
    if err := config.SaveConfig(configPath); err != nil {
        runtime.LogError(ctx, fmt.Sprintf("保存配置失败: %v", err))
    }

    // 初始化数据库
    if err := database.InitDB(appDataPath, config.GetConfig().DebugMode); err != nil {
        runtime.LogError(ctx, fmt.Sprintf("初始化数据库失败: %v", err))
        return
    }

    // 自动迁移数据库表结构
    if err := database.DB.AutoMigrate(&models.Setting{}, &models.CloudLLMModel{}); err != nil {
        runtime.LogError(ctx, fmt.Sprintf("数据库迁移失败: %v", err))
        return
    }

    runtime.LogInfo(ctx, "应用初始化完成")
}

// domReady is called after front-end resources have been loaded
func (a *App) domReady(ctx context.Context) {
    // 前端加载完成后通知
    runtime.LogInfo(ctx, "前端资源加载完成")
    runtime.EventsEmit(ctx, "app:ready", true)
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
    runtime.LogInfo(ctx, "应用即将关闭")
    return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
    // 关闭数据库连接
    if err := database.CloseDB(); err != nil {
        runtime.LogError(ctx, fmt.Sprintf("关闭数据库连接失败: %v", err))
    }
    runtime.LogInfo(ctx, "应用已关闭")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
    return fmt.Sprintf("你好 %s, 欢迎体验 Grove Studio!", name)
}

// GetSetting 获取设置值 (供前端调用)
func (a *App) GetSetting(key string) (string, error) {    return a.settingService.GetSetting(key)
}

// SetSetting 设置值 (供前端调用)
func (a *App) SetSetting(key, value string) error {
    return a.settingService.SetSetting(key, value)
}

// GetAppConfig 获取应用配置 (供前端调用)
func (a *App) GetAppConfig() (*config.AppConfig, error) {
    return config.GetConfig(), nil
}

// ----------------------------- 云端模型相关API -----------------------------

// GetCloudLLMModels 分页获取云端模型列表 (供前端调用)
func (a *App) GetCloudLLMModels(page, size int) (*services.CloudLLMModelPageResult, error) {
    return a.cloudLLMModelService.GetList(page, size)
}

// GetCloudLLMModelByID 获取云端模型详情 (供前端调用)
func (a *App) GetCloudLLMModelByID(id uint) (*models.CloudLLMModel, error) {
    return a.cloudLLMModelService.GetByID(id)
}

// CreateCloudLLMModel 创建云端模型 (供前端调用)
func (a *App) CreateCloudLLMModel(model *models.CloudLLMModel) error {
    return a.cloudLLMModelService.Create(model)
}

// UpdateCloudLLMModel 更新云端模型 (供前端调用)
func (a *App) UpdateCloudLLMModel(model *models.CloudLLMModel) error {
    return a.cloudLLMModelService.Update(model)
}

// DeleteCloudLLMModel 删除云端模型 (供前端调用)
func (a *App) DeleteCloudLLMModel(id uint) error {
    return a.cloudLLMModelService.Delete(id)
}

// ToggleCloudLLMModelEnabled 切换云端模型启用状态 (供前端调用)
func (a *App) ToggleCloudLLMModelEnabled(id uint, enabled bool) error {
    return a.cloudLLMModelService.ToggleEnabled(id, enabled)
}

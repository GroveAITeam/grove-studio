package main

import (
	"embed"
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:             "grove-studio",
		Width:             1024,
		Height:            768,
		MinWidth:          1024,
		MinHeight:         768,
		DisableResize:     false,
		Fullscreen:        false,
		Frameless:         runtime.GOOS != "darwin",
		StartHidden:       false,
		HideWindowOnClose: true,
		AssetServer: &assetserver.Options{
			Assets:     assets,
			Handler:    nil,
			Middleware: nil,
		},
		BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 255},
		Menu:             nil,
		Logger:           nil,
		LogLevel:         logger.DEBUG,
		OnStartup:        app.startup,
		OnDomReady:       app.domReady,
		OnBeforeClose:    app.beforeClose,
		OnShutdown:       app.shutdown,
		WindowStartState: options.Normal,
		Bind: []interface{}{
			app,
		},
		// Windows 平台特殊配置
		Windows: &windows.Options{
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			DisableWindowIcon:    true,
			// DisableFramelessWindowDecorations: false,
			WebviewUserDataPath: "",
			BackdropType:        windows.Acrylic,
		},
		// Mac 平台特殊配置
		Mac: &mac.Options{
			TitleBar:             mac.TitleBarHidden(),
			Appearance:           mac.NSAppearanceNameVibrantLight,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   "grove-studio",
				Message: "本地化AI工具",
				Icon:    icon,
			},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

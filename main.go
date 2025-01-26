package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "vrc_avatar_library_updatter",
		Width:  400,  // 幅を800ピクセルに設定
		Height: 200,  // 高さを600ピクセルに設定
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 255},
		OnStartup:        app.startup,
		DisableResize:    true, // フルスクリーンボタンを無効化
		Windows: &windows.Options{
            CustomTheme: &windows.ThemeSettings{
				LightModeTitleBar:  windows.RGB(255, 255, 255),
                LightModeTitleText: windows.RGB(0, 0, 0),
                LightModeBorder:    windows.RGB(255, 255, 255),
            },
			Theme: windows.Light,
			BackdropType: windows.None,
        },
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

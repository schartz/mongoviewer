package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Open", keys.CmdOrCtrl("o"), func(_ *menu.CallbackData) {
		println("open file")
	})
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(context.Background())
	})

	HelpMenu := AppMenu.AddSubmenu("Help")
	HelpMenu.AddText("About", nil, func(_ *menu.CallbackData) {
		println("About App")
	})

	//if runtime.GOOS == "darwin" {
	//	AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	//}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "some",
		Width:            1024,
		Height:           768,
		Assets:           assets,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Menu:             AppMenu,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

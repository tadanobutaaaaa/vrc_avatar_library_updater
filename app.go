package main

import (
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) domReady(ctx context.Context) {
	if err := DownloadAndReplace(a); err != nil {
		fmt.Println("Error: ", err)
		runtime.EventsEmit(a.ctx, "error", err.Error())
	}
	if err := StartApp(a); err != nil {
		fmt.Println("Error: ", err)
		runtime.EventsEmit(a.ctx, "error", err.Error())
	}
}
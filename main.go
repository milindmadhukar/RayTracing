package main

import (
	"context"

	"fyne.io/fyne/v2/app"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/window"
	// "net/http"
	// _ "net/http/pprof"
)

type Application struct {
	Window *window.Window
	Scene  *scene.Scene
}

func NewApplication() *Application {
	fyneApp := app.New()
	fyneWindow := fyneApp.NewWindow("RayTracing")

	window := &window.Window{
		FyneApp:    fyneApp,
		FyneWindow: fyneWindow,
	}

	scene := scene.NewScene()

	return &Application{
		Window: window,
		Scene:  scene,
	}
}

func main() {
	// TODO: Rendering with sampling without a GUI.
	application := NewApplication()
	application.Window.AutoRender = false
  application.Window.AutoRenderContext, application.Window.AutoRenderContextCancel = context.WithCancel(context.Background())

	application.Window.Create(application.Scene)

  go func() {
    // http.ListenAndServe("localhost:6060", nil)
  }()

	application.Window.FyneWindow.ShowAndRun()
}

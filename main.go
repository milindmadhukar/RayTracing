package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/window"
)

type Application struct {
  Window *window.Window
  Scene *scene.Scene
}

func main() {
	application := NewApplication()
	application.Window.AutoRender = false

	application.Window.Init(application.Scene)

	application.Window.FyneWindow.ShowAndRun()
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
    Scene: scene,
  }
}

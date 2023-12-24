package main

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2/app"
	"github.com/milindmadhukar/RayTracing/models"
	"github.com/milindmadhukar/RayTracing/renderer"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	window := CreateWindow()
	renderer.InitWindow(window)
	window.FyneWindow.ShowAndRun()
}

func CreateWindow() *models.Window {
	fyneApp := app.New()
	fyneWindow := fyneApp.NewWindow("RayTracing")

	return &models.Window{
		FyneApp:    fyneApp,
		FyneWindow: fyneWindow,
	}
}

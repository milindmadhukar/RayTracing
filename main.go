package main

import (
	"image"
	"image/color"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/models"
)

func main() {
	window := CreateWindow()
	initWindow(window)
	window.FyneWindow.ShowAndRun()
}

func initWindow(window *models.Window) {
	canvasImage := canvas.NewImageFromImage(window.Image)

	renderButton := widget.NewButton("Render", func() {
		render(canvasImage, window)
		canvasImage.Refresh()
	})

	settingsContinaer := container.New(layout.NewVBoxLayout(), renderButton)

	mainContainer := container.NewHSplit(canvasImage, settingsContinaer)
	mainContainer.SetOffset(0.8)

	window.RenderContainer = canvasImage
	window.SettingsContainer = settingsContinaer

	go func() {
		for {
      render(canvasImage, window)
		}
	}()

	window.FyneWindow.SetContent(mainContainer)
}

func getImage(size fyne.Size) *image.RGBA {
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	img := image.NewRGBA(image.Rect(0, 0, int(size.Width), int(size.Height)))
	for y := 0; y < int(size.Height); y++ {
		for x := 0; x < int(size.Width); x++ {
			img.Set(x, y, color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255})
		}
	}

	return img
}

func CreateWindow() *models.Window {
	fyneApp := app.New()
	fyneWindow := fyneApp.NewWindow("RayTracing")

	return &models.Window{
		FyneApp:    fyneApp,
		FyneWindow: fyneWindow,
	}
}

func render(canvasImage *canvas.Image, window *models.Window) {
	img := getImage(window.RenderContainer.Size())
	canvasImage.Image = img
	canvasImage.FillMode = canvas.ImageFillContain
	window.Image = img
	canvasImage.Refresh()
}

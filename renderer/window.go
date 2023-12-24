package renderer

import (
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/models"
)

func InitWindow(window *models.Window) {
	window.RenderContainer = canvas.NewRasterWithPixels(getPixelColor)

	window.FPSLabel = widget.NewLabel("FPS: 0 (Render time: 0ms)")

	renderButton := widget.NewButton("Render", func() {
		render(window)
	})

	window.SettingsContainer = container.New(layout.NewVBoxLayout(), window.FPSLabel, renderButton)

	mainContainer := container.NewHSplit(window.RenderContainer, window.SettingsContainer)
	mainContainer.SetOffset(0.8)

	go func() {
		for range time.Tick(time.Second / 60) {
			render(window)
		}
	}()

	window.FyneWindow.SetContent(mainContainer)
}

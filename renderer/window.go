package renderer

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/models"
	"github.com/milindmadhukar/RayTracing/utils"
)

func InitWindow(window *models.Window) {
	// window.RenderContainer = canvas.NewRaster(generateImage)
	window.RenderContainer = canvas.NewRasterWithPixels(func(x, y, w, h int) color.Color {
		return utils.ConvertToRGBA(getPixelColor(x, y, w, h))
	})

	window.FPSLabel = widget.NewLabel("FPS: 0 (Render time: 0ms)")

	renderButton := widget.NewButton("Render", func() {
		render(window)
	})

	window.SettingsContainer = container.New(layout.NewVBoxLayout(), window.FPSLabel, renderButton)

	mainContainer := container.NewHSplit(window.RenderContainer, window.SettingsContainer)
  mainContainer.SetOffset(0.5) // TODO: Change to 0.84 later

	go func() {
		for range time.Tick(time.Second / 60) {
			render(window)
		}
	}()

	window.FyneWindow.SetContent(mainContainer)
}

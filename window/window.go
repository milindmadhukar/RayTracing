package window

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/camera"
	"github.com/milindmadhukar/RayTracing/renderer"
	"github.com/milindmadhukar/RayTracing/scene"
)

type Window struct {
	FyneApp           fyne.App
	FyneWindow        fyne.Window
	SettingsContainer *fyne.Container
	RenderedRaster    *canvas.Raster
	FPSLabel          *widget.Label
	AutoRender        bool
}

func (applicationWindow *Window) GetRenderedImage(camera *camera.Camera) *canvas.Raster {

	return canvas.NewRaster(
		func(w, h int) image.Image {
			now := time.Now()
			camera.OnResize(w, h)
			img := renderer.GenerateImage(w, h, camera)
			renderer.UpdateFPSLabel(applicationWindow.FPSLabel, time.Since(now))
			return img
		},
	)
}

func (applicationWindow *Window) Init(scene *scene.Scene) {
	applicationWindow.FPSLabel = widget.NewLabel("FPS: 0 (Render time: 0ms)")

	renderButton := widget.NewButton("Render", func() {
		applicationWindow.Update()
	})

	applicationWindow.SettingsContainer = container.New(layout.NewVBoxLayout(), applicationWindow.FPSLabel, renderButton)

	applicationWindow.RenderedRaster = applicationWindow.GetRenderedImage(scene.Camera)

	mainContainer := container.NewHSplit(applicationWindow.RenderedRaster, applicationWindow.SettingsContainer)
	mainContainer.SetOffset(0.80)

	if applicationWindow.AutoRender {
		go func() {
			for {
				applicationWindow.Update()
			}
		}()
	}

	applicationWindow.FyneWindow.SetContent(mainContainer)
}

func (applicationWindow *Window) Update() {
	applicationWindow.RenderedRaster.Refresh()
}

package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/camera"
	"github.com/milindmadhukar/RayTracing/scene"
)

type Window struct {
	FyneApp           fyne.App
	FyneWindow        fyne.Window
	SettingsContainer *fyne.Container
	RenderedRaster    *RenderedRaster
	FPSLabel          *widget.Label
	AutoRender        bool
}

func (applicationWindow *Window) Init(scene *scene.Scene) {
	applicationWindow.FPSLabel = widget.NewLabel("FPS: 0 (Render time: 0ms)")

	renderButton := widget.NewButton("Render", func() {
		applicationWindow.Update(scene.Camera)
	})

	applicationWindow.SettingsContainer = container.New(layout.NewVBoxLayout(), applicationWindow.FPSLabel, renderButton)

	applicationWindow.RenderedRaster = NewRenderedRaster(applicationWindow, scene.Camera)

	mainContainer := container.NewHSplit(applicationWindow.RenderedRaster, applicationWindow.SettingsContainer)
	mainContainer.SetOffset(0.80)

	if applicationWindow.AutoRender {
		go func() {
			for {
				applicationWindow.Update(scene.Camera)
			}
		}()
	}

	applicationWindow.FyneWindow.Canvas()

	applicationWindow.FyneWindow.SetContent(mainContainer)
}

func (applicationWindow *Window) Update(camera *camera.Camera) {
	camera.OnUpdate(applicationWindow.RenderedRaster)
	applicationWindow.RenderedRaster.Refresh()
}

package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/scene"
)

type Window struct {
	FyneApp           fyne.App
	FyneWindow        fyne.Window
	SettingsContainer *container.Scroll
	RenderedRaster    *RenderedRaster
	FPSLabel          *widget.Label
	AutoRender        bool
}

func (applicationWindow *Window) Create(scene *scene.Scene) {
	applicationWindow.FPSLabel = widget.NewLabel("FPS: 0 (Render time: 0ms)")

	renderButton := widget.NewButton("Render", func() {
		applicationWindow.Update()
	})

	allSphereContainers := container.New(layout.NewVBoxLayout())

	for _, sphere := range scene.Spheres {
		posContainer := getSpherePositionUI(sphere)
		radiusContainer := getSphereRadiusUI(sphere)
		colorContainer := getSphereColorUI(sphere, applicationWindow)
		sphereContainer := container.New(layout.NewVBoxLayout(), posContainer, radiusContainer, colorContainer)

    allSphereContainers.Add(sphereContainer)
	}

	// applicationWindow.SettingsContainer = container.New(layout.NewVBoxLayout(), applicationWindow.FPSLabel, renderButton, sphereContainer)

	applicationWindow.SettingsContainer = container.NewVScroll(
		container.New(
			layout.NewVBoxLayout(),
			applicationWindow.FPSLabel, renderButton, allSphereContainers,
		),
	)

	applicationWindow.RenderedRaster = NewRenderedRaster(applicationWindow, scene)

	mainContainer := container.NewHSplit(applicationWindow.RenderedRaster, applicationWindow.SettingsContainer)
	mainContainer.SetOffset(0.80)

	if applicationWindow.AutoRender {
		go func() {
			for {
				applicationWindow.Update()
			}
		}()
	}

	applicationWindow.FyneWindow.Canvas()

	applicationWindow.FyneWindow.SetContent(mainContainer)
}

func (applicationWindow *Window) Update() {
	applicationWindow.RenderedRaster.Refresh()
}

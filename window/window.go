package window

import (
	"strconv"

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

  raytracingConfigContainer := getRayTracingConfigUI(scene)

	allSphereContainers := container.New(layout.NewVBoxLayout())

	for index, sphere := range scene.Spheres {
    sphereLabel := widget.NewLabel("Sphere " + strconv.Itoa(index) + " Settings")
		posContainer := getSpherePositionUI(sphere)
		radiusContainer := getSphereRadiusUI(sphere)
		sphereMaterialIndexContainer := getSphereMaterialIndexUI(sphere, applicationWindow, scene)
		sphereContainer := container.New(layout.NewVBoxLayout(), sphereLabel, posContainer, radiusContainer, sphereMaterialIndexContainer)

		allSphereContainers.Add(sphereContainer)
	}

	allMaterialsContainers := container.New(layout.NewVBoxLayout())

	for idx, material := range scene.Materials {
		materialProperties := getMaterialsUI(material, idx, applicationWindow)

		allMaterialsContainers.Add(materialProperties)
	}


	applicationWindow.SettingsContainer = container.NewVScroll(
		container.New(
			layout.NewVBoxLayout(),
			applicationWindow.FPSLabel, renderButton, raytracingConfigContainer, allSphereContainers, allMaterialsContainers,
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

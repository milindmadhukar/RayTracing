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

	stateUIContainer := getSceneStateUI(scene, applicationWindow)

	accumlationOptions := getAccumulationUI(scene, applicationWindow)

	raytracingConfigContainer := getRayTracingConfigUI(scene)

	// FIXME: Camera position in UI doesn't change when changed with kb
	cameraPositionContainer := getCameraPostionUI(scene.Camera)

	allSphereContainers := container.New(layout.NewVBoxLayout())
	for index, sphere := range scene.Spheres {
		sphereContainer := getSphereUI(sphere, index, scene)
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
			applicationWindow.FPSLabel, renderButton, stateUIContainer, accumlationOptions, raytracingConfigContainer, cameraPositionContainer, allSphereContainers, allMaterialsContainers,
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

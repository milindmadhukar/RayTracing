package window

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/scene"
)

func getSceneStateUI(myScene *scene.Scene, applicationWindow *Window) *fyne.Container {

	exampleSceneURI := storage.NewFileURI("example_scenes/")
	list, _ := storage.ListerForURI(exampleSceneURI)

	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			return
		}
		if writer == nil {
			return
		}

		jsonBytes, err := myScene.ToJson()
		if err != nil {
			fmt.Println(err)
			return
		}

		writer.Write(jsonBytes)
	}, applicationWindow.FyneWindow)

	saveDialog.SetLocation(list)

	saveDialog.SetFileName("scene.json")
	saveDialog.SetConfirmText("Save Scene")
	saveDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))

	saveStateButton := widget.NewButton(
		"Save State",
		func() {
			saveDialog.Show()
		},
	)

	openDialog := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err != nil {
			return
		}
		if reader == nil {
			return
		}

		newScene, err := scene.LoadStateFromJson(reader)
		if err != nil {
			fmt.Println(err)
			return
		}

		applicationWindow.AutoRenderContextCancel()

		myScene.ToAccumulate = false

		myScene.MaxRayBounces = newScene.MaxRayBounces
		myScene.RaysPerPixel = newScene.RaysPerPixel
		myScene.MaxRayBounceDistance = newScene.MaxRayBounceDistance

		myScene.SkyColor = newScene.SkyColor

		myScene.Camera = newScene.Camera
		myScene.Camera.RecalculateViewMatrix()

		myScene.Spheres = newScene.Spheres
		myScene.Materials = newScene.Materials

		myScene.FrameIndex = 1

		applicationWindow.AutoRenderContext, applicationWindow.AutoRenderContextCancel = context.WithCancel(context.Background())

		applicationWindow.Create(myScene)

		applicationWindow.SettingsContainer.Refresh()
		applicationWindow.RenderedRaster.Refresh()

	}, applicationWindow.FyneWindow)

	openDialog.SetConfirmText("Load Scene")

	openDialog.SetLocation(list)
	openDialog.SetFilter(storage.NewExtensionFileFilter([]string{".json"}))

	loadStateButton := widget.NewButton(
		"Load State",
		func() {
			openDialog.Show()
		},
	)

	return container.New(layout.NewGridLayout(2), saveStateButton, loadStateButton)

}

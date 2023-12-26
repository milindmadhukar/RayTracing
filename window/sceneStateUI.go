package window

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/scene"
)

func getSceneStateUI(myScene *scene.Scene, applicationWindow *Window) *fyne.Container {

	saveStateButton := widget.NewButton(
		"Save State",
		func() {
			dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
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
		},
	)

	loadStateButton := widget.NewButton(
		"Load State",
		func() {
			dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
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

				myScene.MaxRayBounces = newScene.MaxRayBounces
				myScene.RaysPerPixel = newScene.RaysPerPixel
        myScene.MaxRayBounceDistance = newScene.MaxRayBounceDistance

				myScene.Camera = newScene.Camera
				myScene.Camera.OnResize(myScene.Camera.ViewportWidth, myScene.Camera.ViewportHeight)
				myScene.Camera.RecalculateViewMatrix()

				myScene.Spheres = newScene.Spheres
				myScene.Materials = newScene.Materials

				myScene.FrameIndex = 1

				applicationWindow.Create(myScene)

				applicationWindow.SettingsContainer.Refresh()
				applicationWindow.RenderedRaster.Refresh()

			}, applicationWindow.FyneWindow)
		},
	)

	return container.New(layout.NewGridLayout(2), saveStateButton, loadStateButton)

}

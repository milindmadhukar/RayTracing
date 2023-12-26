package window

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/scene"
)

func getCameraPostionUI(myScene *scene.Scene) *fyne.Container {
  camera := myScene.Camera

	xPos := widget.NewEntry()
	xPos.SetPlaceHolder("X:")
	xPos.MultiLine = false
	xPos.SetText(strconv.FormatFloat(camera.Position.X(), 'f', 2, 64))

	yPos := widget.NewEntry()
	yPos.SetPlaceHolder("Y:")
	yPos.MultiLine = false
	yPos.SetText(strconv.FormatFloat(camera.Position.Y(), 'f', 2, 64))

	zPos := widget.NewEntry()
	zPos.SetPlaceHolder("Z:")
	zPos.MultiLine = false
	zPos.SetText(strconv.FormatFloat(camera.Position.Z(), 'f', 2, 64))

	xPos.OnChanged = func(text string) {
		xPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		camera.Position = glm.Vec3{xPosValue, camera.Position.Y(), camera.Position.Z()}
		myScene.FrameIndex = 1
	}

	yPos.OnChanged = func(text string) {
		yPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		camera.Position = glm.Vec3{camera.Position.X(), yPosValue, camera.Position.Z()}
		myScene.FrameIndex = 1
	}

	zPos.OnChanged = func(text string) {
		zPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		camera.Position = glm.Vec3{camera.Position.X(), camera.Position.Y(), zPosValue}
		myScene.FrameIndex = 1
	}

	posLabel := widget.NewLabel("Camera Position:")
	posContainer := container.New(layout.NewGridLayout(3), xPos, yPos, zPos)
	return container.New(layout.NewVBoxLayout(), posLabel, posContainer)
}

package window

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/utils"
)

func getSpherePositionUI(sphere *scene.Sphere) *fyne.Container {

	xPos := widget.NewEntry()
	xPos.SetPlaceHolder("X:")
	xPos.MultiLine = false
	xPos.SetText(strconv.FormatFloat(sphere.Position.X(), 'f', 2, 64))

	yPos := widget.NewEntry()
	yPos.SetPlaceHolder("Y:")
	yPos.MultiLine = false
	yPos.SetText(strconv.FormatFloat(sphere.Position.Y(), 'f', 2, 64))

	zPos := widget.NewEntry()
	zPos.SetPlaceHolder("Z:")
	zPos.MultiLine = false
	zPos.SetText(strconv.FormatFloat(sphere.Position.Z(), 'f', 2, 64))

	xPos.OnChanged = func(text string) {
		xPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		sphere.Position = glm.Vec3{xPosValue, sphere.Position.Y(), sphere.Position.Z()}
	}

	yPos.OnChanged = func(text string) {
		yPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		sphere.Position = glm.Vec3{sphere.Position.X(), yPosValue, sphere.Position.Z()}
	}

	zPos.OnChanged = func(text string) {
		zPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		sphere.Position = glm.Vec3{sphere.Position.X(), sphere.Position.Y(), zPosValue}
	}

	posLabel := widget.NewLabel("Position:")
	posContainer := container.New(layout.NewGridLayout(3), xPos, yPos, zPos)

	return container.New(layout.NewVBoxLayout(), posLabel, posContainer)
}

func getSphereRadiusUI(sphere *scene.Sphere) *fyne.Container {
	radius := widget.NewEntry()
	radius.SetPlaceHolder("Radius:")
	radius.MultiLine = false
	radius.SetText(strconv.FormatFloat(sphere.Radius, 'f', 2, 64))

	radius.OnChanged = func(text string) {
		radiusValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		sphere.Radius = radiusValue
	}

	radiusLabel := widget.NewLabel("Radius:")
	return container.New(layout.NewVBoxLayout(), radiusLabel, radius)

}

func getSphereColorUI(sphere *scene.Sphere, applicationWindow *Window) *fyne.Container {
	colorRect := canvas.NewRectangle(utils.ConvertToRGBA(sphere.Albedo.Vec4(1.0)))

	colourPicker := dialog.NewColorPicker("Sphere Colour", "Pick a colour for the sphere", func(colour color.Color) {
		r, g, b, _ := colour.RGBA()
		sphere.Albedo = glm.Vec3{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff}
		colorRect.FillColor = colour
		colorRect.Refresh()
	}, applicationWindow.FyneWindow)

	colourPicker.Advanced = true

	colourPickerBtn := widget.NewButton("Sphere Colour", func() {
		colourPicker.Show()
	})

	return container.New(layout.NewGridLayout(2), colorRect, colourPickerBtn)
}

package window

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/utils"
)

func getRayTracingConfigUI(scene *scene.Scene) *fyne.Container {
	maxRayBouncesLabel := widget.NewLabel("Max Ray Bounces:")
  maxRayBouncesBind := binding.BindInt(&scene.MaxRayBounces)
  maxRayBounces := widget.NewEntryWithData(binding.IntToString(maxRayBouncesBind))
  maxRayBounces.SetPlaceHolder("Max Ray Bounces:")
  maxRayBounces.MultiLine = false
  maxRayBounceContainer := container.New(layout.NewGridLayout(2), maxRayBouncesLabel, maxRayBounces)

  raysPerPixelLabale := widget.NewLabel("Rays Per Pixel:")
  raysPerPixelBind := binding.BindInt(&scene.RaysPerPixel)
  raysPerPixel := widget.NewEntryWithData(binding.IntToString(raysPerPixelBind))
  raysPerPixel.SetPlaceHolder("Rays Per Pixel:")
  raysPerPixel.MultiLine = false
  raysPerPixelContainer := container.New(layout.NewGridLayout(2), raysPerPixelLabale, raysPerPixel)

  maxRayDistanceLabel := widget.NewLabel("Max Ray Distance:")
  maxRayDistanceBind := binding.BindFloat(&scene.MaxRayBounceDistance)
  maxRayDistance := widget.NewEntryWithData(binding.FloatToStringWithFormat(maxRayDistanceBind, "%0.2f"))
  maxRayDistance.SetPlaceHolder("Max Ray Distance:")
  maxRayDistance.MultiLine = false
  maxRayDistanceContainer := container.New(layout.NewGridLayout(2), maxRayDistanceLabel, maxRayDistance)

	return container.New(layout.NewVBoxLayout(), maxRayBounceContainer, raysPerPixelContainer, maxRayDistanceContainer)
}

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
	radiusLabel := widget.NewLabel("Radius:")

	radiusBind := binding.BindFloat(&sphere.Radius)
	radius := widget.NewEntryWithData(binding.FloatToStringWithFormat(radiusBind, "%0.2f"))
	radius.SetPlaceHolder("Radius:")
	radius.MultiLine = false

	return container.New(layout.NewGridLayout(2), radiusLabel, radius)

}

func getSphereMaterialIndexUI(sphere *scene.Sphere, applicationWindow *Window, myScene *scene.Scene) *fyne.Container {
	materialLabel := widget.NewLabel("Material:")

	materialEntry := widget.NewEntry()
	materialEntry.SetPlaceHolder("Material Index:")
	materialEntry.MultiLine = false
	materialEntry.SetText(strconv.Itoa(sphere.MaterialIndex))

	materialEntry.OnChanged = func(text string) {
		materialIndex, err := strconv.Atoi(text)
		if err != nil {
			return
		}
		if materialIndex >= len(myScene.Materials) {
			materialIndex = len(myScene.Materials) - 1
			materialEntry.SetText(strconv.Itoa(materialIndex))
		}

		sphere.MaterialIndex = materialIndex
	}

	return container.New(layout.NewGridLayout(2), materialLabel, materialEntry)
}

func getMaterialsUI(material *scene.Material, index int, applicationWindow *Window) *fyne.Container {
	materialLabel := widget.NewLabel("Material " + strconv.Itoa(index) + " Settings")

	colorRect := canvas.NewRectangle(utils.ConvertToRGBA(material.Albedo.Vec4(1.0)))

	colourPicker := dialog.NewColorPicker("Sphere Colour", "Pick a colour for the sphere", func(colour color.Color) {
		r, g, b, _ := colour.RGBA()
		material.Albedo = glm.Vec3{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff}
		colorRect.FillColor = colour
		colorRect.Refresh()
	}, applicationWindow.FyneWindow)

	colourPicker.Advanced = true

	colourPickerBtn := widget.NewButton("Sphere Colour", func() {
		colourPicker.Show()
	})

	colourPickerContainer := container.New(layout.NewGridLayout(2), colorRect, colourPickerBtn)

	roughnessBind := binding.BindFloat(&material.Roughness)
	roughnessLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(roughnessBind, "Roughness: %0.2f"))
	roughnessSlider := widget.NewSliderWithData(0.0, 1.0, roughnessBind)
	roughnessSlider.Step = 0.01

	metallicBind := binding.BindFloat(&material.Metallic)
	metallicLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(metallicBind, "Metallic: %0.2f"))
	metallicSlider := widget.NewSliderWithData(0.0, 1.0, metallicBind)
	metallicSlider.Step = 0.01

	return container.New(layout.NewVBoxLayout(), materialLabel, colourPickerContainer, roughnessLabel, roughnessSlider, metallicLabel, metallicSlider)

}

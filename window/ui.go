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
	"github.com/milindmadhukar/RayTracing/camera"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/utils"
	"golang.design/x/clipboard"
)

func getAccumulationUI(scene *scene.Scene, applicationWindow *Window) *fyne.Container {
	copyImagetoClipboardBtn := widget.NewButton(
		"Copy Image to Clipboard",
		func() {
			clipboard.Write(clipboard.FmtImage, utils.ImageToPNGBytes(scene.FinalImage))
		},
	)
	exportImageBtn := widget.NewButton(
		"Export Image",
		func() {
			dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
				if err != nil {
					return
				}
        if writer == nil {
          return
        }
				writer.Write(utils.ImageToPNGBytes(scene.FinalImage))
			}, applicationWindow.FyneWindow)
		},
	)

  //FIXME: Only works when clicked multiple times
	resetAccumulationBtn := widget.NewButton(
		"Reset Accumulation",
		func() {
			scene.FrameIndex = 1
		},
	)

	return container.New(layout.NewHBoxLayout(), copyImagetoClipboardBtn, exportImageBtn, resetAccumulationBtn)
}

func getRayTracingConfigUI(scene *scene.Scene) *fyne.Container {

	toAccumulateBind := binding.BindBool(&scene.ToAccumulate)
	toAccumulateCheck := widget.NewCheckWithData("Accumulate:", toAccumulateBind)
	toAccumulateContainer := container.New(layout.NewGridLayout(2), toAccumulateCheck)

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

	return container.New(layout.NewVBoxLayout(), toAccumulateContainer, maxRayBounceContainer, raysPerPixelContainer, maxRayDistanceContainer)
}

func getCameraPostionUI(camera *camera.Camera) *fyne.Container {
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
	}

	yPos.OnChanged = func(text string) {
		yPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		camera.Position = glm.Vec3{camera.Position.X(), yPosValue, camera.Position.Z()}
	}

	zPos.OnChanged = func(text string) {
		zPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		camera.Position = glm.Vec3{camera.Position.X(), camera.Position.Y(), zPosValue}
	}

	posLabel := widget.NewLabel("Camera Position:")
	posContainer := container.New(layout.NewGridLayout(3), xPos, yPos, zPos)

	return container.New(layout.NewVBoxLayout(), posLabel, posContainer)

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

	posLabel := widget.NewLabel("Sphere Position:")
	posContainer := container.New(layout.NewGridLayout(3), xPos, yPos, zPos)

	return container.New(layout.NewVBoxLayout(), posLabel, posContainer)
}

func getSphereRadiusUI(sphere *scene.Sphere) *fyne.Container {
	radiusLabel := widget.NewLabel("Radius:")

	radiusBind := binding.BindFloat(&sphere.Radius)
	radius := widget.NewEntryWithData(binding.FloatToString(radiusBind))
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

	albedoRect := canvas.NewRectangle(utils.ConvertToRGBA(material.Albedo.Vec4(1.0)))

	albedoPicker := dialog.NewColorPicker("Material Albedo", "Pick an albedo for the material", func(colour color.Color) {
		r, g, b, _ := colour.RGBA()
		material.Albedo = glm.Vec3{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff}
		albedoRect.FillColor = colour
		albedoRect.Refresh()
	}, applicationWindow.FyneWindow)

	albedoPicker.Advanced = true

	albedoPickerBtn := widget.NewButton("Material Albedo", func() {
		albedoPicker.Show()
	})

	albedoPickerContainer := container.New(layout.NewGridLayout(2), albedoRect, albedoPickerBtn)

	roughnessBind := binding.BindFloat(&material.Roughness)
	roughnessLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(roughnessBind, "Roughness: %0.2f"))
	roughnessSlider := widget.NewSliderWithData(0.0, 1.0, roughnessBind)
	roughnessSlider.Step = 0.01

	metallicBind := binding.BindFloat(&material.Metallic)
	metallicLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(metallicBind, "Metallic: %0.2f"))
	metallicSlider := widget.NewSliderWithData(0.0, 1.0, metallicBind)
	metallicSlider.Step = 0.01

	emissionColorRect := canvas.NewRectangle(utils.ConvertToRGBA(material.EmissionColor.Vec4(1.0)))

	emissionColorPicker := dialog.NewColorPicker("Emission Colour", "Pick a colour for emission", func(colour color.Color) {
		r, g, b, _ := colour.RGBA()
		material.EmissionColor = glm.Vec3{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff}
		emissionColorRect.FillColor = colour
		emissionColorRect.Refresh()
	}, applicationWindow.FyneWindow)

	emissionColorPicker.Advanced = true

	emissionColorPickerBtn := widget.NewButton("Emission Colour", func() {
		emissionColorPicker.Show()
	})

	emissionColorPickerContainer := container.New(layout.NewGridLayout(2), emissionColorRect, emissionColorPickerBtn)

	emissionStrengthBind := binding.BindFloat(&material.EmissionPower)
	emissionStrengthLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(emissionStrengthBind, "Emission Strength: %0.2f"))
	emissionStrengthEntry := widget.NewEntryWithData(binding.FloatToString(emissionStrengthBind))
	emissionStrengthEntry.SetPlaceHolder("Emission Strength:")
	emissionStrengthEntry.MultiLine = false

	return container.New(layout.NewVBoxLayout(), materialLabel, albedoPickerContainer, roughnessLabel, roughnessSlider, metallicLabel, metallicSlider, emissionColorPickerContainer, emissionStrengthLabel, emissionStrengthEntry)

}

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

func getAlbedoPickerUI(material *scene.Material, applicationWindow *Window) *fyne.Container {
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

	return container.New(layout.NewGridLayout(2), albedoRect, albedoPickerBtn)
}

func getRoughnessMetallicUI(material *scene.Material) *fyne.Container {
	roughnessBind := binding.BindFloat(&material.Roughness)
	roughnessLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(roughnessBind, "Roughness: %0.2f"))
	roughnessSlider := widget.NewSliderWithData(0.0, 1.0, roughnessBind)
	roughnessSlider.Step = 0.01

	metallicBind := binding.BindFloat(&material.Metallic)
	metallicLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(metallicBind, "Metallic: %0.2f"))
	metallicSlider := widget.NewSliderWithData(0.0, 1.0, metallicBind)
	metallicSlider.Step = 0.01

	return container.New(layout.NewVBoxLayout(), roughnessLabel, roughnessSlider, metallicLabel, metallicSlider)
}

func getEmissionColorPickerUI(material *scene.Material, applicationWindow *Window) *fyne.Container {
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

	return container.New(layout.NewGridLayout(2), emissionColorRect, emissionColorPickerBtn)
}

func getEmissionStrengthUI(material *scene.Material) *fyne.Container {

	emissionStrengthBind := binding.BindFloat(&material.EmissionPower)
	emissionStrengthLabel := widget.NewLabelWithData(binding.FloatToStringWithFormat(emissionStrengthBind, "Emission Strength: %0.2f"))
	emissionStrengthEntry := widget.NewEntryWithData(binding.FloatToString(emissionStrengthBind))
	emissionStrengthEntry.SetPlaceHolder("Emission Strength:")
	emissionStrengthEntry.MultiLine = false

	return container.New(layout.NewGridLayout(2), emissionStrengthLabel, emissionStrengthEntry)
}

func getMaterialsUI(material *scene.Material, index int, applicationWindow *Window) *fyne.Container {
	materialLabel := widget.NewLabel("Material " + strconv.Itoa(index) + " Settings")

	return container.New(layout.NewVBoxLayout(), materialLabel, getAlbedoPickerUI(material, applicationWindow), getRoughnessMetallicUI(material), getEmissionColorPickerUI(material, applicationWindow), getEmissionStrengthUI(material))
}

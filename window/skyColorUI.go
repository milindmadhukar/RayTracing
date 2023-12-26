package window

import (
	"image/color"

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

func getSkyColorPickerUI(myScene *scene.Scene, applicationWindow *Window) *fyne.Container {
	skyColorRect := canvas.NewRectangle(utils.ConvertToRGBA(myScene.SkyColor.Vec4(1.0)))

	skyColorPicker := dialog.NewColorPicker("Sky Colour", "Pick a colour for sky", func(colour color.Color) {
		r, g, b, _ := colour.RGBA()
		myScene.SkyColor = glm.Vec3{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff}
		skyColorRect.FillColor = colour
		skyColorRect.Refresh()
	}, applicationWindow.FyneWindow)

	skyColorPicker.Advanced = true

	skyColorPickerBtn := widget.NewButton("Sky Colour", func() {
		skyColorPicker.Show()
	})

	return container.New(layout.NewGridLayout(2), skyColorRect, skyColorPickerBtn)
}

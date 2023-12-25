package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/scene"
)

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

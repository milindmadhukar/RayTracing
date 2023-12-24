package window

import (
	"image"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/camera"
	"github.com/milindmadhukar/RayTracing/renderer"
)

type RenderedRaster struct {
	widget.BaseWidget
	rasterImage *canvas.Raster
}

func (applicationWindow *Window) GetRenderedImage(camera *camera.Camera) *canvas.Raster {
	return canvas.NewRaster(
		func(w, h int) image.Image {
			now := time.Now()
			camera.OnResize(w, h)
			img := renderer.GenerateImage(w, h, camera)
			renderer.UpdateFPSLabel(applicationWindow.FPSLabel, time.Since(now))
			return img
		},
	)
}

func NewRenderedRaster(applicationWindow *Window, camera *camera.Camera) *RenderedRaster {
	renderedRaster := &RenderedRaster{
		rasterImage: applicationWindow.GetRenderedImage(camera),
	}
	renderedRaster.ExtendBaseWidget(renderedRaster)

	return renderedRaster
}

func (renderedRaster *RenderedRaster) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(renderedRaster.rasterImage)
}

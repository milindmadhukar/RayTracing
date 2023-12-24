package renderer

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/milindmadhukar/RayTracing/models"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func render(window *models.Window) {
	now := time.Now()
	window.RenderContainer.Refresh()
	frameTime := time.Since(now)
	updateFPSLabel(window.FPSLabel, frameTime)
}

func getPixelColor(x, y, w, h int) color.Color {

  y = h - y

	xcoord := float64(x) / float64(w)
	ycoord := float64(y) / float64(h)

	return color.RGBA{
		R: uint8(xcoord * 255),
		G: uint8(ycoord * 255),
		B: 0,
		A: 255,
	}
}

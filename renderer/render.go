package renderer

import (
	"image"
	"math"
	"math/rand"
	"sync"
	"time"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/models"
	"github.com/milindmadhukar/RayTracing/utils"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func render(window *models.Window) {
	now := time.Now()
	window.RenderContainer.Refresh()
	frameTime := time.Since(now)
	updateFPSLabel(window.FPSLabel, frameTime)
}

func generateImage(width, height int) image.Image {
	var wg sync.WaitGroup

	wg.Add(width * height)

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			go func(x, y int) {
				defer wg.Done()
				img.Set(x, y, utils.ConvertToRGBA(getPixelColor(x, y, width, height)))
			}(x, y)
		}
	}

	wg.Wait()

	return img
}

func getPixelColor(x, y, w, h int) glm.Vec4 {
	y = h - y

	coord := glm.Vec2{
		(float64(x)/float64(w))*2.0 - 1.0,
		(float64(y)/float64(h))*2.0 - 1.0,
	}

	rayOrigin := glm.Vec3{0.0, 0.0, 2.0}
	rayDirection := coord.Vec3(-1.0)

	radius := 0.5

	a := rayDirection.Dot(rayDirection)
	b := 2.0 * rayOrigin.Dot(rayDirection)
	c := rayOrigin.Dot(rayOrigin) - radius*radius

	discriminant := b*b - 4*a*c

	if discriminant <= 0 {
		return glm.Vec4{0.0, 0.0, 0.0, 1.0}
	}

  sphereColour := glm.Vec3{1, 0, 1}

  // t0 := (-b + math.Sqrt(discriminant)) / (2.0 * a)
  closestT := (-b - math.Sqrt(discriminant)) / (2.0 * a)

  hitPoint := rayOrigin.Add(rayDirection.Mul(closestT))

  sphereColour = hitPoint

  return sphereColour.Vec4(1.0)

}

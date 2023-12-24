package renderer

import (
	"image"
	"math"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/camera"
	"github.com/milindmadhukar/RayTracing/utils"
)

type Ray struct {
	Origin    glm.Vec3
	Direction glm.Vec3
}

func GenerateImage(width, height int, camera *camera.Camera) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	ray := Ray{}
	ray.Origin = camera.Position
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ray.Direction = camera.RayDirections[utils.FlattenXY(x, y, width)]
			colour := TraceRay(&ray)
			img.Set(x, height-y, utils.ConvertToRGBA(colour))
		}
	}

	return img
}

func TraceRay(ray *Ray) glm.Vec4 {
	radius := 0.5

	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Origin.Dot(ray.Direction)
	c := ray.Origin.Dot(ray.Origin) - (radius * radius)

	discriminant := (b * b) - (4 * a * c)

	if discriminant <= 0 {
		// Return black if no intersection
		return glm.Vec4{0.0, 0.0, 0.0, 1.0}
	}

	// t0 := (-b + math.Sqrt(discriminant)) / (2.0 * a)
	closestT := (-b - math.Sqrt(discriminant)) / (2.0 * a)

	hitPoint := ray.Origin.Add(ray.Direction.Mul(closestT))
	normal := hitPoint.Normalize()

	lightDir := glm.Vec3{-1.0, -1.0, -1.0}.Normalize()

	d := math.Max(0, normal.Dot(lightDir.Mul(-1.0))) // cosine of angle between normal and light direction

	sphereColour := glm.Vec3{1, 0, 1}.Mul(d)

	return sphereColour.Vec4(1.0)
}

package renderer

import (
	"image"
	"math"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/utils"
)

type Ray struct {
	Origin    glm.Vec3
	Direction glm.Vec3
}

func GenerateImage(width, height int, scene *scene.Scene) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	ray := Ray{}
	ray.Origin = scene.Camera.Position
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			ray.Direction = scene.Camera.RayDirections[utils.FlattenXY(x, y, width)]
			colour := TraceRay(scene, &ray)
			img.Set(x, height-y, utils.ConvertToRGBA(colour))
		}
	}

	return img
}

func TraceRay(myScene *scene.Scene, ray *Ray) glm.Vec4 {
	if len(myScene.Spheres) == 0 {
		return glm.Vec4{0.0, 0.0, 0.0, 1.0}
	}

	var closestSphere *scene.Sphere = nil
  var hitDistance float64 = math.MaxFloat64

	for _, sphere := range myScene.Spheres {
		origin := ray.Origin.Sub(sphere.Position)

		a := ray.Direction.Dot(ray.Direction)
		b := 2.0 * origin.Dot(ray.Direction)
		c := origin.Dot(origin) - (sphere.Radius * sphere.Radius)

		discriminant := (b * b) - (4 * a * c)

		if discriminant <= 0 {
			continue
		}

		// t0 := (-b + math.Sqrt(discriminant)) / (2.0 * a)
		closestT := (-b - math.Sqrt(discriminant)) / (2.0 * a)

    if closestT < hitDistance {
      hitDistance = closestT
      closestSphere = sphere
    }
	}

  if closestSphere == nil {
    return glm.Vec4{0.0, 0.0, 0.0, 1.0}
  }

  origin := ray.Origin.Sub(closestSphere.Position)

	hitPoint := origin.Add(ray.Direction.Mul(hitDistance))
	normal := hitPoint.Normalize()

	lightDir := glm.Vec3{-1.0, -1.0, -1.0}.Normalize()

	lightIntensity := math.Max(0, normal.Dot(lightDir.Mul(-1.0))) // cosine of angle between normal and light direction

	sphereColour := closestSphere.Albedo.Mul(lightIntensity)

	return sphereColour.Vec4(1.0)
}

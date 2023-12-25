package renderer

import (
	"math"

	"github.com/milindmadhukar/RayTracing/scene"
)

func (ray *Ray) ClosestHit(myScene *scene.Scene, hitDistance float64, objectIndex int) *hitPayLoad {
	closestSphere := myScene.Spheres[objectIndex]

	origin := ray.Origin.Sub(closestSphere.Position)

	worldPosition := origin.Add(ray.Direction.Mul(hitDistance))
	worldNormal := worldPosition.Normalize()
	worldPosition = worldPosition.Add(closestSphere.Position)

	return &hitPayLoad{
		HitDistance:   hitDistance,
		WorldPosition: worldPosition,
		WorldNormal:   worldNormal,
		ObjectIndex:   objectIndex,
	}
}

func (ray *Ray) Miss(myScene *scene.Scene) *hitPayLoad {
	return &hitPayLoad{
		HitDistance: -1,
	}
}

func (ray *Ray) TraceRay(myScene *scene.Scene) *hitPayLoad {
	var hitDistance float64 = math.MaxFloat64
	closestSphereIdx := -1

	for idx, sphere := range myScene.Spheres {
		origin := ray.Origin.Sub(sphere.Position)

		a := ray.Direction.Dot(ray.Direction)
		b := 2.0 * origin.Dot(ray.Direction)
		c := origin.Dot(origin) - (sphere.Radius * sphere.Radius)

		discriminant := (b * b) - (4 * a * c)

		if discriminant <= 0 {
			continue
		}

		closestT := (-b - math.Sqrt(discriminant)) / (2.0 * a)

		if closestT > 0 && closestT < hitDistance {
			hitDistance = closestT
			closestSphereIdx = idx
		}
	}

	if closestSphereIdx < 0 {
		return ray.Miss(myScene)
	}

	return ray.ClosestHit(myScene, hitDistance, closestSphereIdx)
}

package scene

import (
	"math/rand"
	"time"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/camera"
)

type Scene struct {
	Camera  *camera.Camera
	Spheres []*Sphere
  Materials []*Material
  Random  *rand.Rand
  MaxRayBounces int
  RaysPerPixel int
  MaxRayBounceDistance float64
}

func NewScene() *Scene {
  scene := &Scene{
		Camera:  camera.NewDefaultCamera(),
	}

  Material1 := &Material{
    Albedo: glm.Vec3{0.1, 0.3, 0.5},
    Roughness: 1.0,
    Metallic: 0.0,
  }

  Material2 := &Material{
    Albedo: glm.Vec3{0.8, 0.3, 0.2},
    Roughness: 0.5,
    Metallic: 0.1,
  }

  scene.Materials = append(scene.Materials, Material1, Material2)

  scene.Random = rand.New(rand.NewSource(time.Now().UnixNano()))

  sphere1 := NewDefaultSphere()
  sphere1.Radius = 1 
  sphere1.MaterialIndex = 0

  sphere2 := NewDefaultSphere()
  sphere2.Position = glm.Vec3{0, -101, -5}
  sphere2.Radius = 100
  sphere2.MaterialIndex = 1

  scene.Spheres = append(scene.Spheres, sphere1, sphere2)

  scene.MaxRayBounces = 2
  scene.RaysPerPixel = 1
  scene.MaxRayBounceDistance = 1000.0

  return scene
}

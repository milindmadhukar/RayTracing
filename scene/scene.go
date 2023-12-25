package scene

import (
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/camera"
)

type Scene struct {
	Camera  *camera.Camera
	Spheres []*Sphere
}

func NewScene() *Scene {
  scene := &Scene{
		Camera:  camera.NewDefaultCamera(),
	}


  sphere1 := NewDefaultSphere()

  sphere1.Albedo = glm.Vec3{0.5, 0.2, 0.9}
  sphere1.Position = glm.Vec3{1.5, 0.0, 0.5}
  sphere1.Radius = 0.8

  sphere2 := NewDefaultSphere()

  scene.Spheres = append(scene.Spheres, sphere1, sphere2)

  return scene
}

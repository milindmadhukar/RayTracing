package scene

import "github.com/milindmadhukar/RayTracing/camera"

type Scene struct {
	Camera *camera.Camera
}

func NewScene() *Scene {
	return &Scene{
		Camera: camera.NewDefaultCamera(),
	}
}

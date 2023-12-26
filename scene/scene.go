package scene

import (
	"image"
	"math/rand"
	"time"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/camera"
)

type Scene struct {
	Camera *camera.Camera `json:"camera"`

	FinalImage       *image.RGBA `json:"-"`
	AccumulatedImage []*glm.Vec4 `json:"-"`
	ToAccumulate     bool        `json:"-"`

	FrameIndex int `json:"frame_index"`

	Spheres []*Sphere `json:"spheres"`

	Materials []*Material `json:"materials"`

	Random               *rand.Rand `json:"-"`
	MaxRayBounces        int        `json:"max_ray_bounces"`
	RaysPerPixel         int        `json:"rays_per_pixel"`
	MaxRayBounceDistance float64    `json:"max_ray_bounce_distance"`
}

func NewScene() *Scene {

	// TODO: make example directory and load example scene from there.

	scene := &Scene{
		Camera: camera.NewDefaultCamera(),
	}

	scene.ToAccumulate = false
	scene.FrameIndex = 1

	material1 := NewDefaultMaterial()

	material2 := NewDefaultMaterial()
	material2.Albedo = glm.Vec3{0.8, 0.3, 0.2}
	material2.Roughness = 0.5
	material2.Metallic = 0.1

	material3 := NewDefaultMaterial()
	material3.Albedo = glm.Vec3{0.2, 0.7, 0.34}
	material3.Roughness = 0.02
	material3.Metallic = 0.9

	material4 := NewDefaultMaterial()
	material4.Albedo = glm.Vec3{0.8, 0.4, 0.6}
	material4.Roughness = 0.2
	material4.Metallic = 0.5
	material4.EmissionColor = material1.Albedo
	material4.EmissionPower = 2.0

	scene.Materials = append(scene.Materials, material1, material2, material3, material4)

	scene.Random = rand.New(rand.NewSource(time.Now().UnixNano()))

	sphere1 := NewDefaultSphere()
	sphere1.Radius = 1
	sphere1.MaterialIndex = 0

	sphere2 := NewDefaultSphere()
	sphere2.Position = glm.Vec3{0, -101, -5}
	sphere2.Radius = 100
	sphere2.MaterialIndex = 1

	sphere3 := NewDefaultSphere()
	sphere3.Position = glm.Vec3{-3.2, 1, 2}
	sphere3.Radius = 1.35
	sphere3.MaterialIndex = 2

	sphere4 := NewDefaultSphere()
	sphere4.Position = glm.Vec3{2.2, 1, 2}
	sphere4.Radius = 1.2
	sphere4.MaterialIndex = 3

	scene.Spheres = append(scene.Spheres, sphere1, sphere2, sphere3, sphere4)

	scene.MaxRayBounces = 3
	scene.RaysPerPixel = 2
	scene.MaxRayBounceDistance = 1000.0 // TODO: Not implemented yet.

	return scene
}

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

	SkyColor glm.Vec3 `json:"sky_color"`

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

	scene.SkyColor = glm.Vec3{0.6, 0.7, 0.9}

	scene.Random = rand.New(rand.NewSource(time.Now().UnixNano()))

	material1 := NewDefaultMaterial()
	scene.Materials = append(scene.Materials, material1)

	sphere1 := NewDefaultSphere()
	scene.Spheres = append(scene.Spheres, sphere1)

	scene.MaxRayBounces = 3
	scene.RaysPerPixel = 2
	scene.MaxRayBounceDistance = 1000.0 // TODO: Not implemented yet.

	return scene
}

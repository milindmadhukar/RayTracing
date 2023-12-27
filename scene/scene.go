package scene

import (
	"image"
	"io"
	"math/rand"
	"os"
	"time"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/camera"
	"github.com/milindmadhukar/RayTracing/utils"
)

type Scene struct {
	Camera *camera.Camera `json:"camera"`

	FinalImage       *image.RGBA `json:"-"`
	AccumulatedImage []*glm.Vec4 `json:"-"`
	ToAccumulate     bool        `json:"-"`
	IsRendering      bool        `json:"-"`

	FrameIndex int `json:"frame_index"`

	Spheres []*Sphere `json:"spheres"`

	Materials []*Material `json:"materials"`

	Random *utils.SafeRandom `json:"-"`

	SkyColor glm.Vec3 `json:"sky_color"`

	MaxRayBounces        int     `json:"max_ray_bounces"`
	RaysPerPixel         int     `json:"rays_per_pixel"`
	MaxRayBounceDistance float64 `json:"max_ray_bounce_distance"`
}

func NewScene() *Scene {
	basic, err := os.Open("example_scenes/rtiow_basic.json")
	if err != nil {
		panic(err)
	}

	reader := io.Reader(basic)
	scene, err := LoadStateFromJson(reader)
	if err != nil {
		panic(err)
	}

  scene.Random = utils.NewSafeRandom()

	// TODO: Max bounces not implemented yet.

	return scene
}

func NewRayTracingInOneWeekendScene() *Scene {

	scene := &Scene{}
	scene.Camera = camera.NewDefaultCamera()
	scene.Camera.VerticalFOV = 20
	scene.Camera.Position = glm.Vec3{13.0, 2.0, 3.0}
	lookAt := glm.Vec3{0.0, 0.0, 0.0}
	scene.Camera.ForwardDirection = utils.CalculateDirection(scene.Camera.Position, lookAt)

	scene.FrameIndex = 1

	scene.SkyColor = glm.Vec3{0.6, 0.7, 0.9}

	scene.MaxRayBounces = 50
	scene.RaysPerPixel = 500

	scene.MaxRayBounceDistance = 1000.0 // TODO: Not implemented yet.

	groundMaterial := NewDefaultMaterial()
	groundMaterial.Albedo = glm.Vec3{0.5, 0.5, 0.5}
	groundMaterial.Roughness = 0.9
	groundMaterial.Metallic = 0.0

	groundSphere := NewDefaultSphere()
	groundSphere.Position = glm.Vec3{0.0, -1000.0, 0.0}
	groundSphere.Radius = 1000.050
	groundSphere.MaterialIndex = 0

	scene.Materials = append(scene.Materials, groundMaterial)
	scene.Spheres = append(scene.Spheres, groundSphere)

	materail1 := NewDefaultMaterial()
	materail1.Albedo = glm.Vec3{0.2, 0.4, 0.6}
	materail1.Roughness = 0.05
	materail1.Metallic = 0.0

	materail2 := NewDefaultMaterial()
	materail2.Albedo = glm.Vec3{0.8, 0.6, 0.2}
	materail2.Roughness = 0.0
	materail2.Metallic = 0.5

	materail3 := NewDefaultMaterial()
	materail3.Albedo = glm.Vec3{0.9, 0.5, 0.5}
	materail3.Roughness = 0.2
	materail3.Metallic = 1.0

	materail4 := NewDefaultMaterial()
	materail4.Albedo = glm.Vec3{0.3, 0.7, 0.5}
	materail4.Roughness = 0.9
	materail4.Metallic = 0.7

	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := random.Float64()
			center := glm.Vec3{float64(a) + 0.9*random.Float64(), 0.2, float64(b) + 0.9*random.Float64()}
			if center.Sub(glm.Vec3{4.0, 0.2, 0.0}).Len() > 0.9 {
				if chooseMaterial < 0.8 { // Diffuse
					albedo := utils.ComponentWiseMultiplication(utils.InUnitSphere(scene.Random), utils.InUnitSphere(scene.Random))
					material := NewMaterial(albedo, 0.5, 0.0, glm.Vec3{0.0, 0.0, 0.0}, 0.0)
					sphere := NewSphere(center, 0.2, len(scene.Materials))
					scene.Materials = append(scene.Materials, material)
					scene.Spheres = append(scene.Spheres, sphere)
				} else if chooseMaterial < 0.95 { // Metal
					albedo := utils.InUnitSphere(scene.Random).Add(glm.Vec3{1.0, 1.0, 1.0}).Mul(0.5)
					roughness := 0.5 * random.Float64()
					material := NewMaterial(albedo, roughness, 1.0, glm.Vec3{0.0, 0.0, 0.0}, 0.0)
					sphere := NewSphere(center, 0.2, len(scene.Materials))
					scene.Materials = append(scene.Materials, material)
					scene.Spheres = append(scene.Spheres, sphere)
				} else { // Glass
					albedo := glm.Vec3{1.0, 1.0, 1.0}
					roughness := 0.0
					material := NewMaterial(albedo, roughness, 1.0, glm.Vec3{0.0, 0.0, 0.0}, 0.0)
					sphere := NewSphere(center, 0.2, len(scene.Materials))
					scene.Materials = append(scene.Materials, material)
					scene.Spheres = append(scene.Spheres, sphere)
				}
			}
		}
	}

	return scene
}

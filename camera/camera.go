package camera

import (
	"math"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/utils"
)

type Camera struct {
	ProjectionMatrix        glm.Mat4 `json:"-"`
	ViewMatrix              glm.Mat4 `json:"-"`
	InverseProjectionMatrix glm.Mat4 `json:"-"`
	InverseViewMatrix       glm.Mat4 `json:"-"`

	VerticalFOV   float64 `json:"vertical_fov"`
	NearClipPlane float64 `json:"near_clip_plane"`
	FarClipPlane  float64 `json:"far_clip_plane"`

	Position         glm.Vec3 `json:"position"`
	ForwardDirection glm.Vec3 `json:"forward_direction"`

	// Cached ray directions
	RayDirections []glm.Vec3 `json:"-"`

	LastMousePosition glm.Vec2 `json:"-"`

	RotationSpeed float64 `json:"rotation_speed"`

	ViewportWidth  int `json:"viewport_width"`
	ViewportHeight int `json:"viewport_height"`
}

func NewDefaultCamera() *Camera {
	camera := NewCamera(45.0, 0.1, 100.0)
	camera.RotationSpeed = 0.3

	return camera
}

func NewCamera(verticalFOV, NearClipPlane, FarClipPlane float64) *Camera {
	camera := &Camera{
		VerticalFOV:      verticalFOV,
		NearClipPlane:    NearClipPlane,
		FarClipPlane:     FarClipPlane,
		ForwardDirection: glm.Vec3{0.0, 0.0, -1.0},
		Position:         glm.Vec3{0.0, 0.0, 5.0},
	}

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if x != y {
				continue
			}
			camera.ProjectionMatrix[y*4+x] = 1.0
			camera.ViewMatrix[y*4+x] = 1.0
			camera.InverseProjectionMatrix[y*4+x] = 1.0
			camera.InverseViewMatrix[y*4+x] = 1.0
		}
	}

	return camera
}

func (camera *Camera) OnResize(width, height int) {
	camera.ViewportWidth = width
	camera.ViewportHeight = height

	camera.RecalculateProjectionMatrix()
	camera.RecalculateRayDirections()
}

func (camera *Camera) RecalculateProjectionMatrix() {
	camera.ProjectionMatrix = glm.Perspective(camera.VerticalFOV*(math.Pi/180), float64(camera.ViewportWidth)/float64(camera.ViewportHeight), camera.NearClipPlane, camera.FarClipPlane)
	camera.InverseProjectionMatrix = camera.ProjectionMatrix.Inv()

}

func (camera *Camera) RecalculateViewMatrix() {
	camera.ViewMatrix = glm.LookAtV(camera.Position, camera.Position.Add(camera.ForwardDirection), glm.Vec3{0.0, 1.0, 0.0})
	camera.InverseViewMatrix = camera.ViewMatrix.Inv()
}

func (camera *Camera) RecalculateRayDirections() {
	camera.RayDirections = make([]glm.Vec3, camera.ViewportWidth*camera.ViewportHeight)

	for y := 0; y < camera.ViewportHeight; y++ {
		for x := 0; x < camera.ViewportWidth; x++ {
			coord := glm.Vec2{
				float64(x)/float64(camera.ViewportWidth)*2.0 - 1.0,
				float64(y)/float64(camera.ViewportHeight)*2.0 - 1.0,
			}

			target := utils.MulMat4WithVec4(camera.InverseProjectionMatrix, glm.Vec4{coord.X(), coord.Y(), 1.0, 1.0})

			multiplier := target.Vec3().Mul(1.0 / target.W()).Normalize().Vec4(0)

			rayDirection := utils.MulMat4WithVec4(camera.InverseViewMatrix, multiplier).Vec3()

			camera.RayDirections[utils.FlattenXY(x, y, camera.ViewportWidth)] = rayDirection
		}
	}
}

package camera

import (
	"math"

	"fyne.io/fyne/v2"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/utils"
)

type Camera struct {
	ProjectionMatrix        glm.Mat4
	ViewMatrix              glm.Mat4
	InverseProjectionMatrix glm.Mat4
	InverseViewMatrix       glm.Mat4

	VerticalFOV   float64
	NearClipPlane float64
	FarClipPlane  float64

	Position         glm.Vec3
	ForwardDirection glm.Vec3

	// Cached ray directions
	RayDirections []glm.Vec3

	LastMousePosition glm.Vec2

	RotationSpeed float64

	ViewportWidth  int
	ViewportHeight int
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
		Position:         glm.Vec3{0.0, 0.0, 3.0},
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

func (camera *Camera) OnUpdate(renderedRaster fyne.CanvasObject) {

}

func (camera *Camera) OnResize(width, height int) {
	if width == camera.ViewportWidth && height == camera.ViewportHeight {
		return
	}

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

package scene

import (
	glm "github.com/go-gl/mathgl/mgl64"
)

type Sphere struct {
	Position glm.Vec3
	Radius   float64
	Albedo   glm.Vec3
}

func NewDefaultSphere() *Sphere {
	return &Sphere{
		Position: glm.Vec3{0, 0, 0},
		Radius:   0.5,
		Albedo:   glm.Vec3{1, 1, 1}, // White
	}
}

func NewSphere(position glm.Vec3, radius float64, albedo glm.Vec3) *Sphere {
	return &Sphere{
		Position: position,
		Radius:   radius,
		Albedo:   albedo,
	}
}

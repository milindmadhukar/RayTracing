package scene

import (
	glm "github.com/go-gl/mathgl/mgl64"
)

type Sphere struct {
	Position glm.Vec3 `json:"position"`
	Radius   float64  `json:"radius"`

	MaterialIndex int `json:"material_index"`
}

func NewDefaultSphere() *Sphere {
	return &Sphere{
		Position:      glm.Vec3{0, 0, 0},
		Radius:        0.5,
		MaterialIndex: 0,
	}
}

func NewSphere(position glm.Vec3, radius float64, materialIndex int) *Sphere {
	return &Sphere{
		Position:      position,
		Radius:        radius,
		MaterialIndex: materialIndex,
	}
}

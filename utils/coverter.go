package utils

import (
	"image/color"

	glm "github.com/go-gl/mathgl/mgl64"
)

func ConvertToRGBA(vecColour glm.Vec4) color.Color {
	// TODO: Clamp vector values to 0-1
	return color.RGBA{
		R: uint8(vecColour.X() * 255),
		G: uint8(vecColour.Y() * 255),
		B: uint8(vecColour.Z() * 255),
		A: uint8(vecColour.W() * 255),
	}
}

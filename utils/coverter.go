package utils

import (
	"image/color"

	glm "github.com/go-gl/mathgl/mgl64"
)

func ConvertToRGBA(vecColour glm.Vec4) color.Color {
	// TODO: Clamp vector values to 0-1

	return color.RGBA{
		R: uint8(glm.Clamp(vecColour.X(), 0, 1) * 255),
    G: uint8(glm.Clamp(vecColour.Y(), 0, 1) * 255),
    B: uint8(glm.Clamp(vecColour.Z(), 0, 1) * 255),
    A: uint8(glm.Clamp(vecColour.W(), 0, 1) * 255),
	}
}

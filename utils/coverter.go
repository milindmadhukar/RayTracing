package utils

import (
	"bytes"
	"image"
	"image/color"
	"image/png"

	glm "github.com/go-gl/mathgl/mgl64"
)

func ConvertToRGBA(vecColour glm.Vec4) color.Color {
	return color.RGBA{
		R: uint8(glm.Clamp(vecColour.X(), 0, 1) * 255),
		G: uint8(glm.Clamp(vecColour.Y(), 0, 1) * 255),
		B: uint8(glm.Clamp(vecColour.Z(), 0, 1) * 255),
		A: uint8(glm.Clamp(vecColour.W(), 0, 1) * 255),
	}
}

func Vec4ArrayToImage(vecColours []*glm.Vec4, width, height int) image.Image {
	finalImage := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			finalImage.Set(x, height-y, ConvertToRGBA(*vecColours[FlattenXY(x, y, width)]))
		}
	}

	return finalImage
}

func ImageToPNGBytes(img image.Image) []byte {
	buf := new(bytes.Buffer)
	if err := png.Encode(buf, img); err != nil {
		return nil
	}

  return buf.Bytes()
}

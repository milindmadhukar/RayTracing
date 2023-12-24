package utils

import (
	glm "github.com/go-gl/mathgl/mgl64"
)

func FlattenXY(x, y, width int) int {
	return y*width + x
}

func MulMat4WithVec4(matrix glm.Mat4, vector glm.Vec4) glm.Vec4 {
	return glm.Vec4{
		matrix[FlattenXY(0, 0, 4)]*vector[0] +
			matrix[FlattenXY(0, 1, 4)]*vector[1] +
			matrix[FlattenXY(0, 2, 4)]*vector[2] +
			matrix[FlattenXY(0, 3, 4)]*vector[3],
		matrix[FlattenXY(1, 0, 4)]*vector[0] +
			matrix[FlattenXY(1, 1, 4)]*vector[1] +
			matrix[FlattenXY(1, 2, 4)]*vector[2] +
			matrix[FlattenXY(1, 3, 4)]*vector[3],
		matrix[FlattenXY(2, 0, 4)]*vector[0] +
			matrix[FlattenXY(2, 1, 4)]*vector[1] +
			matrix[FlattenXY(2, 2, 4)]*vector[2] +
			matrix[FlattenXY(2, 3, 4)]*vector[3],
		matrix[FlattenXY(3, 0, 4)]*vector[0] +
			matrix[FlattenXY(3, 1, 4)]*vector[1] +
			matrix[FlattenXY(3, 2, 4)]*vector[2] +
			matrix[FlattenXY(3, 3, 4)]*vector[3],
	}
}

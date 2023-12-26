package utils

import (
	"math/rand"

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

func QuaternionCrossProduct(q1, q2 glm.Quat) glm.Quat {
	return glm.Quat{
		W: q1.W*q2.W - q1.X()*q2.X() - q1.Y()*q2.Y() - q1.Z()*q2.Z(),
		V: glm.Vec3{
			q1.W*q2.X() + q1.X()*q2.W + q1.Y()*q2.Z() - q1.Z()*q2.Y(),
			q1.W*q2.Y() + q1.Y()*q2.W + q1.Z()*q2.X() - q1.X()*q2.Z(),
			q1.W*q2.Z() + q1.Z()*q2.W + q1.X()*q2.Y() - q1.Y()*q2.X(),
		},
	}
}

func InUnitSphere(random *rand.Rand) glm.Vec3 {
	return glm.Vec3{
		random.Float64(),
		random.Float64(),
		random.Float64(),
	}.Mul(2.0).Sub(glm.Vec3{1.0, 1.0, 1.0})
}

func ComponentWiseMultiplication(vec1, vec2 glm.Vec3) glm.Vec3 {
	return glm.Vec3{
		vec1.X() * vec2.X(),
		vec1.Y() * vec2.Y(),
		vec1.Z() * vec2.Z(),
	}
}

func CalculateDirection(from, to glm.Vec3) glm.Vec3 {
  return to.Sub(from).Normalize()
}

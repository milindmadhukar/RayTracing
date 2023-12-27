package renderer

import (
	"image"

	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/utils"
)

type Ray struct {
	Origin    glm.Vec3
	Direction glm.Vec3
}

type hitPayLoad struct {
	HitDistance   float64
	WorldPosition glm.Vec3
	WorldNormal   glm.Vec3

	ObjectIndex int
}

func AcummulateImage(myScene *scene.Scene) *image.RGBA {
	finalImage := myScene.FinalImage

	width := myScene.Camera.ViewportWidth
	height := myScene.Camera.ViewportHeight

	if myScene.FrameIndex == 1 {
		myScene.AccumulatedImage = make([]*glm.Vec4, width*height)

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				myScene.AccumulatedImage[utils.FlattenXY(x, y, width)] = &glm.Vec4{0.0, 0.0, 0.0, 0.0}
			}
		}
	}

	// TODO: Allow to choose between render styles ie accumulate over real time or per pixel with defined sampling
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colour := PerPixel(x, y, width, height, myScene)
			newColour := myScene.AccumulatedImage[utils.FlattenXY(x, y, width)].Add(colour)
			myScene.AccumulatedImage[utils.FlattenXY(x, y, width)] = &newColour
			accumulatedColor := newColour.Mul(1.0 / float64(myScene.FrameIndex))

			finalImage.Set(x, height-y, utils.ConvertToRGBA(accumulatedColor))
		}
	}

	if myScene.ToAccumulate {
		myScene.FrameIndex++
	} else {
		myScene.FrameIndex = 1
	}

	return finalImage
}

func PerPixel(x, y, width, height int, myScene *scene.Scene) glm.Vec4 { // Ray Gen

	if (myScene.Camera.RayDirections) == nil {
		// HACK: This is a hack to make sure that the camera is initialized before the first render.
		myScene.Camera.OnResize(width, height)
	}

	ray := &Ray{
		Origin:    myScene.Camera.Position,
		Direction: myScene.Camera.RayDirections[utils.FlattenXY(x, y, width)],
	}

	contribution := glm.Vec3{1.0, 1.0, 1.0} // Throughput
	totalLight := glm.Vec3{0.0, 0.0, 0.0}

	// for rayCount := 0; rayCount < myScene.RaysPerPixel; rayCount++ { // NOTE: Seems kinda pointless since its random accumulation over time anyways
	for bounces := 0; bounces < myScene.MaxRayBounces; bounces++ {
		payload := ray.TraceRay(myScene)
		if payload.HitDistance < 0 {
			totalLight = totalLight.Add(utils.ComponentWiseMultiplication(myScene.SkyColor, contribution))
			break
		}

		sphere := myScene.Spheres[payload.ObjectIndex]
		material := myScene.Materials[sphere.MaterialIndex]

		contribution = utils.ComponentWiseMultiplication(contribution, myScene.Materials[payload.ObjectIndex].Albedo)
		totalLight = totalLight.Add(material.GetEmission())

		ray.Origin = payload.WorldPosition.Add(payload.WorldNormal.Mul(0.0001))

		vec1 := ray.Direction

		roughness := myScene.Materials[sphere.MaterialIndex].Roughness

		vec2 := payload.WorldNormal.Add(
			utils.ComponentWiseMultiplication(
				glm.Vec3{roughness, roughness, roughness},
				utils.InUnitSphere(myScene.Random),
			),
		)

		ray.Direction = vec1.Sub(vec2.Mul(2.0 * vec1.Dot(vec2)))
	}
	// }

	return totalLight.Vec4(1.0)
}

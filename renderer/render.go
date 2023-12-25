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

func GenerateImage(width, height int, scene *scene.Scene) image.Image {
	finalImage := image.NewRGBA(image.Rect(0, 0, width, height))

	if scene.FrameIndex == 1 || len(scene.AccumulatedImage) == 0 {
		scene.AccumulatedImage = make([]*glm.Vec4, width*height)

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				scene.AccumulatedImage[utils.FlattenXY(x, y, width)] = &glm.Vec4{0.0, 0.0, 0.0, 0.0}
			}
		}
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			colour := PerPixel(x, y, width, scene)

			newColour := scene.AccumulatedImage[utils.FlattenXY(x, y, width)].Add(colour)
			scene.AccumulatedImage[utils.FlattenXY(x, y, width)] = &newColour

			accumulatedColour := newColour.Mul(1.0 / float64(scene.FrameIndex))

			finalImage.Set(x, height-y, utils.ConvertToRGBA(accumulatedColour))
		}
	}

	if scene.ToAccumulate {
		scene.FrameIndex++
	} else {
		scene.FrameIndex = 1
	}

	scene.FinalImage = finalImage
	return finalImage
}

func PerPixel(x, y, width int, myScene *scene.Scene) glm.Vec4 { // Ray Gen
	ray := Ray{}
	ray.Origin = myScene.Camera.Position
	ray.Direction = myScene.Camera.RayDirections[utils.FlattenXY(x, y, width)]

	contribution := glm.Vec3{1.0, 1.0, 1.0} // Throughput

	totalLight := glm.Vec3{0.0, 0.0, 0.0}

	for rayCount := 0; rayCount < myScene.RaysPerPixel; rayCount++ {
		for bounces := 0; bounces < myScene.MaxRayBounces; bounces++ {
			payload := ray.TraceRay(myScene)
			if payload.HitDistance < 0 {
				skyColor := glm.Vec3{0.6, 0.7, 0.9}
				// skyColor := glm.Vec3{0.0, 0.0, 0.0}
				totalLight = totalLight.Add(utils.ComponentWiseMultiplication(skyColor, contribution))
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
          utils.InUnitSphere(myScene.Random,
          ),
        ),
      )
			
			ray.Direction = vec1.Sub(vec2.Mul(2.0 * vec1.Dot(vec2)))
		}
	}

	return totalLight.Mul(1.0 / float64(myScene.RaysPerPixel)).Vec4(1.0)
}

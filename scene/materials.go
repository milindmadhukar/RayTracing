package scene

import glm "github.com/go-gl/mathgl/mgl64"

type Material struct {
	Albedo    glm.Vec3
	Roughness float64
	Metallic  float64
  EmissionColor glm.Vec3
  EmissionPower float64
}


func NewMaterial(albedo glm.Vec3, roughness, metallic float64, emmissionColor glm.Vec3, EmissionStrength float64) *Material {
	return &Material{
		Albedo:           albedo,
		Roughness:        roughness,
		Metallic:         metallic,
		EmissionColor:    emmissionColor,
		EmissionPower: EmissionStrength,
	}
}

func NewDefaultMaterial() *Material {
	material := Material{}
	material.Albedo = glm.Vec3{0.1, 0.3, 0.5}
	material.Roughness = 1.0
	material.Metallic = 0.0

	material.EmissionColor = glm.Vec3{0.0, 0.0, 0.0}
	material.EmissionPower = 0.0

	return &material
}

func (material *Material) GetEmission() glm.Vec3 {
  return material.EmissionColor.Mul(material.EmissionPower)
}

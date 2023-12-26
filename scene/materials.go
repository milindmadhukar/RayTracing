package scene

import glm "github.com/go-gl/mathgl/mgl64"

type Material struct {
	Albedo        glm.Vec3 `json:"albedo"`
	Roughness     float64  `json:"roughness"`
	Metallic      float64  `json:"metallic"`
	EmissionColor glm.Vec3 `json:"emission_color"`
	EmissionPower float64  `json:"emission_power"`
}

func NewMaterial(albedo glm.Vec3, roughness, metallic float64, emmissionColor glm.Vec3, EmissionStrength float64) *Material {
	return &Material{
		Albedo:        albedo,
		Roughness:     roughness,
		Metallic:      metallic,
		EmissionColor: emmissionColor,
		EmissionPower: EmissionStrength,
	}
}

func NewDefaultMaterial() *Material {
	material := Material{}
	material.Albedo = glm.Vec3{0.1, 0.3, 0.5}
	material.Roughness = 0.5
	material.Metallic = 0.0

	material.EmissionColor = glm.Vec3{0.2, 0.3, 0.4}
	material.EmissionPower = 3.0

	return &material
}

func (material *Material) GetEmission() glm.Vec3 {
	return material.EmissionColor.Mul(material.EmissionPower)
}

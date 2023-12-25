package window

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/scene"
)

func getSpherePositionUI(sphere *scene.Sphere) *fyne.Container {

	xPos := widget.NewEntry()
	xPos.SetPlaceHolder("X:")
	xPos.MultiLine = false
	xPos.SetText(strconv.FormatFloat(sphere.Position.X(), 'f', 2, 64))

	yPos := widget.NewEntry()
	yPos.SetPlaceHolder("Y:")
	yPos.MultiLine = false
	yPos.SetText(strconv.FormatFloat(sphere.Position.Y(), 'f', 2, 64))

	zPos := widget.NewEntry()
	zPos.SetPlaceHolder("Z:")
	zPos.MultiLine = false
	zPos.SetText(strconv.FormatFloat(sphere.Position.Z(), 'f', 2, 64))

	xPos.OnChanged = func(text string) {
		xPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		sphere.Position = glm.Vec3{xPosValue, sphere.Position.Y(), sphere.Position.Z()}
	}

	yPos.OnChanged = func(text string) {
		yPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		sphere.Position = glm.Vec3{sphere.Position.X(), yPosValue, sphere.Position.Z()}
	}

	zPos.OnChanged = func(text string) {
		zPosValue, err := strconv.ParseFloat(text, 64)
		if err != nil {
			return
		}
		sphere.Position = glm.Vec3{sphere.Position.X(), sphere.Position.Y(), zPosValue}
	}

	posLabel := widget.NewLabel("Sphere Position:")
	posContainer := container.New(layout.NewGridLayout(3), xPos, yPos, zPos)

	return container.New(layout.NewVBoxLayout(), posLabel, posContainer)
}

func getSphereRadiusUI(sphere *scene.Sphere) *fyne.Container {
	radiusLabel := widget.NewLabel("Radius:")

	radiusBind := binding.BindFloat(&sphere.Radius)
	radius := widget.NewEntryWithData(binding.FloatToString(radiusBind))
	radius.SetPlaceHolder("Radius:")
	radius.MultiLine = false

	return container.New(layout.NewGridLayout(2), radiusLabel, radius)

}

func getSphereMaterialIndexUI(sphere *scene.Sphere, myScene *scene.Scene) *fyne.Container {
	materialLabel := widget.NewLabel("Material:")

	materialEntry := widget.NewEntry()
	materialEntry.SetPlaceHolder("Material Index:")
	materialEntry.MultiLine = false
	materialEntry.SetText(strconv.Itoa(sphere.MaterialIndex))

	materialEntry.OnChanged = func(text string) {
		materialIndex, err := strconv.Atoi(text)
		if err != nil {
			return
		}
		if materialIndex >= len(myScene.Materials) {
			materialIndex = len(myScene.Materials) - 1
			materialEntry.SetText(strconv.Itoa(materialIndex))
		}

		sphere.MaterialIndex = materialIndex
	}

	return container.New(layout.NewGridLayout(2), materialLabel, materialEntry)
}

func getSphereUI(sphere *scene.Sphere, index int, myScene *scene.Scene) *fyne.Container {
		sphereLabel := widget.NewLabel("Sphere " + strconv.Itoa(index) + " Settings")
		posContainer := getSpherePositionUI(sphere)
		radiusContainer := getSphereRadiusUI(sphere)
		sphereMaterialIndexContainer := getSphereMaterialIndexUI(sphere, myScene)
		return container.New(layout.NewVBoxLayout(), sphereLabel, posContainer, radiusContainer, sphereMaterialIndexContainer)


}

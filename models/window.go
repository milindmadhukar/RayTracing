package models

import (
	"image"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2"
)

type Window struct {
  FyneApp fyne.App
  FyneWindow fyne.Window
  Image image.Image
  SettingsContainer *fyne.Container
  RenderContainer *canvas.Image
}

package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type Window struct {
  FyneApp fyne.App
  FyneWindow fyne.Window
  SettingsContainer *fyne.Container
  RenderContainer *canvas.Raster
  FPSLabel *widget.Label
}

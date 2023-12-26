package window

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/utils"
	"golang.design/x/clipboard"
)

func getAccumulationUI(scene *scene.Scene, applicationWindow *Window) *fyne.Container {
	copyImagetoClipboardBtn := widget.NewButton(
		"Copy Image to Clipboard",
		func() {
			clipboard.Write(clipboard.FmtImage, utils.ImageToPNGBytes(scene.FinalImage))
		},
	)
	exportImageBtn := widget.NewButton(
		"Export Image",
		func() {
			dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
				if err != nil {
					return
				}
				if writer == nil {
					return
				}
				writer.Write(utils.ImageToPNGBytes(scene.FinalImage))
			}, applicationWindow.FyneWindow)
		},
	)

	//FIXME: Only works when clicked multiple times
	resetAccumulationBtn := widget.NewButton(
		"Reset Accumulation",
		func() {
			scene.FrameIndex = 1
		},
	)

	return container.New(layout.NewGridLayout(3), copyImagetoClipboardBtn, exportImageBtn, resetAccumulationBtn)
}

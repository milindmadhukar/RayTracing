package renderer

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2/widget"
)

func updateFPSLabel(fpsLabel *widget.Label, frameTime time.Duration) {
	fps := int(1 / frameTime.Seconds())
	frameTimeMs := strconv.FormatFloat(float64(frameTime.Milliseconds()), 'f', 2, 64)
	fpsLabel.SetText("FPS: " + strconv.Itoa(fps) + " (Render time: " + frameTimeMs + "ms)")
}

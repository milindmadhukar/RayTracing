package window

import (
	"image"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	glm "github.com/go-gl/mathgl/mgl64"
	"github.com/milindmadhukar/RayTracing/camera"
	"github.com/milindmadhukar/RayTracing/renderer"
	"github.com/milindmadhukar/RayTracing/scene"
	"github.com/milindmadhukar/RayTracing/utils"
)

type RenderedRaster struct {
	widget.BaseWidget
	rasterImage *canvas.Raster
	camera      *camera.Camera
	window      *Window
}

func (applicationWindow *Window) accumulatedRender(scene *scene.Scene) image.Image {
	now := time.Now()
	renderedImage := renderer.AcummulateImage(scene)
	frameTime := time.Since(now)
	fps := int(1 / frameTime.Seconds())
	frameTimeMs := strconv.FormatFloat(float64(frameTime.Milliseconds()), 'f', 2, 64)
	applicationWindow.TimeLabel.SetText("FPS: " + strconv.Itoa(fps) + " (Render time: " + frameTimeMs + "ms)")
	return renderedImage
}

func (applicationWindow *Window) perPixelRender(scene *scene.Scene) {
	width := scene.Camera.ViewportWidth
	height := scene.Camera.ViewportHeight

	applicationWindow.TimeLabel.SetText("Rendering...")

	now := time.Now()
	for y := height - 1; y > 0; y-- {
		for x := 0; x < width; x++ {
      // NOTE: Running this in a goroutine makes all pixels render together
			finalColor := glm.Vec4{0.0, 0.0, 0.0, 0.0}
			for i := 0; i < scene.RaysPerPixel; i++ {
				finalColor = finalColor.Add(renderer.PerPixel(x, y, width, height, scene))
			}
			scene.FinalImage.Set(x, height-y, utils.ConvertToRGBA(finalColor.Mul(1.0/float64(scene.RaysPerPixel))))
		}
	}

	renderTime := time.Since(now)
	if renderTime.Seconds() > 0.1 {
		applicationWindow.TimeLabel.SetText("Render time: " + strconv.FormatFloat(renderTime.Seconds(), 'f', 2, 64) + "s")
	} else {
		applicationWindow.TimeLabel.SetText("Render time: " + strconv.Itoa(int(renderTime.Milliseconds())) + "ms")
	}

}

func (applicationWindow *Window) GetRenderedImage(scene *scene.Scene) *canvas.Raster {
	return canvas.NewRaster(
		func(width, height int) image.Image {

			if width != scene.Camera.ViewportWidth || height != scene.Camera.ViewportHeight {
				scene.Camera.OnResize(width, height)
				scene.FrameIndex = 1
				scene.FinalImage = image.NewRGBA(image.Rect(0, 0, width, height))
				scene.IsRendering = false
			}

			if scene.ToAccumulate {
				scene.IsRendering = false
				return applicationWindow.accumulatedRender(scene)
			}

			if !scene.IsRendering {
				scene.FinalImage = image.NewRGBA(image.Rect(0, 0, width, height))
				scene.IsRendering = true
				go applicationWindow.perPixelRender(scene)
			}

			return scene.FinalImage
		},
	)
}

func NewRenderedRaster(applicationWindow *Window, scene *scene.Scene) *RenderedRaster {
	renderedRaster := &RenderedRaster{
		rasterImage: applicationWindow.GetRenderedImage(scene),
	}
	renderedRaster.ExtendBaseWidget(renderedRaster)
	renderedRaster.camera = scene.Camera
	renderedRaster.window = applicationWindow

	return renderedRaster
}

func (renderedRaster *RenderedRaster) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(renderedRaster.rasterImage)
}

func (renderedRaster *RenderedRaster) MouseDown(mouseEvent *desktop.MouseEvent) {
	if mouseEvent.Button != desktop.MouseButtonSecondary {
		return
	}

	var moved bool = false

	// FIXME: Its broken
	mousePosition := glm.Vec2{float64(mouseEvent.AbsolutePosition.X), float64(mouseEvent.AbsolutePosition.Y)}
	delta := mousePosition.Sub(renderedRaster.camera.LastMousePosition).Mul(0.0008)
	renderedRaster.camera.LastMousePosition = mousePosition

	if delta.X() != 0.0 || delta.Y() != 0.0 {
		upDirection := glm.Vec3{0.0, 1.0, 0.0}
		rightDirection := renderedRaster.camera.ForwardDirection.Cross(upDirection)

		pitchDelta := delta.Y() * renderedRaster.camera.RotationSpeed
		yawDelta := delta.X() * renderedRaster.camera.RotationSpeed

		q := glm.QuatRotate(-pitchDelta, rightDirection).Mul(glm.QuatRotate(-yawDelta, upDirection)).Normalize()

		// q := utils.QuaternionCrossProduct(glm.QuatRotate(-pitchDelta, rightDirection), glm.QuatRotate(-yawDelta, upDirection)).Normalize()

		renderedRaster.camera.ForwardDirection = q.Rotate(renderedRaster.camera.ForwardDirection)

		moved = true
	}

	if moved {
		renderedRaster.camera.RecalculateViewMatrix()
		renderedRaster.camera.RecalculateRayDirections()
		renderedRaster.window.Update()
		// FIXME: This did not work, maybe I have to update the binds
		renderedRaster.window.CameraPositionContainer.Refresh()
		// TODO: Set frame index to 1
	}
}

func (renderedRaster *RenderedRaster) MouseUp(mouseEvent *desktop.MouseEvent) {
}

func (renderedRaster *RenderedRaster) Tapped(tapEvent *fyne.PointEvent) {
	renderedRaster.window.FyneWindow.Canvas().Focus(renderedRaster)
}

func (renderedRaster *RenderedRaster) KeyDown(keyEvent *fyne.KeyEvent) {
	var moved bool = false

	upDirection := glm.Vec3{0.0, 1.0, 0.0}
	rightDirection := renderedRaster.camera.ForwardDirection.Cross(upDirection)

	speed := 0.1

	// TODO: Maybe add some keyboard shortcuts.
	// Maybe a top bar for import export help github etc.

	switch keyEvent.Name {
	case fyne.KeyW:
		renderedRaster.camera.Position = renderedRaster.camera.Position.Add(renderedRaster.camera.ForwardDirection.Mul(speed))
		moved = true
	case fyne.KeyS:
		renderedRaster.camera.Position = renderedRaster.camera.Position.Sub(renderedRaster.camera.ForwardDirection.Mul(speed))
		moved = true
	case fyne.KeyA:
		renderedRaster.camera.Position = renderedRaster.camera.Position.Sub(rightDirection.Mul(speed))
		moved = true
	case fyne.KeyD:
		renderedRaster.camera.Position = renderedRaster.camera.Position.Add(rightDirection.Mul(speed))
		moved = true
	case fyne.KeyQ:
		renderedRaster.camera.Position = renderedRaster.camera.Position.Sub(upDirection.Mul(speed))
		moved = true
	case fyne.KeyE:
		renderedRaster.camera.Position = renderedRaster.camera.Position.Add(upDirection.Mul(speed))
		moved = true
	case fyne.KeyBackTick:
		// TODO: Find a way to hide the settings panel
	case fyne.KeyEscape:
		renderedRaster.window.FyneApp.Quit()
	}

	if moved {
		renderedRaster.camera.RecalculateViewMatrix()
		renderedRaster.camera.RecalculateRayDirections()
		renderedRaster.window.Update()
		// FIXME: This did not work, maybe I have to update the binds
		renderedRaster.window.CameraPositionContainer.Refresh()
		// TODO: Set frame index to 1
	}
}

func (renderedRaster *RenderedRaster) KeyUp(keyEvent *fyne.KeyEvent) {
}

func (renderedRaster *RenderedRaster) FocusGained() {
}

func (renderedRaster *RenderedRaster) FocusLost() {
}

func (renderedRaster *RenderedRaster) TypedRune(r rune) {
}

func (renderedRaster *RenderedRaster) TypedKey(keyEvent *fyne.KeyEvent) {
}

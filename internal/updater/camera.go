package updater

import (
	"ant-sim/internal/state"

	"github.com/hajimehoshi/ebiten/v2"
)

func UpdateCamera(w *state.World) {
	// 1. Keyboard Movement (Existing)
	const camSpeed = 10.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		w.CameraPosition.Y -= camSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		w.CameraPosition.Y += camSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		w.CameraPosition.X -= camSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		w.CameraPosition.X += camSpeed
	}

	// 2. Mouse Drag (Middle button or Left button)
	currX, currY := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		dx := float64(currX - w.LastMouseX)
		dy := float64(currY - w.LastMouseY)

		w.CameraPosition.X -= dx / w.Zoom
		w.CameraPosition.Y -= dy / w.Zoom
	}
	w.LastMouseX, w.LastMouseY = currX, currY

	// 3. Scroll Wheel Zoom
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		oldZoom := w.Zoom
		w.Zoom += wheelY * 0.1
		if w.Zoom < 0.1 {
			w.Zoom = 0.1
		}
		if w.Zoom > 5.0 {
			w.Zoom = 5.0
		}

		w.CameraPosition.X += (float64(currX) / oldZoom) - (float64(currX) / w.Zoom)
		w.CameraPosition.Y += (float64(currY) / oldZoom) - (float64(currY) / w.Zoom)
	}
}

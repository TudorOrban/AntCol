package mapgen

import (
	"ant-sim/internal/shared"
	"ant-sim/internal/state"
	"math/rand/v2"
)

func generateObstacles(w *state.World) {
	for _ = range w.Config.Map.NumObstacles {
		posX := rand.Float64() * float64(w.Width)
		posY := rand.Float64() * float64(w.Height)
		width := rand.Float64() * w.Config.Map.ObstacleMaxLength

		rect := shared.Rectangle{
			X:      posX,
			Y:      posY,
			Width:  width,
			Height: 50,
		}
		addWall(w, rect)
	}
}

func addWall(w *state.World, rect shared.Rectangle) {
	w.WallRects = append(w.WallRects, rect)

	for y := int(rect.Y); y < int(rect.Y+rect.Height); y++ {
		for x := int(rect.X); x < int(rect.X+rect.Width); x++ {
			if x >= 0 && x < w.Width && y >= 0 && y < w.Height {
				idx := (y/shared.GridScale)*w.GridWidth + x/shared.GridScale
				w.Obstacles[idx] = true
			}
		}
	}
}

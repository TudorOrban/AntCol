package world

import "ant-sim/internal/shared"

func (w *World) AddWall(rect shared.Rectangle) {
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

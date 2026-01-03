package world

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"

	"github.com/hajimehoshi/ebiten/v2"
)

func (w *World) Draw(screen *ebiten.Image) {
	w.drawHome(screen, w.HomePosition)
	w.drawPheromones(screen)

	for _, foodSource := range w.FoodSources {
		w.drawFoodSource(screen, foodSource)
	}

	for _, ant := range w.Ants {
		w.drawAnt(screen, ant)
	}
}

func (w *World) drawHome(screen *ebiten.Image, homePosition shared.Position) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-HomeRadius/2, -HomeRadius/2)

	opts.GeoM.Translate(homePosition.X, homePosition.Y)

	screen.DrawImage(w.HomeImage, opts)
}

func (w *World) drawAnt(screen *ebiten.Image, ant ant.Ant) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-AntLength/2, -AntWidth/2)

	opts.GeoM.Rotate(ant.AngleRadians)

	opts.GeoM.Translate(ant.Position.X, ant.Position.Y)

	screen.DrawImage(w.AntImage, opts)
}

func (w *World) drawFoodSource(screen *ebiten.Image, foodSource shared.FoodSource) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-MaxFoodSourceRadius/2, -MaxFoodSourceRadius/2)

	opts.GeoM.Translate(foodSource.Position.X, foodSource.Position.Y)

	screen.DrawImage(w.FoodSourceImage, opts)
}

func (w *World) drawPheromones(screen *ebiten.Image) {
	for i := 0; i < len(w.HomePheromones); i++ {
		pixIdx := i * 4
		home := w.HomePheromones[i]
		food := w.FoodPheromones[i]

		strength := clamp((home + food) * 255)

		w.PixelBuffer[pixIdx] = uint8(clamp(food * 255))   // Red component
		w.PixelBuffer[pixIdx+1] = 0                        // Green
		w.PixelBuffer[pixIdx+2] = uint8(clamp(home * 255)) // Blue component
		w.PixelBuffer[pixIdx+3] = uint8(strength)
	}

	w.PheromoneImage.WritePixels(w.PixelBuffer)

	screen.DrawImage(w.PheromoneImage, nil)
}

// Utils
func clamp(v float64) float64 {
	if v > 255 {
		return 255
	}
	if v < 0 {
		return 0
	}
	return v
}

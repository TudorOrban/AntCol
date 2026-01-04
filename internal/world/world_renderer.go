package world

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"

	"github.com/hajimehoshi/ebiten/v2"
)

func (w *World) Draw(screen *ebiten.Image) {
	w.drawPheromones(screen)
	w.drawHome(screen, w.HomePosition)

	for _, foodSource := range w.FoodSources {
		w.drawFoodSource(screen, foodSource)
	}

	for _, ant := range w.Ants {
		w.drawAnt(screen, ant)
	}
}

func (w *World) drawHome(screen *ebiten.Image, homePosition shared.Position) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-HomeRadius, -HomeRadius)

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

	opts.GeoM.Translate(-MaxFoodSourceRadius, -MaxFoodSourceRadius)

	opts.GeoM.Translate(foodSource.Position.X, foodSource.Position.Y)

	screen.DrawImage(w.FoodSourceImage, opts)
}

func (w *World) drawPheromones(screen *ebiten.Image) {
	for i := 0; i < len(w.HomePheromones); i++ {
		pixIdx := i * 4
		home := w.HomePheromones[i]
		food := w.FoodPheromones[i]

		strength := clamp((home + food) * 255)

		w.PixelBuffer[pixIdx] = uint8(clamp(food * 255))
		w.PixelBuffer[pixIdx+1] = 0
		w.PixelBuffer[pixIdx+2] = uint8(clamp(home * 255))
		w.PixelBuffer[pixIdx+3] = uint8(strength)
	}

	w.PheromoneImage.WritePixels(w.PixelBuffer)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(float64(shared.GridScale), float64(shared.GridScale))

	screen.DrawImage(w.PheromoneImage, opts)
}

func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}

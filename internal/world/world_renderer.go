package world

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (w *World) Draw(screen *ebiten.Image) {
	screen.DrawImage(w.GrassBackground, nil)
	w.drawPheromones(screen)
	w.drawHome(screen, w.HomePosition)

	for i := range w.FoodSources {
		w.drawFoodSource(screen, &w.FoodSources[i])
	}

	for i := range w.Ants {
		w.drawAnt(screen, &w.Ants[i])
	}

	for i := range w.WallRects {
		w.drawWall(screen, &w.WallRects[i])
	}
}

func (w *World) drawHome(screen *ebiten.Image, homePosition shared.Position) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-HomeRadius, -HomeRadius)

	opts.GeoM.Translate(homePosition.X, homePosition.Y)

	w.applyCamera(opts)

	screen.DrawImage(w.HomeImage, opts)
}

func (w *World) drawAnt(screen *ebiten.Image, ant *ant.Ant) {
	opts := &ebiten.DrawImageOptions{}
	imgW, imgH := w.AntImage.Bounds().Dx(), w.AntImage.Bounds().Dy()

	opts.GeoM.Translate(-float64(imgW)/2, -float64(imgH)/2)

	opts.GeoM.Rotate(ant.AngleRadians + math.Pi/2)

	opts.GeoM.Translate(ant.Position.X, ant.Position.Y)

	w.applyCamera(opts)

	screen.DrawImage(w.AntImage, opts)
}

func (w *World) drawFoodSource(screen *ebiten.Image, foodSource *shared.FoodSource) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-MaxFoodSourceRadius, -MaxFoodSourceRadius)

	opts.GeoM.Translate(foodSource.Position.X, foodSource.Position.Y)

	w.applyCamera(opts)

	screen.DrawImage(w.FoodSourceImage, opts)
}

func (w *World) drawPheromones(screen *ebiten.Image) {
	for i := 0; i < len(w.HomePheromones); i++ {
		pixIdx := i * 4
		home := w.HomePheromones[i]
		food := w.FoodPheromones[i]

		strength := clamp((home + food) * 255)

		w.PixelBuffer[pixIdx] = uint8(clamp(food * 255))
		w.PixelBuffer[pixIdx+1] = 80
		w.PixelBuffer[pixIdx+2] = uint8(clamp(home * 255))
		w.PixelBuffer[pixIdx+3] = uint8(strength)
	}

	w.PheromoneImage.WritePixels(w.PixelBuffer)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(float64(shared.GridScale), float64(shared.GridScale))

	w.applyCamera(opts)

	screen.DrawImage(w.PheromoneImage, opts)
}

func (w *World) drawWall(screen *ebiten.Image, rect *shared.Rectangle) {
	opts := &ebiten.DrawImageOptions{}
	sw, sh := w.ObstacleImage.Bounds().Dx(), w.ObstacleImage.Bounds().Dy()

	opts.GeoM.Scale(rect.Width/float64(sw), rect.Height/float64(sh))

	opts.GeoM.Translate(rect.X, rect.Y)

	w.applyCamera(opts)

	screen.DrawImage(w.ObstacleImage, opts)
}

func (w *World) applyCamera(opts *ebiten.DrawImageOptions) {
	opts.GeoM.Translate(-w.CameraPosition.X, -w.CameraPosition.Y)

	opts.GeoM.Scale(w.Zoom, w.Zoom)
}

// Utils
func clamp(v float64) float64 {
	if v < 0 {
		return 0
	}
	if v > 255 {
		return 255
	}
	return v
}

func (w *World) ScreenToWorld(screenX, screenY int) (float64, float64) {
	worldX := (float64(screenX) / w.Zoom) + w.CameraPosition.X
	worldY := (float64(screenY) / w.Zoom) + w.CameraPosition.Y
	return worldX, worldY
}

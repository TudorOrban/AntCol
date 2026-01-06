package renderer

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
	"ant-sim/internal/state"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func Draw(screen *ebiten.Image, w *state.World) {
	screen.DrawImage(w.GrassBackground, nil)
	drawPheromones(screen, w)
	drawHome(screen, w)

	for i := range w.FoodSources {
		drawFoodSource(screen, w, &w.FoodSources[i])
	}

	for i := range w.Ants {
		drawAnt(screen, w, &w.Ants[i])
	}

	for i := range w.WallRects {
		drawWall(screen, w, &w.WallRects[i])
	}
}

func drawHome(screen *ebiten.Image, w *state.World) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-w.Config.Map.HomeRadius, -w.Config.Map.HomeRadius)

	opts.GeoM.Translate(w.HomePosition.X, w.HomePosition.Y)

	applyCamera(w, opts)

	screen.DrawImage(w.HomeImage, opts)
}

func drawAnt(screen *ebiten.Image, w *state.World, ant *ant.Ant) {
	opts := &ebiten.DrawImageOptions{}
	imgW, imgH := w.AntImage.Bounds().Dx(), w.AntImage.Bounds().Dy()

	opts.GeoM.Translate(-float64(imgW)/2, -float64(imgH)/2)

	opts.GeoM.Rotate(ant.AngleRadians + math.Pi/2)

	opts.GeoM.Translate(ant.Position.X, ant.Position.Y)

	applyCamera(w, opts)

	screen.DrawImage(w.AntImage, opts)
}

func drawFoodSource(screen *ebiten.Image, w *state.World, foodSource *shared.FoodSource) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-w.Config.Food.MaxFoodSourceRadius, -w.Config.Food.MaxFoodSourceRadius)

	opts.GeoM.Translate(foodSource.Position.X, foodSource.Position.Y)

	applyCamera(w, opts)

	screen.DrawImage(w.FoodSourceImage, opts)
}

func drawPheromones(screen *ebiten.Image, w *state.World) {
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

	applyCamera(w, opts)

	screen.DrawImage(w.PheromoneImage, opts)
}

func drawWall(screen *ebiten.Image, w *state.World, rect *shared.Rectangle) {
	opts := &ebiten.DrawImageOptions{}
	sw, sh := w.ObstacleImage.Bounds().Dx(), w.ObstacleImage.Bounds().Dy()

	opts.GeoM.Scale(rect.Width/float64(sw), rect.Height/float64(sh))

	opts.GeoM.Translate(rect.X, rect.Y)

	applyCamera(w, opts)

	screen.DrawImage(w.ObstacleImage, opts)
}

func applyCamera(w *state.World, opts *ebiten.DrawImageOptions) {
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

func ScreenToWorld(w *state.World, screenX, screenY int) (float64, float64) {
	worldX := (float64(screenX) / w.Zoom) + w.CameraPosition.X
	worldY := (float64(screenY) / w.Zoom) + w.CameraPosition.Y
	return worldX, worldY
}

package mapgen

import (
	"github.com/hajimehoshi/ebiten/v2"

	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
	"ant-sim/internal/state"
)

func GenerateWorld(w, h int) *state.World {
	gw, gh := w/shared.GridScale, h/shared.GridScale

	grassTexture, antImage, homeImage, foodSourceImage, obstacleImage := shared.LoadAssets()

	world := &state.World{
		Width:           w,
		Height:          h,
		Config:          state.DefaultConfig(),
		GridWidth:       gw,
		GridHeight:      gh,
		HomePosition:    shared.Position{X: float64(w / 2), Y: float64(h / 2)},
		Ants:            []ant.Ant{},
		FoodSources:     []shared.FoodSource{},
		WallRects:       []shared.Rectangle{},
		Obstacles:       make([]bool, gw*gh),
		HomePheromones:  make([]float64, gw*gh),
		FoodPheromones:  make([]float64, gw*gh),
		HomeTemp:        make([]float64, gw*gh),
		FoodTemp:        make([]float64, gw*gh),
		CameraPosition:  shared.Position{X: 0, Y: 0},
		Zoom:            1,
		GrassBackground: tileBackground(w, h, grassTexture),
		HomeImage:       homeImage,
		AntImage:        antImage,
		FoodSourceImage: foodSourceImage,
		PheromoneImage:  ebiten.NewImage(gw, gh),
		ObstacleImage:   obstacleImage,
		PixelBuffer:     make([]byte, gw*gh*4),
		FoodCollected:   0,
		TotalTicks:      0,
	}

	generateAnts(world)
	generateFoodSources(world)
	generateObstacles(world)

	return world
}

func tileBackground(w, h int, grassTexture *ebiten.Image) *ebiten.Image {
	bakedBackground := ebiten.NewImage(w, h)

	tileW, tileH := grassTexture.Bounds().Dx(), grassTexture.Bounds().Dy()

	for y := 0; y < w; y += tileH {
		for x := 0; x < w; x += tileW {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			bakedBackground.DrawImage(grassTexture, op)
		}
	}

	return bakedBackground
}

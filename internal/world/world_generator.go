package world

import (
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"

	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
)

func GenerateWorld(w, h int) *World {
	gw, gh := w/shared.GridScale, h/shared.GridScale

	homePosition := shared.Position{X: float64(w / 2), Y: float64(h / 2)}
	ants := generateAnts(w, h, homePosition)
	foodSources := generateFoodSources(w, h, homePosition)

	grassTexture, antImage, homeImage, foodSourceImage := shared.LoadAssets()

	tiledBackground := tileBackground(w, h, grassTexture)

	return &World{
		Width:           w,
		Height:          h,
		GridWidth:       gw,
		GridHeight:      gh,
		HomePosition:    homePosition,
		Ants:            ants,
		FoodSources:     foodSources,
		HomePheromones:  make([]float64, gw*gh),
		FoodPheromones:  make([]float64, gw*gh),
		HomeTemp:        make([]float64, gw*gh),
		FoodTemp:        make([]float64, gw*gh),
		GrassBackground: tiledBackground,
		HomeImage:       homeImage,
		AntImage:        antImage,
		FoodSourceImage: foodSourceImage,
		PheromoneImage:  ebiten.NewImage(gw, gh),
		PixelBuffer:     make([]byte, gw*gh*4),
		FoodCollected:   0,
		TotalTicks:      0,
	}
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

func generateAnts(w, h int, homePosition shared.Position) []ant.Ant {
	ants := []ant.Ant{}

	for _ = range NumberOfAnts {
		posX := homePosition.X + rand.Float64()*HomeRadius
		posY := homePosition.Y + rand.Float64()*HomeRadius
		angle := rand.Float64() * 2 * math.Pi

		ant := ant.Ant{
			Position: shared.Position{
				X: posX,
				Y: posY,
			},
			AngleRadians: angle,
			State:        ant.SearchingForFood,
			Scent:        InitialScentStrength,
			GatheredFood: 0,
		}
		ants = append(ants, ant)
	}

	return ants
}

func generateFoodSources(w, h int, homePosition shared.Position) []shared.FoodSource {
	foodSources := []shared.FoodSource{}

	for _ = range NumberOfFoodSources {
		radius := rand.Float64() * MaxFoodSourceRadius

		posX, posY, distanceToHome := 0.0, 0.0, 0.0
		minDistanceToHome := 300.0

		for distanceToHome < minDistanceToHome {
			posX = rand.Float64() * float64(w)
			posY = rand.Float64() * float64(h)

			distanceToHome = shared.GetDistance(homePosition, shared.Position{X: posX, Y: posY})
		}

		foodSource := shared.FoodSource{
			Position: shared.Position{
				X: posX,
				Y: posY,
			},
			Radius: radius,
		}
		foodSources = append(foodSources, foodSource)
	}

	return foodSources
}

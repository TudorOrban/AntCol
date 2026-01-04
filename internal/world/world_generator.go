package world

import (
	"image/color"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"

	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
)

func GenerateWorld(w, h int) *World {
	gw, gh := w/shared.GridScale, h/shared.GridScale

	homePosition := shared.Position{X: float64(w / 2), Y: float64(h / 2)}
	ants := GenerateAnts(w, h, homePosition)
	foodSources := GenerateFoodSources(w, h, homePosition)

	antImage, homeImage := shared.LoadAssets()

	foodSourceImage := ebiten.NewImage(MaxFoodSourceRadius*2, MaxFoodSourceRadius*2)
	foodSourceImage.Fill(color.RGBA{230, 10, 15, 230})

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
		HomeImage:       homeImage,
		AntImage:        antImage,
		FoodSourceImage: foodSourceImage,
		PheromoneImage:  ebiten.NewImage(gw, gh),
		PixelBuffer:     make([]byte, gw*gh*4),
		FoodCollected:   0,
		TotalTicks:      0,
	}
}

func GenerateAnts(w, h int, homePosition shared.Position) []ant.Ant {
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
		}
		ants = append(ants, ant)
	}

	return ants
}

func GenerateFoodSources(w, h int, homePosition shared.Position) []shared.FoodSource {
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

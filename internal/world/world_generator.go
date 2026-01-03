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
	homePosition := shared.Position{X: 100, Y: float64(h / 2)}
	ants := GenerateAnts(w, h, homePosition)
	foodSources := GenerateFoodSources(w, h)

	antImage := ebiten.NewImage(AntLength, AntWidth)
	antImage.Fill(color.Black)

	foodSourceImage := ebiten.NewImage(MaxFoodSourceRadius, MaxFoodSourceRadius)
	foodSourceImage.Fill(color.RGBA{0, 255, 0, 0})

	homeImage := ebiten.NewImage(HomeRadius, HomeRadius)
	homeImage.Fill(color.RGBA{0, 255, 255, 0})

	return &World{
		Width:           w,
		Height:          h,
		HomePosition:    homePosition,
		HomePheromones:  make([]float64, w*h),
		FoodPheromones:  make([]float64, w*h),
		Ants:            ants,
		FoodSources:     foodSources,
		HomeImage:       homeImage,
		AntImage:        antImage,
		FoodSourceImage: foodSourceImage,
		PheromoneImage:  ebiten.NewImage(w, h),
		PixelBuffer:     make([]byte, w*h*4),
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
		}
		ants = append(ants, ant)
	}

	return ants
}

func GenerateFoodSources(w, h int) []shared.FoodSource {
	foodSources := []shared.FoodSource{}

	for _ = range NumberOfFoodSources {
		posX := rand.Float64() * float64(w)
		posY := rand.Float64() * float64(h)
		radius := MaxFoodSourceRadius

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

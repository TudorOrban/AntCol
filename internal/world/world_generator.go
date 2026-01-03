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
	homePosition := shared.Position{X: 100, Y: 100}
	ants := GenerateAnts(w, h)
	foodSources := GenerateFoodSources(w, h)

	antImage := ebiten.NewImage(AntLength, AntWidth)
	antImage.Fill(color.Black)

	foodSourceImage := ebiten.NewImage(MaxFoodSourceRadius, MaxFoodSourceRadius)
	foodSourceImage.Fill(color.RGBA{0, 255, 0, 0})

	return &World{
		Width:           w,
		Height:          h,
		HomePosition:    homePosition,
		HomePheromones:  make([]float64, w*h),
		FoodPheromones:  make([]float64, w*h),
		Ants:            ants,
		FoodSources:     foodSources,
		AntImage:        antImage,
		FoodSourceImage: foodSourceImage,
		PheromoneImage:  ebiten.NewImage(w, h),
		PixelBuffer:     make([]byte, w*h*4),
	}
}

func GenerateAnts(w, h int) []ant.Ant {
	ants := []ant.Ant{}

	for _ = range NumberOfAnts {
		posX := rand.Float64() * float64(w)
		posY := rand.Float64() * float64(h)
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
		radius := rand.Float64() * MaxFoodSourceRadius

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

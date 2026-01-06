package mapgen

import (
	"ant-sim/internal/shared"
	"ant-sim/internal/state"
	"math/rand/v2"
)

func generateFoodSources(w *state.World) {
	foodSources := []shared.FoodSource{}

	for _ = range w.Config.Map.NumFoodSources {
		radius := rand.Float64() * w.Config.Food.MaxFoodSourceRadius

		posX, posY, distanceToHome := 0.0, 0.0, 0.0
		minDistanceToHome := 300.0

		for distanceToHome < minDistanceToHome {
			posX = rand.Float64() * float64(w.Width)
			posY = rand.Float64() * float64(w.Height)

			distanceToHome = shared.GetDistance(w.HomePosition, shared.Position{X: posX, Y: posY})
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

	w.FoodSources = foodSources
}

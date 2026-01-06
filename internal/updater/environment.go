package updater

import (
	"ant-sim/internal/state"
	"math/rand/v2"
)

func UpdateEnvironment(w *state.World) {
	if rand.Float32() < 0.2 {
		applyPheromoneDiffusion(w)
	}
}

func applyPheromoneDiffusion(w *state.World) {
	if len(w.HomeTemp) == 0 {
		return
	}

	for y := 1; y < w.GridHeight-1; y++ {
		for x := 1; x < w.GridWidth-1; x++ {
			idx := y*w.GridWidth + x

			sumHome := (w.HomePheromones[idx-1] + w.HomePheromones[idx+1] +
				w.HomePheromones[idx-w.GridWidth] + w.HomePheromones[idx+w.GridWidth] +
				w.HomePheromones[idx]) / w.Config.Pheromone.DiffusionStrength

			newValHome := sumHome * w.Config.Pheromone.Decay
			if newValHome > w.Config.Pheromone.Cap {
				newValHome = w.Config.Pheromone.Cap
			}
			w.HomeTemp[idx] = newValHome

			sumFood := (w.FoodPheromones[idx-1] + w.FoodPheromones[idx+1] +
				w.FoodPheromones[idx-w.GridWidth] + w.FoodPheromones[idx+w.GridWidth] +
				w.FoodPheromones[idx]) / w.Config.Pheromone.DiffusionStrength

			newValFood := sumFood * w.Config.Pheromone.Decay
			if newValFood > w.Config.Pheromone.Cap {
				newValFood = w.Config.Pheromone.Cap
			}
			w.FoodTemp[idx] = newValFood
		}
	}

	w.HomePheromones, w.HomeTemp = w.HomeTemp, w.HomePheromones
	w.FoodPheromones, w.FoodTemp = w.FoodTemp, w.FoodPheromones
}

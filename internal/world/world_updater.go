package world

import (
	"ant-sim/internal/ant"
	"math"
	"math/rand/v2"
)

func (w *World) UpdateEnvironment() {
	if len(w.HomeTemp) == 0 {
		return
	}
	if rand.Float32() > 0.2 {
		return
	}

	// Apply diffusion
	for y := 1; y < w.GridHeight-1; y++ {
		for x := 1; x < w.GridWidth-1; x++ {
			idx := y*w.GridWidth + x

			sumHome := (w.HomePheromones[idx-1] + w.HomePheromones[idx+1] +
				w.HomePheromones[idx-w.GridWidth] + w.HomePheromones[idx+w.GridWidth] +
				w.HomePheromones[idx]) / DiffusionStrength

			newValHome := sumHome * PheromoneDecay
			if newValHome > PheromoneCap {
				newValHome = PheromoneCap
			}
			w.HomeTemp[idx] = newValHome

			sumFood := (w.FoodPheromones[idx-1] + w.FoodPheromones[idx+1] +
				w.FoodPheromones[idx-w.GridWidth] + w.FoodPheromones[idx+w.GridWidth] +
				w.FoodPheromones[idx]) / DiffusionStrength

			newValFood := sumFood * PheromoneDecay
			if newValFood > PheromoneCap {
				newValFood = PheromoneCap
			}
			w.FoodTemp[idx] = newValFood
		}
	}

	w.HomePheromones, w.HomeTemp = w.HomeTemp, w.HomePheromones
	w.FoodPheromones, w.FoodTemp = w.FoodTemp, w.FoodPheromones
}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		currentAnt := &w.Ants[i]

		// Steer & Move
		var guiding []float64
		if currentAnt.State == ant.SearchingForFood {
			guiding = w.FoodPheromones
		} else {
			guiding = w.HomePheromones
		}

		currentAnt.ApplySteering(w.GridWidth, w.GridHeight, guiding)
		currentAnt.Move(float64(w.Width), float64(w.Height))

		// Check destination
		if currentAnt.State == ant.SearchingForFood {
			if currentAnt.IsAtFoodSource(w.FoodSources) {
				currentAnt.State = ant.ReturningHome
				currentAnt.AngleRadians += math.Pi // Turn around
				currentAnt.Scent = InitialScentStrength
			}
		} else if currentAnt.State == ant.ReturningHome {
			if currentAnt.IsAtHome(w.HomePosition, HomeRadius) {
				currentAnt.State = ant.SearchingForFood
				currentAnt.AngleRadians += math.Pi // Head back out

				// Stats
				currentAnt.GatheredFood++
				w.FoodCollected++
				w.updateLeaderboard()
			}
		}

		// Deposit
		var release []float64
		if currentAnt.State == ant.SearchingForFood {
			release = w.HomePheromones
		} else {
			release = w.FoodPheromones
		}

		currentAnt.DepositPheromone(w.GridWidth, w.GridHeight, release, w.FoodSources)
	}
}

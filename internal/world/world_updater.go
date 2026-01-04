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

	for y := 1; y < w.Height-1; y++ {
		for x := 1; x < w.Width-1; x++ {
			idx := y*w.Width + x

			// Average the 3x3 nieghbors
			sumHome := (w.HomePheromones[idx-1] + w.HomePheromones[idx+1] +
				w.HomePheromones[idx-w.Width] + w.HomePheromones[idx+w.Width] +
				w.HomePheromones[idx]) / DiffusionStrength

			sumFood := (w.FoodPheromones[idx-1] + w.FoodPheromones[idx+1] +
				w.FoodPheromones[idx-w.Width] + w.FoodPheromones[idx+w.Width] +
				w.FoodPheromones[idx]) / DiffusionStrength

			newValHome := sumHome * PheromoneDecay
			if newValHome > PheromoneCap {
				newValHome = PheromoneCap
			}
			w.HomeTemp[idx] = newValHome

			newValFood := sumFood * PheromoneDecay
			if newValFood > PheromoneCap {
				newValFood = PheromoneCap
			}
			w.FoodTemp[idx] = newValFood
		}
	}

	// Swap buffers
	w.HomePheromones, w.HomeTemp = w.HomeTemp, w.HomePheromones
	w.FoodPheromones, w.FoodTemp = w.FoodTemp, w.FoodPheromones
}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		currentAnt := &w.Ants[i]

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
				w.FoodCollected++
			}
		}

		var guiding, release []float64
		if currentAnt.State == ant.SearchingForFood {
			guiding = w.FoodPheromones
			release = w.HomePheromones
		} else {
			guiding = w.HomePheromones
			release = w.FoodPheromones
		}

		currentAnt.ApplySteering(w.Width, w.Height, guiding)

		currentAnt.Move(float64(w.Width), float64(w.Height))

		currentAnt.DepositPheromone(w.Width, release, w.FoodSources)
	}
}

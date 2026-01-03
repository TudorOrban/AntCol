package world

import (
	"ant-sim/internal/ant"
	"math"
)

func (w *World) UpdateEnvironment() {
	for i := 0; i < len(w.HomePheromones); i++ {
		w.HomePheromones[i] *= PheromoneDecay
		w.FoodPheromones[i] *= PheromoneDecay
	}
}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		currentAnt := &w.Ants[i]

		if currentAnt.State == ant.SearchingForFood {
			if currentAnt.IsAtFoodSource(w.FoodSources) {
				currentAnt.State = ant.ReturningHome
				currentAnt.AngleRadians += math.Pi // Turn around
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

package world

import (
	"ant-sim/internal/ant"
)

func (w *World) UpdateEnvironment() {
	for i := 0; i < len(w.HomePheromones); i++ {
		w.HomePheromones[i] *= PheromoneDecay
	}
}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		currentAnt := &w.Ants[i]

		var releasePheromones, guidingPheromones []float64
		if currentAnt.State == ant.SearchingForFood {
			releasePheromones = w.FoodPheromones
			guidingPheromones = w.HomePheromones
		} else {
			releasePheromones = w.HomePheromones
			guidingPheromones = w.FoodPheromones
		}

		currentAnt.ApplySteering(w.Width, w.Height, releasePheromones)

		currentAnt.Move(float64(w.Width), float64(w.Height))

		currentAnt.DepositPheromone(w.Width, guidingPheromones, w.FoodSources)
	}
}

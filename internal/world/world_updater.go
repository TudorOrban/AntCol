package world

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
)

// Update
func (w *World) UpdateEnvironment() {
	for i := 0; i < len(w.HomePheromones); i++ {
		w.HomePheromones[i] *= PheromoneDecay
	}
}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		w.Ants[i].Move(float64(w.Width), float64(w.Height))
		w.DepositPheromone(i)
	}
}

func (w *World) DepositPheromone(i int) {
	currentAnt := &w.Ants[i]

	isInFoodSourceRange := isAtFoodSource(currentAnt.Position, w.FoodSources)
	if isInFoodSourceRange {
		currentAnt.State = ant.ReturningHome
	}

	if currentAnt.State == ant.SearchingForFood {
		w.SetHomePheromone(currentAnt.Position, 1.0)
	} else {
		w.SetFoodPheromone(currentAnt.Position, 1.0)
	}
}

func (w *World) SetHomePheromone(position shared.Position, value float64) {
	idx := w.getIndex(position)
	if idx >= 0 && idx < len(w.HomePheromones) {
		w.HomePheromones[idx] = value
	}
}

func (w *World) SetFoodPheromone(position shared.Position, value float64) {
	idx := w.getIndex(position)
	if idx >= 0 && idx < len(w.FoodPheromones) {
		w.FoodPheromones[idx] = value
	}
}

func (w *World) getIndex(position shared.Position) int {
	return int(position.Y)*w.Width + int(position.X)
}

func isAtFoodSource(antPosition shared.Position, foodSources []FoodSource) bool {
	for _, foodSource := range foodSources {
		dx := antPosition.X - foodSource.Position.X
		dy := antPosition.Y - foodSource.Position.Y
		distSq := dx*dx + dy*dy

		if distSq < foodSource.Radius*foodSource.Radius {
			return true
		}
	}
	return false
}

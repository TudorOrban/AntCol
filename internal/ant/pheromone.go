package ant

import (
	"ant-sim/internal/shared"
)

func (a *Ant) DepositPheromone(worldWidth int, pheromones []float64, foodSources []shared.FoodSource) {
	idx := getPheromoneIndex(worldWidth, a.Position)
	if idx >= 0 && idx < len(pheromones) {
		pheromones[idx] += 1.0
	}
}

func getPheromoneIndex(worldWidth int, position shared.Position) int {
	return int(position.Y)*worldWidth + int(position.X)
}
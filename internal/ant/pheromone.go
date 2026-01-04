package ant

import (
	"ant-sim/internal/shared"
)

func (a *Ant) DepositPheromone(worldWidth int, pheromones []float64, foodSources []shared.FoodSource) {
	idx := GetPheromoneIndex(worldWidth, a.Position)
	if idx >= 0 && idx < len(pheromones) {
		pheromones[idx] += a.Scent * DepositStrength
	}
}

func GetPheromoneIndex(worldWidth int, position shared.Position) int {
	return int(position.Y)*worldWidth + int(position.X)
}

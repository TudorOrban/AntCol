package ant

import (
	"ant-sim/internal/shared"
)

func (a *Ant) DepositPheromone(gridWidth, gridHeight int, pheromones []float64, foodSources []shared.FoodSource) {
	idx := GetGridIndex(gridWidth, gridHeight, a.Position.X, a.Position.Y)
	if idx >= 0 && idx < len(pheromones) {
		pheromones[idx] += a.Scent * DepositStrength
	}
}

func GetGridIndex(gridWidth, gridHeight int, x, y float64) int {
	gx := int(x) / shared.GridScale
	gy := int(y) / shared.GridScale

	if gx < 0 || gx >= gridWidth || gy < 0 || gy >= gridHeight {
		return -1
	}
	return gy*gridWidth + gx
}

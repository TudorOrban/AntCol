package ant

import "ant-sim/internal/shared"

func (a *Ant) DepositPheromone(worldWidth int, pheromones []float64, foodSources []shared.FoodSource) {
	isInFoodSourceRange := isAtFoodSource(a.Position, foodSources)
	if isInFoodSourceRange {
		a.State = ReturningHome
	}

	setPheromone(worldWidth, pheromones, a.Position, 1)
}

func setPheromone(worldWidth int, pheromones []float64, position shared.Position, value float64) {
	idx := getPheromoneIndex(worldWidth, position)
	if idx >= 0 && idx < len(pheromones) {
		pheromones[idx] = value
	}
}

func getPheromoneIndex(worldWidth int, position shared.Position) int {
	return int(position.Y)*worldWidth + int(position.X)
}

func isAtFoodSource(antPosition shared.Position, foodSources []shared.FoodSource) bool {
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

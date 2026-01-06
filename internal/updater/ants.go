package updater

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/state"
	"ant-sim/internal/statistics"
	"math"
)

func UpdateAnts(w *state.World) {
	liveCount := 0

	for i := 0; i < len(w.Ants); i++ {
		currentAnt := &w.Ants[i]

		moveAnt(w, currentAnt)
		if currentAnt.IsDead() {
			continue
		}

		actOnDestination(w, currentAnt)
		depositPheromones(w, currentAnt)

		if i != liveCount {
			w.Ants[liveCount] = w.Ants[i]
		}
		liveCount++
	}

	w.Ants = w.Ants[:liveCount]
}

func moveAnt(w *state.World, currentAnt *ant.Ant) {
	var guiding []float64
	if currentAnt.State == ant.SearchingForFood {
		guiding = w.FoodPheromones
	} else {
		guiding = w.HomePheromones
	}

	currentAnt.ApplySteering(w.GridWidth, w.GridHeight, guiding)
	currentAnt.Move(float64(w.Width), float64(w.Height), w.Obstacles, w.GridWidth)
}

func actOnDestination(w *state.World, currentAnt *ant.Ant) {
	if currentAnt.State == ant.SearchingForFood && currentAnt.IsAtFoodSource(w.FoodSources) {
		actOnFoodArrival(w, currentAnt)
	} 
	if currentAnt.State == ant.ReturningHome && currentAnt.IsAtHome(w.HomePosition, w.Config.Map.HomeRadius) {
		actOnHomeArrival(w, currentAnt)
	}
}

func actOnFoodArrival(w *state.World, currentAnt *ant.Ant) {
	currentAnt.State = ant.ReturningHome
	currentAnt.AngleRadians += math.Pi // Turn around
	currentAnt.Scent = w.Config.Pheromone.InitialScentStrength
}

func actOnHomeArrival(w *state.World, currentAnt *ant.Ant) {
	currentAnt.State = ant.SearchingForFood

	currentAnt.AngleRadians += math.Pi // Head back out

	// Stats
	currentAnt.GatheredFood++
	w.FoodCollected++
	statistics.UpdateLeaderboard(w)
}

func depositPheromones(w *state.World, currentAnt *ant.Ant) {
	var release []float64
	if currentAnt.State == ant.SearchingForFood {
		release = w.HomePheromones
	} else {
		release = w.FoodPheromones
	}

	currentAnt.DepositPheromone(w.GridWidth, w.GridHeight, release, w.FoodSources)
}

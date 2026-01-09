package updater

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
	"ant-sim/internal/state"
	"ant-sim/internal/statistics"
	"math"
	"math/rand/v2"
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

	ReproduceAnts(w)
}

func ReproduceAnts(w *state.World) {
	if rand.Float32() > float32(w.Config.Reproduction.ReproductionRate)*(1/60) {
		return
	}

	newAnt := ant.Ant{
		Position:     w.HomePosition,
		AngleRadians: rand.Float64() * math.Pi,
		State:        ant.SearchingForFood,
		Scent:        w.Config.Pheromone.ScentDecay,
		GatheredFood: 0,
		CurrentFood:  w.Config.Food.MaxFood,
	}
	w.Ants = append(w.Ants, newAnt)
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
	if currentAnt.State == ant.SearchingForFood {
		if source := currentAnt.GetFoodSourceAt(w.FoodSources); source != nil && source.TotalFood > 0 {
			actOnFoodArrival(w, currentAnt, source)
		}
	}
	if currentAnt.State == ant.ReturningHome && currentAnt.IsAtHome(w.HomePosition, w.Config.Map.HomeRadius) {
		actOnHomeArrival(w, currentAnt)
	}
}

func actOnFoodArrival(w *state.World, currentAnt *ant.Ant, source *shared.FoodSource) {
	// Grab food
	amountToTake := 0.0
	if source.TotalFood > w.Config.Food.FoodPerGrab {
		amountToTake = w.Config.Food.FoodPerGrab
	} else {
		amountToTake = source.TotalFood
	}

	source.TotalFood -= amountToTake
	currentAnt.CarriedFood = amountToTake

	// Go back home
	currentAnt.State = ant.ReturningHome
	currentAnt.AngleRadians += math.Pi
	currentAnt.Scent = w.Config.Pheromone.InitialScentStrength
}

func actOnHomeArrival(w *state.World, currentAnt *ant.Ant) {
	// Deposit food
	w.HomeFoodSupply += currentAnt.CarriedFood
	currentAnt.CarriedFood = 0

	// Eat if necessary
	if currentAnt.CurrentFood < 30.0 {
		amountToEat := math.Min(w.HomeFoodSupply, w.Config.Food.MaxFood-currentAnt.CurrentFood)
		w.HomeFoodSupply -= amountToEat
		currentAnt.CurrentFood += amountToEat
	}

	// Head back out
	currentAnt.State = ant.SearchingForFood
	currentAnt.AngleRadians += math.Pi

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

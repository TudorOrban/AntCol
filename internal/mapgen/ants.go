package mapgen

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
	"ant-sim/internal/state"
	"math"
	"math/rand/v2"
)


func generateAnts(w *state.World) {
	ants := []ant.Ant{}

	for _ = range w.Config.Map.NumAnts {
		posX := w.HomePosition.X + rand.Float64()*w.Config.Map.HomeRadius
		posY := w.HomePosition.Y + rand.Float64()*w.Config.Map.HomeRadius
		angle := rand.Float64() * 2 * math.Pi

		ant := ant.Ant{
			Position: shared.Position{
				X: posX,
				Y: posY,
			},
			AngleRadians: angle,
			State:        ant.SearchingForFood,
			Scent:        w.Config.Pheromone.InitialScentStrength,
			GatheredFood: 0,
			CurrentFood:  w.Config.Food.MaxFood,
		}
		ants = append(ants, ant)
	}

	w.Ants = ants
}

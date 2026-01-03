package ant

import (
	"ant-sim/internal/shared"
	"math"
	"math/rand"
)

type AntState int

const (
	StateIdle AntState = iota
	SearchingForFood
	ReturningHome
)

const (
	AntSpeed     = 2
	AntTurnSpeed = 0.2
	SensorAngle  = 0.4
	SensorDist   = 15
)

var stateName = map[AntState]string{
	StateIdle:        "idle",
	SearchingForFood: "searching for food",
	ReturningHome:    "returning home",
}

func (as AntState) String() string {
	return stateName[as]
}

type Ant struct {
	Position     shared.Position
	AngleRadians float64
	State        AntState
}

func (a *Ant) Move(worldWidth, worldHeight float64) {
	if rand.Float32() < 0.1 {
		a.AngleRadians += (float64)(rand.Float32()-0.5) * 0.2
	}

	a.Position.X += math.Cos(a.AngleRadians) * AntSpeed
	a.Position.Y += math.Sin(a.AngleRadians) * AntSpeed

	// Check screen boundaries
	if a.Position.X < 0 || a.Position.X >= worldWidth {
		a.AngleRadians = math.Pi - a.AngleRadians
		a.Position.X = math.Max(0, math.Min(a.Position.X, worldWidth-1))
	}
	if a.Position.Y < 0 || a.Position.Y >= worldHeight {
		a.AngleRadians = -a.AngleRadians
		a.Position.Y = math.Max(0, math.Min(a.Position.Y, worldHeight-1))
	}
}

func (a *Ant) ApplySteering(worldWidth, worldHeight int, pheromones []float64) {
	vF := a.sense(worldWidth, worldHeight, pheromones, 0, SensorDist)
	vL := a.sense(worldWidth, worldHeight, pheromones, SensorAngle, SensorDist)
	vR := a.sense(worldWidth, worldHeight, pheromones, -SensorAngle, SensorDist)

	if vF > vL && vF > vR {
		// Path is strongest ahead, do nothing
	} else if vL > vR {
		a.AngleRadians += AntTurnSpeed
	} else if vR > vL {
		a.AngleRadians -= AntTurnSpeed
	} else if vF < 0 {
		a.AngleRadians += math.Pi
	}
}

func (a *Ant) sense(
	worldWidth, worldHeight int,
	pheromones []float64,
	sensorAngle float64, sensorDist float64,
) float64 {
	angle := a.AngleRadians + sensorAngle
	sensorX := a.Position.X + math.Cos(angle)*sensorDist
	sensorY := a.Position.Y + math.Sin(angle)*sensorDist

	x, y := int(sensorX), int(sensorY)

	if x >= 0 && x < worldWidth && y >= 0 && y < worldHeight {
		index := getPheromoneIndex(worldWidth, shared.Position{X: sensorX, Y: sensorY})
		return pheromones[index]
	}
	return -1
}

func (a *Ant) IsAtFoodSource(foodSources []shared.FoodSource) bool {
	for _, foodSource := range foodSources {
		dx := a.Position.X - foodSource.Position.X
		dy := a.Position.Y - foodSource.Position.Y
		distSq := dx*dx + dy*dy

		if distSq < foodSource.Radius*foodSource.Radius {
			return true
		}
	}
	return false
}

func (a *Ant) IsAtHome(homePosition shared.Position, homeRadius float64) bool {
	dx := a.Position.X - homePosition.X
	dy := a.Position.Y - homePosition.Y
	distSq := dx*dx + dy*dy
	return distSq < homeRadius*homeRadius
}

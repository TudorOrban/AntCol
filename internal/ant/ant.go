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
	AntSpeed        = 2
	AntTurnSpeed    = 0.15
	SensorAngle     = 0.4
	SensorDist      = 35
	SensorThreshold = 0.05
	DepositStrength = 5
	ScentDecay      = 0.995
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
	Scent        float64
	GatheredFood int
}

func (a *Ant) Move(worldWidth, worldHeight float64) {
	if rand.Float32() < 0.2 {
		a.AngleRadians += (float64)(rand.Float32()-0.5) * 0.2
	}

	a.Position.X += math.Cos(a.AngleRadians) * AntSpeed
	a.Position.Y += math.Sin(a.AngleRadians) * AntSpeed

	// Check screen boundaries
	margin := 20.0
	if a.Position.X < margin || a.Position.X >= worldWidth-margin {
		a.AngleRadians = math.Pi - a.AngleRadians
		a.Position.X = math.Max(margin, math.Min(a.Position.X, worldWidth-margin))
	}
	if a.Position.Y < margin || a.Position.Y >= worldHeight-margin {
		a.AngleRadians = -a.AngleRadians
		a.Position.Y = math.Max(margin, math.Min(a.Position.Y, worldHeight-margin))
	}

	a.Scent *= ScentDecay
}

func (a *Ant) ApplySteering(gridWidth, gridHeight int, pheromones []float64) {
	vF := a.sense(gridWidth, gridHeight, pheromones, 0, SensorDist)
	vL := a.sense(gridWidth, gridHeight, pheromones, SensorAngle, SensorDist)
	vR := a.sense(gridWidth, gridHeight, pheromones, -SensorAngle, SensorDist)

	if vF < 0 {
		if vL > vR {
			a.AngleRadians += AntTurnSpeed * 2
		} else {
			a.AngleRadians -= AntTurnSpeed * 2
		}
		return
	}

	if vL > SensorThreshold || vR > SensorThreshold || vF > SensorThreshold {
		const inertiaRatio = 1.1

		if vL > vR && vL > vF*inertiaRatio {
			if rand.Float32() > 0.1 {
				a.AngleRadians += AntTurnSpeed
			}
		} else if vR > vL && vR > vF*inertiaRatio {
			if rand.Float32() > 0.1 {
				a.AngleRadians -= AntTurnSpeed
			}
		}
	}
}

func (a *Ant) sense(
	gridWidth, gridHeight int,
	pheromones []float64,
	sensorAngle float64, sensorDist float64,
) float64 {
	angle := a.AngleRadians + sensorAngle
	sensorX := a.Position.X + math.Cos(angle)*sensorDist
	sensorY := a.Position.Y + math.Sin(angle)*sensorDist

	gx := int(sensorX) / shared.GridScale
	gy := int(sensorY) / shared.GridScale

	if gx >= 0 && gx < gridWidth && gy >= 0 && gy < gridHeight {
		index := GetGridIndex(gridWidth, gridHeight, sensorX, sensorY)
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

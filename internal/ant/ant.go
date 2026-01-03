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
	AntSpeed     = 1
	AntTurnSpeed = 0.1
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

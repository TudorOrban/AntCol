package main

import (
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

type Position struct {
	x, y float64
}

type Ant struct {
	Position     Position
	AngleRadians float64
	State        AntState
}

func (a *Ant) Move(worldWidth, worldHeight float64) {
	if rand.Float32() < 0.1 {
		a.AngleRadians += (float64)(rand.Float32()-0.5) * 0.2
	}

	a.Position.x += math.Cos(a.AngleRadians) * AntSpeed
	a.Position.y += math.Sin(a.AngleRadians) * AntSpeed

	// Check screen boundaries
	if a.Position.x < 0 || a.Position.x >= worldWidth {
		a.AngleRadians = math.Pi - a.AngleRadians
		a.Position.x = math.Max(0, math.Min(a.Position.x, worldWidth-1))
	}
	if a.Position.y < 0 || a.Position.y >= worldHeight {
		a.AngleRadians = -a.AngleRadians
		a.Position.y = math.Max(0, math.Min(a.Position.y, worldHeight-1))
	}
}

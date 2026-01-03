package main

import "math"

type AntState int

const (
	StateIdle AntState = iota
	SearchingForFood
	ReturningHome
)

const (
	AntSpeed     = 0.25
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

func (a *Ant) Move() {
	newX := a.Position.x + math.Cos(a.AngleRadians)*AntSpeed
	newY := a.Position.y + math.Sin(a.AngleRadians)*AntSpeed

	a.Position = Position{newX, newY}
}

package main

type AntState int

const (
	StateIdle AntState = iota
	SearchingForFood
	ReturningHome
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
	x, y float32
}

type Ant struct {
	Position     Position
	AngleRadians float32
	State        AntState
}

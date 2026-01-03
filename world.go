package main

import (
	"fmt"
	"sync"
)

type World struct {
	Width, Height  int
	HomePheromones []float32
	FoodPheromones []float32
	Ants           []Ant
	mu             sync.RWMutex
}

func NewWorld(w, h int) *World {
	ants := []Ant{
		{
			Position: Position{
				x: 10,
				y: 10,
			},
			AngleRadians: 1,
			State:        SearchingForFood,
		},
		{
			Position: Position{
				x: 20,
				y: 30,
			},
			AngleRadians: 2,
			State:        SearchingForFood,
		},
	}

	return &World{
		Width:          w,
		Height:         h,
		HomePheromones: make([]float32, w*h),
		FoodPheromones: make([]float32, w*h),
		Ants:           ants,
	}
}

func GenerateWorld(w, h int) *World {
	return &World{
		Width:          w,
		Height:         h,
		HomePheromones: make([]float32, w*h),
		FoodPheromones: make([]float32, w*h),
	}
}

func (w *World) getIndex(position Position) int {
	return int(position.y)*w.Width + int(position.x)
}

func (w *World) SetPheromone(position Position, value float32) {
	idx := w.getIndex(position)
	if idx >= 0 && idx < len(w.HomePheromones) {
		w.HomePheromones[idx] = value
	}
}

func (w *World) UpdateEnvironment() {

}

func (w *World) UpdateAnts() {
	for i, ant := range w.Ants {
		fmt.Printf("Ant %d is at position %v\n", i, ant.Position)
	}
}

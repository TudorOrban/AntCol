package main

import (
	"image/color"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	AntLength = 50
	AntWidth  = 20
)

type World struct {
	Width, Height  int
	HomePheromones []float32
	FoodPheromones []float32
	Ants           []Ant
	AntImage       *ebiten.Image
	mu             sync.RWMutex
}

func NewWorld(w, h int) *World {
	ants := []Ant{
		{
			Position: Position{
				x: 40,
				y: 40,
			},
			AngleRadians: 1,
			State:        SearchingForFood,
		},
		{
			Position: Position{
				x: 120,
				y: 300,
			},
			AngleRadians: 2,
			State:        SearchingForFood,
		},
		{
			Position: Position{
				x: 400,
				y: 80,
			},
			AngleRadians: 2,
			State:        SearchingForFood,
		},
	}

	antImage := ebiten.NewImage(AntLength, AntWidth)
	antImage.Fill(color.Black)

	return &World{
		Width:          w,
		Height:         h,
		HomePheromones: make([]float32, w*h),
		FoodPheromones: make([]float32, w*h),
		Ants:           ants,
		AntImage:       antImage,
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

func (w *World) Draw(screen *ebiten.Image) {
	for _, ant := range w.Ants {
		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(-AntLength/2, -AntWidth/2)

		opts.GeoM.Rotate(ant.AngleRadians)

		opts.GeoM.Translate(ant.Position.x, ant.Position.y)

		screen.DrawImage(w.AntImage, opts)
	}
}

func (w *World) UpdateEnvironment() {

}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		w.Ants[i].Move()
	}
}

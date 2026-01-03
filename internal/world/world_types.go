package world

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	AntLength           = 30
	AntWidth            = 10
	NumberOfAnts        = 50
	NumberOfFoodSources = 5
	MaxFoodSourceRadius = 50.0
	PheromoneDecay      = 0.995
)

type World struct {
	Width, Height  int
	HomePosition   shared.Position
	HomePheromones []float64
	FoodPheromones []float64
	Ants           []ant.Ant
	FoodSources    []shared.FoodSource

	AntImage        *ebiten.Image
	FoodSourceImage *ebiten.Image
	PheromoneImage  *ebiten.Image
	PixelBuffer     []byte
	mu              sync.RWMutex
}

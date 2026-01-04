package world

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	AntLength            = 30
	AntWidth             = 10
	NumberOfAnts         = 100
	NumberOfFoodSources  = 5
	HomeRadius           = 100
	MaxFoodSourceRadius  = 50.0
	PheromoneDecay       = 0.98
	PheromoneCap         = 10.0
	DiffusionStrength    = 5.0
	InitialScentStrength = 10.0
)

type World struct {
	Width, Height int
	HomePosition  shared.Position
	Ants          []ant.Ant
	FoodSources   []shared.FoodSource

	HomePheromones []float64
	FoodPheromones []float64
	HomeTemp       []float64
	FoodTemp       []float64

	HomeImage       *ebiten.Image
	AntImage        *ebiten.Image
	FoodSourceImage *ebiten.Image
	PheromoneImage  *ebiten.Image
	PixelBuffer     []byte
	mu              sync.RWMutex

	FoodCollected int
	TotalTicks    int
}

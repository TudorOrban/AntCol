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
	NumberOfAnts        = 150
	NumberOfFoodSources = 10
	NumberOfObstacles   = 4
	ObstacleMaxLength   = 400
	HomeRadius          = 50.0
	MaxFoodSourceRadius = 40.0

	CameraSpeed = 5.0

	PheromoneDecay       = 0.99
	PheromoneCap         = 10.0
	DiffusionStrength    = 5.0
	InitialScentStrength = 10.0
)

type World struct {
	Width, Height         int
	GridWidth, GridHeight int
	HomePosition          shared.Position
	Ants                  []ant.Ant
	FoodSources           []shared.FoodSource
	Obstacles             []bool
	WallRects             []shared.Rectangle

	HomePheromones []float64
	FoodPheromones []float64
	HomeTemp       []float64
	FoodTemp       []float64

	CameraPosition shared.Position
	Zoom           float64
	lastMouseX     int
	lastMouseY     int

	GrassBackground *ebiten.Image
	HomeImage       *ebiten.Image
	AntImage        *ebiten.Image
	FoodSourceImage *ebiten.Image
	PheromoneImage  *ebiten.Image
	ObstacleImage   *ebiten.Image
	PixelBuffer     []byte
	mu              sync.RWMutex

	FoodCollected int
	TotalTicks    int
	TopAnts       []*ant.Ant
}

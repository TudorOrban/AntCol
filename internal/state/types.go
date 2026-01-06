package state

import (
	"ant-sim/internal/ant"
	"ant-sim/internal/shared"

	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	Width, Height         int
	GridWidth, GridHeight int
	Config                WorldConfig

	HomePosition shared.Position
	Ants         []ant.Ant
	FoodSources  []shared.FoodSource
	Obstacles    []bool
	WallRects    []shared.Rectangle

	HomePheromones []float64
	FoodPheromones []float64
	HomeTemp       []float64
	FoodTemp       []float64

	HomeFoodSupply float64

	CameraPosition shared.Position
	Zoom           float64
	LastMouseX     int
	LastMouseY     int

	GrassBackground *ebiten.Image
	HomeImage       *ebiten.Image
	AntImage        *ebiten.Image
	FoodSourceImage *ebiten.Image
	PheromoneImage  *ebiten.Image
	ObstacleImage   *ebiten.Image
	PixelBuffer     []byte

	FoodCollected int
	TotalTicks    int
	TopAnts       []*ant.Ant
}

type WorldConfig struct {
	Map struct {
		Width               int
		Height              int
		NumAnts             int
		NumFoodSources      int
		NumObstacles        int
		ObstacleMaxLength   float64
		HomeRadius          float64
	}
	Ant struct {
		Length           float64
		Width            float64
		Speed            float64
		TurnSpeed        float64
		SensorAngle      float64
		SensorDist       float64
		SensorThreshold  float64
		MovementFoodCost float64
	}
	Pheromone struct {
		Decay                float64
		Cap                  float64
		DiffusionStrength    float64
		InitialScentStrength float64
		DepositStrength      float64
		ScentDecay           float64
	}
	Food struct {
		MaxFood          float64
		MaxFoodSourceRadius float64
	}
	UI struct {
		CameraSpeed float64
	}
}

func DefaultConfig() WorldConfig {
	c := WorldConfig{}
	// Map
	c.Map.NumAnts = 150
	c.Map.NumFoodSources = 10
	c.Map.NumObstacles = 4
	c.Map.ObstacleMaxLength = 400
	c.Map.HomeRadius = 50.0

	// Ant
	c.Ant.Length = 30
	c.Ant.Width = 10
	c.Ant.Speed = 2
	c.Ant.TurnSpeed = 0.15
	c.Ant.SensorAngle = 0.4
	c.Ant.SensorDist = 35
	c.Ant.SensorThreshold = 0.05
	c.Ant.MovementFoodCost = 0.005

	// Pheromone
	c.Pheromone.Decay = 0.99
	c.Pheromone.Cap = 10.0
	c.Pheromone.DiffusionStrength = 5.0
	c.Pheromone.DepositStrength = 5
	c.Pheromone.InitialScentStrength = 10.0
	c.Pheromone.ScentDecay = 0.995

	// UI
	c.UI.CameraSpeed = 5.0

	// Food
	c.Food.MaxFoodSourceRadius = 40.0
	c.Food.MaxFood = 100.0

	return c
}

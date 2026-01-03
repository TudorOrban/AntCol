package main

import (
	"image/color"
	"math"
	"math/rand"
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
	HomePosition   Position
	HomePheromones []float64
	FoodPheromones []float64
	Ants           []Ant
	FoodSources    []FoodSource

	AntImage        *ebiten.Image
	FoodSourceImage *ebiten.Image
	PheromoneImage  *ebiten.Image
	PixelBuffer     []byte
	mu              sync.RWMutex
}

type FoodSource struct {
	Position Position
	Radius   float64
}

// Initialization
func NewWorld(w, h int) *World {
	homePosition := Position{100, 100}
	ants := GenerateAnts(w, h)
	foodSources := GenerateFoodSources(w, h)

	antImage := ebiten.NewImage(AntLength, AntWidth)
	antImage.Fill(color.Black)

	foodSourceImage := ebiten.NewImage(MaxFoodSourceRadius, MaxFoodSourceRadius)
	foodSourceImage.Fill(color.RGBA{0, 255, 0, 0})

	return &World{
		Width:           w,
		Height:          h,
		HomePosition:    homePosition,
		HomePheromones:  make([]float64, w*h),
		FoodPheromones:  make([]float64, w*h),
		Ants:            ants,
		FoodSources:     foodSources,
		AntImage:        antImage,
		FoodSourceImage: foodSourceImage,
		PheromoneImage:  ebiten.NewImage(w, h),
		PixelBuffer:     make([]byte, w*h*4),
	}
}

func GenerateAnts(w, h int) []Ant {
	ants := []Ant{}

	for _ = range NumberOfAnts {
		posX := rand.Float64() * float64(w)
		posY := rand.Float64() * float64(h)
		angle := rand.Float64() * 2 * math.Pi

		ant := Ant{
			Position: Position{
				x: posX,
				y: posY,
			},
			AngleRadians: angle,
			State:        SearchingForFood,
		}
		ants = append(ants, ant)
	}

	return ants
}

func GenerateFoodSources(w, h int) []FoodSource {
	foodSources := []FoodSource{}

	for _ = range NumberOfFoodSources {
		posX := rand.Float64() * float64(w)
		posY := rand.Float64() * float64(h)
		radius := rand.Float64() * MaxFoodSourceRadius

		foodSource := FoodSource{
			Position: Position{
				x: posX,
				y: posY,
			},
			Radius: radius,
		}
		foodSources = append(foodSources, foodSource)
	}

	return foodSources
}

// Drawing
func (w *World) Draw(screen *ebiten.Image) {
	w.DrawPheromones(screen)

	for _, foodSource := range w.FoodSources {
		w.DrawFoodSource(screen, foodSource)
	}

	for _, ant := range w.Ants {
		w.DrawAnt(screen, ant)
	}
}

func (w *World) DrawAnt(screen *ebiten.Image, ant Ant) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-AntLength/2, -AntWidth/2)

	opts.GeoM.Rotate(ant.AngleRadians)

	opts.GeoM.Translate(ant.Position.x, ant.Position.y)

	screen.DrawImage(w.AntImage, opts)
}

func (w *World) DrawFoodSource(screen *ebiten.Image, foodSource FoodSource) {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(-MaxFoodSourceRadius/2, -MaxFoodSourceRadius/2)

	opts.GeoM.Translate(foodSource.Position.x, foodSource.Position.y)

	screen.DrawImage(w.FoodSourceImage, opts)
}

func (w *World) DrawPheromones(screen *ebiten.Image) {
	for i := 0; i < len(w.HomePheromones); i++ {
		pixIdx := i * 4
		home := w.HomePheromones[i]
		food := w.FoodPheromones[i]

		strength := clamp((home + food) * 255)

		w.PixelBuffer[pixIdx] = uint8(clamp(food * 255))   // Red component
		w.PixelBuffer[pixIdx+1] = 0                        // Green
		w.PixelBuffer[pixIdx+2] = uint8(clamp(home * 255)) // Blue component
		w.PixelBuffer[pixIdx+3] = uint8(strength)
	}

	w.PheromoneImage.WritePixels(w.PixelBuffer)

	screen.DrawImage(w.PheromoneImage, nil)
}

// Update
func (w *World) UpdateEnvironment() {
	for i := 0; i < len(w.HomePheromones); i++ {
		w.HomePheromones[i] *= PheromoneDecay
	}
}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		w.Ants[i].Move(float64(w.Width), float64(w.Height))
		w.DepositPheromone(i)
	}
}

func (w *World) DepositPheromone(i int) {
	pos := w.Ants[i].Position
	if w.Ants[i].State == SearchingForFood {
		w.SetHomePheromone(pos, 1.0)
	} else {
		w.SetFoodPheromone(pos, 1.0)
	}
}

func (w *World) SetHomePheromone(position Position, value float64) {
	idx := w.getIndex(position)
	if idx >= 0 && idx < len(w.HomePheromones) {
		w.HomePheromones[idx] = value
	}
}

func (w *World) SetFoodPheromone(position Position, value float64) {
	idx := w.getIndex(position)
	if idx >= 0 && idx < len(w.FoodPheromones) {
		w.FoodPheromones[idx] = value
	}
}

func (w *World) getIndex(position Position) int {
	return int(position.y)*w.Width + int(position.x)
}

// Utils
func clamp(v float64) float64 {
	if v > 255 {
		return 255
	}
	if v < 0 {
		return 0
	}
	return v
}

package main

import (
	"ant-sim/internal/world"
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 1800
	screenHeight = 1200
)

type Game struct {
	world *world.World
}

func (g *Game) Update() error {
	g.world.TotalTicks++
	g.world.UpdateEnvironment()
	g.world.UpdateAnts()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{34, 139, 34, 255})

	g.world.Draw(screen)

	// Display info
	debugBar := ebiten.NewImage(200, 40)
	debugBar.Fill(color.RGBA{0, 0, 0, 150})
	screen.DrawImage(debugBar, nil)

	minutesPassed := float64(g.world.TotalTicks) / 3600.0
	avgPerMin := 0.0
	if minutesPassed > 0 {
		avgPerMin = float64(g.world.FoodCollected) / minutesPassed
	}

	stats := fmt.Sprintf("Food Collected: %d\nAvg Food/Min: %.2f", g.world.FoodCollected, avgPerMin)
	ebitenutil.DebugPrint(screen, stats)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		world: world.GenerateWorld(screenWidth, screenHeight),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ant Colony Simulation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

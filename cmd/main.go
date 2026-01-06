package main

import (
	"ant-sim/internal/world"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1800
	screenHeight = 1200
	mapWidth     = 3000
	mapHeight    = 2000
)

type Game struct {
	world *world.World
}

func (g *Game) Update() error {
	g.world.TotalTicks++
	g.world.UpdateEnvironment()
	g.world.UpdateAnts()
	g.world.UpdateCamera()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw(screen)

	g.world.RenderStats(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		world: world.GenerateWorld(mapWidth, mapHeight),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ant Colony Simulation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

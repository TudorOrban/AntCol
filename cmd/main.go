package main

import (
	"ant-sim/internal/mapgen"
	"ant-sim/internal/renderer"
	"ant-sim/internal/state"
	"ant-sim/internal/statistics"
	"ant-sim/internal/updater"
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
	world *state.World
}

func (g *Game) Update() error {
	g.world.TotalTicks++

	updater.UpdateEnvironment(g.world)
	updater.UpdateAnts(g.world)
	updater.UpdateCamera(g.world)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	renderer.Draw(screen, g.world)

	statistics.RenderStats(screen, g.world)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		world: mapgen.GenerateWorld(mapWidth, mapHeight),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ant Colony Simulation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

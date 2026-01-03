package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1800
	screenHeight = 1200
)

type Game struct {
	world *World
}

func (g *Game) Update() error {
	g.world.UpdateEnvironment()

	g.world.UpdateAnts()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{200, 200, 200, 255})

	g.world.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := &Game{
		world: NewWorld(screenWidth, screenHeight),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Ant Colony Simulation")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

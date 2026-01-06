package world

import (
	"ant-sim/internal/ant"
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

func (w *World) UpdateEnvironment() {
	if len(w.HomeTemp) == 0 {
		return
	}
	if rand.Float32() > 0.2 {
		return
	}

	// Apply diffusion
	for y := 1; y < w.GridHeight-1; y++ {
		for x := 1; x < w.GridWidth-1; x++ {
			idx := y*w.GridWidth + x

			sumHome := (w.HomePheromones[idx-1] + w.HomePheromones[idx+1] +
				w.HomePheromones[idx-w.GridWidth] + w.HomePheromones[idx+w.GridWidth] +
				w.HomePheromones[idx]) / DiffusionStrength

			newValHome := sumHome * PheromoneDecay
			if newValHome > PheromoneCap {
				newValHome = PheromoneCap
			}
			w.HomeTemp[idx] = newValHome

			sumFood := (w.FoodPheromones[idx-1] + w.FoodPheromones[idx+1] +
				w.FoodPheromones[idx-w.GridWidth] + w.FoodPheromones[idx+w.GridWidth] +
				w.FoodPheromones[idx]) / DiffusionStrength

			newValFood := sumFood * PheromoneDecay
			if newValFood > PheromoneCap {
				newValFood = PheromoneCap
			}
			w.FoodTemp[idx] = newValFood
		}
	}

	w.HomePheromones, w.HomeTemp = w.HomeTemp, w.HomePheromones
	w.FoodPheromones, w.FoodTemp = w.FoodTemp, w.FoodPheromones
}

func (w *World) UpdateAnts() {
	for i := range w.Ants {
		currentAnt := &w.Ants[i]

		// Steer & Move
		var guiding []float64
		if currentAnt.State == ant.SearchingForFood {
			guiding = w.FoodPheromones
		} else {
			guiding = w.HomePheromones
		}

		currentAnt.ApplySteering(w.GridWidth, w.GridHeight, guiding)
		currentAnt.Move(float64(w.Width), float64(w.Height), w.Obstacles, w.GridWidth)

		// Check destination
		if currentAnt.State == ant.SearchingForFood {
			if currentAnt.IsAtFoodSource(w.FoodSources) {
				currentAnt.State = ant.ReturningHome
				currentAnt.AngleRadians += math.Pi // Turn around
				currentAnt.Scent = InitialScentStrength
			}
		} else if currentAnt.State == ant.ReturningHome {
			if currentAnt.IsAtHome(w.HomePosition, HomeRadius) {
				currentAnt.State = ant.SearchingForFood
				currentAnt.AngleRadians += math.Pi // Head back out

				// Stats
				currentAnt.GatheredFood++
				w.FoodCollected++
				w.updateLeaderboard()
			}
		}

		// Deposit
		var release []float64
		if currentAnt.State == ant.SearchingForFood {
			release = w.HomePheromones
		} else {
			release = w.FoodPheromones
		}

		currentAnt.DepositPheromone(w.GridWidth, w.GridHeight, release, w.FoodSources)
	}
}

func (w *World) UpdateCamera() {
	// 1. Keyboard Movement
	const camSpeed = 10.0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		w.CameraPosition.Y -= camSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		w.CameraPosition.Y += camSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		w.CameraPosition.X -= camSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		w.CameraPosition.X += camSpeed
	}

	// 2. Mouse Drag (Middle button or Left button)
	currX, currY := ebiten.CursorPosition()
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		dx := float64(currX - w.lastMouseX)
		dy := float64(currY - w.lastMouseY)

		w.CameraPosition.X -= dx / w.Zoom
		w.CameraPosition.Y -= dy / w.Zoom
	}
	w.lastMouseX, w.lastMouseY = currX, currY

	// 3. Scroll Wheel Zoom
	_, wheelY := ebiten.Wheel()
	if wheelY != 0 {
		oldZoom := w.Zoom
		w.Zoom += wheelY * 0.1
		if w.Zoom < 0.1 {
			w.Zoom = 0.1
		}
		if w.Zoom > 5.0 {
			w.Zoom = 5.0
		}

		w.CameraPosition.X += (float64(currX) / oldZoom) - (float64(currX) / w.Zoom)
		w.CameraPosition.Y += (float64(currY) / oldZoom) - (float64(currY) / w.Zoom)
	}
}

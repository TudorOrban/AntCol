package world

import (
	"ant-sim/internal/ant"
	"fmt"
	"image/color"
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (w *World) RenderStats(screen *ebiten.Image) {
	debugBar := ebiten.NewImage(200, 400)
	debugBar.Fill(color.RGBA{0, 0, 0, 150})
	screen.DrawImage(debugBar, nil)

	minutesPassed := float64(w.TotalTicks) / 3600.0
	avgPerMin := 0.0
	if minutesPassed > 0 {
		avgPerMin = float64(w.FoodCollected) / minutesPassed
	}

	stats := fmt.Sprintf("Food Collected: %d\nAvg Food/Min: %.2f", w.FoodCollected, avgPerMin)
	ebitenutil.DebugPrint(screen, stats)

	yOffset := 60
	ebitenutil.DebugPrintAt(screen, "--- TOP PRODUCERS ---", 10, yOffset)
	for i, a := range w.TopAnts {
		yOffset += 20
		msg := fmt.Sprintf("Ant #%d: %d food", i, a.GatheredFood)
		ebitenutil.DebugPrintAt(screen, msg, 10, yOffset)
	}
}

func (w *World) updateLeaderboard() {
	allAnts := make([]*ant.Ant, len(w.Ants))
	for i := range w.Ants {
		allAnts[i] = &w.Ants[i]
	}

	sort.Slice(allAnts, func(i, j int) bool {
		return allAnts[i].GatheredFood > allAnts[j].GatheredFood
	})

	limit := 5
	if len(allAnts) < limit {
		limit = len(allAnts)
	}

	w.TopAnts = allAnts[:limit]
}

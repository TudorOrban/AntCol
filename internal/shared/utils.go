package shared

import "math"

const (
	GridScale = 2
)

func GetDistance(src Position, dest Position) float64 {
	return math.Sqrt((dest.X-src.X)*(dest.X-src.X) + (dest.Y-src.Y)*(dest.Y-src.Y))
}

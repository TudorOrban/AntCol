package shared

const (
	GridScale = 2
)

type Position struct {
	X, Y float64
}

type FoodSource struct {
	Position Position
	Radius   float64
}

type Rectangle struct {
	X, Y          float64
	Width, Height float64
}

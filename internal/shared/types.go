package shared

type Position struct {
	X, Y float64
}

type FoodSource struct {
	Position Position
	Radius   float64
}

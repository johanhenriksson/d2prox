package d2prox

import (
	"fmt"
	"math"
)

type Vec2 struct {
	X int
	Y int
}

func (v *Vec2) String() string {
	return fmt.Sprintf("(%d,%d)", v.X, v.Y)
}

func (v *Vec2) Distance(to *Vec2) int {
	return Distance(v, to)
}

func Distance(v1, v2 *Vec2) int {
	dx := float64(v1.X - v2.X)
	dy := float64(v1.Y - v2.Y)
	return int(math.Sqrt(dx*dx + dy*dy))
}

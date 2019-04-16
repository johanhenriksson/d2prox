package d2prox

type WarpMap map[int]*Warp

type Warp struct {
	ID       int
	Type     int
	ClassID  int
	Position Vec2
}

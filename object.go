package d2prox

type ObjectMap map[int]*Object

type Object struct {
	ID       int
	Type     int
	Code     int
	Position Vec2
}

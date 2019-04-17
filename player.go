package d2prox

type PlayerMap map[int]*Player

type Player struct {
	ID       int
	Name     string
	Class    PlayerClass
	Position Vec2
	Health   int
	Mana     int
	Stats    map[int]int
}

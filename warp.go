package d2prox

type WarpMap map[int]*Warp

type Warp struct {
	ID       int
	Type     int
	ClassID  int
	Position Vec2
}

type LevelWarp struct {
	Name       string
	ID         int
	SelectX    int
	SelectY    int
	SelectDX   int
	SelectDY   int
	ExitWalkX  int
	ExitWalkY  int
	OffsetX    int
	OffsetY    int
	LitVersion int
	Tiles      int
	Direction  int
	Beta       int
}

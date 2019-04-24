package d2prox

var Levels = LevelsMap{}

func init() {
	var err error
	Levels, err = ReadLevels("./data/Levels.txt")
	if err != nil {
		panic(err)
	}
}

type LevelsMap map[int]*Level

func ReadLevels(path string) (LevelsMap, error) {
	levels := make(LevelsMap)
	err := ParseTsv(path, func(record []string) error {
		level := &Level{}
		if err := UnmarshalCsv(record, level); err != nil {
			return err
		}
		levels[level.ID] = level
		return nil
	})
	if err != nil {
		return nil, err
	}
	return levels, err
}

type Level struct {
	Name         string
	ID           int
	Pal          int
	Act          int
	Layer        int
	SizeX        int
	SizeY        int
	OffsetX      int
	OffsetY      int
	Depend       int
	Rain         int
	Mud          int
	NoPer        int
	LOSDraw      int
	FloorFilter  int
	BlankScreen  int
	DrawEdges    int
	IsInside     int
	DrlgType     int
	LevelType    int
	SubType      int
	SubTheme     int
	SubWaypoint  int
	SubShrine    int
	Vis0         int
	Vis1         int
	Vis2         int
	Vis3         int
	Vis4         int
	Vis5         int
	Vis6         int
	Vis7         int
	Warp0        int
	Warp1        int
	Warp2        int
	Warp3        int
	Warp4        int
	Warp5        int
	Warp6        int
	Warp7        int
	Intensity    int
	Red          int
	Green        int
	Blue         int
	Portal       int
	Position     int
	SaveMonsters int
	Quest        int
	WarpDist     int
	MonLvl1      int
	MonLvl2      int
	MonLvl3      int
	MonDen       int
	MonUMin      int
	MonUMax      int
	MonWndr      int
	MonSpcWalk   int
	Mtot         int
	M1           int
	M2           int
	M3           int
	M4           int
	M5           int
	M6           int
	M7           int
	M8           int
	M9           int
	M10          int
	M11          int
	M12          int
	M13          int
	M14          int
	M15          int
	M16          int
	M17          int
	M18          int
	M19          int
	M20          int
	M21          int
	M22          int
	M23          int
	M24          int
	M25          int
	S1           int
	S2           int
	S3           int
	S4           int
	S5           int
	S6           int
	S7           int
	S8           int
	S9           int
	S10          int
	S11          int
	S12          int
	S13          int
	S14          int
	S15          int
	S16          int
	S17          int
	S18          int
	S19          int
	S20          int
	S21          int
	S22          int
	S23          int
	S24          int
	S25          int
	Utot         int
	U1           int
	U2           int
	U3           int
	U4           int
	U5           int
	U6           int
	U7           int
	U8           int
	U9           int
	U10          int
	U11          int
	U12          int
	U13          int
	U14          int
	U15          int
	U16          int
	U17          int
	U18          int
	U19          int
	U20          int
	U21          int
	U22          int
	U23          int
	U24          int
	U25          int
	C1           int
	C2           int
	C3           int
	C4           int
	C5           int
	CA1          int
	CA2          int
	CA3          int
	CA4          int
	CA5          int
	CD1          int
	CD2          int
	CD3          int
	CD4          int
	CD5          int
	Themes       int
	SoundEnv     int
	Waypoint     int
	LevelName    string
	LevelWarp    string
	EntryFile    string
	ObjGrp0      int
	ObjGrp1      int
	ObjGrp2      int
	ObjGrp3      int
	ObjGrp4      int
	ObjGrp5      int
	ObjGrp6      int
	ObjGrp7      int
	ObjPrb0      int
	ObjPrb1      int
	ObjPrb2      int
	ObjPrb3      int
	ObjPrb4      int
	ObjPrb5      int
	ObjPrb6      int
	ObjPrb7      int
	Beta         int
}

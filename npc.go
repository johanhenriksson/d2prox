package d2prox

import "fmt"

type NPCMap map[int]*NPC

// NPC instance
type NPC struct {
	*NPCType
	ID       int
	Class    int
	Life     int
	Position Vec2
}

// Health returns npc health as a fraction
func (npc *NPC) Health() float64 {
	return float64(npc.Life) / 128.0
}

// String returns a string representation
func (npc *NPC) String() string {
	return fmt.Sprintf("[%x] %s %.f%% at %s", npc.ID, npc.Name, npc.Health(), npc.Position.String())
}

// NPCTypeIDs maps npc class ids to NPC Types
var NPCTypeIDs = map[int]*NPCType{}

func init() {
	// initial setup of the type mapping
	for _, npc := range NPCTypes {
		NPCTypeIDs[npc.ClassID] = npc
	}
}

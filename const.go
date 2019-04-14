package d2prox

import "fmt"

//
// enums & constants
//

const JoinGameOk = 0x00

//
// battle.net server packets
//

const SidAuthInfo byte = 0x50
const SidLogonResponse2 byte = 0x3A
const SidLogonRealmEx byte = 0x3E

//
// realm server packets
//

const McpJoinGame byte = 0x04
const McpStartup byte = 0x01

//
//
//

type UnitType int

const UnitTypePlayer = UnitType(0x00)
const UnitTypeNPC = UnitType(0x01)
const UnitTypeObject = UnitType(0x02)

type Quality int

const QualityNone = Quality(0x00)
const QualityInferior = Quality(0x01)
const QualityNormal = Quality(0x02)
const QualitySuperior = Quality(0x03)
const QualityMagic = Quality(0x04)
const QualitySet = Quality(0x05)
const QualityRare = Quality(0x06)
const QualityUnique = Quality(0x07)
const QualityCraft = Quality(0x08)

func (q Quality) String() string {
	switch q {
	case QualityInferior:
		return "Inferior"
	case QualityNormal:
		return ""
	case QualitySuperior:
		return "Superior"
	case QualityMagic:
		return "Magic"
	case QualitySet:
		return "Set"
	case QualityRare:
		return "Rare"
	case QualityUnique:
		return "Unique"
	case QualityCraft:
		return "Crafted"
	}
	return fmt.Sprintf("None [%d]", q)
}

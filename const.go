package d2prox

import "fmt"

//
// enums & constants
//

const JoinGameOk = 0x00

//
// battle.net server packets
//

const SidAuthInfo = 0x50
const SidAuthCheck = 0x51
const SidLogonResponse2 = 0x3A
const SidLogonRealmEx = 0x3E

//
// realm server packets
//

const McpStartup = 0x01
const McpCreateGame = 0x03
const McpJoinGame = 0x04

//
// unit type
//

type UnitType int

const UnitTypePlayer = UnitType(0x00)
const UnitTypeNPC = UnitType(0x01)
const UnitTypeObject = UnitType(0x02)
const UnitTypeMissile = UnitType(0x03)
const UnitTypeItem = UnitType(0x04)
const UnitTypeWarp = UnitType(0x05)

//
// player class
//

type PlayerClass int

const ClassAmazon = PlayerClass(0x00)
const ClassSorceress = PlayerClass(0x01)
const ClassNecromancer = PlayerClass(0x2)
const ClassPaladin = PlayerClass(0x03)
const ClassBarbarian = PlayerClass(0x4)
const ClassDruid = PlayerClass(0x5)
const ClassAssassin = PlayerClass(0x6)

func (c PlayerClass) String() string {
	switch c {
	case ClassAmazon:
		return "Amazon"
	case ClassSorceress:
		return "Sorceress"
	case ClassNecromancer:
		return "Necromancer"
	case ClassPaladin:
		return "Paladin"
	case ClassBarbarian:
		return "Barbarian"
	case ClassDruid:
		return "Druid"
	case ClassAssassin:
		return "Assassin"
	}
	return fmt.Sprintf("Unknown (%d)", c)
}

//
// item quality
//

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
	return fmt.Sprintf("None [%d]", int(q))
}

//
// item actions
//

type ItemAction int

const ItemActionNewGround = ItemAction(0x00)
const ItemActionPickup = ItemAction(0x01)
const ItemActionDrop = ItemAction(0x02)
const ItemActionOldGround = ItemAction(0x03)
const ItemActionToStorage = ItemAction(0x04)
const ItemActionFromStorage = ItemAction(0x05)
const ItemActionEquip = ItemAction(0x06)
const ItemActionIndirectSwapBody = ItemAction(0x07)
const ItemActionUnequip = ItemAction(0x08)
const ItemActionSwapBody = ItemAction(0x09)
const ItemActionAddQuantity = ItemAction(0x0a)
const ItemActionToStore = ItemAction(0x0b)
const ItemActionFromStore = ItemAction(0x0c)
const ItemActionSwapInContainer = ItemAction(0x0d)
const ItemActionPlaceBelt = ItemAction(0x0e)
const ItemActionRemoveBelt = ItemAction(0x0f)
const ItemActionSwapBelt = ItemAction(0x10)
const ItemActionAutoUnequip = ItemAction(0x11)
const ItemActionToCursor = ItemAction(0x12)
const ItemActionItemInSocket = ItemAction(0x13)
const ItemActionUpdateStats = ItemAction(0x15)
const ItemActionWeaponSwitch = ItemAction(0x17)

func (a ItemAction) String() string {
	switch a {
	case ItemActionDrop:
		return "Drop"
	case ItemActionEquip:
		return "Equip"
	case ItemActionPickup:
		return "Pickup"
	case ItemActionToCursor:
		return "ToCursor"
	}
	return fmt.Sprintf("ItemAction(%X)", int(a))
}

//
// item group
//

const ItemGroupAllArmor = 0x20000000
const ItemGroupAllWeapons = 0x40000000

package d2prox

import (
	"fmt"
	"strings"
)

type Item struct {
	ID             int
	Type           *ItemType
	Category       int
	Owner          int
	Code           string
	Equipped       bool
	InSocket       bool
	Identified     bool
	Broken         bool
	Socketed       bool
	Ethereal       bool
	Runeword       bool
	Simple         bool
	Gamble         bool
	ForSale        bool
	Ground         bool
	IsArmor        bool
	IsWeapon       bool
	Indestructible bool
	Version        int
	Container      int
	X              int
	Y              int
	Amount         int
	UsedSockets    int
	Level          int
	Quality        Quality
	Graphic        int
	Color          int
	Prefix         int
	Suffix         int
	Set            int
	UniqueCode     int
	Defense        int
	Durability     int
	MaxDurability  int
	Sockets        int
	Prefixes       []int
	Suffixes       []int
	Stats          []Stat
}

func (i *Item) String() string {
	parts := []string{}
	if i.Ethereal {
		parts = append(parts, "Ethereal")
	}
	if i.Quality != QualityNormal {
		parts = append(parts, i.Quality.String())
	}
	parts = append(parts, i.Type.Name)
	if i.Level > 1 {
		parts = append(parts, fmt.Sprintf("(%d)", i.Level))
	}
	return strings.Join(parts, " ")
}

func ParseItem(packet Packet) *Item {
	//fmt.Println(hex.Dump(packet))

	item := &Item{
		Level:   1,
		Quality: QualityNormal,
	}

	r := NewBitField(packet)
	msgID := r.Byte() // message id
	r.Skip(8)         // action
	r.Skip(8)         // packet length
	item.Category = r.Byte()
	item.ID = r.Bits(32)

	// 0x9d has extra data here
	if msgID == GsItemActionOwned {
		r.Bits(8)
		item.Owner = r.Bits(32) // owner id
	}

	item.Equipped = r.Bool() // 64
	r.Bit()                  // just bought 65
	if r.Bit() != 0 {        // unknown 1 66
		fmt.Println("expected unknown bit 1 to be zero")
	}
	item.InSocket = r.Bool()   // 67
	item.Identified = r.Bool() // 68

	if r.Bit() != 0 { // unknown 2 69
		fmt.Println("expected unknown bit 2 to be zero")
	}

	r.Bit()                // switched in 70
	r.Bit()                // switched out 71
	item.Broken = r.Bool() // broken 72

	if r.Bit() != 0 { // unknown 3 73
		fmt.Println("expected unknown bit 3 to be zero")
	}

	r.Bool() // potion thingy?
	item.Socketed = r.Bool()

	if r.Bit() != 0 { // unknown 4
		fmt.Println("expected unknown bit 4 to be zero")
	}

	item.ForSale = r.Bool()
	notInSocket := r.Bool()
	if (notInSocket || item.InSocket) && notInSocket == item.InSocket {
		fmt.Println("Item is both in and out of socket o_O")
	}
	r.Bit() // unknown

	ear := r.Bool()
	r.Bool()  // start item
	r.Bits(3) // unknown
	item.Simple = r.Bool()
	item.Ethereal = r.Bool()
	r.Bit() // has magic stats

	personalized := r.Bool() // personalized
	item.Gamble = r.Bool()
	item.Runeword = r.Bool()
	r.Bits(5) // unknown

	item.Version = r.Byte()

	r.Bits(2) // unknown
	destination := r.Bits(3)

	ground := destination == 0x03
	if ground {
		item.X = r.Bits(16)
		item.Y = r.Bits(16)
		fmt.Printf("Ground at %d,%d\n", item.X, item.Y)
	} else {
		directory := r.Bits(4)
		item.X = r.Bits(4)
		item.Y = r.Bits(4)
		item.Container = r.Bits(3)
		fmt.Printf("Stored %d,%d dir: %d, container: %d\n", item.X, item.Y, directory, item.Container)
	}

	// ear special case
	if ear {
		item.Code = "ear"
		return item
	}

	// item code
	codebytes := make([]byte, 3)
	for i := 0; i < 3; i++ {
		codebytes[i] = byte(r.Byte())
	}
	item.Code = string(codebytes)
	r.Skip(8) // 4 bytes? :S

	item.Type = ItemTypes[item.Code]

	// gold
	if item.Code == "gld" {
		bigPile := r.Bool()
		if bigPile {
			item.Amount = r.Bits(32)
		} else {
			item.Amount = r.Bits(12)
		}
		return item
	}

	item.UsedSockets = r.Bits(3)

	if item.Simple || item.Gamble {
		return item
	}

	item.Level = r.Bits(7)
	item.Quality = Quality(r.Bits(4))

	hasGraphic := r.Bool()
	if hasGraphic {
		item.Graphic = r.Bits(3)
	}

	hasColor := r.Bool()
	if hasColor {
		item.Color = r.Bits(11)
	}

	if item.Identified {
		switch item.Quality {
		case QualityInferior:
			item.Prefix = r.Bits(3)

		/*
			case ITEM_QUALITY_SUPERIOR:
				item->superiority = static_cast<unsigned int>(reader.read(3));
				break;
		*/

		case QualityMagic:
			item.Prefix = r.Bits(11)
			item.Suffix = r.Bits(11)

		case QualityCraft:
			item.Prefix = r.Bits(8) - 156
			item.Suffix = r.Bits(8) - 1

		case QualityRare:
			item.Prefix = r.Bits(8) - 156
			item.Suffix = r.Bits(8) - 1

		case QualitySet:
			item.Set = r.Bits(12)

		case QualityUnique:
			if item.Code == "std" {
				break
			}
			item.UniqueCode = r.Bits(12)
		}
	}

	if item.Quality == QualityRare || item.Quality == QualityCraft {
		for i := 0; i < 3; i++ {
			if r.Bool() {
				item.Prefixes = append(item.Prefixes, r.Bits(11))
			}
			if r.Bool() {
				item.Suffixes = append(item.Suffixes, r.Bits(11))
			}
		}
	}

	if item.Runeword {
		r.Skip(12) // runeword id
		r.Skip(4)  // runeword param
	}

	if personalized {
		// skip personalized name
		for i := 0; i < 16; i++ {
			letter := r.Bits(7)
			if letter == 0 {
				break
			}
		}
	}

	item.IsArmor = (item.Type.Flags & ItemGroupAllArmor) > 0
	item.IsWeapon = (item.Type.Flags & ItemGroupAllWeapons) > 0

	// defense
	if item.IsArmor {
		item.Defense = r.Bits(11) - 10
	}

	// durability
	if item.Code == "7cr" {
		// indestructible phase blade
		r.Skip(8)
	} else if item.IsArmor || item.IsWeapon {
		item.MaxDurability = r.Bits(8)
		item.Indestructible = item.MaxDurability == 0
		item.Durability = r.Bits(8)
		r.Skip(1)
	}

	// sockets
	if item.Socketed {
		item.Sockets = r.Bits(4)
	}

	// if the item is unidentified, we're all done
	if !item.Identified {
		return item
	}

	// todo: convert stackable & usable to bool
	if item.Type.Stackable > 0 {
		if item.Type.Useable > 0 {
			r.Skip(5)
		}
		item.Amount = r.Bits(9)
	}

	// set modifiers
	if item.Quality == QualitySet {
		// set mods
		r.Bits(5)
	}

	// magic stats
	for {
		statID := r.Bits(9)
		if statID == 0x1ff {
			break
		}

		stat := ParseStat(statID, r)
		item.Stats = append(item.Stats, stat)
	}

	return item
}

type Stat struct {
	*StatProperty
	ID         int
	Skill      int
	Chance     int
	Value      int
	Monster    int
	Level      int
	Charges    int
	MaxCharges int
	SkillTab   int
	Class      int
	MinDmg     int
	MaxDmg     int
	Length     int
}

func ParseStat(id int, r *BitField) Stat {
	p := StatProperties[id]
	stat := Stat{
		StatProperty: p,
		ID:           id,
	}

	if p.SaveParamBits > 0 {
		switch id {
		case STAT_CLASSSKILLS:
			stat.Class = r.Bits(p.SaveParamBits)
			stat.Value = r.Bits(p.SaveBits)
		case STAT_NONCLASSSKILL:
		case STAT_SINGLESKILL:
			stat.Skill = r.Bits(p.SaveParamBits)
			stat.Value = r.Bits(p.SaveBits)
		case STAT_ELEMENTALSKILLS:
			r.Bits(p.SaveParamBits) // element
			stat.Value = r.Bits(p.SaveBits)
		case STAT_AURA:
			stat.Skill = r.Bits(p.SaveParamBits)
			stat.Value = r.Bits(p.SaveBits)
		case STAT_REANIMATE:
			stat.Monster = r.Bits(p.SaveParamBits)
			stat.Value = r.Bits(p.SaveBits)
		case STAT_SKILLTAB:
			stat.SkillTab = r.Bits(3)
			stat.Class = r.Bits(3)
			r.Skip(10) // unknown
			stat.Value = r.Bits(p.SaveBits)
		case STAT_SKILLONDEATH:
			fallthrough
		case STAT_SKILLONHIT:
			fallthrough
		case STAT_SKILLONKILL:
			fallthrough
		case STAT_SKILLONLEVELUP:
			fallthrough
		case STAT_SKILLONSTRIKING:
			fallthrough
		case STAT_SKILLWHENSTRUCK:
			stat.Level = r.Bits(6)
			stat.Skill = r.Bits(10)
			stat.Chance = r.Bits(p.SaveBits)
		case STAT_CHARGED:
			stat.Level = r.Bits(6)
			stat.Skill = r.Bits(10)
			stat.Charges = r.Bits(8)
			stat.MaxCharges = r.Bits(8)
		case STAT_STATE:
			fallthrough
		case STAT_ATTCKRTNGVSMONSTERTYPE:
			fallthrough
		case STAT_DAMAGETOMONSTERTYPE:
			// For some reason heroin_glands doesn't read these, even though
			// they have saveParamBits; maybe they don't occur in practice?
			stat.Value = r.Bits(p.SaveBits) - p.SaveAdd
		default:
			r.Skip(p.SaveParamBits)
			r.Skip(p.SaveBits)
		}
		return stat
	}

	switch id {
	case STAT_ENHANCEDMAXIMUMDAMAGE:
		fallthrough
	case STAT_ENHANCEDMINIMUMDAMAGE:
		stat.MinDmg = r.Bits(p.SaveBits)
		stat.MaxDmg = r.Bits(p.SaveBits)
	case STAT_MINIMUMFIREDAMAGE:
		stat.MinDmg = r.Bits(p.SaveBits)
		stat.MaxDmg = r.Bits(StatProperties[STAT_MAXIMUMFIREDAMAGE].SaveBits)
	case STAT_MINIMUMLIGHTNINGDAMAGE:
		stat.MinDmg = r.Bits(p.SaveBits)
		stat.MaxDmg = r.Bits(StatProperties[STAT_MAXIMUMLIGHTNINGDAMAGE].SaveBits)
	case STAT_MINIMUMMAGICALDAMAGE:
		stat.MinDmg = r.Bits(p.SaveBits)
		stat.MaxDmg = r.Bits(StatProperties[STAT_MAXIMUMMAGICALDAMAGE].SaveBits)
	case STAT_MINIMUMCOLDDAMAGE:
		stat.MinDmg = r.Bits(p.SaveBits)
		stat.MaxDmg = r.Bits(StatProperties[STAT_MAXIMUMCOLDDAMAGE].SaveBits)
		stat.Length = r.Bits(StatProperties[STAT_COLDDAMAGELENGTH].SaveBits)
	case STAT_MINIMUMPOISONDAMAGE:
		stat.MinDmg = r.Bits(p.SaveBits)
		stat.MaxDmg = r.Bits(StatProperties[STAT_MAXIMUMPOISONDAMAGE].SaveBits)
		stat.Length = r.Bits(StatProperties[STAT_POISONDAMAGELENGTH].SaveBits)
	case STAT_REPAIRSDURABILITY:
		fallthrough
	case STAT_REPLENISHESQUANTITY:
		stat.Value = r.Bits(p.SaveBits)
	default:
		stat.Value = r.Bits(p.SaveBits) - p.SaveAdd
	}

	return stat
}

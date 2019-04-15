package d2prox

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
)

// GamePort is the default game server port
const GamePort = 4000

var gameTargets = make(map[string]string)

// GameProxy implements the game server proxy
type GameProxy struct {
	ProxyServer
}

// NewGame creates a new game server proxy
func NewGame() *GameProxy {
	return &GameProxy{
		ProxyServer{
			Name: "game",
			port: GamePort,
		},
	}
}

// Accept a new connection
func (p *GameProxy) Accept(conn net.Conn) {
	c := &GameClient{
		ProxyClient: &ProxyClient{
			Proxy:  p,
			client: conn,
		},
		Game: &Game{
			Players: make(PlayerMap),
			NPCs:    make(NPCMap),
			Objects: make(ObjectMap),
			Items:   make(ItemMap),
			Warps:   make(WarpMap),
		},
	}
	HandleProxySession(p, c, PacketReader(gsClientPacketLength), PacketReader(gsServerPacketLength))
}

// GameClient implements the game server proxy client
type GameClient struct {
	*ProxyClient
	GameCount int
	Game      *Game
	Player    *Player
}

type ObjectMap map[int]*Object
type PlayerMap map[int]*Player
type NPCMap map[int]*NPC
type WarpMap map[int]*Warp
type ItemMap map[int]*Item

type Game struct {
	Difficulty int
	Ladder     bool
	Hardcore   bool
	Expansion  bool
	ExpGained  int
	Players    PlayerMap
	NPCs       NPCMap
	Objects    ObjectMap
	Items      ItemMap
	Warps      WarpMap
}

type Object struct {
	ID   int
	Type int
	Code int
	X    int
	Y    int
}

type Player struct {
	ID     int
	Name   string
	Class  int
	X      int
	Y      int
	Health int
	Mana   int
	Stats  map[int]int
}

type NPC struct {
	ID    int
	Class int
	X     int
	Y     int
	Life  int
}

type Warp struct {
	ID      int
	Type    int
	ClassID int
	X       int
	Y       int
}

// OnAccept is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *GameClient) OnAccept() {
	// client wont send game info until the server sends GsConnectionInfo (0xAF)
	// we'll send it manually and silence it later.
	c.WriteClient(Packet{GsConnectionInfo, 0x00})
}

//
// server -> client
//

// HandleServer packets
func (c *GameClient) HandleServer(packet Packet) Packet {
	pb := PacketBuffer(packet)

	switch packet.GsMsgID() {
	case GsConnectionInfo:
		// silence ConnectionInfo, since we send it manually in Connect()
		return nil

	case GsGameFlags:
		c.Game.Difficulty = pb.Byte(1)
		if pb.Byte(6) != 0 {
			c.Game.Expansion = true
		}
		if pb.Byte(7) != 0 {
			c.Game.Ladder = true
		}
		fmt.Println("hardcore:", pb.Uint16(4))
		return packet

	case GsHandshake:
		id := pb.Uint32(2)
		// set local player reference
		c.Player = c.Game.Players[id]
		c.GameCount++
		c.SendChat(fmt.Sprintf("game %d joined", c.GameCount))
		c.Proxy.Log("I am player %x", id)
		return packet

	//
	// player updates
	//

	case GsAssignPlayer:
		player := &Player{
			ID:    pb.Uint32(1),
			Name:  pb.NullString(6),
			Class: pb.Byte(5),
			X:     pb.Uint16(22),
			Y:     pb.Uint16(24),
			Stats: make(map[int]int),
		}
		c.Game.Players[player.ID] = player
		return packet

	case GsPlayerMove:
		id := pb.Uint32(2)
		if player, exists := c.Game.Players[id]; exists {
			player.X = pb.Uint16(12)
			player.Y = pb.Uint16(14)
		}
		return packet

	case GsPlayerStop:
		id := pb.Uint32(2)
		if player, exists := c.Game.Players[id]; exists {
			player.X = pb.Uint16(7)
			player.Y = pb.Uint16(9)
		}
		return packet

	//
	// npc updates
	//

	case GsAssignNPC:
		npc := &NPC{
			ID:    pb.Uint32(1),
			Class: pb.Uint16(5),
			X:     pb.Uint16(7),
			Y:     pb.Uint16(9),
			Life:  pb.Byte(11),
		}
		c.Game.NPCs[npc.ID] = npc
		fmt.Println("Add NPC", npc)
		return packet

	case GsNPCMove:
		id := pb.Uint32(1)
		if npc, exists := c.Game.NPCs[id]; exists {
			npc.X = pb.Uint16(6)
			npc.Y = pb.Uint16(8)
		} else {
			fmt.Println("NPCMove: Unknown NPC")
		}
		return packet // silent

	case GsNPCStop:
		id := pb.Uint32(1)
		if npc, exists := c.Game.NPCs[id]; exists {
			npc.X = pb.Uint16(5)
			npc.Y = pb.Uint16(7)
			npc.Life = pb.Byte(9)
		} else {
			fmt.Println("NPCMove: Unknown NPC")
		}
		return packet // silent

	case GsNPCHit:
		id := pb.Uint32(2)
		if npc, exists := c.Game.NPCs[id]; exists {
			npc.Life = pb.Byte(8)
		} else {
			fmt.Println("NPCMove: Unknown NPC")
		}
		return packet

	//
	// object updates
	//

	case GsAssignObject:
		object := &Object{
			ID:   pb.Uint32(2),
			Type: pb.Byte(1),
			Code: pb.Uint16(6),
			X:    pb.Uint16(8),
			Y:    pb.Uint16(10),
		}
		c.Game.Objects[object.ID] = object
		return packet

	case GsRemoveObject:
		kind := UnitType(pb.Byte(1))
		id := pb.Uint32(2)
		switch kind {
		case UnitTypePlayer:
			delete(c.Game.Players, id)
		case UnitTypeNPC:
			delete(c.Game.NPCs, id)
		case UnitTypeObject:
			delete(c.Game.Objects, id)
		case UnitTypeItem:
			delete(c.Game.Items, id)
		default:
			c.Proxy.Log("Remove item %d (type: %d)", id, kind)
		}
		return packet

	case GsReportKill:
		kind := UnitType(pb.Byte(1))
		id := pb.Uint32(2)
		fmt.Println("Kill type", kind, "id:", id)
		switch kind {
		case UnitTypeNPC:
			if npc, exists := c.Game.NPCs[id]; exists {
				npc.Life = 0
			}
		}
		return packet

	case GsPlayerLeft:
		id := pb.Uint32(1)
		delete(c.Game.Players, id)

	// set attributes
	case GsSetAttr8:
		attr := pb.Byte(1)
		value := pb.Byte(2)
		c.Player.Stats[attr] = value
		if name, exists := Attrs[attr]; exists {
			fmt.Println(name, "=", value)
			return packet
		}
		return packet
	case GsSetAttr16:
		attr := pb.Byte(1)
		value := pb.Uint16(2)
		c.Player.Stats[attr] = value
		if name, exists := Attrs[attr]; exists {
			fmt.Println(name, "=", value)
			return packet
		}
	case GsSetAttr32:
		attr := pb.Byte(1)
		value := pb.Uint32(2)
		c.Player.Stats[attr] = value
		if name, exists := Attrs[attr]; exists {
			fmt.Println(name, "=", value)
			return packet
		}

	case GsLifeManaUpdate:
		src := pb.Uint32(1)
		c.Player.Health = (src & 0x00007FFF)
		c.Player.Mana = (src & 0x3FFF8000) >> 15
		return packet

	// experience
	case GsAddExp8:
		exp := pb.Byte(1)
		c.Game.ExpGained += exp
	case GsAddExp16:
		exp := pb.Uint16(1)
		c.Game.ExpGained += exp
	case GsAddExp32:
		exp := pb.Uint32(1)
		c.Game.ExpGained += exp

	//
	// items
	//

	case GsItemActionOwned:
		item := ParseItem(packet)
		action := ItemAction(pb.Byte(1))
		c.Proxy.Log("OwnedItem - %s action: %s", item, action)
		return packet

	case GsItemActionWorld:
		item := ParseItem(packet)
		action := ItemAction(pb.Byte(1))
		c.Proxy.Log("WorldItem - %s action: %s", item, action)
		switch action {
		case ItemActionDrop:
			c.Game.Items[item.ID] = item
		}
		return packet

	case GsWardenRequest:
		c.Proxy.Log("Warden Request")

	//
	// other
	//

	case GsRelator1:
		return packet

	case GsRelator2:
		return packet

	case GsPong:
		return packet // silent
	}

	c.Proxy.Log("S->C: %s", GsServerPacketName(packet))
	return packet
}

//
// client -> server
//

// HandleBuffered packets
func (c *GameClient) HandleBuffered(packet Packet) Packet {
	switch packet.GsMsgID() {
	case GsGameLogon:
		logon := GsGameLogonPacket(packet)
		return Packet(c.handleGameLogon(logon))
	}

	c.Proxy.Log("C->S: %s", GsClientPacketName(packet))
	return packet
}

func (c *GameClient) handleGameLogon(packet GsGameLogonPacket) GsGameLogonPacket {
	// the D2GS_GAMELOGON packet contains the token data we need to look up the
	// cached game server ip stored by the realm server proxy. once we have it, we
	// can connect to the game server.

	// lookup the target server associated with the token
	token := packet.Token()
	target, exists := gameTargets[token]
	if !exists {
		// we dont have it - drop the client
		c.Proxy.Log("game target %s not found", token)
		c.Close()
		return nil
	}

	// one time use only - delete token/target pair from cache
	delete(gameTargets, token)

	// manually buffer packet so that it will be available on connect.
	// this packet will be silenced to avoid possible duplication
	c.BufferPacket(Packet(packet))

	// connect to target game server
	if err := c.Connect(target); err != nil {
		c.Proxy.Log("error connecting to game server: %s", err)
	}

	c.Proxy.Log("token %s proxied to game server %s", token, target)
	return nil
}

// HandleClient packets
func (c *GameClient) HandleClient(packet Packet) Packet {
	pb := PacketBuffer(packet)

	switch packet.GsMsgID() {
	case GsChatMessage:
		return Packet(c.handleChatMessage(GsChatMessagePacket(packet)))

	case GsPickupItem:
		id := pb.Uint32(5)
		c.Proxy.Log("Pickup item %x", id)
		fmt.Println(hex.Dump(packet))
		return packet

	case GsDropItem:
		id := pb.Uint32(1)
		c.Proxy.Log("Drop item %x", id)

	case GsPing:
		return packet // silent
	}

	c.Proxy.Log("C->S: %s", GsClientPacketName(packet))
	return packet
}

func (c *GameClient) handleChatMessage(packet GsChatMessagePacket) GsChatMessagePacket {
	// chat commands
	msg := packet.Message()
	if msg[0] == '.' {
		command := strings.ToLower(msg[1:])
		switch command {
		case "exp":
			c.SendChat(fmt.Sprintf("Experience gained: %d", c.Game.ExpGained))
		}
		return nil
	}
	return packet
}

// SendChat writes a chat message packet to the client
func (c *GameClient) SendChat(message string) {
	msgbytes := []byte(message)
	nickbytes := []byte("d2prox")

	length := 10 + len(nickbytes) + 1 + len(msgbytes) + 1
	packet := make(PacketBuffer, length)
	packet.PutByte(0, GsGameChat)
	packet.PutByte(1, 0x01) // chat type
	packet.PutByte(2, 0x00) // locale id
	packet.PutByte(3, 0x02) // unit type
	packet.PutUint32(4, 0)  // unit id
	packet.PutByte(8, 0x00) // chat color
	packet.PutByte(9, 0x62) // subtype
	packet.Put(10, nickbytes)
	packet.Put(10+len(nickbytes)+1, msgbytes)

	c.WriteClient(Packet(packet))
}

var Attrs = map[int]string{
	0:  "strength",
	1:  "energy",
	2:  "dexterity",
	3:  "vitality",
	4:  "statpts",
	5:  "newskills",
	6:  "hitpoints",
	7:  "maxhp",
	8:  "mana",
	9:  "maxmana",
	10: "stamina",
	11: "maxstamina",
	12: "level",
	13: "experience",
}

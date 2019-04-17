package d2prox

import (
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"time"
)

// GamePort is the default game server port
const GamePort = 4000

var gameSessions = map[string]*GameSession{}

// Game represents the state of a game session
type Game struct {
	Start      time.Time
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
	Session *GameSession
	Game    *Game
	Player  *Player
	Ready   bool
}

// OnAccept is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *GameClient) OnAccept() {
	// client wont send game info until the server sends GsConnectionInfo (0xAF)
	// we'll send it manually and silence it later.
	c.WriteClient(Packet{GsConnectionInfo, 0x00})
}

// OnClose is fired immediately after a session is closed
func (c *GameClient) OnClose() {
	// invalidate game reference
	c.Session.Game = nil
}

//
// server -> client
//

// HandleServer packets
func (c *GameClient) HandleServer(packet Packet) Packet {
	pb := PacketBuffer(packet)

	switch packet.GsMsgID() {
	case GsConnectionInfo:
		// silence ConnectionInfo, since we send it manually in OnAccept()
		return nil

	case GsGameFlags:
		c.Game.Start = time.Now()
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
		c.Session.Game = c.Game
		c.Session.Games = append(c.Session.Games, c.Game)
		c.SendChat(fmt.Sprintf("game %d joined", len(c.Session.Games)))
		return packet

	case GsPlayerInGame:
		id := pb.Uint32(3)
		class := pb.Byte(7)
		name := pb.NullString(8)
		c.Proxy.Log("player %s [%x] joined game. class: %d", name, id, class)
		if id == c.Player.ID {
			// its me! this means we're all ready
			c.Ready = true
			c.Game.ExpGained = 0
		}
		return packet

	//
	// player updates
	//

	case GsAssignPlayer:
		player := &Player{
			ID:    pb.Uint32(1),
			Name:  pb.NullString(6),
			Class: PlayerClass(pb.Byte(5)),
			Stats: make(map[int]int),
			Position: Vec2{
				X: pb.Uint16(22),
				Y: pb.Uint16(24),
			},
		}
		c.Game.Players[player.ID] = player
		return packet

	case GsPlayerMove:
		id := pb.Uint32(2)
		if player, exists := c.Game.Players[id]; exists {
			player.Position.X = pb.Uint16(12)
			player.Position.Y = pb.Uint16(14)
		}
		return packet

	case GsPlayerStop:
		id := pb.Uint32(2)
		if player, exists := c.Game.Players[id]; exists {
			player.Position.X = pb.Uint16(7)
			player.Position.Y = pb.Uint16(9)
		}
		return packet

	//
	// npc updates
	//

	case GsAssignNPC:
		npc := &NPC{
			ID:    pb.Uint32(1),
			Class: pb.Uint16(5),
			Life:  pb.Byte(11),
			Position: Vec2{
				X: pb.Uint16(7),
				Y: pb.Uint16(9),
			},
		}
		npc.NPCType = NPCTypeIDs[npc.Class]
		c.Game.NPCs[npc.ID] = npc
		return packet

	case GsNPCMove:
		id := pb.Uint32(1)
		if npc, exists := c.Game.NPCs[id]; exists {
			npc.Position.X = pb.Uint16(6)
			npc.Position.Y = pb.Uint16(8)
		} else {
			fmt.Println("NPCMove: Unknown NPC")
		}
		return packet // silent

	case GsNPCStop:
		id := pb.Uint32(1)
		if npc, exists := c.Game.NPCs[id]; exists {
			npc.Position.X = pb.Uint16(5)
			npc.Position.Y = pb.Uint16(7)
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
			Position: Vec2{
				X: pb.Uint16(8),
				Y: pb.Uint16(10),
			},
		}
		c.Game.Objects[object.ID] = object
		return packet

	case GsAssignLevelWarp:
		warp := &Warp{
			ID:      pb.Uint32(2),
			Type:    pb.Byte(1),
			ClassID: pb.Byte(6),
			Position: Vec2{
				X: pb.Uint16(7),
				Y: pb.Uint16(9),
			},
		}
		c.Game.Warps[warp.ID] = warp
		fmt.Println("Assign warp", warp)
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
			c.Proxy.Log("remove item %d (type: %d)", id, kind)
		}
		return packet

	case GsReportKill:
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
		c.SendChat("warden request detected")
		c.Proxy.Log("Warden Request")
		fmt.Println(hex.Dump(packet))
		return packet

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

	if c.Session.Debug {
		c.Proxy.Log("S->C: %s", GsServerPacketName(packet))
	}
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

	if c.Session.Debug {
		c.Proxy.Log("C->S: %s", GsClientPacketName(packet))
	}
	return packet
}

func (c *GameClient) handleGameLogon(packet GsGameLogonPacket) GsGameLogonPacket {
	// the D2GS_GAMELOGON packet contains the token data we need to look up the
	// cached game server ip stored by the realm server proxy. once we have it, we
	// can connect to the game server.

	// lookup the target server associated with the token
	token := packet.Token()
	session, exists := gameSessions[token]
	if !exists {
		// we dont have it - drop the client
		c.Proxy.Log("game session %s not found", token)
		c.Close()
		return nil
	}

	// store in client
	c.Session = session

	// one time use only - delete token/target pair from cache
	delete(gameSessions, token)

	// manually buffer packet so that it will be available on connect.
	// this packet will be silenced to avoid possible duplication
	c.BufferPacket(Packet(packet))

	// connect to target game server
	if err := c.Connect(session.GameHost); err != nil {
		c.Proxy.Log("error connecting to game server: %s", err)
	}

	c.Proxy.Log("token %s proxied to game server %s", token, session.GameHost)
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
		return packet

	case GsDropItem:
		id := pb.Uint32(1)
		c.Proxy.Log("Drop item %x", id)
		return packet

	case GsPing:
		return packet // silent
	}

	if c.Session.Debug {
		c.Proxy.Log("C->S: %s", GsClientPacketName(packet))
	}
	return packet
}

func (c *GameClient) handleChatMessage(packet GsChatMessagePacket) GsChatMessagePacket {
	// chat commands
	msg := packet.Message()
	if msg[0] == '.' {
		command := strings.Trim(strings.ToLower(msg[1:]), " \t")
		if space := strings.Index(command, " "); space > 0 {
			command = command[:space]
		}

		switch command {
		case "exp":
			c.SendChat(fmt.Sprintf("Experience gained: %d", c.Game.ExpGained))
		case "debug":
			c.Session.Debug = !c.Session.Debug
			if c.Session.Debug {
				c.SendChat("Debug mode enabled")
			} else {
				c.SendChat("Debug mode disabled")
			}
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

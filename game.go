package d2prox

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// GamePort is the default game server port
const GamePort = 4000

var gameSessions = map[string]*GameSession{}

// GameList is a slice of game pointers
type GameList []*Game

// Game represents the state of a game session
type Game struct {
	Start time.Time
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
		Game: &Game{},
	}
	HandleProxySession(p, c, PacketReader(gsClientPacketLength), PacketReader(gsServerPacketLength))
}

// GameClient implements the game server proxy client
type GameClient struct {
	*ProxyClient
	Session *GameSession
	Game    *Game
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
	switch packet.GsMsgID() {
	case GsConnectionInfo:
		// silence ConnectionInfo, since we send it manually in OnAccept()
		return nil

	case GsHandshake:
		c.Game.Start = time.Now()
		c.Session.Game = c.Game
		c.Session.Games = append(c.Session.Games, c.Game)
		c.SendChat(fmt.Sprintf("game %d joined", len(c.Session.Games)))
		return packet
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
	switch packet.GsMsgID() {
	case GsChatMessage:
		return Packet(c.handleChatMessage(GsChatMessagePacket(packet)))

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

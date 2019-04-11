package d2prox

import (
	"encoding/hex"
	"fmt"
	"net"
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
	}
	HandleProxySession(p, c)
}

// GameClient implements the game server proxy client
type GameClient struct {
	*ProxyClient
}

// OnAccept is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *GameClient) OnAccept() {
	// client wont send game info until the server sends D2GS_STARTLOGON (0xAF)
	// we'll send it manually and silence it later.
	c.WriteClient(Packet{GsStartLogon, 0x00})
}

//
// server -> client
//

// HandleServer packets
func (c *GameClient) HandleServer(packet Packet) Packet {
	fmt.Println("GS S->C")
	fmt.Println(hex.Dump(packet))
	return packet
}

//
// client -> server
//

// HandleBuffered packets
func (c *GameClient) HandleBuffered(packet Packet) Packet {
	fmt.Println("GS C->S (B)")
	fmt.Println(hex.Dump(packet))

	switch packet.GsMsgID() {
	case GsGameLogon:
		logon := GsGameLogonPacket(packet)
		return Packet(c.handleGameLogon(logon))
	}

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
	fmt.Println("GS C->S")
	fmt.Println(hex.Dump(packet))

	switch packet.GsMsgID() {
	case GsStartLogon:
		// silence D2GS_STARTLOGON, since we send it manually in Connect()
		fmt.Println("Silenced D2GS_STARTLOGON")
		return nil
	}

	return packet
}

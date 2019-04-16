package d2prox

import (
	"encoding/hex"
	"fmt"
	"net"

	"github.com/johanhenriksson/d2prox/ip"
)

// RealmPort is the default realm server port
const RealmPort = 6113

var realmSessions = make(map[string]*GameSession)

// RealmProxy implements the realm proxy server
type RealmProxy struct {
	ProxyServer
}

// NewRealm creates a new realm proxy server
func NewRealm() *RealmProxy {
	return &RealmProxy{
		ProxyServer{
			Name: "realm",
			port: RealmPort,
		},
	}
}

// Accept a new connection
func (p *RealmProxy) Accept(conn net.Conn) {
	c := &RealmClient{
		ProxyClient: &ProxyClient{
			Proxy:  p,
			client: conn,
		},
	}
	HandleProxySession(p, c, PacketReader(realmPacketLength), PacketReader(realmPacketLength))
}

// bnetPacketLength computes the length of the next packet in the buffer
func realmPacketLength(buffer PacketBuffer, offset, length int) (int, error) {
	return buffer.Uint16(offset), nil
}

// RealmClient implements the realm proxy client
type RealmClient struct {
	*ProxyClient
	Session *GameSession
}

// OnAccept is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *RealmClient) OnAccept() {
	// read the game byte 0x01 and put it on the output buffer
	// this simplifies handling of the first packet
	b := Packet{0}
	c.Client().Read(b)
	c.BufferPacket(b)
}

//
// client -> server messages
//

// HandleBuffered packet
func (c *RealmClient) HandleBuffered(packet Packet) Packet {
	switch packet.RealmMsgID() {
	case McpStartup:
		startup := McpStartupPacket(packet)
		c.handleMcpStartup(startup)
	}

	return packet
}

func (c *RealmClient) handleMcpStartup(packet McpStartupPacket) {
	// extract token
	token := packet.Token()
	c.Proxy.Log("realm token: %s", token[8:16])

	// find session
	session, exists := realmSessions[token]
	if !exists {
		c.Proxy.Log("unknown token: %s", token)
		return
	}
	c.Session = session

	// clear target
	delete(realmSessions, token)

	// manually buffer packet so that it will be available on connect.
	// this packet will be silenced to avoid possible duplication
	c.BufferPacket(Packet(packet))

	// connect to realm server
	c.Proxy.Log("realm target: %s", session.RealmHost)
	if err := c.Connect(session.RealmHost); err != nil {
		c.Proxy.Log("error connecting to realm target:", session.RealmHost)
		c.Proxy.Log("%s", err)
		return
	}
}

func (c *RealmClient) HandleClient(packet Packet) Packet {
	switch packet.RealmMsgID() {
	case McpJoinGame:
		c.handleJoinGame(packet)
	}
	return packet
}

func (c *RealmClient) handleJoinGame(packet Packet) {
	pb := PacketBuffer(packet)
	name := pb.NullString(5)
	pass := pb.NullString(5 + len(name) + 1)
	c.Proxy.Log("Joining game %s//%s", name, pass)
}

//
// server -> client messages
//

// HandleServer packet
func (c *RealmClient) HandleServer(packet Packet) Packet {
	switch packet.RealmMsgID() {
	case McpJoinGame:
		join := McpJoinedGamePacket(packet)
		c.handleJoinedGame(join)
	case McpCreateGame:
		c.handleCreateGame(packet)
	}

	return packet
}

func (c *RealmClient) handleJoinedGame(packet McpJoinedGamePacket) {
	if packet.Status() != JoinGameOk {
		// join game failed, do nothing.
		c.Proxy.Log("failed to join game: %x", packet.Status())
		return
	}

	// create a unique game token
	token := make([]byte, 6)
	copy(token[0:4], packet.Hash())
	copy(token[4:6], packet.Token())
	tokenStr := hex.EncodeToString(token)

	// store game target
	gameIP := packet.GameIP()
	target := fmt.Sprintf("%s:%d", gameIP, GamePort)
	c.Session.GameHost = target

	// store session
	gameSessions[tokenStr] = c.Session

	c.Proxy.Log("joining game. ip: %s, token: %s", gameIP, tokenStr)

	// rewrite game server ip
	proxyIP := ip.Public()
	packet.SetGameIP(proxyIP)
}

func (c *RealmClient) handleCreateGame(packet Packet) {
	pb := PacketBuffer(packet)
	result := pb.Uint32(9)
	if result == 0x00 {
		// create game ok
	}
}

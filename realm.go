package d2prox

import (
	"encoding/hex"
	"fmt"
	"net"

	"github.com/johanhenriksson/d2prox/ip"
)

// RealmPort is the default realm server port
const RealmPort = 6113

var realmTargets = make(map[string]string)

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
	HandleProxySession(p, c)
}

// RealmClient implements the realm proxy client
type RealmClient struct {
	*ProxyClient
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

	// find target
	target, exists := realmTargets[token]
	if !exists {
		c.Proxy.Log("unknown token: %s", token)
		return
	}

	// clear target
	delete(realmTargets, token)

	// manually buffer packet so that it will be available on connect.
	// this packet will be silenced to avoid possible duplication
	c.BufferPacket(Packet(packet))

	// connect to realm server
	c.Proxy.Log("realm target: %s", target)
	if err := c.Connect(target); err != nil {
		c.Proxy.Log("error connecting to realm target:", target)
		c.Proxy.Log("%s", err)
		return
	}
}

//
// server -> client messages
//

// HandleServer packet
func (c *RealmClient) HandleServer(packet Packet) Packet {
	switch packet.RealmMsgID() {
	case McpJoinGame:
		join := McpJoinGamePacket(packet)
		c.handleJoinGame(join)
	}
	return packet
}

func (c *RealmClient) handleJoinGame(packet McpJoinGamePacket) {
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
	gameTargets[tokenStr] = target

	c.Proxy.Log("joining game. ip: %s, token: %s", gameIP, tokenStr)

	// rewrite game server ip
	proxyIP := ip.Public()
	packet.SetGameIP(proxyIP)
}

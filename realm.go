package d2prox

import (
	"encoding/hex"
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
			Name:     "realm",
			OnAccept: acceptRealm,
			port:     RealmPort,
		},
	}
}

func acceptRealm(server Proxy, base *ProxyClient) Client {
	return &RealmClient{
		ProxyClient: base,
	}
}

// RealmClient implements the realm proxy client
type RealmClient struct {
	*ProxyClient
}

// Connect to the target realm server
func (c *RealmClient) Connect(target string) error {
	// send 0x01 game byte on connect
	// (its removed to simplify packet handling)
	c.ProxyClient.outBuffer = append(
		[][]byte{[]byte{0x01}},
		c.ProxyClient.outBuffer...)

	return c.ProxyClient.Connect(target)
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
	ip := packet.GameIP()
	gameTargets[tokenStr] = ip

	c.Proxy.Log("joining game. ip: %s, token: %s", ip, tokenStr)

	// rewrite game server ip
	packet[9] = 127
	packet[10] = 0
	packet[11] = 0
	packet[12] = 1
}

package d2prox

import (
	"fmt"

	"github.com/johanhenriksson/d2prox/ip"
)

// BnetPort is the default battle.net port
const BnetPort = 6112

// BnetProxy is the battle.net proxy server implementation
type BnetProxy struct {
	ProxyServer
	RealmHost string
}

// NewBnet creates a new battle.net proxy
func NewBnet(hostname string) *BnetProxy {
	return &BnetProxy{
		ProxyServer: ProxyServer{
			Name:     "bnet",
			OnAccept: acceptBnet,
			port:     BnetPort,
		},
		RealmHost: hostname,
	}
}

func acceptBnet(server Proxy, base *ProxyClient) Client {
	return &BnetClient{
		ProxyClient: base,
	}
}

// BnetClient is the battle.net proxy client implementation
type BnetClient struct {
	*ProxyClient
	AccountName string
	Token       string
}

// OnConnect is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *BnetClient) OnConnect() {
	// read the game byte 0x01 and put it on the output buffer
	// this simplifies handling of the first packet
	b := []byte{0}
	c.Read(b)
	c.ProxyClient.outBuffer = [][]byte{b}
}

//
// server -> client
//

// HandleBuffered packet
func (c *BnetClient) HandleBuffered(packet Packet) Packet {
	// we know the realm server ip, so we can connect immediately
	// todo: rename OnConnect() to something better and put the call there

	bnet := c.Proxy.(*BnetProxy)

	// todo: configurable realm
	// resolve battle.net server ip using external dns (in case the hosts file is modified)
	bnetIP, err := ip.ResolveHost(bnet.RealmHost)
	if err != nil {
		c.Proxy.Log("Unable to resovle battle.net hostname '%s': %s", bnet.RealmHost, err)
	}

	// connect to battle.net server
	target := fmt.Sprintf("%s:%d", bnetIP, BnetPort)
	if err := c.Connect(target); err != nil {
		c.Proxy.Log("battle.net connect() error: %s", err)
		c.Close()
	}

	return packet
}

// HandleServer packet
func (c *BnetClient) HandleServer(packet Packet) Packet {
	switch packet.BnetMsgID() {
	case SidLogonRealmEx:
		logon := LogonRealmExPacket(packet)
		c.handleLogonRealmEx(logon)
	}
	return packet
}

func (c *BnetClient) handleLogonRealmEx(packet LogonRealmExPacket) {
	target := packet.RealmTarget()
	token := packet.Token()

	// intercept ip & port
	realmIP := ip.Public()
	packet.SetRealmIP(realmIP)
	packet.SetRealmPort(RealmPort)

	// store realm target
	realmTargets[token] = target

	c.Proxy.Log("realm logon for %s - token: %s realm: %s", c.AccountName, token[8:16], target)
}

//
// client -> server
//

// HandleClient packet
func (c *BnetClient) HandleClient(packet Packet) Packet {
	switch packet.BnetMsgID() {
	case SidLogonResponse2:
		name := string(packet[32:])
		c.Proxy.Log("account name: %s", name)
		c.AccountName = name
	}
	return packet
}

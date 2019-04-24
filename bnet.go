package d2prox

import (
	"fmt"
	"net"
	"time"

	"github.com/johanhenriksson/d2prox/ip"
)

// BnetPort is the default battle.net port
const BnetPort = 6112

// Keeps track of account names -> game sessions
var accountSessions = map[string]*GameSession{}

// BnetProxy is the battle.net proxy server implementation
type BnetProxy struct {
	ProxyServer
	RealmHost string
}

// NewBnet creates a new battle.net proxy
func NewBnet(hostname string) *BnetProxy {
	return &BnetProxy{
		ProxyServer: ProxyServer{
			Name: "bnet",
			port: BnetPort,
		},
		RealmHost: hostname,
	}
}

// Accept a new connection
func (p *BnetProxy) Accept(conn net.Conn) {
	c := &BnetClient{
		ProxyClient: &ProxyClient{
			Proxy:  p,
			client: conn,
		},
		Session: NewGameSession(),
		Bnet:    p,
	}

	// read first protocol byte to choose session handler
	protocolPacket := PacketBuffer{0}
	if _, err := c.Client().Read(protocolPacket); err != nil {
		c.Proxy.Log("error reading protocol byte")
		conn.Close()
		return
	}

	// add it back to the outgoing buffer
	c.BufferPacket(Packet(protocolPacket))

	// handle session depending on protocol type
	switch protocolPacket.Byte(0) {
	case 0x01: // Battle.net Chat
		c.Proxy.Log("bncs session")
		HandleProxySession(p, c, PacketReader(bnetPacketLength), PacketReader(bnetPacketLength))
	case 0x02: // BNFTP
		c.Proxy.Log("bnftp session")
		HandleProxySession(p, &BnetFtpClient{BnetClient: c}, StreamReader, StreamReader)
	}
}

// bnetPacketLength computes the length of the next packet in the buffer
func bnetPacketLength(buffer PacketBuffer, offset, length int) (int, error) {
	// packets should start with 0xFF
	if buffer.Byte(offset) != 0xFF {
		return 0, fmt.Errorf("Expected packet to start with 0xFF")
	}
	// return packet length
	return buffer.Uint16(offset + 2), nil
}

// BnetClient is the battle.net proxy client implementation
type BnetClient struct {
	*ProxyClient
	Bnet    *BnetProxy
	Session *GameSession
}

// OnAccept is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *BnetClient) OnAccept() {
	// todo: configurable realm
	// resolve battle.net server ip using external dns (in case the hosts file is modified)
	bnetIP, err := ip.ResolveHost(c.Bnet.RealmHost)
	if err != nil {
		c.Proxy.Log("Unable to resovle battle.net hostname '%s': %s", c.Bnet.RealmHost, err)
	}

	// connect to battle.net server
	target := fmt.Sprintf("%s:%d", bnetIP, BnetPort)
	if err := c.Connect(target); err != nil {
		c.Proxy.Log("battle.net connect() error: %s", err)
		c.Close()
	}
}

//
// server -> client
//

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

	// store realm host in session
	c.Session.RealmHost = target

	// store realm session
	realmSessions[token] = c.Session

	c.Proxy.Log("realm logon for %s - token: %s realm: %s", c.Session.AccountName, token[8:16], target)
}

//
// client -> server
//

// HandleClient packet
func (c *BnetClient) HandleClient(packet Packet) Packet {
	switch packet.BnetMsgID() {
	case SidAuthCheck:
		// extract cdkey hash
		auth := BnetAuthCheckPacket(packet)
		keyhash := auth.KeysHash()
		c.Session.KeyHash = keyhash
		c.Proxy.Log("key hash: %s", keyhash)

	case SidLogonResponse2:
		// extract account name
		name := string(packet[32:])

		// look up or create new game session
		if session, exists := accountSessions[name]; exists {
			// copy current key hash
			session.KeyHash = c.Session.KeyHash

			// update client session
			c.Session = session
			c.Proxy.Log("retrieved previous session for %s", name)

			past1h := 0
			past12h := 0
			for _, game := range c.Session.Games {
				since := time.Now().Sub(game.Start)
				if since < time.Hour {
					past1h++
				}
				if since < 12*time.Hour {
					past12h++
				}
			}
		} else {
			// store account name
			c.Session.AccountName = name

			// store session in account-session mapping
			accountSessions[name] = c.Session

			c.Proxy.Log("started new session for %s", name)
		}
	}
	return packet
}

// BnetFtpClient represents a BNFTP session. It derives from BnetClient since it shares a common
// OnAccept method, but it does not read or interfere with any packets
type BnetFtpClient struct {
	*BnetClient
}

// HandleClient nop
func (c *BnetFtpClient) HandleClient(packet Packet) Packet { return packet }

// HandleServer nop
func (c *BnetFtpClient) HandleServer(packet Packet) Packet { return packet }

// HandleBuffered nop
func (c *BnetFtpClient) HandleBuffered(packet Packet) Packet { return packet }

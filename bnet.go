package d2prox

// BnetPort is the default battle.net port
const BnetPort = 6112

// BnetProxy is the battle.net proxy server implementation
type BnetProxy struct {
	ProxyServer
}

// NewBnet creates a new battle.net proxy
func NewBnet() *BnetProxy {
	return &BnetProxy{
		ProxyServer{
			Name:     "bnet",
			OnAccept: acceptBnet,
			port:     BnetPort,
		},
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

// Connect to battle.net server
func (c *BnetClient) Connect(target string) error {
	// send 0x01 game byte on connect
	// (its removed to simplify packet handling)
	c.ProxyClient.outBuffer = append(
		[][]byte{[]byte{0x01}},
		c.ProxyClient.outBuffer...)

	return c.ProxyClient.Connect(target)
}

//
// server -> client
//

// HandleBuffered packet
func (c *BnetClient) HandleBuffered(packet Packet) Packet {
	// we always know the realm server ip, so we can connect immediately
	// todo: rename OnConnect() to something better and put the call there

	// todo: configurable battle.net server ip (this is EU)
	if err := c.Connect("5.42.181.16:6112"); err != nil {
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
	target := packet.RealmIP()
	token := packet.Token()

	// intercept ip
	packet[20] = 127
	packet[21] = 0
	packet[22] = 0
	packet[23] = 1

	// set port 6113
	packet[24] = 0x17
	packet[25] = 0xe1
	packet[26] = 0
	packet[27] = 0

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

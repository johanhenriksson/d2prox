package d2prox

const GamePort = 4000

var gameTargets = make(map[string]string)

type GameProxy struct {
	ProxyServer
}

func NewGame() *GameProxy {
	return &GameProxy{
		ProxyServer{
			Name:     "game",
			OnAccept: acceptGame,
			port:     GamePort,
		},
	}
}

func acceptGame(server Proxy, base *ProxyClient) Client {
	return &GameClient{
		ProxyClient: base,
	}
}

type GameClient struct {
	*ProxyClient
}

func (c *GameClient) OnConnect() {
	// client wont send game info until the server sends D2GS_STARTLOGON (0xAF)
	// we'll send it manually and silence it later.
	c.Write([]byte{
		GsStartLogon, 0x00,
	})
}

//
// server -> client
//

func (c *GameClient) HandleServer(packet Packet) Packet {
	/*
		fmt.Println("GS S->C")
		fmt.Println(hex.Dump(packet))
	*/
	return packet
}

//
// client -> server
//

func (c *GameClient) HandleBuffered(packet Packet) Packet {
	/*
		fmt.Println("GS C->S (B)")
		fmt.Println(hex.Dump(packet))
	*/

	switch packet.GsMsgID() {
	case GsGameLogon:
		logon := GsGameLogonPacket(packet)
		c.handleGameLogon(logon)
	}

	return packet
}

func (c *GameClient) handleGameLogon(packet GsGameLogonPacket) {
	token := packet.Token()

	target, exists := gameTargets[token]
	if !exists {
		c.Proxy.Log("game target %s not found", token)
		// todo: what to we dooo?
		c.Close()
		return
	}

	c.Proxy.Log("token %s proxied to game server %s", token, target)

	// one time use only
	delete(gameTargets, token)

	if err := c.Connect(target); err != nil {
		c.Proxy.Log("error connecting to game server: %s", err)
	}
}

func (c *GameClient) HandleClient(packet Packet) Packet {
	/*
		fmt.Println("GS C->S")
		fmt.Println(hex.Dump(packet))
	*/

	switch packet.GsMsgID() {
	case GsStartLogon:
		// silence D2GS_STARTLOGON, since we send it manually in Connect()
		return nil
	}

	return packet
}

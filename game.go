package d2prox

import (
	"encoding/hex"
	"fmt"
)

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
	// we need to send some init stuff to the client first
	c.Write([]byte{
		0xAF, 0x00,
	})
}

func (c *GameClient) HandleServer(packet Packet) Packet {
	/*
		fmt.Println("GS S->C")
		fmt.Println(hex.Dump(packet))
	*/
	return packet
}

func (c *GameClient) HandleBuffered(packet Packet) Packet {
	/*
		fmt.Println("GS C->S (B)")
		fmt.Println(hex.Dump(packet))
	*/

	if packet[0] == 0x68 {
		token := hex.EncodeToString(packet[1:7])
		fmt.Println("Game token:", token)

		target, exists := gameTargets[token]
		if !exists {
			fmt.Println("No game target found")
			return packet
		}

		// one time use only
		delete(gameTargets, token)

		if err := c.Connect(target); err != nil {
			fmt.Println("Game connect error:", err)
		}
	}

	return packet
}

func (c *GameClient) HandleClient(packet Packet) Packet {
	/*
		fmt.Println("GS C->S")
		fmt.Println(hex.Dump(packet))
	*/

	if packet[0] == 0xAF {
		c.Proxy.Log("Blocked C->S packet 0xAF")
		return nil
	}

	return packet
}

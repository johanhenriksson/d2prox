package d2prox

import (
	"encoding/hex"
	"fmt"
)

const RealmPort = 6113

var realmTargets = make(map[string]string)

type RealmProxy struct {
	ProxyServer
}

func NewRealm() *RealmProxy {
	return &RealmProxy{
		ProxyServer{
			Name:     "realm",
			OnAccept: AcceptRealm,
			port:     RealmPort,
		},
	}
}

func AcceptRealm(server Proxy, base *ProxyClient) Client {
	return &RealmClient{
		ProxyClient: base,
	}
}

type RealmClient struct {
	*ProxyClient
}

func (c *RealmClient) Connect(target string) error {
	c.ProxyClient.outBuffer = append(
		[][]byte{[]byte{0x01}},
		c.ProxyClient.outBuffer...)
	return c.ProxyClient.Connect(target)
}

func (c *RealmClient) HandleBuffered(packet []byte) []byte {
	fmt.Println("handle buffered realm packet")

	if packet[2] == 0x01 {
		// extract token
		token := hex.EncodeToString(packet[3:67])
		fmt.Println("Realm token:", token)

		// find target
		target, exists := realmTargets[token]
		if !exists {
			fmt.Println("Unknown token")
			return packet
		}

		fmt.Println("Realm target:", target)
		if err := c.Connect(target); err != nil {
			fmt.Println("error connecting to realm target:", target)
		}
	}

	return packet
}

func (c *RealmClient) HandleServer(packet []byte) []byte {

	if packet[2] == McpJoinGame {
		// dump it
		// intercept join game
		ip := fmt.Sprintf("%d.%d.%d.%d:4000", packet[9], packet[10], packet[11], packet[12])

		c.Proxy.Log("Intercepted MCP_JOINGAME. Game ip: %s", ip)
		fmt.Println(hex.Dump(packet))

		token := make([]byte, 6)
		copy(token[0:4], packet[13:17])
		copy(token[4:6], packet[5:7])

		tokenStr := hex.EncodeToString(token)
		fmt.Println("Token:", tokenStr)
		gameTargets[tokenStr] = ip

		packet[9] = 127
		packet[10] = 0
		packet[11] = 0
		packet[12] = 1
	}

	return packet
}

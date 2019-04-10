package d2prox

import (
	"encoding/hex"
	"fmt"
)

const BnetPort = 6112

type BnetProxy struct {
	ProxyServer
}

func NewBnet() *BnetProxy {
	return &BnetProxy{
		ProxyServer{
			Name:     "bnet",
			OnAccept: AcceptBnet,
			port:     BnetPort,
		},
	}
}

func AcceptBnet(server Proxy, base *ProxyClient) Client {
	return &BnetClient{
		ProxyClient: base,
	}
}

type BnetClient struct {
	*ProxyClient
}

func (c *BnetClient) Connect(target string) error {
	c.ProxyClient.outBuffer = append(
		[][]byte{[]byte{0x01}},
		c.ProxyClient.outBuffer...)
	return c.ProxyClient.Connect(target)
}

func (c *BnetClient) HandleBuffered(packet []byte) []byte {
	// europe hard coded atm
	if err := c.Connect("5.42.181.16:6112"); err != nil {
		fmt.Println("bnet connect error", err)
	}
	return packet
}

func (c *BnetClient) HandleServer(packet []byte) []byte {
	if packet[1] == SidLogonRealmEx {
		fmt.Println("SID_LOGON_REALMEX")

		token := make([]byte, 64)
		copy(token[0:16], packet[4:20])
		copy(token[16:64], packet[28:76])

		ip := fmt.Sprintf("%d.%d.%d.%d", packet[20], packet[21], packet[22], packet[23])
		fmt.Println("Realm IP:", ip)

		// intercept
		packet[20] = 127
		packet[21] = 0
		packet[22] = 0
		packet[23] = 1

		packet[24] = 0x17
		packet[25] = 0xe1
		packet[26] = 0
		packet[27] = 0

		tokenStr := hex.EncodeToString(token)

		// store realm target
		target := fmt.Sprintf("%s:6112", ip)
		realmTargets[tokenStr] = target

		fmt.Println("MCP Token:", tokenStr)
	}

	return packet
}

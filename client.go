package d2prox

import (
	"fmt"
	"net"
)

type Client interface {
	HandleBuffered([]byte) []byte
	HandleClient([]byte) []byte
	HandleServer([]byte) []byte

	Connect(target string) error
	Close()
	Connected() bool

	OnConnect()
}

type ProxyClient struct {
	net.Conn
	Proxy         Proxy
	outBuffer     [][]byte
	server        net.Conn
	errors        chan error
	clientPackets PacketStream
	serverPackets PacketStream
}

func (c *ProxyClient) Connected() bool {
	return c.server != nil
}

func (c *ProxyClient) Connect(target string) error {
	if c.Connected() {
		return nil
	}

	c.Proxy.Log("connecting to %s", target)

	conn, err := net.Dial("tcp", target)
	if err != nil {
		return fmt.Errorf("Error connecting to realm server: %s", err)
	}

	c.server = conn
	c.serverPackets = StreamReader("realm server", conn, c.errors)

	c.Proxy.Log("outbound socket connected to %s", target)

	for _, packet := range c.outBuffer {
		// send buffered packets
		conn.Write(packet)
	}
	c.outBuffer = nil

	return nil
}

func (c *ProxyClient) OnConnect() {
	// nop
}

func (c *ProxyClient) Close() {

}

func (c *ProxyClient) HandleBuffered(packet []byte) []byte {
	return packet
}

func (c *ProxyClient) HandleClient(packet []byte) []byte {
	return packet
}

func (c *ProxyClient) HandleServer(packet []byte) []byte {
	return packet
}

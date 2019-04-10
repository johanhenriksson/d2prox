package d2prox

import (
	"fmt"
	"net"
)

type Client interface {
	HandleBuffered(Packet) Packet
	HandleClient(Packet) Packet
	HandleServer(Packet) Packet

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
		return fmt.Errorf("Error connecting to target: %s", err)
	}

	c.server = conn
	c.serverPackets = StreamReader(conn, c.errors)

	c.Proxy.Log("outbound socket connected to %s", target)

	// send buffered packets
	for _, packet := range c.outBuffer {
		// send buffered packets
		if _, err := conn.Write(packet); err != nil {
			return err
		}
	}
	c.outBuffer = nil

	return nil
}

func (c *ProxyClient) OnConnect() {
	// nop
}

func (c *ProxyClient) Close() {
	c.Conn.Close()
	c.server.Close()
	c.server = nil
	c.Proxy.Log("connection closed")
}

func (c *ProxyClient) HandleBuffered(packet Packet) Packet { return packet }
func (c *ProxyClient) HandleClient(packet Packet) Packet   { return packet }
func (c *ProxyClient) HandleServer(packet Packet) Packet   { return packet }

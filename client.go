package d2prox

import (
	"fmt"
	"net"
)

// Client describes a proxy session implementation
type Client interface {
	Client() net.Conn
	Server() net.Conn

	Close()
	Connect(target string) error
	WriteClient(Packet) error
	WriteServer(Packet) error

	BufferPacket(Packet)
	HandleBuffered(Packet) Packet
	HandleClient(Packet) Packet
	HandleServer(Packet) Packet

	OnAccept()
	OnConnect()
}

// ProxyClient represents a generic proxy session
type ProxyClient struct {
	Proxy     Proxy
	outBuffer []Packet
	client    net.Conn
	server    net.Conn
}

func (c *ProxyClient) Client() net.Conn { return c.client }
func (c *ProxyClient) Server() net.Conn { return c.server }

func (c *ProxyClient) WriteClient(p Packet) error {
	_, err := c.client.Write(p)
	return err
}

func (c *ProxyClient) WriteServer(p Packet) error {
	_, err := c.server.Write(p)
	return err
}

func (c *ProxyClient) BufferPacket(p Packet) {
	c.outBuffer = append(c.outBuffer, p)
}

// Connect to a remote server. Will automatically send buffered packets once connected.
func (c *ProxyClient) Connect(target string) error {
	if c.server != nil {
		return nil
	}

	conn, err := net.Dial("tcp", target)
	if err != nil {
		return fmt.Errorf("error connecting to target: %s", err)
	}

	c.server = conn
	c.Proxy.Log("outbound socket connected to %s", target)

	// send buffered packets
	for _, packet := range c.outBuffer {
		// send buffered packets
		fmt.Println("write buffered packet len", len(packet))
		if _, err := conn.Write(packet); err != nil {
			return err
		}
	}
	c.outBuffer = nil

	return nil
}

// OnAccept is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *ProxyClient) OnAccept() {}

// OnConnect is fired immediately after a connection to the remote server is established
// Should only be called by the proxy session handler
func (c *ProxyClient) OnConnect() {}

// Close the proxy session
func (c *ProxyClient) Close() {
	if c.client != nil {
		c.client.Close()
		c.client = nil
	}
	if c.server != nil {
		c.server.Close()
		c.server = nil
	}
	c.Proxy.Log("connection closed")
}

//
// nop packet handlers
//

// HandleBuffered packets
func (c *ProxyClient) HandleBuffered(packet Packet) Packet { return packet }

// HandleClient packets
func (c *ProxyClient) HandleClient(packet Packet) Packet { return packet }

// HandleServer packets
func (c *ProxyClient) HandleServer(packet Packet) Packet { return packet }

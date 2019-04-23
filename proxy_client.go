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
	OnClose()
}

// ProxyClient represents a generic proxy session
type ProxyClient struct {
	Proxy     Proxy
	outBuffer []Packet
	client    net.Conn
	server    net.Conn
}

// Client returns a reference to the client socket
func (c *ProxyClient) Client() net.Conn { return c.client }

// Server returns a reference to the server socket
func (c *ProxyClient) Server() net.Conn { return c.server }

// WriteClient sends a packet to the client
func (c *ProxyClient) WriteClient(p Packet) error {
	_, err := c.client.Write(p)
	return err
}

// WriteServer sends a packet to the server
func (c *ProxyClient) WriteServer(p Packet) error {
	_, err := c.server.Write(p)
	return err
}

// BufferPacket appends a packet to the output buffer, to be sent when connected to the server
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
		if _, err := conn.Write(packet); err != nil {
			return err
		}
	}
	c.outBuffer = nil

	return nil
}

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
	c.OnClose()
	c.Proxy.Log("connection closed")
}

//
// event handlers
//

// OnAccept is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *ProxyClient) OnAccept() {}

// OnConnect is fired immediately after a connection to the remote server is established
// Should only be called by the proxy session handler
func (c *ProxyClient) OnConnect() {}

// OnClose is fired immediately after a session is disconnected from the server or client
func (c *ProxyClient) OnClose() {}

// HandleBuffered packets
func (c *ProxyClient) HandleBuffered(packet Packet) Packet { return packet }

// HandleClient packets
func (c *ProxyClient) HandleClient(packet Packet) Packet { return packet }

// HandleServer packets
func (c *ProxyClient) HandleServer(packet Packet) Packet { return packet }

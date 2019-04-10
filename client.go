package d2prox

import (
	"fmt"
	"net"
)

// Client describes a proxy session implementation
type Client interface {
	HandleBuffered(Packet) Packet
	HandleClient(Packet) Packet
	HandleServer(Packet) Packet

	Connect(target string) error
	Close()
	Connected() bool

	OnConnect()
}

// ProxyClient represents a generic proxy session
type ProxyClient struct {
	net.Conn
	Proxy         Proxy
	outBuffer     [][]byte
	server        net.Conn
	errors        chan error
	clientPackets PacketStream
	serverPackets PacketStream
}

// Connected returns true if the proxy session is connected to the remote server
func (c *ProxyClient) Connected() bool {
	return c.server != nil
}

// Connect to a remote server. Will automatically send buffered packets once connected.
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

// OnConnect is fired immediately after a client connects to the proxy
// Should only be called by the server Accept() function
func (c *ProxyClient) OnConnect() {}

// Close the proxy session
func (c *ProxyClient) Close() {
	c.Conn.Close()
	c.server.Close()
	c.server = nil
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

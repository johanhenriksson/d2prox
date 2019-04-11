package d2prox

import (
	"fmt"
	"net"
	"time"
)

// Proxy describes the base proxy server implementation
type Proxy interface {
	Port() int
	Log(string, ...interface{})
	Accept(net.Conn)
}

// AcceptHandler is a function that instantiates a custom proxy session
type AcceptHandler func(Proxy, *ProxyClient) Client

// ProxyServer is the base proxy implementation
type ProxyServer struct {
	Name     string
	OnAccept AcceptHandler
	port     int
}

// Serve a proxy
func Serve(p Proxy) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", p.Port()))
	if err != nil {
		p.Log("failed to listen on port %d", p.Port())
		return
	}
	defer listener.Close()

	p.Log("listening on port %d", p.Port())

	for {
		in, err := listener.Accept()
		if err != nil {
			p.Log("failed to accept socket")
			return
		}
		p.Log("accepted client")
		go p.Accept(in)
	}
}

// Port returns the listening port
func (p *ProxyServer) Port() int {
	return p.port
}

// Log a message
func (p *ProxyServer) Log(format string, args ...interface{}) {
	fmt.Printf("%s | %s\n", p.Name, fmt.Sprintf(format, args...))
}

// Accept and handle a proxy session. Should be called as a goroutine
func (p *ProxyServer) Accept(conn net.Conn) {
	c := &ProxyClient{
		Proxy:  p,
		client: conn,
	}
	HandleProxySession(p, c)
}

func HandleProxySession(p Proxy, c Client) {
	defer c.Close()

	// fire client connected event
	c.OnAccept()

	errs := make(chan error)
	clientPackets := StreamReader(c.Client(), errs)

	// while not connected to server, buffer up client messages
	for c.Server() == nil {
		select {
		// abort on errors
		case err := <-errs:
			if err != nil {
				p.Log("Read error: %s", err)
			}
			return

		// receive client -> server packets
		case packet, more := <-clientPackets:
			if !more {
				return
			}
			packet = c.HandleBuffered(packet)
			if packet != nil {
				c.BufferPacket(packet)
			}

		case <-time.After(100 * time.Millisecond):
			// periodic timeout to check if we've connected to the server
		}
	}

	// fire server connected event
	c.OnConnect()

	serverPackets := StreamReader(c.Server(), errs)
	for {
		select {
		// abort on errors
		case err := <-errs:
			if err != nil {
				p.Log("Read error: %s", err)
			}
			return

		// receive client -> server packets
		case packet, more := <-clientPackets:
			if !more {
				return
			}

			packet = c.HandleClient(packet)
			if packet == nil {
				continue // skip silenced packets
			}

			// forward to remote server
			if err := c.WriteServer(packet); err != nil {
				p.Log("server: packet write failed")
				return
			}

		// receive server -> client packets
		case packet, more := <-serverPackets:
			if !more {
				return
			}

			packet = c.HandleServer(packet)
			if packet == nil {
				continue // skip silenced packets
			}

			// forward to client
			if err := c.WriteClient(packet); err != nil {
				p.Log("client: packet write failed")
				return
			}
		}
	}
}

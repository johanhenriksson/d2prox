package d2prox

import (
	"fmt"
	"net"
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
	errs := make(chan error)
	clientPackets := StreamReader(conn, errs)

	base := &ProxyClient{
		Proxy:         p,
		Conn:          conn,
		errors:        errs,
		clientPackets: clientPackets,
	}

	c := p.OnAccept(p, base)

	c.OnConnect()

	defer c.Close()

	for {
		select {
		// abort on errors
		case err := <-errs:
			if err != nil {
				p.Log("Read error: %s", err)
			}
			return

		// receive client -> server packets
		case packet, more := <-base.clientPackets:
			if !more {
				return
			}

			if !c.Connected() {
				// store packets in a buffer until the remote is connected
				// handle them separately. buffered packets cannot be easily silenced
				base.outBuffer = append(base.outBuffer, packet)
				c.HandleBuffered(packet)
			} else {
				packet = c.HandleClient(packet)
				if packet == nil {
					continue // skip silenced packets
				}

				// forward to remote server
				if _, err := base.server.Write(packet); err != nil {
					p.Log("server: packet write failed")
					return
				}
			}

		// receive server -> client packets
		case packet, more := <-base.serverPackets:
			if !more {
				return
			}

			packet = c.HandleServer(packet)
			if packet == nil {
				continue // skip silenced packets
			}

			// forward to client
			if _, err := base.Write(packet); err != nil {
				p.Log("client: packet write failed")
				return
			}
		}
	}
}

package d2prox

import (
	"fmt"
	"net"
)

type Proxy interface {
	Port() int
	Log(string, ...interface{})
	Accept(net.Conn)
}

type AcceptHandler func(Proxy, *ProxyClient) Client

type ProxyServer struct {
	Name     string
	OnAccept AcceptHandler
	port     int
}

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

func (p *ProxyServer) Port() int {
	return p.port
}

func (p *ProxyServer) Log(format string, args ...interface{}) {
	fmt.Printf("%s | %s\n", p.Name, fmt.Sprintf(format, args...))
}

func (p *ProxyServer) Accept(conn net.Conn) {
	errs := make(chan error)
	clientPackets := StreamReader("realm client", conn, errs)

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
		case err := <-errs:
			if err == nil {
				p.Log("socket closed")
			} else {
				p.Log("read error: %s", err)
			}
			return

		case packet := <-base.clientPackets:
			// skip 0x01 game byte
			if packet[0] == 0x01 {
				if len(packet) == 1 {
					continue
				}
				packet = packet[1:]
			}

			if !c.Connected() {
				base.outBuffer = append(base.outBuffer, packet)
				c.HandleBuffered(packet)
			} else {
				packet = c.HandleClient(packet)
				if packet == nil {
					continue // skip silenced packets
				}

				if _, err := base.server.Write(packet); err != nil {
					p.Log("server: packet write failed")
					return
				}
			}

		case packet := <-base.serverPackets:
			packet = c.HandleServer(packet)
			if packet == nil {
				continue // skip silenced packets
			}

			if _, err := base.Write(packet); err != nil {
				p.Log("client: packet write failed")
				return
			}
		}
	}
}

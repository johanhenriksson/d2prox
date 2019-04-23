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
	fmt.Printf("%-6v| %s\n", p.Name, fmt.Sprintf(format, args...))
}

// Accept and handle a proxy session. Should be called as a goroutine
func (p *ProxyServer) Accept(conn net.Conn) {
	c := &ProxyClient{
		Proxy:  p,
		client: conn,
	}
	HandleProxySession(p, c, StreamReader, StreamReader)
}

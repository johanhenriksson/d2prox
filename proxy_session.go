package d2prox

import (
	"io"
	"time"
)

// HandleProxySession is the main
func HandleProxySession(p Proxy, c Client, clientReader, serverReader PacketStreamReader) {
	defer c.Close()

	// fire client connected event
	c.OnAccept()

	errs := make(chan error)
	clientPackets := clientReader(c.Client(), errs)

	// while not connected to server, buffer up client messages
	for c.Server() == nil {
		select {
		// abort on errors
		case err := <-errs:
			if err != io.EOF {
				// dump non-eof errors
				p.Log("read error: %s", err)
			}
			return

		// receive client -> server packets
		case packet := <-clientPackets:
			packet = c.HandleBuffered(packet)
			if packet == nil {
				continue // skip silenced packets
			}
			c.BufferPacket(packet)

		case <-time.After(100 * time.Millisecond):
			// periodic timeout to check if we've connected to the server
		}
	}

	// fire server connected event
	serverPackets := serverReader(c.Server(), errs)
	c.OnConnect()

	for {
		select {
		// abort on errors
		case err := <-errs:
			if err != io.EOF {
				// dump non-eof errors
				p.Log("read error: %s", err)
			}
			return

		// receive client -> server packets
		case packet := <-clientPackets:
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
		case packet := <-serverPackets:
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

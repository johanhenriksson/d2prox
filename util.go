package d2prox

import (
	"net"
)

// BufferSize is the default receive buffer size
const BufferSize = 1024

// PacketStream is a channel for packets
type PacketStream chan Packet

//type PacketStreamReader func(net.Conn, chan error) PacketStream

// StreamReader reads packets in a separate goroutine and writes them to a packet channel
func StreamReader(sck net.Conn, errs chan error) PacketStream {
	buffer := make([]byte, BufferSize)
	stream := make(PacketStream)

	// ideally, this method should be replaced with 3 separate variations for each proxy type.
	// this would allow it to use knowledge about the protocols to properly merge/split packets
	// before passing them on to the handlers.
	// to implement, a StreamReader function could be passed as an argument to HandleProxySession

	go func() {
		defer close(stream)

		for {
			len, err := sck.Read(buffer)
			if err != nil {
				errs <- err
				return
			}

			// does this ever happen?
			if len == 0 {
				errs <- nil
				return
			}

			packet := make([]byte, len)
			copy(packet, buffer[:len])
			stream <- Packet(packet)
		}
	}()

	return stream
}

package d2prox

import (
	"net"
)

// BufferSize is the default receive buffer size
const BufferSize = 1024

// PacketStream is a channel for packets
type PacketStream chan Packet

// StreamReader reads packets in a separate goroutine and writes them to a packet channel
func StreamReader(sck net.Conn, errs chan error) PacketStream {
	buffer := make([]byte, BufferSize)
	stream := make(PacketStream)

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

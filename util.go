package d2prox

import (
	"net"
)

const BufferSize = 1024

type PacketStream chan Packet

func StreamReader(sck net.Conn, errs chan error) PacketStream {
	buffer := make([]byte, 1024)
	stream := make(PacketStream)
	go func() {
		defer close(stream)

		for {
			len, err := sck.Read(buffer)
			if err != nil {
				errs <- err
				return
			}
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

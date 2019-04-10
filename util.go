package d2prox

import (
	"fmt"
	"net"
)

const BufferSize = 1024

type PacketTransform func([]byte) []byte
type PacketStream chan []byte

func StreamReader(name string, sck net.Conn, errs chan error) PacketStream {
	buffer := make([]byte, 1024)
	stream := make(PacketStream)
	go func() {
		for {
			len, err := sck.Read(buffer)
			if err != nil {
				fmt.Println(name, "read error", err)
				errs <- err
				return
			}
			if len == 0 {
				fmt.Println(name, "closed")
				errs <- nil
				return
			}

			packet := make([]byte, len)
			copy(packet, buffer[:len])
			stream <- packet
		}
	}()
	return stream
}

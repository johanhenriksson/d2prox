package d2prox

import (
	"net"
)

// BufferSize is the default receive buffer size
const BufferSize = 4 * 1024 // 4k

// PacketStream is a channel for packets
type PacketStream chan Packet

type PacketStreamReader func(net.Conn, chan error) PacketStream

type PacketLengthFunc func(PacketBuffer, int, int) (int, error)

// PacketReader is a base for separating packets 
func PacketReader(lengthFunc PacketLengthFunc) PacketStreamReader {
	return func(sck net.Conn, errs chan error) PacketStream {
		buffer := make(PacketBuffer, BufferSize)
		stream := make(PacketStream)

		go func() {
			defer close(stream)

			remain := 0
			for {
				len, err := sck.Read(buffer[remain:])
				if err != nil {
					errs <- err
					return
				}

				offset := 0
				for offset < len {
					plen, err := lengthFunc(buffer, offset, len)
					if err != nil {
						errs <- err
						return
					}

					// check if we have all of it
					if offset+plen > len {
						break // we need more data
					}

					// extract and send packet
					packet := buffer.Extract(offset, plen)
					stream <- packet

					// move offset forward
					offset += plen
				}

				// keep remaining data
				remain = len - offset
				copy(buffer[:remain], buffer[offset:len])
			}
		}()

		return stream
	}
}

// StreamReader is the simplest possible implementation of a packet reader.
// Packets are written to the stream as they are read from the socket without inspection.
var StreamReader = PacketReader(bufferLengthFunc)

func bufferLengthFunc(buffer PacketBuffer, offset, length int) (int, error) {
	return length, nil // length of the packet is simply all the read bytes
}

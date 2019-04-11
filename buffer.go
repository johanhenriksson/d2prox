package d2prox

import "encoding/binary"

// PacketBuffer aids in reading data from binary buffers
type PacketBuffer []byte

// Byte returns the byte at the given offset
func (pb PacketBuffer) Byte(offset int) int {
	return int(pb[offset])
}

// Uint16 returns the Uint16 at the given offset
func (pb PacketBuffer) Uint16(offset int) int {
	return int(binary.LittleEndian.Uint16(pb[offset : offset+2]))
}

// Uint32 returns the Uint32 at the given offset
func (pb PacketBuffer) Uint32(offset int) int {
	return int(binary.LittleEndian.Uint32(pb[offset : offset+4]))
}

// Extract a part of the buffer as a Packet
func (pb PacketBuffer) Extract(offset, length int) Packet {
	packet := make(Packet, length)
	copy(packet, pb[offset:offset+length])
	return packet
}

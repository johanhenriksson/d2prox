package d2prox

import "encoding/binary"

// PacketBuffer aids in reading data from binary buffers
type PacketBuffer []byte

// Put writes an array of bytes into the buffer, starting at the given offset
func (pb PacketBuffer) Put(offset int, value []byte) {
	copy(pb[offset:offset+len(value)], value)
}

// Byte returns the byte at the given offset
func (pb PacketBuffer) Byte(offset int) int {
	return int(pb[offset])
}

// PutByte writes a byte at the given offset
func (pb PacketBuffer) PutByte(offset, value int) {
	pb[offset] = byte(value)
}

// Uint16 returns the Uint16 at the given offset
func (pb PacketBuffer) Uint16(offset int) int {
	return int(binary.LittleEndian.Uint16(pb[offset : offset+2]))
}

// PutUint16 writes an Uin16 at the given offset
func (pb PacketBuffer) PutUint16(offset, value int) {
	binary.LittleEndian.PutUint16(pb[offset:offset+4], uint16(value))
}

// Uint32 returns the Uint32 at the given offset
func (pb PacketBuffer) Uint32(offset int) int {
	return int(binary.LittleEndian.Uint32(pb[offset : offset+4]))
}

// PutUint32 writes an Uint32 at the given offset
func (pb PacketBuffer) PutUint32(offset, value int) {
	binary.LittleEndian.PutUint32(pb[offset:offset+4], uint32(value))
}

// Extract a part of the buffer as a Packet
func (pb PacketBuffer) Extract(offset, length int) Packet {
	packet := make(Packet, length)
	copy(packet, pb[offset:offset+length])
	return packet
}

// IndexOf returns the next index of a given byte in the buffer, starting from startIndex.
func (pb PacketBuffer) IndexOf(val byte, startIndex int) int {
	for i := startIndex; i < len(pb); i++ {
		if pb[i] == val {
			return i
		}
	}
	return -1
}

// NullString reads a null-terminated string starting at the given offset
func (pb PacketBuffer) NullString(offset int) string {
	end := pb.IndexOf(0x00, offset)
	if end <= offset {
		return ""
	}
	return string(pb[offset:end])
}

package d2prox

// BitField is a tool for reading individual bits from a byte slice
type BitField struct {
	Offset int
	data   []byte
}

// NewBitField creates a new BitField from a byte slice
func NewBitField(data []byte) *BitField {
	return &BitField{
		Offset: 0,
		data:   data,
	}
}

// Length returns the total bit length
func (bs *BitField) Length() int {
	return 8 * len(bs.data)
}

// Skip a number of bits
func (bs *BitField) Skip(n int) {
	bs.Offset += n
}

// BitAt returns the bit at a given offset
func (bs *BitField) BitAt(pos int) int {
	c := int(bs.data[pos>>3])
	bitmask := 1 << uint(pos&7)
	return (c & bitmask) >> uint(pos&7)
}

// BitsAt returns a number of bits starting at a given offset
func (bs *BitField) BitsAt(pos, count int) int {
	bits := 0
	for i := 0; i < count; i++ {
		bits |= bs.BitAt(pos+i) << uint(i)
	}
	return bits
}

// Bit returns the next bit and advances the offset by 1
func (bs *BitField) Bit() int {
	v := bs.BitAt(bs.Offset)
	bs.Offset++
	return v
}

// Bits returns N bits and advances the offset by N
func (bs *BitField) Bits(n int) int {
	v := bs.BitsAt(bs.Offset, n)
	bs.Offset += n
	return v
}

// Byte returns the next 8 bits and advances the offset by 8
func (bs *BitField) Byte() int {
	return bs.Bits(8)
}

// Bool returns the next bit casted to boolean, and advances the offset by 1
func (bs *BitField) Bool() bool {
	return bs.Bit() > 0
}

package d2prox

type BitField struct {
	Offset int
	data   []byte
}

func NewBitField(data []byte) *BitField {
	return &BitField{
		Offset: 0,
		data:   data,
	}
}

func (bs *BitField) Skip(n int) {
	bs.Offset += n
}

func (bs *BitField) BitAt(pos int) int {
	c := uint(bs.data[pos>>3])
	bitmask := uint(1 << (uint(pos) & 7))
	if c&bitmask > 0 {
		return 1
	}
	return 0
}

func (bs *BitField) BitsAt(pos, count int) int {
	i := uint(0)
	bits := uint(0)
	for currentbit := pos; currentbit < pos+count; currentbit++ {
		bits = bits | uint(bs.BitAt(currentbit)<<i)
		i++
	}
	return int(bits)
}

func (bs *BitField) Bit() int {
	v := bs.BitAt(bs.Offset)
	bs.Offset++
	return v
}

func (bs *BitField) Bits(n int) int {
	v := bs.BitsAt(bs.Offset, n)
	bs.Offset += n
	return v
}

func (bs *BitField) Byte() int {
	return bs.Bits(8)
}

func (bs *BitField) Bool() bool {
	return bs.Bit() > 0
}

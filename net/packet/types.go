package packet

import (
	"io"
	"math"
)

type Field interface {
	FieldEncoder
	FieldDecoder
}

type FieldEncoder interface {
	Encode() []byte
}

type FieldDecoder interface {
	Decode(r io.ByteReader) error
}

type (
	//Boolean of True is encoded as 0x01, false as 0x00.
	Boolean bool
	//Byte is signed 8-bit integer, two's complement
	Byte int8
	//UnsignedByte is unsigned 8-bit integer
	UnsignedByte uint8
	//Short is signed 16-bit integer, two's complement
	Short int16
	//UnsignedShort is unsigned 16-bit integer
	UnsignedShort uint16
	//Int is signed 32-bit integer, two's complement
	Int int32
	//Long is signed 64-bit integer, two's complement
	Long int64
	//A Float is a single-precision 32-bit IEEE 754 floating point number
	Float float32
	//A Double is a double-precision 64-bit IEEE 754 floating point number
	Double float64
	//String is sequence of Unicode scalar values
	String string

	//Chat is encoded as a String with max length of 32767.
	Chat = String
	//Identifier is encoded as a String with max length of 32767.
	Identifier = String

	//VarInt is variable-length data encoding a two's complement signed 32-bit integer
	VarInt int32
	//VarLong is variable-length data encoding a two's complement signed 64-bit integer
	VarLong int64

	//Position x as a 26-bit integer, followed by y as a 12-bit integer, followed by z as a 26-bit integer (all signed, two's complement)
	Position struct {
		X, Y, Z int
	}

	//Angle is rotation angle in steps of 1/256 of a full turn
	Angle int8

	//UUID encoded as an unsigned 128-bit integer
	UUID [16]byte
)

//ReadNBytes read N bytes from bytes.Reader
func ReadNBytes(r io.ByteReader, n int) (bs []byte, err error) {
	bs = make([]byte, n)
	for i := 0; i < n; i++ {
		bs[i], err = r.ReadByte()
		if err != nil {
			return
		}
	}
	return
}

//Encode a Boolean
func (b Boolean) Encode() []byte {
	if b {
		return []byte{0x01}
	}
	return []byte{0x00}
}

//Decode a Boolean
func (b *Boolean) Decode(r io.ByteReader) error {
	v, err := r.ReadByte()
	if err != nil {
		return err
	}

	*b = Boolean(v != 0)
	return nil
}

// Encode a String
func (s String) Encode() (p []byte) {
	byteString := []byte(s)
	p = append(p, VarInt(len(byteString)).Encode()...) //len
	p = append(p, byteString...)                       //data
	return
}

//Decode a String
func (s *String) Decode(r io.ByteReader) error {
	var l VarInt //String length
	if err := l.Decode(r); err != nil {
		return err
	}

	bs, err := ReadNBytes(r, int(l))
	if err != nil {
		return err
	}

	*s = String(bs)
	return nil
}

//Encode a Byte
func (b Byte) Encode() []byte {
	return []byte{byte(b)}
}

//Decode a Byte
func (b *Byte) Decode(r io.ByteReader) error {
	v, err := r.ReadByte()
	if err != nil {
		return err
	}
	*b = Byte(v)
	return nil
}

//Encode a UnsignedByte
func (ub UnsignedByte) Encode() []byte {
	return []byte{byte(ub)}
}

//Decode a UnsignedByte
func (ub *UnsignedByte) Decode(r io.ByteReader) error {
	v, err := r.ReadByte()
	if err != nil {
		return err
	}
	*ub = UnsignedByte(v)
	return nil
}

// Encode a Signed Short
func (s Short) Encode() []byte {
	n := uint16(s)
	return []byte{
		byte(n >> 8),
		byte(n),
	}
}

//Decode a Short
func (s *Short) Decode(r io.ByteReader) error {
	bs, err := ReadNBytes(r, 2)
	if err != nil {
		return err
	}

	*s = Short(int16(bs[0])<<8 | int16(bs[1]))
	return nil
}

// Encode a Unsigned Short
func (us UnsignedShort) Encode() []byte {
	n := uint16(us)
	return []byte{
		byte(n >> 8),
		byte(n),
	}
}

//Decode a UnsignedShort
func (us *UnsignedShort) Decode(r io.ByteReader) error {
	bs, err := ReadNBytes(r, 2)
	if err != nil {
		return err
	}

	*us = UnsignedShort(int16(bs[0])<<8 | int16(bs[1]))
	return nil
}

// Encode a Int
func (i Int) Encode() []byte {
	n := uint32(i)
	return []byte{
		byte(n >> 24), byte(n >> 16),
		byte(n >> 8), byte(n),
	}
}

//Decode a Int
func (i *Int) Decode(r io.ByteReader) error {
	bs, err := ReadNBytes(r, 4)
	if err != nil {
		return err
	}

	*i = Int(int32(bs[0])<<24 | int32(bs[1])<<16 | int32(bs[2])<<8 | int32(bs[3]))
	return nil
}

// Encode a Long
func (l Long) Encode() []byte {
	n := uint64(l)
	return []byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n),
	}
}

//Decode a Long
func (l *Long) Decode(r io.ByteReader) error {
	bs, err := ReadNBytes(r, 8)
	if err != nil {
		return err
	}

	*l = Long(int64(bs[0])<<56 | int64(bs[1])<<48 | int64(bs[2])<<40 | int64(bs[3])<<32 |
		int64(bs[4])<<24 | int64(bs[5])<<16 | int64(bs[6])<<8 | int64(bs[7]))
	return nil
}

//Encode a VarInt
func (v VarInt) Encode() (vi []byte) {
	num := uint32(v)
	for {
		b := num & 0x7F
		num >>= 7
		if num != 0 {
			b |= 0x80
		}
		vi = append(vi, byte(b))
		if num == 0 {
			break
		}
	}
	return
}

//Decode a VarInt
func (v *VarInt) Decode(r io.ByteReader) error {
	var n uint32
	for i := 0; i < 5; i++ { //读数据前的长度标记
		sec, err := r.ReadByte()
		if err != nil {
			return err
		}

		n |= (uint32(sec&0x7F) << uint32(7*i))

		if sec&0x80 == 0 {
			break
		}
	}

	*v = VarInt(n)
	return nil
}

//Encode a Position
func (p Position) Encode() []byte {
	b := make([]byte, 8)
	position := (uint64(p.X&0x3FFFFFF)<<38 | uint64((p.Z&0x3FFFFFF)<<12) | uint64(p.Y&0xFFF))
	for i := 7; i >= 0; i-- {
		b[i] = byte(position)
		position >>= 8
	}
	return b
}

// Decode a Position
func (p *Position) Decode(r io.ByteReader) error {
	var v Long
	if err := v.Decode(r); err != nil {
		return err
	}

	x := int(v >> 38)
	y := int(v & 0xFFF)
	z := int(v << 26 >> 38)

	//处理负数
	if x >= 1<<25 {
		x -= 1 << 26
	}
	if y >= 1<<11 {
		y -= 1 << 12
	}
	if z >= 1<<25 {
		z -= 1 << 26
	}

	p.X, p.Y, p.Z = x, y, z
	return nil
}

//Encode a Float
func (f Float) Encode() []byte {
	return Int(math.Float32bits(float32(f))).Encode()
}

// Decode a Float
func (f *Float) Decode(r io.ByteReader) error {
	var v Int
	if err := v.Decode(r); err != nil {
		return err
	}

	*f = Float(math.Float32frombits(uint32(v)))
	return nil
}

//Encode a Double
func (d Double) Encode() []byte {
	return Long(math.Float64bits(float64(d))).Encode()
}

// Decode a Double
func (d *Double) Decode(r io.ByteReader) error {
	var v Long
	if err := v.Decode(r); err != nil {
		return err
	}

	*d = Double(math.Float64frombits(uint64(v)))
	return nil
}

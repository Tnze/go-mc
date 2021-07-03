package packet

import (
	"bytes"
	"errors"
	"io"
	"math"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/nbt"
)

// A Field is both FieldEncoder and FieldDecoder
type Field interface {
	FieldEncoder
	FieldDecoder
}

// A FieldEncoder can be encode as minecraft protocol used.
type FieldEncoder io.WriterTo

// A FieldDecoder can Decode from minecraft protocol
type FieldDecoder io.ReaderFrom

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
	// Deprecated: Use chat.Message
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
	Angle Byte

	//UUID encoded as an unsigned 128-bit integer
	UUID uuid.UUID

	//ByteArray is []byte with prefix VarInt as length
	ByteArray []byte

	//PluginMessageData is only used in LoginPlugin,and it will read all left bytes
	PluginMessageData []byte
)

const MaxVarIntLen = 5
const MaxVarLongLen = 10

//Encode a Boolean
func (b Boolean) WriteTo(w io.Writer) (int64, error) {
	var v byte
	if b {
		v = 0x01
	} else {
		v = 0x00
	}
	nn, err := w.Write([]byte{v})
	return int64(nn), err
}

//Decode a Boolean
func (b *Boolean) ReadFrom(r io.Reader) (n int64, err error) {
	v, err := readByte(r)
	if err != nil {
		return 1, err
	}

	*b = v != 0
	return 1, nil
}

// Encode a String
func (s String) WriteTo(w io.Writer) (int64, error) {
	byteStr := []byte(s)
	n1, err := VarInt(len(byteStr)).WriteTo(w)
	if err != nil {
		return n1, err
	}
	n2, err := w.Write(byteStr)
	return n1 + int64(n2), err
}

//Decode a String
func (s *String) ReadFrom(r io.Reader) (n int64, err error) {
	var l VarInt //String length

	nn, err := l.ReadFrom(r)
	if err != nil {
		return nn, err
	}
	n += nn

	bs := make([]byte, l)
	if _, err := io.ReadFull(r, bs); err != nil {
		return n, err
	}
	n += int64(l)

	*s = String(bs)
	return n, nil
}

// readByte read one byte from io.Reader
func readByte(r io.Reader) (byte, error) {
	if r, ok := r.(io.ByteReader); ok {
		return r.ReadByte()
	}
	var v [1]byte
	_, err := io.ReadFull(r, v[:])
	return v[0], err
}

//Encode a Byte
func (b Byte) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write([]byte{byte(b)})
	return int64(nn), err
}

//Decode a Byte
func (b *Byte) ReadFrom(r io.Reader) (n int64, err error) {
	v, err := readByte(r)
	if err != nil {
		return 0, err
	}
	*b = Byte(v)
	return 1, nil
}

//Encode a UnsignedByte
func (u UnsignedByte) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write([]byte{byte(u)})
	return int64(nn), err
}

//Decode a UnsignedByte
func (u *UnsignedByte) ReadFrom(r io.Reader) (n int64, err error) {
	v, err := readByte(r)
	if err != nil {
		return 0, err
	}
	*u = UnsignedByte(v)
	return 1, nil
}

// Encode a Signed Short
func (s Short) WriteTo(w io.Writer) (int64, error) {
	n := uint16(s)
	nn, err := w.Write([]byte{byte(n >> 8), byte(n)})
	return int64(nn), err
}

//Decode a Short
func (s *Short) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [2]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*s = Short(int16(bs[0])<<8 | int16(bs[1]))
	return
}

// Encode a Unsigned Short
func (us UnsignedShort) WriteTo(w io.Writer) (int64, error) {
	n := uint16(us)
	nn, err := w.Write([]byte{byte(n >> 8), byte(n)})
	return int64(nn), err
}

//Decode a UnsignedShort
func (us *UnsignedShort) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [2]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*us = UnsignedShort(int16(bs[0])<<8 | int16(bs[1]))
	return
}

// Encode a Int
func (i Int) WriteTo(w io.Writer) (int64, error) {
	n := uint32(i)
	nn, err := w.Write([]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
	return int64(nn), err
}

//Decode a Int
func (i *Int) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [4]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*i = Int(int32(bs[0])<<24 | int32(bs[1])<<16 | int32(bs[2])<<8 | int32(bs[3]))
	return
}

// Encode a Long
func (l Long) WriteTo(w io.Writer) (int64, error) {
	n := uint64(l)
	nn, err := w.Write([]byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n),
	})
	return int64(nn), err
}

//Decode a Long
func (l *Long) ReadFrom(r io.Reader) (n int64, err error) {
	var bs [8]byte
	if nn, err := io.ReadFull(r, bs[:]); err != nil {
		return int64(nn), err
	} else {
		n += int64(nn)
	}

	*l = Long(int64(bs[0])<<56 | int64(bs[1])<<48 | int64(bs[2])<<40 | int64(bs[3])<<32 |
		int64(bs[4])<<24 | int64(bs[5])<<16 | int64(bs[6])<<8 | int64(bs[7]))
	return
}

//Encode a VarInt
func (v VarInt) WriteTo(w io.Writer) (n int64, err error) {
	var vi = make([]byte, 0, MaxVarIntLen)
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
	nn, err := w.Write(vi)
	return int64(nn), err
}

//Decode a VarInt
func (v *VarInt) ReadFrom(r io.Reader) (n int64, err error) {
	var V uint32
	for sec := byte(0x80); sec&0x80 != 0; n++ {
		if n > MaxVarIntLen {
			return n, errors.New("VarInt is too big")
		}

		sec, err = readByte(r)
		if err != nil {
			return n, err
		}

		V |= uint32(sec&0x7F) << uint32(7*n)
	}

	*v = VarInt(V)
	return
}

//Encode a VarLong
func (v VarLong) WriteTo(w io.Writer) (n int64, err error) {
	var vi = make([]byte, 0, MaxVarLongLen)
	num := uint64(v)
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
	nn, err := w.Write(vi)
	return int64(nn), err
}

//Decode a VarLong
func (v *VarLong) ReadFrom(r io.Reader) (n int64, err error) {
	var V uint64
	for sec := byte(0x80); sec&0x80 != 0; n++ {
		if n >= MaxVarLongLen {
			return n, errors.New("VarLong is too big")
		}
		sec, err = readByte(r)
		if err != nil {
			return
		}

		V |= uint64(sec&0x7F) << uint64(7*n)
	}

	*v = VarLong(V)
	return
}

//Encode a Position
func (p Position) WriteTo(w io.Writer) (n int64, err error) {
	var b [8]byte
	position := uint64(p.X&0x3FFFFFF)<<38 | uint64((p.Z&0x3FFFFFF)<<12) | uint64(p.Y&0xFFF)
	for i := 7; i >= 0; i-- {
		b[i] = byte(position)
		position >>= 8
	}
	nn, err := w.Write(b[:])
	return int64(nn), err
}

// Decode a Position
func (p *Position) ReadFrom(r io.Reader) (n int64, err error) {
	var v Long
	nn, err := v.ReadFrom(r)
	if err != nil {
		return nn, err
	}
	n += nn

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
	return
}

// ToDeg convert Angle to Degree
func (a Angle) ToDeg() float64 {
	return 360 * float64(a) / 256
}

// ToRad convert Angle to Radian
func (a Angle) ToRad() float64 {
	return 2 * math.Pi * float64(a) / 256
}

func (a Angle) WriteTo(w io.Writer) (int64, error) {
	return Byte(a).WriteTo(w)
}

func (a *Angle) ReadFrom(r io.Reader) (int64, error) {
	return (*Byte)(a).ReadFrom(r)
}

//Encode a Float
func (f Float) WriteTo(w io.Writer) (n int64, err error) {
	return Int(math.Float32bits(float32(f))).WriteTo(w)
}

// Decode a Float
func (f *Float) ReadFrom(r io.Reader) (n int64, err error) {
	var v Int

	n, err = v.ReadFrom(r)
	if err != nil {
		return
	}

	*f = Float(math.Float32frombits(uint32(v)))
	return
}

//Encode a Double
func (d Double) WriteTo(w io.Writer) (n int64, err error) {
	return Long(math.Float64bits(float64(d))).WriteTo(w)
}

// Decode a Double
func (d *Double) ReadFrom(r io.Reader) (n int64, err error) {
	var v Long
	n, err = v.ReadFrom(r)
	if err != nil {
		return
	}

	*d = Double(math.Float64frombits(uint64(v)))
	return
}

//NBT encode a value as Named Binary Tag
func NBT(v interface{}, optionalTagName ...string) Field {
	if len(optionalTagName) > 0 {
		return nbtField{V: v, FieldName: optionalTagName[0]}
	}
	return nbtField{V: v}
}

type nbtField struct {
	V         interface{}
	FieldName string
}

func (n nbtField) WriteTo(w io.Writer) (int64, error) {
	var buf bytes.Buffer
	if n.V == nil {
		buf.WriteByte(nbt.TagEnd)
	} else if err := nbt.NewEncoder(&buf).Encode(n.V, n.FieldName); err != nil {
		return 0, err
	}
	return buf.WriteTo(w)
}

func (n nbtField) ReadFrom(r io.Reader) (int64, error) {
	// LimitReader is used to count reader length
	lr := &io.LimitedReader{R: r, N: math.MaxInt64}
	err := nbt.NewDecoder(lr).Decode(n.V)
	if err != nil && errors.Is(err, nbt.ErrEND) {
		err = nil
	}
	return math.MaxInt64 - lr.N, err
}

// Encode a ByteArray
func (b ByteArray) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := VarInt(len(b)).WriteTo(w)
	if err != nil {
		return n1, err
	}
	n2, err := w.Write(b)
	return n1 + int64(n2), err
}

// Decode a ByteArray
func (b *ByteArray) ReadFrom(r io.Reader) (n int64, err error) {
	var Len VarInt
	n1, err := Len.ReadFrom(r)
	if err != nil {
		return n1, err
	}
	buf := bytes.NewBuffer(*b)
	buf.Reset()
	n2, err := io.CopyN(buf, r, int64(Len))
	*b = buf.Bytes()
	return n1 + n2, err
}

// Encode a UUID
func (u UUID) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write(u[:])
	return int64(nn), err
}

// Decode a UUID
func (u *UUID) ReadFrom(r io.Reader) (n int64, err error) {
	nn, err := io.ReadFull(r, (*u)[:])
	return int64(nn), err
}

// Encode a PluginsMessageData
func (p *PluginMessageData) WriteTo(w io.Writer) (n int64, err error) {
	nn, err := w.Write(*p)
	return int64(nn), err
}

// Decode a PluginsMessageData
func (p *PluginMessageData) ReadFrom(r io.Reader) (n int64, err error) {
	*p, err = io.ReadAll(r)
	return int64(len(*p)), err
}

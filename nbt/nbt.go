// Package nbt implement the Named Binary Tag format of Minecraft.
// It provides api like encoding/xml package.
package nbt

import (
	"io"
)

//Tag type IDs
const (
	TagEnd byte = iota
	TagByte
	TagShort
	TagInt
	TagLong
	TagFloat
	TagDouble
	TagByteArray
	TagString
	TagList
	TagCompound
	TagIntArray
	TagLongArray
	TagNone = 0xFF
)

func IsArrayTag(ty byte) bool {
	return ty == TagByteArray || ty == TagIntArray || ty == TagLongArray
}

type DecoderReader = interface {
	io.ByteReader
	io.Reader
}
type Decoder struct {
	r DecoderReader
}

func NewDecoder(r io.Reader) *Decoder {
	d := new(Decoder)
	if br, ok := r.(DecoderReader); ok {
		d.r = br
	} else {
		d.r = reader{r}
	}
	return d
}

type reader struct {
	io.Reader
}

func (r reader) ReadByte() (byte, error) {
	var b [1]byte
	n, err := r.Read(b[:])
	if n == 1 {
		return b[0], nil
	}
	return 0, err
}

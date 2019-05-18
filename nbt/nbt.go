package nbt

import (
	"bufio"
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
)

type Decoder struct {
	r interface {
		io.ByteReader
		io.Reader
	}
}

func NewDecoder(r io.Reader) *Decoder {
	d := new(Decoder)
	if br, ok := r.(interface {
		io.ByteReader
		io.Reader
	}); ok {
		d.r = br
	} else {
		d.r = bufio.NewReader(r)
	}
	return d
}

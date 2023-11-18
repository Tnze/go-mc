// Package nbt implement the Named Binary Tag format of Minecraft.
// It provides api like encoding/json package.
package nbt

import (
	"io"
)

// Tag type IDs
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

type DecoderReader = interface {
	io.ByteReader
	io.Reader
}
type Decoder struct {
	r                     DecoderReader
	disallowUnknownFields bool
	networkFormat         bool
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

// DisallowUnknownFields makes the decoder return an error when unmarshalling a compound
// tag item that has a tag name not present in the destination struct.
func (d *Decoder) DisallowUnknownFields() {
	d.disallowUnknownFields = true
}

// NetworkFormat controls wether the decoder parsing nbt in "network format".
// Means it haven't a tag name for root tag.
//
// It is disabled by default.
func (d *Decoder) NetworkFormat(enable bool) {
	d.networkFormat = enable
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

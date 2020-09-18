// Package nbt implement the Named Binary Tag format of Minecraft.
// It provides api like encoding/xml package.
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
	TagNone = 0xFF
)

func IsArrayTag(ty byte) bool {
	return ty == TagByteArray || ty == TagIntArray || ty == TagLongArray
}

type DecoderReader = interface {
	io.ByteScanner
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
		d.r = bufio.NewReader(r)
	}
	return d
}

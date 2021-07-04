package nbt

import "io"

type NBTDecoder interface {
	Decode(tagType byte, r DecoderReader) error
}

type NBTEncoder interface {
	TagType() byte
	Encode(w io.Writer) error
}

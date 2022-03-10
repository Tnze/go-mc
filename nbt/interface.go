package nbt

import "io"

type Unmarshaler interface {
	UnmarshalNBT(tagType byte, r DecoderReader) error
}

type Marshaler interface {
	TagType() byte
	MarshalNBT(w io.Writer) error
}

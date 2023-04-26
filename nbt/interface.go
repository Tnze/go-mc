package nbt

import "io"

type Unmarshaler interface {
	UnmarshalNBT(tagType byte, r DecoderReader) error
}

type Marshaler interface {
	TagType() byte
	MarshalNBT(w io.Writer) error
}

// FieldsUnmarshaler is a type can hold many Tags just like a TagCompound.
//
// If and only if a type which implements this interface is used as an anonymous field of a struct,
// and didn't set a struct tag, the content it holds will be considered as in the outer struct.
type FieldsUnmarshaler interface {
	UnmarshalField(tagType byte, tagName string, r DecoderReader) (ok bool, err error)
}

// FieldsMarshaler is similar to FieldsUnmarshaler, but for marshaling.
type FieldsMarshaler interface {
	MarshalFields(w io.Writer) (ok bool, err error)
}

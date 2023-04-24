package fastnbt

import (
	"errors"
	"io"

	"github.com/Tnze/go-mc/nbt"
)

func (v *Value) TagType() byte { return v.tag }

func (v *Value) MarshalNBT(w io.Writer) (err error) {
	switch v.tag {
	case nbt.TagEnd:
		_, err = w.Write([]byte{0})
		if err != nil {
			return err
		}

	case nbt.TagByte, nbt.TagShort, nbt.TagInt, nbt.TagLong, nbt.TagFloat, nbt.TagDouble,
		nbt.TagByteArray, nbt.TagString, nbt.TagIntArray, nbt.TagLongArray:
		_, err = w.Write(v.data)

	case nbt.TagList:
		// Take a look at the first element's tag.
		// If length == 0, use TagEnd
		elemType := nbt.TagEnd
		length := len(v.list)
		if length > 0 {
			elemType = v.list[0].tag
		}

		_, err = w.Write([]byte{elemType})
		if err != nil {
			return
		}

		err = writeInt32(w, int32(length))
		if err != nil {
			return
		}

		for _, val := range v.list {
			err = val.MarshalNBT(w)
			if err != nil {
				return
			}
		}

	case nbt.TagCompound:
		for _, field := range v.comp.kvs {
			err = writeTag(w, field.v.tag, field.tag)
			if err != nil {
				return
			}

			err = field.v.MarshalNBT(w)
			if err != nil {
				return
			}
		}

		_, err = w.Write([]byte{nbt.TagEnd})
		if err != nil {
			return
		}

	default:
		err = errors.New("internal: unknown tag")
	}
	return
}

func writeTag(w io.Writer, tagType byte, tagName string) error {
	if _, err := w.Write([]byte{tagType}); err != nil {
		return err
	}
	bName := []byte(tagName)
	if err := writeInt16(w, int16(len(bName))); err != nil {
		return err
	}
	_, err := w.Write(bName)
	return err
}

func writeInt16(w io.Writer, n int16) error {
	_, err := w.Write([]byte{byte(n >> 8), byte(n)})
	return err
}

func writeInt32(w io.Writer, n int32) error {
	_, err := w.Write([]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
	return err
}

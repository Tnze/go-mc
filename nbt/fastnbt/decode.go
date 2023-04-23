package fastnbt

import (
	"errors"
	"fmt"
	"io"

	"github.com/Tnze/go-mc/nbt"
)

//func (v *Value) Parse(data []byte) {
//	// TODO
//}

func (v *Value) UnmarshalNBT(tagType byte, r nbt.DecoderReader) error {
	v.tag = tagType
	var buf [8]byte
	switch tagType {
	case nbt.TagEnd:
	case nbt.TagByte:
		n, err := r.ReadByte()
		if err != nil {
			return err
		}
		v.data = append(v.data[:0], n)

	case nbt.TagShort:
		if _, err := r.Read(buf[:2]); err != nil {
			return err
		}
		v.data = append(v.data[:0], buf[:2]...)

	case nbt.TagInt, nbt.TagFloat:
		if _, err := r.Read(buf[:4]); err != nil {
			return err
		}
		v.data = append(v.data[:0], buf[:4]...)

	case nbt.TagLong, nbt.TagDouble:
		if _, err := r.Read(buf[:]); err != nil {
			return err
		}
		v.data = append(v.data[:0], buf[:]...)

	case nbt.TagByteArray:
		n, err := readInt32(r)
		if err != nil {
			return err
		}

		v.data = append(v.data[:0], make([]byte, 4+n)...)
		v.data[0], v.data[1], v.data[2], v.data[3] = byte(n>>24), byte(n>>16), byte(n>>8), byte(n)

		_, err = io.ReadFull(r, v.data[4:])
		if err != nil {
			return err
		}

	case nbt.TagString:
		n, err := readInt16(r)
		if err != nil {
			return err
		}

		v.data = append(v.data[:0], make([]byte, 2+n)...)
		v.data[0], v.data[1] = byte(n>>8), byte(n)

		_, err = io.ReadFull(r, v.data[2:])
		if err != nil {
			return err
		}

	case nbt.TagList:
		t, err := r.ReadByte()
		if err != nil {
			return err
		}

		length, err := readInt32(r)
		if err != nil {
			return err
		}

		v.list = v.list[:0]

		for i := int32(0); i < length; i++ {
			field := new(Value)
			err = field.UnmarshalNBT(t, r)
			if err != nil {
				return err
			}

			v.list = append(v.list, field)
		}

	case nbt.TagCompound:
		for {
			t, name, err := readTag(r)
			if err != nil {
				return err
			}

			if t == nbt.TagEnd {
				break
			}

			field := new(Value)
			err = field.UnmarshalNBT(t, r)
			if err != nil {
				return decodeErr{name, err}
			}

			v.comp.kvs = append(v.comp.kvs, kv{tag: name, v: field})
		}
	case nbt.TagIntArray:
		n, err := readInt32(r)
		if err != nil {
			return err
		}

		v.data = append(v.data[:0], make([]byte, 4+n*4)...)
		v.data[0], v.data[1], v.data[2], v.data[3] = byte(n>>24), byte(n>>16), byte(n>>8), byte(n)

		_, err = io.ReadFull(r, v.data[4:])
		if err != nil {
			return err
		}

	case nbt.TagLongArray:
		n, err := readInt32(r)
		if err != nil {
			return err
		}

		v.data = append(v.data[:0], make([]byte, 4+n*8)...)
		v.data[0], v.data[1], v.data[2], v.data[3] = byte(n>>24), byte(n>>16), byte(n>>8), byte(n)

		_, err = io.ReadFull(r, v.data[4:])
		if err != nil {
			return err
		}
	}
	return nil
}

func readTag(r nbt.DecoderReader) (tagType byte, tagName string, err error) {
	tagType, err = r.ReadByte()
	if err != nil {
		return
	}

	switch tagType {
	// case 0x1f, 0x78:
	case nbt.TagEnd:
	default: // Read Tag
		tagName, err = readString(r)
	}
	return
}

func readInt16(r nbt.DecoderReader) (int16, error) {
	var data [2]byte
	_, err := io.ReadFull(r, data[:])
	return int16(data[0])<<8 | int16(data[1]), err
}

func readInt32(r nbt.DecoderReader) (int32, error) {
	var data [4]byte
	_, err := io.ReadFull(r, data[:])
	return int32(data[0])<<24 | int32(data[1])<<16 |
		int32(data[2])<<8 | int32(data[3]), err
}

func readString(r nbt.DecoderReader) (string, error) {
	length, err := readInt16(r)
	if err != nil {
		return "", err
	} else if length < 0 {
		return "", errors.New("string length less than 0")
	}

	var str string
	if length > 0 {
		buf := make([]byte, length)
		_, err = io.ReadFull(r, buf)
		str = string(buf)
	}
	return str, err
}

type decodeErr struct {
	decoding string
	err      error
}

func (d decodeErr) Error() string {
	return fmt.Sprintf("fail to decode tag %q: %v", d.decoding, d.err)
}

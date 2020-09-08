package nbt

import (
	"errors"
	"io"
	"math"
	"reflect"
)

func Marshal(w io.Writer, v interface{}) error {
	return NewEncoder(w).Encode(v)
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v interface{}) error {
	val := reflect.ValueOf(v)
	return e.marshal(val, "")
}

func (e *Encoder) marshal(val reflect.Value, tagName string) error {
	switch vk := val.Kind(); vk {
	default:
		return errors.New("unknown type " + vk.String())

	case reflect.Uint8:
		if err := e.writeTag(TagByte, tagName); err != nil {
			return err
		}
		_, err := e.w.Write([]byte{byte(val.Uint())})
		return err

	case reflect.Int16, reflect.Uint16:
		if err := e.writeTag(TagShort, tagName); err != nil {
			return err
		}
		return e.writeInt16(int16(val.Int()))

	case reflect.Int32, reflect.Uint32:
		if err := e.writeTag(TagInt, tagName); err != nil {
			return err
		}
		return e.writeInt32(int32(val.Int()))

	case reflect.Float32:
		if err := e.writeTag(TagFloat, tagName); err != nil {
			return err
		}
		return e.writeInt32(int32(math.Float32bits(float32(val.Float()))))

	case reflect.Int64, reflect.Uint64:
		if err := e.writeTag(TagLong, tagName); err != nil {
			return err
		}
		return e.writeInt64(int64(val.Int()))

	case reflect.Float64:
		if err := e.writeTag(TagDouble, tagName); err != nil {
			return err
		}
		return e.writeInt64(int64(math.Float64bits(val.Float())))

	case reflect.Array, reflect.Slice:
		n := val.Len()
		switch val.Type().Elem().Kind() {
		case reflect.Uint8: // []byte
			if err := e.writeTag(TagByteArray, tagName); err != nil {
				return err
			}
			if err := e.writeInt32(int32(val.Len())); err != nil {
				return err
			}
			_, err := e.w.Write(val.Bytes())
			return err

		case reflect.Int32:
			if err := e.writeTag(TagIntArray, tagName); err != nil {
				return err
			}
			if err := e.writeInt32(int32(n)); err != nil {
				return err
			}
			for i := 0; i < n; i++ {
				if err := e.writeInt32(int32(val.Index(i).Int())); err != nil {
					return err
				}
			}

		case reflect.Int64:
			if err := e.writeTag(TagLongArray, tagName); err != nil {
				return err
			}
			if err := e.writeInt32(int32(n)); err != nil {
				return err
			}
			for i := 0; i < n; i++ {
				if err := e.writeInt64(val.Index(i).Int()); err != nil {
					return err
				}
			}

		case reflect.Int16:
			if err := e.writeListHeader(TagShort, tagName, val.Len()); err != nil {
				return err
			}
			for i := 0; i < n; i++ {
				if err := e.writeInt16(int16(val.Index(i).Int())); err != nil {
					return err
				}
			}

		case reflect.Float32:
			if err := e.writeListHeader(TagFloat, tagName, val.Len()); err != nil {
				return err
			}
			for i := 0; i < n; i++ {
				if err := e.writeInt32(int32(math.Float32bits(float32(val.Index(i).Float())))); err != nil {
					return err
				}
			}

		case reflect.Float64:
			if err := e.writeListHeader(TagDouble, tagName, val.Len()); err != nil {
				return err
			}
			for i := 0; i < n; i++ {
				if err := e.writeInt64(int64(math.Float64bits(val.Index(i).Float()))); err != nil {
					return err
				}
			}

		case reflect.String:
			if err := e.writeListHeader(TagString, tagName, n); err != nil {
				return err
			}
			for i := 0; i < n; i++ {
				// Write length of this string
				s := val.Index(i).String()
				if err := e.writeInt16(int16(len(s))); err != nil {
					return err
				}
				// Write string
				if _, err := e.w.Write([]byte(s)); err != nil {
					return err
				}
			}
		case reflect.Struct, reflect.Interface:
			if err := e.writeListHeader(TagCompound, tagName, n); err != nil {
				return err
			}
			for i := 0; i < n; i++ {
				elemVal := val.Index(i)
				if val.Type().Elem().Kind() == reflect.Interface {
					elemVal = reflect.ValueOf(elemVal.Interface())
				}
				err := e.marshal(elemVal, "")
				if err != nil {
					return err
				}
			}
		default:
			return errors.New("unknown type " + val.Type().String() + " slice")
		}

	case reflect.String:
		if err := e.writeTag(TagString, tagName); err != nil {
			return err
		}
		if err := e.writeInt16(int16(val.Len())); err != nil {
			return err
		}
		_, err := e.w.Write([]byte(val.String()))
		return err

	case reflect.Struct:
		if err := e.writeTag(TagCompound, ""); err != nil {
			return err
		}

		n := val.NumField()
		for i := 0; i < n; i++ {
			f := val.Type().Field(i)
			tag := f.Tag.Get("nbt")
			if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
				continue // Private field
			}

			tagName := f.Name
			if tag != "" {
				tagName = tag
			}

			err := e.marshal(val.Field(i), tagName)
			if err != nil {
				return err
			}
		}
		_, err := e.w.Write([]byte{TagEnd})
		return err
	}
	return nil
}

func (e *Encoder) writeTag(tagType byte, tagName string) error {
	if _, err := e.w.Write([]byte{tagType}); err != nil {
		return err
	}
	bName := []byte(tagName)
	if err := e.writeInt16(int16(len(bName))); err != nil {
		return err
	}
	_, err := e.w.Write(bName)
	return err
}

func (e *Encoder) writeListHeader(elementType byte, tagName string, n int) (err error) {
	if err = e.writeTag(TagList, tagName); err != nil {
		return
	}
	if _, err = e.w.Write([]byte{elementType}); err != nil {
		return
	}
	// Write length of strings
	if err = e.writeInt32(int32(n)); err != nil {
		return
	}
	return nil
}

func (e *Encoder) writeNamelessTag(tagType byte, tagName string) error {
	_, err := e.w.Write([]byte{tagType})
	return err
}

func (e *Encoder) writeInt16(n int16) error {
	_, err := e.w.Write([]byte{byte(n >> 8), byte(n)})
	return err
}

func (e *Encoder) writeInt32(n int32) error {
	_, err := e.w.Write([]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
	return err
}

func (e *Encoder) writeInt64(n int64) error {
	_, err := e.w.Write([]byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
	return err
}

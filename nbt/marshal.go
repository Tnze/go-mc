package nbt

import (
	"errors"
	"io"
	"math"
	"reflect"
	"strings"
)

var (
	ErrMustBeStruct = errors.New("a compound can only be a struct")
)

func Marshal(w io.Writer, v interface{}) error {
	return NewEncoder(w).Encode(v)
}

func MarshalCompound(w io.Writer, v interface{}, rootTagName string) error {
	enc := NewEncoder(w)
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Struct {
		return ErrMustBeStruct
	}
	return enc.marshal(val, TagCompound, rootTagName)
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(v interface{}) error {
	val := reflect.ValueOf(v)
	return e.marshal(val, getTagType(val.Type()), "")
}

func (e *Encoder) marshal(val reflect.Value, tagType byte, tagName string) error {
	if err := e.writeHeader(val, tagType, tagName); err != nil {
		return err
	}
	return e.writeValue(val, tagType)
}

func (e *Encoder) writeHeader(val reflect.Value, tagType byte, tagName string) (err error) {
	if tagType == TagList {
		eleType := getTagType(val.Type().Elem())
		err = e.writeListHeader(eleType, tagName, val.Len())
	} else {
		err = e.writeTag(tagType, tagName)
	}
	return err
}

func (e *Encoder) writeValue(val reflect.Value, tagType byte) error {
	switch tagType {
	default:
		return errors.New("unsupported type " + val.Type().Kind().String())
	case TagByte:
		_, err := e.w.Write([]byte{byte(val.Uint())})
		return err
	case TagShort:
		return e.writeInt16(int16(val.Int()))
	case TagInt:
		return e.writeInt32(int32(val.Int()))
	case TagFloat:
		return e.writeInt32(int32(math.Float32bits(float32(val.Float()))))
	case TagLong:
		return e.writeInt64(val.Int())
	case TagDouble:
		return e.writeInt64(int64(math.Float64bits(val.Float())))
	case TagByteArray, TagIntArray, TagLongArray:
		n := val.Len()
		if err := e.writeInt32(int32(n)); err != nil {
			return err
		}

		if tagType == TagByteArray {
			_, err := e.w.Write(val.Bytes())
			return err
		} else {
			for i := 0; i < n; i++ {
				v := val.Index(i).Int()

				var err error
				if tagType == TagIntArray {
					err = e.writeInt32(int32(v))
				} else if tagType == TagLongArray {
					err = e.writeInt64(v)
				}
				if err != nil {
					return err
				}
			}
		}

	case TagList:
		for i := 0; i < val.Len(); i++ {
			arrVal := val.Index(i)
			err := e.writeValue(arrVal, getTagType(arrVal.Type()))
			if err != nil {
				return err
			}
		}

	case TagString:
		if err := e.writeInt16(int16(val.Len())); err != nil {
			return err
		}
		_, err := e.w.Write([]byte(val.String()))
		return err

	case TagCompound:
		if val.Kind() == reflect.Interface {
			val = reflect.ValueOf(val.Interface())
		}

		n := val.NumField()
		for i := 0; i < n; i++ {
			f := val.Type().Field(i)
			tag := f.Tag.Get("nbt")
			if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
				continue // Private field
			}

			tagProps := parseTag(f, tag)
			err := e.marshal(val.Field(i), tagProps.Type, tagProps.Name)
			if err != nil {
				return err
			}
		}
		_, err := e.w.Write([]byte{TagEnd})
		return err
	}
	return nil
}

func getTagType(vk reflect.Type) byte {
	switch vk.Kind() {
	case reflect.Uint8:
		return TagByte
	case reflect.Int16, reflect.Uint16:
		return TagShort
	case reflect.Int32, reflect.Uint32:
		return TagInt
	case reflect.Float32:
		return TagFloat
	case reflect.Int64, reflect.Uint64:
		return TagLong
	case reflect.Float64:
		return TagDouble
	case reflect.String:
		return TagString
	case reflect.Struct, reflect.Interface:
		return TagCompound
	case reflect.Array, reflect.Slice:
		switch vk.Elem().Kind() {
		case reflect.Uint8: // Special types for these values
			return TagByteArray
		case reflect.Int32:
			return TagIntArray
		case reflect.Int64:
			return TagLongArray
		default:
			return TagList
		}
	default:
		return TagNone
	}
}

type tagProps struct {
	Name string
	Type byte
}

func parseTag(f reflect.StructField, tagName string) tagProps {
	result := tagProps{}
	result.Name = tagName
	if result.Name == "" {
		result.Name = f.Name
	}

	nbtType := f.Tag.Get("nbt_type")
	result.Type = getTagType(f.Type)
	if strings.Contains(nbtType, "noarray") {
		if IsArrayTag(result.Type) {
			result.Type = TagList // for expanding the array to a standard list
		} else {
			panic("noarray is only supported for array types (byte, int, long)")
		}
	}

	return result
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

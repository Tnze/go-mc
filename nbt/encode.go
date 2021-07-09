package nbt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// Marshal is the shortcut of NewEncoder().Encode() with empty tag name.
// Notices that repeatedly init buffers is low efficiency.
// Using Encoder and Reset the buffer in each times is recommended in that cases.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := NewEncoder(&buf).Encode(v, "")
	return buf.Bytes(), err
}

type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// Encode encodes v into the writer inside Encoder with the root tag named tagName.
// In most cases, the root tag typed TagCompound and the tag name is empty string,
// but any other type is allowed just because there is valid technically. Once if
// you should pass an string into this, you should get a TagString.
//
// Normally, any slice or array typed Go value will be encoded as TagList,
// expect `[]int8`, `[]int32`, `[]int64`, `[]uint8`, `[]uint32` and `[]uint64`,
// which TagByteArray, TagIntArray and TagLongArray.
// To force encode them as TagList, add a struct field tag.
// You haven't ability to encode them as TagList as root element at this time,
// issue or pull-request is welcome.
func (e *Encoder) Encode(v interface{}, tagName string) error {
	val := reflect.ValueOf(v)
	return e.marshal(val, getTagType(val), tagName)
}

func (e *Encoder) marshal(val reflect.Value, tagType byte, tagName string) error {
	if err := e.writeTag(tagType, tagName); err != nil {
		return err
	}
	if val.CanInterface() {
		if encoder, ok := val.Interface().(NBTEncoder); ok {
			return encoder.Encode(e.w)
		}
	}
	return e.writeValue(val, tagType)
}

func (e *Encoder) writeValue(val reflect.Value, tagType byte) error {
	switch tagType {
	default:
		return errors.New("unsupported type 0x" + strconv.FormatUint(uint64(tagType), 16))
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
		var eleType byte
		if val.Len() > 0 {
			eleType = getTagType(val.Index(0))
		} else {
			eleType = getTagTypeByType(val.Type().Elem())
		}
		if err := e.writeListHeader(eleType, val.Len()); err != nil {
			return err
		}

		for i := 0; i < val.Len(); i++ {
			arrVal := val.Index(i)
			err := e.writeValue(arrVal, getTagType(arrVal))
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
		for val.Kind() == reflect.Interface {
			val = val.Elem()
		}

		switch val.Kind() {
		case reflect.Struct:
			n := val.NumField()
			for i := 0; i < n; i++ {
				f := val.Type().Field(i)
				v := val.Field(i)
				tag := f.Tag.Get("nbt")
				if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
					continue // Private field
				}

				tagProps := parseTag(f, v, tag)
				if err := e.marshal(val.Field(i), tagProps.Type, tagProps.Name); err != nil {
					return err
				}
			}
		case reflect.Map:
			r := val.MapRange()
			for r.Next() {
				var tagName string
				if tn, ok := r.Key().Interface().(fmt.Stringer); ok {
					tagName = tn.String()
				} else {
					tagName = r.Key().String()
				}
				tagValue := r.Value()
				tagType := getTagType(tagValue)
				if tagType == TagNone {
					return errors.New("unsupported value " + tagValue.String())
				}

				if err := e.marshal(tagValue, tagType, tagName); err != nil {
					return err
				}
			}
		}

		_, err := e.w.Write([]byte{TagEnd})
		return err
	}
	return nil
}

func getTagType(v reflect.Value) byte {
	if v.CanInterface() {
		if encoder, ok := v.Interface().(NBTEncoder); ok {
			return encoder.TagType()
		}
	}
	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		var elemType byte
		if v.Len() > 0 {
			elemType = getTagType(v.Index(0))
		} else {
			elemType = getTagTypeByType(v.Type().Elem())
		}
		switch elemType {
		case TagByte: // Special types for these values
			return TagByteArray
		case TagInt:
			return TagIntArray
		case TagLong:
			return TagLongArray
		default:
			return TagList
		}
	default:
		return getTagTypeByType(v.Type())
	}
}

func getTagTypeByType(vk reflect.Type) byte {
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
	case reflect.Struct, reflect.Interface, reflect.Map:
		return TagCompound
	default:
		return TagNone
	}
}

type tagProps struct {
	Name string
	Type byte
}

func parseTag(f reflect.StructField, v reflect.Value, tagName string) (result tagProps) {
	if tagName != "" {
		result.Name = tagName
	} else {
		result.Name = f.Name
	}

	nbtType := f.Tag.Get("nbt_type")
	result.Type = getTagType(v)
	if strings.Contains(nbtType, "list") {
		if IsArrayTag(result.Type) {
			result.Type = TagList // for expanding the array to a standard list
		} else {
			panic("list is only supported for array types ([]byte, []int, []long)")
		}
	}

	return
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

func (e *Encoder) writeListHeader(elementType byte, n int) (err error) {
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

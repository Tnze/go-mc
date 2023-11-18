package nbt

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

// Marshal is the shortcut of NewEncoder().Encode() with empty tag name.
// Notices that repeatedly init buffers is low efficiency.
// Using Encoder and Reset the buffer in each time is recommended in that cases.
func Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	err := NewEncoder(&buf).Encode(v, "")
	return buf.Bytes(), err
}

type Encoder struct {
	w             io.Writer
	networkFormat bool
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

// NetworkFormat controls wether encoder encoding nbt in "network format".
// Means it haven't a tag name for root tag.
//
// It is disabled by default.
func (e *Encoder) NetworkFormat(enable bool) {
	e.networkFormat = enable
}

// Encode encodes v into the writer inside Encoder with the root tag named tagName.
// In most cases, the root tag typed TagCompound and the tag name is empty string,
// but any other type is allowed just because there is valid technically. Once if
// you should pass a string into this, you should get a TagString.
//
// Normally, any slice or array typed Go value will be encoded as TagList,
// expect `[]int8`, `[]int32`, `[]int64`, `[]uint8`, `[]uint32` and `[]uint64`,
// which TagByteArray, TagIntArray and TagLongArray.
// To force encode them as TagList, add a struct field tag.
func (e *Encoder) Encode(v any, tagName string) (err error) {
	t, val := getTagType(reflect.ValueOf(v))
	if e.networkFormat {
		_, err = e.w.Write([]byte{t})
	} else {
		err = writeTag(e.w, t, tagName)
	}
	if err != nil {
		return err
	}
	return e.marshal(val, t)
}

func (e *Encoder) marshal(val reflect.Value, tagType byte) error {
	if val.CanInterface() {
		if encoder, ok := val.Interface().(Marshaler); ok {
			return encoder.MarshalNBT(e.w)
		}
	}
	return e.writeValue(val, tagType)
}

func (e *Encoder) writeValue(val reflect.Value, tagType byte) error {
	switch tagType {
	default:
		return errors.New("unsupported type 0x" + strconv.FormatUint(uint64(tagType), 16))
	case TagByte:
		var err error
		switch val.Kind() {
		case reflect.Bool:
			var b byte
			if val.Bool() {
				b = 1
			}
			_, err = e.w.Write([]byte{b})
		case reflect.Int8:
			_, err = e.w.Write([]byte{byte(val.Int())})
		case reflect.Uint8:
			_, err = e.w.Write([]byte{byte(val.Uint())})
		}
		return err
	case TagShort:
		return writeInt16(e.w, int16(val.Int()))
	case TagInt:
		return writeInt32(e.w, int32(val.Int()))
	case TagFloat:
		return writeInt32(e.w, int32(math.Float32bits(float32(val.Float()))))
	case TagLong:
		return writeInt64(e.w, val.Int())
	case TagDouble:
		return writeInt64(e.w, int64(math.Float64bits(val.Float())))
	case TagByteArray, TagIntArray, TagLongArray:
		n := val.Len()
		if err := writeInt32(e.w, int32(n)); err != nil {
			return err
		}

		if tagType == TagByteArray {
			var data []byte
			switch val.Type().Elem().Kind() {
			case reflect.Bool:
				data = make([]byte, val.Len())
				for i := range data {
					if val.Index(i).Bool() {
						data[i] = 1
					} else {
						data[i] = 0
					}
				}
			case reflect.Uint8:
				data = val.Bytes()
			case reflect.Int8:
				data = unsafe.Slice((*byte)(val.UnsafePointer()), val.Len())
			}
			_, err := e.w.Write(data)
			return err
		} else {
			for i := 0; i < n; i++ {
				elem := val.Index(i)
				for elem.Kind() == reflect.Interface {
					elem = elem.Elem()
				}
				var err error
				var v int64
				switch elem.Kind() {
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					v = elem.Int()
				case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
					v = int64(elem.Uint())
				default:
					return errors.New("value typed " + elem.Type().String() + "is not allowed in Tag 0x" + strconv.FormatUint(uint64(tagType), 16))
				}
				if tagType == TagIntArray {
					err = writeInt32(e.w, int32(v))
				} else if tagType == TagLongArray {
					err = writeInt64(e.w, v)
				}
				if err != nil {
					return err
				}
			}
		}

	case TagList:
		var eleType byte
		if val.Len() > 0 {
			eleType, _ = getTagType(val.Index(0))
		} else {
			eleType = getTagTypeByType(val.Type().Elem())
		}
		if err := e.writeListHeader(eleType, val.Len()); err != nil {
			return err
		}

		for i := 0; i < val.Len(); i++ {
			arrType, arrVal := getTagType(val.Index(i))
			err := e.writeValue(arrVal, arrType)
			if err != nil {
				return err
			}
		}

	case TagString:
		var str []byte
		if val.NumMethod() > 0 && val.CanInterface() {
			if t, ok := val.Interface().(encoding.TextMarshaler); ok {
				var err error
				str, err = t.MarshalText()
				if err != nil {
					return err
				}
			}
		} else {
			str = []byte(val.String())
		}
		if err := writeInt16(e.w, int16(len(str))); err != nil {
			return err
		}
		_, err := e.w.Write(str)
		return err

	case TagCompound:
		for val.Kind() == reflect.Interface {
			val = val.Elem()
		}

		switch val.Kind() {
		case reflect.Struct:
			fields := cachedTypeFields(val.Type())
		FieldLoop:
			for i := range fields.list {
				t := &fields.list[i]

				v := val
				for _, i := range t.index {
					if v.Kind() == reflect.Pointer {
						if v.IsNil() {
							continue FieldLoop
						}
						v = v.Elem()
					}
					v = v.Field(i)
				}

				if t.omitEmpty && isEmptyValue(v) {
					continue
				}
				typ, v := getTagType(v)
				if typ == TagEnd {
					return fmt.Errorf("encode %q error: unsupport type %v", t.name, v.Type())
				}

				if t.asList {
					switch typ {
					case TagByteArray, TagIntArray, TagLongArray:
						typ = TagList // override the parsed type
					default:
						return fmt.Errorf("invalid use of ,list struct tag, trying to encode %v as TagList", v.Type())
					}
				}

				if err := writeTag(e.w, typ, t.name); err != nil {
					return err
				}
				if err := e.marshal(v, typ); err != nil {
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
				tagType, tagValue := getTagType(r.Value())
				if tagType == TagEnd {
					return fmt.Errorf("encoding %q error: unsupport type %v", tagName, tagValue.Type())
				}

				if err := writeTag(e.w, tagType, tagName); err != nil {
					return err
				}
				if err := e.marshal(tagValue, tagType); err != nil {
					return err
				}
			}
		}

		_, err := e.w.Write([]byte{TagEnd})
		return err
	}
	return nil
}

func getTagType(v reflect.Value) (byte, reflect.Value) {
	for {
		// Load value from interface
		if v.Kind() == reflect.Interface && !v.IsNil() {
			v = v.Elem()
			continue
		}

		if v.Kind() != reflect.Ptr {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v any
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		if v.Type().NumMethod() > 0 && v.CanInterface() {
			i := v.Interface()
			if u, ok := i.(Marshaler); ok {
				return u.TagType(), v
			} else if _, ok := i.(encoding.TextMarshaler); ok {
				return TagString, v
			}
		}

		v = v.Elem()
	}

	if v.Type().NumMethod() > 0 && v.CanInterface() {
		i := v.Interface()
		if u, ok := i.(Marshaler); ok {
			return u.TagType(), v
		} else if _, ok := i.(encoding.TextMarshaler); ok {
			return TagString, v
		}
	}

	switch v.Kind() {
	case reflect.Array, reflect.Slice:
		var elemType byte
		if v.Len() > 0 {
			elemType, _ = getTagType(v.Index(0))
		} else {
			elemType = getTagTypeByType(v.Type().Elem())
		}
		switch elemType {
		case TagByte: // Special types for these values
			return TagByteArray, v
		case TagInt:
			return TagIntArray, v
		case TagLong:
			return TagLongArray, v
		default:
			return TagList, v
		}

	default:
		return getTagTypeByType(v.Type()), v
	}
}

func getTagTypeByType(vk reflect.Type) byte {
	switch vk.Kind() {
	case reflect.Bool, reflect.Int8, reflect.Uint8:
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
	case reflect.Struct, reflect.Map:
		return TagCompound
	default:
		return TagEnd
	}
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

func (e *Encoder) writeListHeader(elementType byte, n int) (err error) {
	if _, err = e.w.Write([]byte{elementType}); err != nil {
		return
	}
	// Write length of strings
	if err = writeInt32(e.w, int32(n)); err != nil {
		return
	}
	return nil
}

func writeInt16(w io.Writer, n int16) error {
	_, err := w.Write([]byte{byte(n >> 8), byte(n)})
	return err
}

func writeInt32(w io.Writer, n int32) error {
	_, err := w.Write([]byte{byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n)})
	return err
}

func writeInt64(w io.Writer, n int64) error {
	_, err := w.Write([]byte{
		byte(n >> 56), byte(n >> 48), byte(n >> 40), byte(n >> 32),
		byte(n >> 24), byte(n >> 16), byte(n >> 8), byte(n),
	})
	return err
}

// Copied from encoding/json/encode.go
func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Pointer:
		return v.IsNil()
	}
	return false
}

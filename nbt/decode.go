package nbt

import (
	"bytes"
	"encoding"
	"errors"
	"fmt"
	"io"
	"math"
	"reflect"
)

// Unmarshal decode binary NBT data and fill into v
// This is a shortcut to `NewDecoder(bytes.NewReader(data)).Decode(v)`.
func Unmarshal(data []byte, v interface{}) error {
	_, err := NewDecoder(bytes.NewReader(data)).Decode(v)
	return err
}

// Decode method decodes an NBT value from the reader underline the Decoder into v.
// Internally try to handle all possible v by reflection,
// but the type of v must matches the NBT value logically.
// For example, you can decode an NBT value which root tag is TagCompound(0x0a)
// into a struct or map, but not a string.
//
// If v implement Unmarshaler, the method will be called and override the default behavior.
// Else if v implement encoding.TextUnmarshaler, the value will be encoded as TagString.
//
// This method also return tag name of the root tag.
// In real world, it is often empty, but the API should allow you to get it when ever you want.
func (d *Decoder) Decode(v interface{}) (string, error) {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return "", errors.New("nbt: non-pointer passed to Decode")
	}
	//start read NBT
	tagType, tagName, err := d.readTag()
	if err != nil {
		return tagName, fmt.Errorf("nbt: %w", err)
	}

	if c := d.checkCompressed(tagType); c != "" {
		return tagName, fmt.Errorf("nbt: unknown Tag, maybe compressed by %s, uncompress it first", c)
	}

	// We decode val not val.Elem because the Unmarshaler interface
	// test must be applied at the top level of the value.
	err = d.unmarshal(val, tagType)
	if err != nil {
		return tagName, fmt.Errorf("nbt: fail to decode tag %q: %w", tagName, err)
	}
	return tagName, nil
}

// checkCompressed check if the first byte is compress head
func (d *Decoder) checkCompressed(head byte) (compress string) {
	switch head {
	case 0x1f:
		return "gzip"
	case 0x78:
		return "zlib"
	default:
		return ""
	}
}

// ErrEND error will be returned when reading a NBT with only Tag_End
var ErrEND = errors.New("unexpected TAG_End")

func (d *Decoder) unmarshal(val reflect.Value, tagType byte) error {
	u, t, val, assign := indirect(val, tagType == TagEnd)
	if assign != nil {
		defer assign()
	}
	if u != nil {
		return u.UnmarshalNBT(tagType, d.r)
	}

	switch tagType {
	default:
		return fmt.Errorf("unknown Tag 0x%02x", tagType)
	case TagEnd:
		return ErrEND

	case TagByte:
		value, err := d.readByte()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagByte as " + vk.String())
		case reflect.Bool:
			val.SetBool(value != 0)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val.SetUint(uint64(value))
		case reflect.Interface:
			val.Set(reflect.ValueOf(value))
		}

	case TagShort:
		value, err := d.readShort()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagShort as " + vk.String())
		case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val.SetUint(uint64(value))
		case reflect.Interface:
			val.Set(reflect.ValueOf(value))
		}

	case TagInt:
		value, err := d.readInt()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagInt as " + vk.String())
		case reflect.Int, reflect.Int32, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint32, reflect.Uint64:
			val.SetUint(uint64(value))
		case reflect.Interface:
			val.Set(reflect.ValueOf(value))
		}

	case TagFloat:
		vInt, err := d.readInt()
		if err != nil {
			return err
		}
		value := math.Float32frombits(uint32(vInt))
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagFloat as " + vk.String())
		case reflect.Float32:
			val.Set(reflect.ValueOf(value))
		case reflect.Float64:
			val.Set(reflect.ValueOf(float64(value)))
		case reflect.Interface:
			val.Set(reflect.ValueOf(value))
		}

	case TagLong:
		value, err := d.readLong()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagLong as " + vk.String())
		case reflect.Int, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint64:
			val.SetUint(uint64(value))
		case reflect.Interface:
			val.Set(reflect.ValueOf(value))
		}

	case TagDouble:
		vInt, err := d.readLong()
		if err != nil {
			return err
		}
		value := math.Float64frombits(uint64(vInt))

		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagDouble as " + vk.String())
		case reflect.Float64:
			val.Set(reflect.ValueOf(value))
		case reflect.Interface:
			val.Set(reflect.ValueOf(value))
		}

	case TagString:
		s, err := d.readString()
		if err != nil {
			return err
		}
		if t != nil {
			err := t.UnmarshalText([]byte(s))
			if err != nil {
				return err
			}
		} else {
			switch vk := val.Kind(); vk {
			default:
				return errors.New("cannot parse TagString as " + vk.String())
			case reflect.String:
				val.SetString(s)
			case reflect.Interface:
				val.Set(reflect.ValueOf(s))
			}
		}

	case TagByteArray:
		aryLen, err := d.readInt()
		if err != nil {
			return err
		}
		if aryLen < 0 {
			return errors.New("byte array len less than 0")
		}
		ba := make([]byte, aryLen)
		if _, err = io.ReadFull(d.r, ba); err != nil {
			return err
		}

		vt := val.Type()
		if vt == reflect.TypeOf(ba) {
			val.SetBytes(ba)
		} else if vt.Kind() == reflect.Slice {
			switch ve := vt.Elem(); ve.Kind() {
			case reflect.Int8, reflect.Uint8:
				length := int(aryLen)
				if val.Cap() < length {
					val.Set(reflect.MakeSlice(vt, length, length))
				}
				val.SetLen(length)
				switch ve.Kind() {
				case reflect.Int8:
					for i := 0; i < length; i++ {
						val.Index(i).Set(reflect.ValueOf(int8(ba[i])))
					}
				case reflect.Uint8:
					for i := 0; i < length; i++ {
						val.Index(i).Set(reflect.ValueOf(ba[i]))
					}
				}
			default:
				return errors.New("cannot parse TagByteArray to slice of" + ve.String())
			}
		} else if vt.Kind() == reflect.Interface {
			val.Set(reflect.ValueOf(ba))
		} else {
			return errors.New("cannot parse TagByteArray to " + vt.String())
		}

	case TagIntArray:
		aryLen, err := d.readInt()
		if err != nil {
			return err
		}
		vt := val.Type() //receiver must be []int or []int32
		if vt.Kind() == reflect.Interface {
			vt = reflect.TypeOf([]int32{}) // pass
		} else if vt.Kind() == reflect.Array && vt.Len() != int(aryLen) {
			return errors.New("cannot parse TagIntArray to " + vt.String() + ", length not match")
		} else if k := vt.Kind(); k != reflect.Slice && k != reflect.Array {
			return errors.New("cannot parse TagIntArray to " + vt.String() + ", it must be a slice")
		} else if tk := val.Type().Elem().Kind(); tk != reflect.Int && tk != reflect.Int32 {
			return errors.New("cannot parse TagIntArray to " + vt.String())
		}

		buf := val
		if vt.Kind() == reflect.Slice {
			buf = reflect.MakeSlice(vt, int(aryLen), int(aryLen))
		}
		for i := 0; i < int(aryLen); i++ {
			value, err := d.readInt()
			if err != nil {
				return err
			}
			buf.Index(i).SetInt(int64(value))
		}
		if vt.Kind() == reflect.Slice {
			val.Set(buf)
		}

	case TagLongArray:
		aryLen, err := d.readInt()
		if err != nil {
			return err
		}
		vt := val.Type() //receiver must be []int or []int64
		if vt.Kind() == reflect.Interface {
			vt = reflect.TypeOf([]int64{}) // pass
		} else if vt.Kind() != reflect.Slice {
			return errors.New("cannot parse TagLongArray to " + vt.String() + ", it must be a slice")
		}
		switch vt.Elem().Kind() {
		case reflect.Int64:
			buf := reflect.MakeSlice(vt, int(aryLen), int(aryLen))
			for i := 0; i < int(aryLen); i++ {
				value, err := d.readLong()
				if err != nil {
					return err
				}
				buf.Index(i).SetInt(value)
			}
			val.Set(buf)
		case reflect.Uint64:
			buf := reflect.MakeSlice(vt, int(aryLen), int(aryLen))
			for i := 0; i < int(aryLen); i++ {
				value, err := d.readLong()
				if err != nil {
					return err
				}
				buf.Index(i).SetUint(uint64(value))
			}
			val.Set(buf)
		default:
			return errors.New("cannot parse TagLongArray to " + vt.String())
		}

	case TagList:
		listType, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		listLen, err := d.readInt()
		if err != nil {
			return err
		}
		if listLen < 0 {
			return errors.New("list length less than 0")
		}

		// If we need parse TAG_List into slice, make a new with right length.
		// Otherwise, if we need parse into array, we check if len(array) are enough.
		var buf reflect.Value
		vk := val.Kind()
		switch vk {
		default:
			return errors.New("cannot parse TagList as " + vk.String())
		case reflect.Interface:
			buf = reflect.ValueOf(make([]interface{}, listLen))
		case reflect.Slice:
			buf = reflect.MakeSlice(val.Type(), int(listLen), int(listLen))
		case reflect.Array:
			if vl := val.Len(); vl < int(listLen) {
				return fmt.Errorf(
					"TagList has len %d, but array %v only has len %d",
					listLen, val.Type(), vl)
			}
			buf = val
		}
		for i := 0; i < int(listLen); i++ {
			if err := d.unmarshal(buf.Index(i), listType); err != nil {
				return err
			}
		}

		if vk != reflect.Array {
			val.Set(buf)
		}

	case TagCompound:
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagCompound as " + vk.String())
		case reflect.Struct:
			tinfo := getTypeInfo(val.Type())
			for {
				tt, tn, err := d.readTag()
				if err != nil {
					return err
				}
				if tt == TagEnd {
					break
				}
				field := tinfo.findIndexByName(tn)
				if field != -1 {
					err = d.unmarshal(val.Field(field), tt)
					if err != nil {
						return fmt.Errorf("fail to decode tag %q: %w", tn, err)
					}
				} else {
					if err := d.rawRead(tt); err != nil {
						return err
					}
				}
			}
		case reflect.Map:
			if val.Type().Key().Kind() != reflect.String {
				return errors.New("cannot parse TagCompound as " + val.Type().String())
			}
			if val.IsNil() {
				val.Set(reflect.MakeMap(val.Type()))
			}
			for {
				tt, tn, err := d.readTag()
				if err != nil {
					return err
				}
				if tt == TagEnd {
					break
				}
				v := reflect.New(val.Type().Elem())
				if err = d.unmarshal(v.Elem(), tt); err != nil {
					return fmt.Errorf("fail to decode tag %q: %w", tn, err)
				}
				val.SetMapIndex(reflect.ValueOf(tn), v.Elem())
			}
		case reflect.Interface:
			buf := make(map[string]interface{})
			for {
				tt, tn, err := d.readTag()
				if err != nil {
					return err
				}
				if tt == TagEnd {
					break
				}
				var value interface{}
				if err = d.unmarshal(reflect.ValueOf(&value).Elem(), tt); err != nil {
					return fmt.Errorf("fail to decode tag %q: %w", tn, err)
				}
				buf[tn] = value
			}
			val.Set(reflect.ValueOf(buf))
		}
	}

	return nil
}

// indirect walks down v allocating pointers as needed,
// until it gets to a non-pointer.
// If it encounters an Unmarshaler, indirect stops and returns that.
// If decodingNull is true, indirect stops at the first settable pointer, so it
// can be set to nil.
//
// This function is copied and modified from encoding/json
func indirect(v reflect.Value, decodingNull bool) (Unmarshaler, encoding.TextUnmarshaler, reflect.Value, func()) {
	v0 := v
	haveAddr := false
	var assign func()

	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Ptr && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		// Otherwise, try init a new value
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Ptr && !e.IsNil() && (!decodingNull || e.Elem().Kind() == reflect.Ptr) {
				haveAddr = false
				v = e
				continue
			} else if v.CanSet() {
				e = reflect.New(e.Type())
				cv := v
				assign = func() { cv.Set(e.Elem()) }
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Ptr {
			break
		}

		if decodingNull && v.CanSet() {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
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
			if u, ok := i.(Unmarshaler); ok {
				return u, nil, reflect.Value{}, assign
			}
			if u, ok := i.(encoding.TextUnmarshaler); ok {
				return nil, u, v, assign
			}
		}

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return nil, nil, v, assign
}

// rawRead read and discard a value
func (d *Decoder) rawRead(tagType byte) error {
	var buf [8]byte
	switch tagType {
	default:
		return fmt.Errorf("unknown to read 0x%02x", tagType)
	case TagByte:
		_, err := d.readByte()
		return err
	case TagString:
		_, err := d.readString()
		return err
	case TagShort:
		_, err := io.ReadFull(d.r, buf[:2])
		return err
	case TagInt, TagFloat:
		_, err := io.ReadFull(d.r, buf[:4])
		return err
	case TagLong, TagDouble:
		_, err := io.ReadFull(d.r, buf[:8])
		return err
	case TagByteArray:
		aryLen, err := d.readInt()
		if err != nil {
			return err
		}

		if _, err = io.CopyN(io.Discard, d.r, int64(aryLen)); err != nil {
			return err
		}
	case TagIntArray:
		aryLen, err := d.readInt()
		if err != nil {
			return err
		}
		for i := 0; i < int(aryLen); i++ {
			if _, err := d.readInt(); err != nil {
				return err
			}
		}

	case TagLongArray:
		aryLen, err := d.readInt()
		if err != nil {
			return err
		}
		for i := 0; i < int(aryLen); i++ {
			if _, err := d.readLong(); err != nil {
				return err
			}
		}

	case TagList:
		listType, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		listLen, err := d.readInt()
		if err != nil {
			return err
		}
		for i := 0; i < int(listLen); i++ {
			if err := d.rawRead(listType); err != nil {
				return err
			}
		}
	case TagCompound:
		for {
			tt, _, err := d.readTag()
			if err != nil {
				return err
			}
			if tt == TagEnd {
				break
			}
			err = d.rawRead(tt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d *Decoder) readTag() (tagType byte, tagName string, err error) {
	tagType, err = d.r.ReadByte()
	if err != nil {
		return
	}

	if tagType != TagEnd { //Read Tag
		tagName, err = d.readString()
	}
	return
}

func (d *Decoder) readByte() (int8, error) {
	b, err := d.r.ReadByte()
	// TagByte is signed byte (that's what in Java), so we need to convert to int8
	return int8(b), err
}

func (d *Decoder) readShort() (int16, error) {
	var data [2]byte
	_, err := io.ReadFull(d.r, data[:])
	return int16(data[0])<<8 | int16(data[1]), err
}

func (d *Decoder) readInt() (int32, error) {
	var data [4]byte
	_, err := io.ReadFull(d.r, data[:])
	return int32(data[0])<<24 | int32(data[1])<<16 |
		int32(data[2])<<8 | int32(data[3]), err
}

func (d *Decoder) readLong() (int64, error) {
	var data [8]byte
	_, err := io.ReadFull(d.r, data[:])
	return int64(data[0])<<56 | int64(data[1])<<48 |
		int64(data[2])<<40 | int64(data[3])<<32 |
		int64(data[4])<<24 | int64(data[5])<<16 |
		int64(data[6])<<8 | int64(data[7]), err
}

func (d *Decoder) readString() (string, error) {
	length, err := d.readShort()
	if err != nil {
		return "", err
	} else if length < 0 {
		return "", errors.New("string length less than 0")
	}

	var str string
	if length > 0 {
		buf := make([]byte, length)
		_, err = io.ReadFull(d.r, buf)
		str = string(buf)
	}
	return str, err
}

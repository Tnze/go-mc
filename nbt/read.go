package nbt

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"reflect"
)

func Unmarshal(data []byte, v interface{}) error {
	return NewDecoder(bytes.NewReader(data)).Decode(v)
}

func (d *Decoder) Decode(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("nbt: non-pointer passed to Unmarshal")
	}

	//start read NBT
	tagType, tagName, err := d.readTag()
	if err != nil {
		return fmt.Errorf("nbt: %w", err)
	}

	if c := d.checkCompressed(tagType); c != "" {
		return fmt.Errorf("nbt: unknown Tag, maybe need %s", c)
	}

	err = d.unmarshal(val.Elem(), tagType, tagName)
	if err != nil {
		return fmt.Errorf("nbt: fail to decode tag %q: %w", tagName, err)
	}
	return nil
}

// check the first byte and return if it use compress
func (d *Decoder) checkCompressed(head byte) (compress string) {
	if head == 0x1f { //gzip
		compress = "gzip"
	} else if head == 0x78 { //zlib
		compress = "zlib"
	}
	return
}

// ErrEND error will be returned when reading a NBT with only Tag_End
var ErrEND = errors.New("unexpected TAG_End")

func (d *Decoder) unmarshal(val reflect.Value, tagType byte, tagName string) error {
	if val.CanInterface() {
		if i, ok := val.Interface().(Unmarshaler); ok {
			return i.Unmarshal(tagType, tagName, d.r)
		}
	}

	switch tagType {
	default:
		return fmt.Errorf("unknown Tag 0x%02x", tagType)
	case TagEnd:
		return ErrEND

	case TagByte:
		value, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagByte as " + vk.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val.SetUint(uint64(value))
		case reflect.Interface:
			val.Set(reflect.ValueOf(value))
		}

	case TagShort:
		value, err := d.readInt16()
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
		value, err := d.readInt32()
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
		vInt, err := d.readInt32()
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
		value, err := d.readInt64()
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
		vInt, err := d.readInt64()
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
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagString as " + vk.String())
		case reflect.String:
			val.SetString(s)
		case reflect.Interface:
			val.Set(reflect.ValueOf(s))
		}

	case TagByteArray:
		aryLen, err := d.readInt32()
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

		switch vt := val.Type(); {
		default:
			return errors.New("cannot parse TagByteArray to " + vt.String() + ", use []byte in this instance")
		case vt == reflect.TypeOf(ba):
			val.SetBytes(ba)
		case vt.Kind() == reflect.Interface:
			val.Set(reflect.ValueOf(ba))
		}

	case TagIntArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		vt := val.Type() //receiver must be []int or []int32
		if vt.Kind() == reflect.Interface {
			vt = reflect.TypeOf([]int32{}) // pass
		} else if vt.Kind() != reflect.Slice {
			return errors.New("cannot parse TagIntArray to " + vt.String() + ", it must be a slice")
		} else if tk := val.Type().Elem().Kind(); tk != reflect.Int && tk != reflect.Int32 {
			return errors.New("cannot parse TagIntArray to " + vt.String())
		}

		buf := reflect.MakeSlice(vt, int(aryLen), int(aryLen))
		for i := 0; i < int(aryLen); i++ {
			value, err := d.readInt32()
			if err != nil {
				return err
			}
			buf.Index(i).SetInt(int64(value))
		}
		val.Set(buf)

	case TagLongArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		vt := val.Type() //receiver must be []int or []int64
		if vt.Kind() == reflect.Interface {
			vt = reflect.TypeOf([]int64{}) // pass
		} else if vt.Kind() != reflect.Slice {
			return errors.New("cannot parse TagLongArray to " + vt.String() + ", it must be a slice")
		} else if val.Type().Elem().Kind() != reflect.Int64 {
			return errors.New("cannot parse TagLongArray to " + vt.String())
		}

		buf := reflect.MakeSlice(vt, int(aryLen), int(aryLen))
		for i := 0; i < int(aryLen); i++ {
			value, err := d.readInt64()
			if err != nil {
				return err
			}
			buf.Index(i).SetInt(value)
		}
		val.Set(buf)

	case TagList:
		listType, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		listLen, err := d.readInt32()
		if err != nil {
			return err
		}
		if listLen < 0 {
			return errors.New("list length less than 0")
		}

		// If we need parse TAG_List into slice, make a new with right length.
		// Otherwise if we need parse into array, we check if len(array) are enough.
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
					"TagList %s has len %d, but array %v only has len %d",
					tagName, listLen, val.Type(), vl)
			}
			buf = val
		}
		for i := 0; i < int(listLen); i++ {
			if err := d.unmarshal(buf.Index(i), listType, ""); err != nil {
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
					err = d.unmarshal(val.Field(field), tt, tn)
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
				if err = d.unmarshal(v.Elem(), tt, tn); err != nil {
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
				if err = d.unmarshal(reflect.ValueOf(&value).Elem(), tt, tn); err != nil {
					return fmt.Errorf("fail to decode tag %q: %w", tn, err)
				}
				buf[tn] = value
			}
			val.Set(reflect.ValueOf(buf))
		}
	}

	return nil
}

func (d *Decoder) rawRead(tagType byte) error {
	var buf [8]byte
	switch tagType {
	default:
		return fmt.Errorf("unknown to read 0x%02x", tagType)
	case TagByte:
		_, err := d.r.ReadByte()
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
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}

		if _, err = io.CopyN(ioutil.Discard, d.r, int64(aryLen)); err != nil {
			return err
		}
	case TagIntArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		for i := 0; i < int(aryLen); i++ {
			if _, err := d.readInt32(); err != nil {
				return err
			}
		}

	case TagLongArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		for i := 0; i < int(aryLen); i++ {
			if _, err := d.readInt64(); err != nil {
				return err
			}
		}

	case TagList:
		listType, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		listLen, err := d.readInt32()
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

func (d *Decoder) readInt16() (int16, error) {
	var data [2]byte
	_, err := io.ReadFull(d.r, data[:])
	return int16(data[0])<<8 | int16(data[1]), err
}

func (d *Decoder) readInt32() (int32, error) {
	var data [4]byte
	_, err := io.ReadFull(d.r, data[:])
	return int32(data[0])<<24 | int32(data[1])<<16 |
		int32(data[2])<<8 | int32(data[3]), err
}

func (d *Decoder) readInt64() (int64, error) {
	var data [8]byte
	_, err := io.ReadFull(d.r, data[:])
	return int64(data[0])<<56 | int64(data[1])<<48 |
		int64(data[2])<<40 | int64(data[3])<<32 |
		int64(data[4])<<24 | int64(data[5])<<16 |
		int64(data[6])<<8 | int64(data[7]), err
}

func (d *Decoder) readString() (string, error) {
	length, err := d.readInt16()
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

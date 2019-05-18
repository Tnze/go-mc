package nbt

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"reflect"
)

func Unmarshal(data []byte, v interface{}) error {
	return NewDecoder(bytes.NewReader(data)).Decode(v)
}

func (d *Decoder) Decode(v interface{}) error {
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("non-pointer passed to Unmarshal")
	}
	tagType, tagName, err := d.readTag()
	if err != nil {
		return err
	}
	return d.unmarshal(val.Elem(), tagType, tagName)
}

func (d *Decoder) unmarshal(val reflect.Value, tagType byte, tagName string) error {
	switch tagType {
	default:
		return fmt.Errorf("unknown Tag 0x%02x", tagType)

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
		}

	case TagShort:
		value, err := d.readInt16()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagShort as " + vk.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val.SetUint(uint64(value))
		}

	case TagInt:
		value, err := d.readInt32()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagInt as " + vk.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val.SetUint(uint64(value))
		}

	case TagFloat:
		if vk := val.Kind(); vk != reflect.Float32 {
			return errors.New("cannot parse TagFloat as " + vk.String())
		}
		vInt, err := d.readInt32()
		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(math.Float32frombits(uint32(vInt))))

	case TagLong:
		value, err := d.readInt64()
		if err != nil {
			return err
		}
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagLong as " + vk.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val.SetInt(int64(value))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val.SetUint(uint64(value))
		}

	case TagDouble:
		if vk := val.Kind(); vk != reflect.Float64 {
			return errors.New("cannot parse TagDouble as " + vk.String())
		}
		vInt, err := d.readInt64()
		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(math.Float64frombits(uint64(vInt))))

	case TagString:
		if vk := val.Kind(); vk != reflect.String {
			return errors.New("cannot parse TagString as " + vk.String())
		}
		s, err := d.readString()
		if err != nil {
			return err
		}
		val.SetString(s)

	case TagByteArray:
		var ba []byte
		if vt := val.Type(); vt != reflect.TypeOf(ba) {
			return errors.New("cannot parse TagByteArray to " + vt.String() + ", use []byte in this instance")
		}
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		if ba, err = d.readNByte(int(aryLen)); err != nil {
			return err
		}
		val.SetBytes(ba)

	case TagList:
		listType, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		listLen, err := d.readInt32()
		if err != nil {
			return err
		}
		// If we need parse TAG_List into slice, make a new with right length.
		// Otherwise if we need parse into array, we check if len(array) are enough.
		switch vk := val.Kind(); vk {
		default:
			return errors.New("cannot parse TagList as " + vk.String())
		case reflect.Slice:
			val.Set(reflect.MakeSlice(val.Type(), int(listLen), int(listLen)))
		case reflect.Array:
			if vl := val.Len(); vl < int(listLen) {
				return fmt.Errorf(
					"TagList %s has len %d, but array %v only has len %d",
					tagName, listLen, val.Type(), vl)
			}
		}
		for i := 0; i < int(listLen); i++ {
			if err := d.unmarshal(val.Index(i), listType, ""); err != nil {
				return err
			}
		}

	case TagCompound:
		if vk := val.Kind(); vk != reflect.Struct {
			return errors.New("cannot parse TagCompound as " + vk.String())
		}
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
					return err
				}
			} else {
				if err := d.skip(tt); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (d *Decoder) skip(tagType byte) error {
	switch tagType {
	default:
		return fmt.Errorf("unknown to skip 0x%02x", tagType)
	case TagByte:
		_, err := d.r.ReadByte()
		return err
	case TagString:
		_, err := d.readString()
		return err
	case TagShort:
		_, err := d.readNByte(2)
		return err
	case TagInt, TagFloat:
		_, err := d.readNByte(4)
		return err
	case TagLong, TagDouble:
		_, err := d.readNByte(8)
		return err
	case TagByteArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		if _, err = d.readNByte(int(aryLen)); err != nil {
			return err
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
			if err := d.skip(listType); err != nil {
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
			err = d.skip(tt)
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

func (d *Decoder) readNByte(n int) (buf []byte, err error) {
	buf = make([]byte, n)
	_, err = d.r.Read(buf) //what happend if (returned n) != (argument n) ?
	return
}

func (d *Decoder) readInt16() (int16, error) {
	data, err := d.readNByte(2)
	return int16(data[0])<<8 | int16(data[1]), err
}

func (d *Decoder) readInt32() (int32, error) {
	data, err := d.readNByte(4)
	return int32(data[0])<<24 | int32(data[1])<<16 |
		int32(data[2])<<8 | int32(data[3]), err
}

func (d *Decoder) readInt64() (int64, error) {
	data, err := d.readNByte(8)
	return int64(data[0])<<56 | int64(data[1])<<48 |
		int64(data[2])<<40 | int64(data[3])<<32 |
		int64(data[4])<<24 | int64(data[5])<<16 |
		int64(data[6])<<8 | int64(data[7]), err
}

func (d *Decoder) readString() (string, error) {
	length, err := d.readInt16()
	if err != nil {
		return "", err
	}
	buf, err := d.readNByte(int(length))
	return string(buf), err
}

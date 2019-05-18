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
	case TagFloat:
		if vk := val.Kind(); vk != reflect.Float32 {
			return errors.New("cannot parse TagString as " + vk.String())
		}
		vInt, err := d.readInt32()
		if err != nil {
			return err
		}
		val.Set(reflect.ValueOf(math.Float32frombits(uint32(vInt))))

	case TagString:
		if vk := val.Kind(); vk != reflect.String {
			return errors.New("cannot parse TagString as " + vk.String())
		}
		s, err := d.readString()
		if err != nil {
			return err
		}
		val.SetString(s)
	case TagCompound:
		if vk := val.Kind(); vk != reflect.Struct {
			return errors.New("cannot parse TagCompound as " + vk.String())
		}
		fmt.Println("TagName:", tagName)
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

	if tagType != TagEnd && !d.nameless { //Read Tag
		tagName, err = d.readString()
	}
	return
}

func (d *Decoder) readNByte(n int) (buf []byte, err error) {
	buf = make([]byte, n)
	_, err = d.r.Read(buf) //what happend if (returned n) != (argument n) ?
	// for i := 0; i < n; i++ {
	// 	buf[i], err = d.r.ReadByte()
	// 	if err != nil {
	// 		return
	// 	}
	// }
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

func (d *Decoder) readString() (string, error) {
	length, err := d.readInt16()
	if err != nil {
		return "", err
	}
	buf, err := d.readNByte(int(length))
	return string(buf), err
}

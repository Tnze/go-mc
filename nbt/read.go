package nbt

import (
	"bytes"
	"errors"
	"fmt"
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
	return d.unmarshal(val.Elem())
}

// func (d *Decoder) unmarshalInt(val reflect.Value) error {

// }

func (d *Decoder) unmarshal(val reflect.Value) error {
	tagType, err := d.r.ReadByte()
	if err != nil {
		return err
	}

	var tagName string
	if !d.nameless { //Read Tag
		if tagName, err = d.readString(); err != nil {
			return err
		}
	}

	switch tagType {
	default:
		return fmt.Errorf("unknown tag 0x%2x", tagType)
	case TagString:
		if val.Kind() != reflect.String {
			return errors.New("cannot unmarshal TAG_String into " + val.Kind().String())
		}

		s, err := d.readString()
		if err != nil {
			return err
		}

		val.SetString(s)
	case TagCompound:
		if val.Kind() != reflect.Struct {
			return errors.New("cannot unmarshal TAG_Compound into " + val.Kind().String())
		}
		fmt.Println("TagName:", tagName)
	}

	return nil
}

func (d *Decoder) readNByte(n int) (buf []byte, err error) {
	buf = make([]byte, n)
	_, err = d.r.Read(buf) //what happend if (returned n) != (argument n) ?
	return
}

func (d *Decoder) readInt16() (int16, error) {
	data, err := d.readNByte(2)
	return int16(data[0])<<4 | int16(data[1]), err
}

func (d *Decoder) readString() (string, error) {
	length, err := d.readInt16()
	if err != nil {
		return "", err
	}
	buf, err := d.readNByte(int(length))
	return string(buf), err
}

package packet

import (
	"reflect"
)

type Ary struct {
	Len Field
	Ary interface{}
}

func (a Ary) Encode() (data []byte) {
	length := int(reflect.ValueOf(a.Len).Int())
	array := reflect.ValueOf(a.Ary).Elem()
	for i := 0; i < length; i++ {
		elem := array.Index(i)
		data = append(data, elem.Interface().(FieldEncoder).Encode()...)
	}
	return
}

func (a Ary) Decode(r DecodeReader) error {
	length := int(reflect.ValueOf(a.Len).Int())
	array := reflect.ValueOf(a.Ary).Elem()
	for i := 0; i < length; i++ {
		elem := array.Index(i)
		if err := elem.Interface().(FieldDecoder).Decode(r); err != nil {
			return err
		}
	}
	return nil
}

type Opt struct {
	Has   func() bool
	Field interface{}
}

func (o Opt) Encode() []byte {
	if o.Has() {
		return nil
	}
	return o.Field.(FieldEncoder).Encode()
}

func (o Opt) Decode(r DecodeReader) error {
	if o.Has() {
		return nil
	}
	return o.Field.(FieldDecoder).Decode(r)
}

package packet

import (
	"io"
	"reflect"
)

type Ary struct {
	Len Field       // Pointer of VarInt, VarLong, Int or Long
	Ary interface{} // Slice of FieldEncoder, FieldDecoder or both (Field)
}

func (a Ary) WriteTo(r io.Writer) (n int64, err error) {
	length := int(reflect.ValueOf(a.Len).Int())
	array := reflect.ValueOf(a.Ary)
	for i := 0; i < length; i++ {
		elem := array.Index(i)
		nn, err := elem.Interface().(FieldEncoder).WriteTo(r)
		n += nn
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (a Ary) ReadFrom(r io.Reader) (n int64, err error) {
	length := int(reflect.ValueOf(a.Len).Elem().Int())
	array := reflect.ValueOf(a.Ary).Elem()
	if array.Cap() < length {
		array.Set(reflect.MakeSlice(array.Type(), length, length))
	}
	for i := 0; i < length; i++ {
		elem := array.Index(i)
		nn, err := elem.Addr().Interface().(FieldDecoder).ReadFrom(r)
		n += nn
		if err != nil {
			return n, err
		}
	}
	return n, err
}

type Opt struct {
	Has   *Boolean
	Field interface{} // FieldEncoder, FieldDecoder or both (Field)
}

func (o Opt) WriteTo(w io.Writer) (int64, error) {
	if *o.Has {
		return o.Field.(FieldEncoder).WriteTo(w)
	}
	return 0, nil
}

func (o Opt) ReadFrom(r io.Reader) (int64, error) {
	if *o.Has {
		return o.Field.(FieldDecoder).ReadFrom(r)
	}
	return 0, nil
}

type Tuple []interface{} // FieldEncoder, FieldDecoder or both (Field)

// WriteTo write Tuple to io.Writer, panic when any of filed don't implement FieldEncoder
func (t Tuple) WriteTo(w io.Writer) (n int64, err error) {
	for _, v := range t {
		nn, err := v.(FieldEncoder).WriteTo(w)
		if err != nil {
			return n, err
		}
		n += nn
	}
	return
}

// ReadFrom read Tuple from io.Reader, panic when any of field don't implement FieldDecoder
func (t Tuple) ReadFrom(r io.Reader) (n int64, err error) {
	for _, v := range t {
		nn, err := v.(FieldDecoder).ReadFrom(r)
		if err != nil {
			return n, err
		}
		n += nn
	}
	return
}

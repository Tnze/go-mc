package packet

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

// Ary is used to send or receive the packet field like "Array of X"
// which has a count must be known from the context.
//
// Typically, you must decode an integer representing the length. Then
// receive the corresponding amount of data according to the length.
// In this case, the field Len should be a pointer of integer type so
// the value can be updating when Packet.Scan() method is decoding the
// previous field.
// In some special cases, you might want to read an "Array of X" with a fix length.
// So it's allowed to directly set an integer type Len, but not a pointer.
//
// Note that Ary DO read or write the Len. You aren't need to do so by your self.
type Ary[T VarInt | VarLong | Byte | UnsignedByte | Short | UnsignedShort | Int | Long] struct {
	Ary interface{} // Slice or Pointer of Slice of FieldEncoder, FieldDecoder or both (Field)
}

func (a Ary[T]) WriteTo(w io.Writer) (n int64, err error) {
	array := reflect.ValueOf(a.Ary)
	for array.Kind() == reflect.Ptr {
		array = array.Elem()
	}
	Len := T(array.Len())
	if nn, err := any(&Len).(FieldEncoder).WriteTo(w); err != nil {
		return n, err
	} else {
		n += nn
	}
	for i := 0; i < array.Len(); i++ {
		elem := array.Index(i)
		nn, err := elem.Interface().(FieldEncoder).WriteTo(w)
		n += nn
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (a Ary[T]) ReadFrom(r io.Reader) (n int64, err error) {
	var Len T
	if nn, err := any(&Len).(FieldDecoder).ReadFrom(r); err != nil {
		return nn, err
	} else {
		n += nn
	}

	array := reflect.ValueOf(a.Ary)
	for array.Kind() == reflect.Ptr {
		array = array.Elem()
	}
	if !array.CanAddr() {
		panic(errors.New("the contents of the Ary are not addressable"))
	}
	if array.Cap() < int(Len) {
		array.Set(reflect.MakeSlice(array.Type(), int(Len), int(Len)))
	} else {
		array.Slice(0, int(Len))
	}
	for i := 0; i < int(Len); i++ {
		elem := array.Index(i)
		nn, err := elem.Addr().Interface().(FieldDecoder).ReadFrom(r)
		n += nn
		if err != nil {
			return n, err
		}
	}
	return n, err
}

func Array(ary any) Field {
	return Ary[VarInt]{Ary: ary}
}

type Opt struct {
	Has   interface{} // Pointer of bool, or `func() bool`
	Field interface{} // FieldEncoder, FieldDecoder or both (Field)
}

func (o Opt) has() bool {
	v := reflect.ValueOf(o.Has)
	for {
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		case reflect.Bool:
			return v.Bool()
		case reflect.Func:
			return v.Interface().(func() bool)()
		default:
			panic(errors.New("unsupported Has value"))
		}
	}
}

func (o Opt) WriteTo(w io.Writer) (int64, error) {
	if o.has() {
		return o.Field.(FieldEncoder).WriteTo(w)
	}
	return 0, nil
}

func (o Opt) ReadFrom(r io.Reader) (int64, error) {
	if o.has() {
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
	for i, v := range t {
		nn, err := v.(FieldDecoder).ReadFrom(r)
		if err != nil {
			return n, fmt.Errorf("decode tuple[%d] error: %w", i, err)
		}
		n += nn
	}
	return
}

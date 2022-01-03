package packet

import (
	"errors"
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
// Note that Ary DO NOT read or write the Len. You are controlling it manually.
type Ary struct {
	Len interface{} // Value or Pointer of any integer type, only needed in ReadFrom
	Ary interface{} // Slice or Pointer of Slice of FieldEncoder, FieldDecoder or both (Field)
}

func (a Ary) WriteTo(r io.Writer) (n int64, err error) {
	array := reflect.ValueOf(a.Ary)
	for array.Kind() == reflect.Ptr {
		array = array.Elem()
	}
	for i := 0; i < array.Len(); i++ {
		elem := array.Index(i)
		nn, err := elem.Interface().(FieldEncoder).WriteTo(r)
		n += nn
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

func (a Ary) length() int {
	v := reflect.ValueOf(a.Len)
	for {
		switch v.Kind() {
		case reflect.Ptr:
			v = v.Elem()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return int(v.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return int(v.Uint())
		default:
			panic(errors.New("unsupported Len value: " + v.Type().String()))
		}
	}
}

func (a Ary) ReadFrom(r io.Reader) (n int64, err error) {
	length := a.length()
	array := reflect.ValueOf(a.Ary)
	for array.Kind() == reflect.Ptr {
		array = array.Elem()
	}
	if !array.CanAddr() {
		panic(errors.New("the contents of the Ary are not addressable"))
	}
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

// Array return an Ary but handled the previous Length field
//
// Warning: unstable API, may change in later version
func Array(array interface{}) Field {
	var length VarInt

	value := reflect.ValueOf(array)
	for value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	if array != nil {
		length = VarInt(value.Len())
	}
	return Tuple{
		&length,
		Ary{
			Len: &length,
			Ary: array,
		},
	}
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
	for _, v := range t {
		nn, err := v.(FieldDecoder).ReadFrom(r)
		if err != nil {
			return n, err
		}
		n += nn
	}
	return
}

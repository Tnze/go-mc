package packet

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

var (
	_ Field = (*Option[VarInt, *VarInt])(nil)
	_ Field = (*Ary[VarInt])(nil)
	_ Field = Tuple(nil)
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
type Ary[LEN VarInt | VarLong | Byte | UnsignedByte | Short | UnsignedShort | Int | Long] struct {
	Ary any // Slice or Pointer of Slice of FieldEncoder, FieldDecoder or both (Field)
}

func (a Ary[LEN]) WriteTo(w io.Writer) (n int64, err error) {
	array := reflect.ValueOf(a.Ary)
	for array.Kind() == reflect.Ptr {
		array = array.Elem()
	}
	Len := LEN(array.Len())
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

func (a Ary[LEN]) ReadFrom(r io.Reader) (n int64, err error) {
	var Len LEN
	if nn, err := any(&Len).(FieldDecoder).ReadFrom(r); err != nil {
		return nn, err
	} else {
		n += nn
	}
	if Len < 0 {
		return n, errors.New("array length less than zero")
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

// Opt is an optional [Field] which sending/receiving or not is depending on its Has field.
// When calling `WriteTo()` or `ReadFrom()`, if Has is true, the Field's `WriteTo` or `ReadFrom()` is called.
// Otherwise, it does nothing and return 0 and nil.
//
// The different between [Opt] and [Option] is that [Opt] does NOT read or write the Has field for you.
// Which should be cared.
type Opt struct {
	Has   any // Pointer of bool, or `func() bool`
	Field any // FieldEncoder, FieldDecoder, `func() FieldEncoder`, `func() FieldDecoder` or `func() Field`
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
		switch field := o.Field.(type) {
		case FieldEncoder:
			return field.WriteTo(w)
		case func() FieldEncoder:
			return field().WriteTo(w)
		case func() Field:
			return field().WriteTo(w)
		default:
			panic("unsupported Field type: " + reflect.TypeOf(o.Field).String())
		}
	}
	return 0, nil
}

func (o Opt) ReadFrom(r io.Reader) (int64, error) {
	if o.has() {
		switch field := o.Field.(type) {
		case FieldDecoder:
			return field.ReadFrom(r)
		case func() FieldDecoder:
			return field().ReadFrom(r)
		case func() Field:
			return field().ReadFrom(r)
		default:
			panic("unsupported Field type: " + reflect.TypeOf(o.Field).String())
		}
	}
	return 0, nil
}

type fieldPointer[T any] interface {
	*T
	FieldDecoder
}

// Option is a helper type for encoding/decoding these kind of packet:
//
//	+-----------+------------+----------------------------------- +
//	| Name      | Type       | Notes                              |
//	+-----------+------------+------------------------------------+
//	| Has Value | Boolean    | Whether the Value should be sent.  |
//	+-----------+------------+------------------------------------+
//	| Value     | Optional T | Only exist when Has Value is true. |
//	+-----------+------------+------------------------------------+
//
// # Usage
//
// `Option[T]` implements [FieldEncoder] and `*Option[T]` implements [FieldDecoder].
// That is, you can call `WriteTo()` and `ReadFrom()` methods on it.
//
//	var optStr Option[String]
//	n, err := optStr.ReadFrom(r)
//	if err != nil {
//		// ...
//	}
//	if optStr.Has {
//		fmt.Println(optStr.Val)
//	}
//
// # Notes
//
// Currently we have to repeat T in the type arguments.
//
//	var opt Option[String, *String]
//
// Constraint type will inference makes it less awkward in the future.
// See: https://github.com/golang/go/issues/54469
type Option[T FieldEncoder, P fieldPointer[T]] struct {
	Has Boolean
	Val T
}

func (o Option[T, P]) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := o.Has.WriteTo(w)
	if err != nil || !o.Has {
		return n1, err
	}
	n2, err := o.Val.WriteTo(w)
	return n1 + n2, err
}

func (o *Option[T, P]) ReadFrom(r io.Reader) (n int64, err error) {
	n1, err := o.Has.ReadFrom(r)
	if err != nil || !o.Has {
		return n1, err
	}
	n2, err := P(&o.Val).ReadFrom(r)
	return n1 + n2, err
}

// Pointer returns the pointer of Val if Has is true, otherwise return nil.
func (o *Option[T, P]) Pointer() (p *T) {
	if o.Has {
		p = &o.Val
	}
	return
}

// OptionDecoder is basically same with [Option], but support [FieldDecoder] only.
// This allowed wrapping a [FieldDecoder] type (which isn't a [FieldEncoder]) to an Option.
type OptionDecoder[T any, P fieldPointer[T]] struct {
	Has Boolean
	Val T
}

func (o *OptionDecoder[T, P]) ReadFrom(r io.Reader) (n int64, err error) {
	n1, err := o.Has.ReadFrom(r)
	if err != nil || !o.Has {
		return n1, err
	}
	n2, err := P(&o.Val).ReadFrom(r)
	return n1 + n2, err
}

// OptionEncoder is basically same with [Option], but support [FieldEncoder] only.
// This allowed wrapping a [FieldEncoder] type (which isn't a [FieldDecoder]) to an Option.
type OptionEncoder[T FieldEncoder] struct {
	Has Boolean
	Val T
}

func (o OptionEncoder[T]) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := o.Has.WriteTo(w)
	if err != nil || !o.Has {
		return n1, err
	}
	n2, err := o.Val.WriteTo(w)
	return n1 + n2, err
}

type Tuple []any // FieldEncoder, FieldDecoder or both (Field)

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
			return n, fmt.Errorf("decode tuple[%d] %T error: %w", i, v, err)
		}
		n += nn
	}
	return
}

func CreateByteReader(reader io.Reader) io.ByteReader {
	if byteReader, isByteReader := reader.(io.ByteReader); isByteReader {
		return byteReader
	}
	return byteReaderWrapper{reader}
}

type byteReaderWrapper struct {
	io.Reader
}

var _ io.ByteReader = byteReaderWrapper{}

func (r byteReaderWrapper) ReadByte() (byte, error) {
	var buf [1]byte
	_, err := io.ReadFull(r.Reader, buf[:])
	return buf[0], err
}

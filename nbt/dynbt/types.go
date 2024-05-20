// Package dynbt is a library that provides dynamic NBT operation APIs.
//
// Dynamically represented NBT value is useful in many cases, for example,
//   - You want to store custom structural values at runtime.
//   - You want to query or modify the data later. (Otherwise use the [nbt.RawMessage])
//   - You don't know what type the data is at compile time. (Otherwise use the [nbt.RawMessage] too)
//
// The [*Value] provides a group of APIs on top of the [nbt] package.
// [*Value] implements [nbt.Marshaler] and [nbt.Unmarshaler] interfaces.
// It can be used as a field of struct, or element of slice, map, etc.
// The pointer type should always be used, unless used as fields for structures
//
// Notice that querying Tags in Compound use a linear search, so it's not recommended to use it in a large Compound.
// The better choice is map[string]*Value for dynamic accessing a large Compound.
//
// This package tries its best to not copy data if possible.
// It returns the underlying data in some cases. Don't modify them!
package dynbt

import (
	"encoding/binary"
	"math"

	"github.com/Tnze/go-mc/nbt"
)

type Value struct {
	comp Compound
	list []*Value
	data []byte
	tag  byte // nbt.Tag*
}

func NewBoolean(v bool) *Value {
	data := byte(0)
	if v {
		data = 1
	}
	return &Value{tag: nbt.TagByte, data: []byte{data}}
}

func NewByte(v int8) *Value {
	return &Value{tag: nbt.TagByte, data: []byte{byte(v)}}
}

func NewShort(v int16) *Value {
	data := make([]byte, 2)
	binary.BigEndian.PutUint16(data, uint16(v))
	return &Value{tag: nbt.TagShort, data: data}
}

func NewInt(v int32) *Value {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, uint32(v))
	return &Value{tag: nbt.TagInt, data: data}
}

func NewLong(v int64) *Value {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, uint64(v))
	return &Value{tag: nbt.TagLong, data: data}
}

func NewFloat(f float32) *Value {
	data := make([]byte, 4)
	binary.BigEndian.PutUint32(data, math.Float32bits(f))
	return &Value{tag: nbt.TagFloat, data: data}
}

func NewDouble(d float64) *Value {
	data := make([]byte, 8)
	binary.BigEndian.PutUint64(data, math.Float64bits(d))
	return &Value{tag: nbt.TagDouble, data: data}
}

func NewByteArray(v []byte) *Value {
	data := make([]byte, 4, 4+len(v))
	binary.BigEndian.PutUint32(data, uint32(len(v)))
	return &Value{tag: nbt.TagByteArray, data: append(data, v...)}
}

func NewIntArray(v []int32) *Value {
	data := make([]byte, 4+len(v)*4)
	binary.BigEndian.PutUint32(data, uint32(len(v)))
	for i, j := 0, 4; i < len(v); i, j = i+1, j+4 {
		binary.BigEndian.PutUint32(data[j:], uint32(v[i]))
	}
	return &Value{tag: nbt.TagIntArray, data: data}
}

func NewLongArray(v []int64) *Value {
	data := make([]byte, 4+len(v)*8)
	binary.BigEndian.PutUint32(data, uint32(len(v)))
	for i, j := 0, 4; i < len(v); i, j = i+1, j+8 {
		binary.BigEndian.PutUint64(data[j:], uint64(v[i]))
	}
	return &Value{tag: nbt.TagLongArray, data: data}
}

func NewString(str string) *Value {
	data := make([]byte, 2, 2+len(str))
	binary.BigEndian.PutUint16(data, uint16(len(str)))
	return &Value{tag: nbt.TagString, data: append(data, str...)}
}

func NewList(elems ...*Value) *Value {
	return &Value{tag: nbt.TagList, list: elems}
}

func NewCompound() *Value {
	return &Value{tag: nbt.TagCompound}
}

func (v *Value) Boolean() bool {
	if v.tag != nbt.TagByte {
		return false
	}
	return v.data[0] != 0
}

func (v *Value) Byte() int8 {
	if v.tag != nbt.TagByte {
		return 0
	}
	return int8(v.data[0])
}

func (v *Value) Short() int16 {
	if v.tag != nbt.TagShort {
		return 0
	}
	return int16(binary.BigEndian.Uint16(v.data))
}

func (v *Value) Int() int32 {
	if v.tag != nbt.TagInt {
		return 0
	}
	return int32(binary.BigEndian.Uint32(v.data))
}

func (v *Value) Long() int64 {
	if v.tag != nbt.TagLong {
		return 0
	}
	return int64(binary.BigEndian.Uint64(v.data))
}

func (v *Value) Float() float32 {
	if v.tag != nbt.TagFloat {
		return float32(math.NaN())
	}
	return math.Float32frombits(binary.BigEndian.Uint32(v.data))
}

func (v *Value) Double() float64 {
	if v.tag != nbt.TagDouble {
		return math.NaN()
	}
	return math.Float64frombits(binary.BigEndian.Uint64(v.data))
}

func (v *Value) List() []*Value {
	if v.tag != nbt.TagList {
		return nil
	}
	return v.list
}

func (v *Value) Compound() *Compound {
	if v.tag != nbt.TagCompound {
		return nil
	}
	return &v.comp
}

func (v *Value) ByteArray() []byte {
	if v.tag != nbt.TagByteArray {
		return nil
	}
	return v.data[4:]
}

func (v *Value) IntArray() []int32 {
	if v.tag != nbt.TagIntArray {
		return nil
	}
	length := binary.BigEndian.Uint32(v.data)
	ret := make([]int32, length)
	for i, j := uint32(0), 4; i < length; i, j = i+1, j+4 {
		ret[i] = int32(binary.BigEndian.Uint32(v.data[j:]))
	}
	return ret
}

func (v *Value) LongArray() []int64 {
	if v.tag != nbt.TagLongArray {
		return nil
	}
	length := binary.BigEndian.Uint32(v.data)
	ret := make([]int64, length)
	for i, j := uint32(0), 4; i < length; i, j = i+1, j+8 {
		ret[i] = int64(binary.BigEndian.Uint64(v.data[j:]))
	}
	return ret
}

func (v *Value) String() string {
	if v.tag != nbt.TagString {
		return ""
	}
	return string(v.data[2:])
}

type Compound struct {
	kvs []kv
}

type kv struct {
	tag string
	v   *Value
}

func (c *Compound) Visit(f func(tag string, v *Value)) {
	if c == nil {
		return
	}
	for _, kv := range c.kvs {
		f(kv.tag, kv.v)
	}
}

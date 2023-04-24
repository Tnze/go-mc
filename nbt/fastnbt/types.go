package fastnbt

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
	return v.list
}

func (v *Value) Compound() *Compound {
	return &v.comp
}

func (v *Value) ByteArray() []byte {
	if v.tag != nbt.TagByteArray {
		return nil
	}
	return v.data[4:]
}

func (v *Value) IntArray() []int32 {
	length := binary.BigEndian.Uint32(v.data)
	ret := make([]int32, length)
	for i := uint32(0); i < length; i += 4 {
		ret[i] = int32(binary.BigEndian.Uint32(v.data[i:]))
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

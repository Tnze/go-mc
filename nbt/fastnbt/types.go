package fastnbt

import (
	"math"

	"github.com/Tnze/go-mc/nbt"
)

type Value struct {
	comp Compound
	list []*Value
	data []byte
	tag  byte // nbt.Tag*
}

func (v *Value) Bool() bool {
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
	return int16(v.data[0])<<8 | int16(v.data[1])
}

func (v *Value) Int() int32 {
	if v.tag != nbt.TagInt {
		return 0
	}
	return int32(v.data[0])<<24 | int32(v.data[1])<<16 |
		int32(v.data[2])<<8 | int32(v.data[3])
}

func (v *Value) Long() int64 {
	if v.tag != nbt.TagLong {
		return 0
	}
	return int64(v.data[0])<<56 | int64(v.data[1])<<48 |
		int64(v.data[2])<<40 | int64(v.data[3])<<32 |
		int64(v.data[4])<<24 | int64(v.data[5])<<16 |
		int64(v.data[6])<<8 | int64(v.data[7])
}

func (v *Value) Float() float32 {
	if v.tag != nbt.TagFloat {
		return 0
	}
	return math.Float32frombits(
		uint32(v.data[0])<<24 | uint32(v.data[1])<<16 |
			uint32(v.data[2])<<8 | uint32(v.data[3]))
}

func (v *Value) Double() float64 {
	if v.tag != nbt.TagDouble {
		return 0
	}
	return math.Float64frombits(
		uint64(v.data[0])<<56 | uint64(v.data[1])<<48 |
			uint64(v.data[2])<<40 | uint64(v.data[3])<<32 |
			uint64(v.data[4])<<24 | uint64(v.data[5])<<16 |
			uint64(v.data[6])<<8 | uint64(v.data[7]))
}

func (v *Value) List() []*Value {
	return v.list
}

func (v *Value) ByteArray() []byte {
	if v.tag != nbt.TagByteArray {
		return nil
	}
	return v.data[4:]
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

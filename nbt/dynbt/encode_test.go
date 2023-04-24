package dynbt

import (
	"bytes"
	"math/rand"
	"reflect"
	"testing"

	"github.com/Tnze/go-mc/nbt"
)

func TestValue_new(t *testing.T) {
	if val := NewBoolean(true); val.Boolean() != true {
		t.Error("encode bool error")
	}
	if val := NewBoolean(false); val.Boolean() != false {
		t.Error("encode bool error")
	}
	if val := NewByte(127); val.Byte() != 127 {
		t.Error("encode byte error")
	}
	if val := NewShort(32767); val.Short() != 32767 {
		t.Error("encode short error")
	}
	if val := NewInt(2147483647); val.Int() != 2147483647 {
		t.Error("encode int error")
	}
	if val := NewLong(9223372036854775807); val.Long() != 9223372036854775807 {
		t.Error("encode long error")
	}
	if val := NewString("HELLO WORLD THIS IS A TEST STRING ÅÄÖ!"); val.String() != "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!" {
		t.Error("encode string error")
	}
	if val := NewFloat(0.49823147); val.Float() != 0.49823147 {
		t.Error("encode float error")
	}
	if val := NewDouble(0.4931287132182315); val.Double() != 0.4931287132182315 {
		t.Error("encode double error")
	}

	byteArray := make([]byte, 1000)
	for n := range byteArray {
		byteArray[n] = byte((n*n*255 + n*7) % 100)
	}
	if val := NewByteArray(byteArray); !bytes.Equal(byteArray, val.ByteArray()) {
		t.Error("encode byteArray error")
	}

	intArray := make([]int32, 250)
	for n := range intArray {
		intArray[n] = rand.Int31()
	}
	if val := NewIntArray(intArray); !reflect.DeepEqual(intArray, val.IntArray()) {
		t.Error("encode intArray error")
	}

	longArray := make([]int64, 125)
	for n := range longArray {
		longArray[n] = rand.Int63()
	}
	if val := NewLongArray(longArray); !reflect.DeepEqual(longArray, val.LongArray()) {
		t.Error("encode longArray error")
	}

	val := NewCompound()
	val.Set("a", NewString("tnze"))
	if val.Get("a").String() != "tnze" || val.Compound().Len() != 1 {
		t.Error("encode compound error")
	}
}

func TestValue_bigTest(t *testing.T) {
	data, err := nbt.Marshal(nbt.StringifiedMessage(bigTestSNBT))
	if err != nil {
		t.Fatal(err)
	}

	var val Value
	err = nbt.Unmarshal(data, &val)
	if err != nil {
		t.Fatal(err)
	}

	var data2 []byte
	data2, err = nbt.Marshal(&val)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, data2) {
		t.Fail()
	}
}

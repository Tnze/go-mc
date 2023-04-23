package fastnbt

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/Tnze/go-mc/nbt"
)

//go:embed bigTest_test.snbt
var bigTestSNBT string

func TestValue_UnmarshalNBT(t *testing.T) {
	data, err := nbt.Marshal(nbt.StringifiedMessage(bigTestSNBT))
	if err != nil {
		t.Fatal(err)
	}

	var val Value
	err = nbt.Unmarshal(data, &val)
	if err != nil {
		t.Fatal(err)
	}

	if v := val.Get("longTest"); v == nil {
		t.Fail()
	} else if got, want := v.Long(), int64(9223372036854775807); got != want {
		t.Errorf("expect %v, got: %v", want, got)
	}

	if v := val.Get("shortTest"); v == nil {
		t.Fail()
	} else if got, want := v.Short(), int16(32767); got != want {
		t.Errorf("expect %v, got: %v", want, got)
	}

	if v := val.Get("stringTest"); v == nil {
		t.Fail()
	} else if got, want := v.String(), "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!"; got != want {
		t.Errorf("expect %s, got: %s", want, got)
	}

	if v := val.Get("floatTest"); v == nil {
		t.Fail()
	} else if got, want := v.Float(), float32(0.49823147); got != want {
		t.Errorf("expect %v, got: %v", want, got)
	}

	if v := val.Get("byteTest"); v == nil {
		t.Fail()
	} else if got, want := v.Byte(), int8(127); got != want {
		t.Errorf("expect %v, got: %v", want, got)
	}

	if v := val.Get("intTest"); v == nil {
		t.Fail()
	} else if got, want := v.Int(), int32(2147483647); got != want {
		t.Errorf("expect %v, got: %v", want, got)
	}

	if v := val.Get("nested compound test"); v == nil {
		t.Fail()
	} else if v = v.Get("ham"); v == nil {
		t.Fail()
	} else if v = v.Get("name"); v == nil {
		t.Fail()
	} else if got, want := v.String(), "Hampus"; got != want {
		t.Errorf("expect %v, got: %v", want, got)
	}

	if v := val.Get("nested compound test", "ham", "name"); v == nil {
		t.Fail()
	} else if got, want := v.String(), "Hampus"; got != want {
		t.Errorf("expect %v, got: %v", want, got)
	}

	if v := val.Get("listTest (long)"); v == nil {
		t.Fail()
	} else if list := v.List(); list == nil {
		t.Fail()
	} else if len(list) != 5 {
		t.Fail()
	} else if list[0].Long() != 11 || list[1].Long() != 12 || list[2].Long() != 13 || list[3].Long() != 14 || list[4].Long() != 15 {
		t.Fail()
	}

	want := make([]byte, 1000)
	for n := 0; n < 1000; n++ {
		want[n] = byte((n*n*255 + n*7) % 100)
	}
	if v := val.Get("byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"); v == nil {
		t.Fail()
	} else if got := v.ByteArray(); !bytes.Equal(got, want) {
		t.Errorf("expect %v", want)
		t.Errorf("  got: %v", got)
	}
}

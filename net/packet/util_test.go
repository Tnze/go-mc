package packet_test

import (
	"bytes"
	"fmt"
	pk "github.com/Tnze/go-mc/net/packet"
	"testing"
)

func ExampleAry_WriteTo() {
	data := []pk.Int{0, 1, 2, 3, 4, 5, 6}
	// Len is completely ignored by WriteTo method.
	// The length is inferred from the length of Ary.
	pk.Marshal(
		0x00,
		// It's important to remember that
		// typically the responsibility of
		// sending the length field
		// is on you.
		pk.VarInt(len(data)),
		pk.Ary{
			Len: len(data), // this line can be removed
			Ary: data,
		},
	)
}

func ExampleAry_ReadFrom() {
	var length pk.VarInt
	var data []pk.String

	var p pk.Packet // = conn.ReadPacket()
	if err := p.Scan(

		&length, // decode length first
		pk.Ary{ // then decode Ary according to length
			Len: &length,
			Ary: &data,
		},
	); err != nil {
		panic(err)
	}
}

func TestAry_ReadFrom(t *testing.T) {
	var num pk.Int = 2
	var ary []pk.String
	var bin = []byte{
		4, 'T', 'n', 'z', 'e',
		0,
	}
	var data = pk.Ary{Len: &num, Ary: &ary}
	if _, err := data.ReadFrom(bytes.NewReader(bin)); err != nil {
		t.Fatal(err)
	}
	if len(ary) != int(num) {
		t.Fatalf("length not match: %d != %d", len(ary), num)
	}
	for i, v := range []string{"Tnze", ""} {
		if string(ary[i]) != v {
			t.Errorf("want %q, get %q", v, ary[i])
		}
	}
}

func TestAry_WriteTo(t *testing.T) {
	var buf bytes.Buffer
	want := []byte{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}
	int3, long3, varint3, varlong3 := pk.Int(3), pk.Long(3), pk.VarInt(3), pk.VarLong(3)
	for _, item := range [...]pk.Ary{
		{Len: 3, Ary: []pk.Int{1, 2, 3}},
		{Len: int3, Ary: []pk.Int{1, 2, 3}},
		{Len: long3, Ary: []pk.Int{1, 2, 3}},
		{Len: varint3, Ary: []pk.Int{1, 2, 3}},
		{Len: varlong3, Ary: []pk.Int{1, 2, 3}},
		{Len: &int3, Ary: []pk.Int{1, 2, 3}},
		{Len: &long3, Ary: []pk.Int{1, 2, 3}},
		{Len: &varint3, Ary: []pk.Int{1, 2, 3}},
		{Len: &varlong3, Ary: []pk.Int{1, 2, 3}},
	} {
		_, err := item.WriteTo(&buf)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(buf.Bytes(), want) {
			t.Fatalf("Ary encoding error: got %#v, want %#v", buf.Bytes(), want)
		}
		buf.Reset()
	}
}

func TestAry_WriteTo_pointer(t *testing.T) {
	var buf bytes.Buffer
	want := []byte{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}
	data := pk.Ary{Ary: &[]pk.Int{1, 2, 3}}

	_, err := data.WriteTo(&buf)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(buf.Bytes(), want) {
		t.Fatalf("Ary encoding error: got %#v, want %#v", buf.Bytes(), want)
	}
}

func ExampleOpt_ReadFrom() {
	var has pk.Boolean

	var data pk.String
	p1 := pk.Packet{Data: []byte{
		0x01,                  // pk.Boolean(true)
		4, 'T', 'n', 'z', 'e', // pk.String
	}}
	if err := p1.Scan(
		&has,
		pk.Opt{
			Has: &has, Field: &data,
		},
	); err != nil {
		panic(err)
	}
	fmt.Println(data)

	var data2 pk.String = "WILL NOT BE READ, WILL NOT BE COVERED"
	p2 := pk.Packet{Data: []byte{
		0x00, // pk.Boolean(false)
		// empty
	}}
	if err := p2.Scan(
		&has,
		pk.Opt{
			Has: &has, Field: &data2,
		},
	); err != nil {
		panic(err)
	}
	fmt.Println(data2)

	// Output:
	// Tnze
	// WILL NOT BE READ, WILL NOT BE COVERED
}

func ExampleOpt_ReadFrom_func() {
	// As an example, we define this packet as this:
	// +------+-----------------+----------------------------------+
	// | Name | Type            | Notes                            |
	// +------+-----------------+----------------------------------+
	// | Flag | Unsigned Byte   | Odd if the following is present. |
	// +------+-----------------+----------------------------------+
	// | User | Optional String | The player's name.               |
	// +------+-----------------+----------------------------------+
	// So we need a function to decide if the User field is present.
	var flag pk.Byte
	var data pk.String
	p := pk.Packet{Data: []byte{
		0b_0010_0011,          // pk.Byte(flag)
		4, 'T', 'n', 'z', 'e', // pk.String
	}}
	if err := p.Scan(
		&flag,
		pk.Opt{
			Has: func() bool {
				return flag&1 != 0
			},
			Field: &data,
		},
	); err != nil {
		panic(err)
	}
	fmt.Println(data)

	// Output: Tnze
}

func ExampleTuple_ReadFrom() {
	// When you need to read an "Optional Array of X":
	var has pk.Boolean
	var arylen pk.Int
	var ary []pk.String

	var p pk.Packet // = conn.ReadPacket()
	if err := p.Scan(
		&has,
		pk.Opt{
			Has: &has,
			Field: pk.Tuple{
				&arylen,
				pk.Ary{
					Len: &arylen,
					Ary: &ary,
				},
			},
		},
	); err != nil {
		panic(err)
	}
}

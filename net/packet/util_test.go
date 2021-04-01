package packet_test

import (
	"bytes"
	"testing"

	pk "github.com/Tnze/go-mc/net/packet"
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

func TestAry_WriteTo(t *testing.T) {
	var buf bytes.Buffer
	want := []byte{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}
	int3, long3, varint3, varlong3 := pk.Int(3), pk.Long(3), pk.VarInt(3), pk.VarLong(3)
	for _, item := range [...]pk.Ary{
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

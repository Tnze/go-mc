package packet_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

var VarInts = []pk.VarInt{0, 1, 2, 127, 128, 255, 2147483647, -1, -2147483648}

var PackedVarInts = [][]byte{
	{0x00},
	{0x01},
	{0x02},
	{0x7f},
	{0x80, 0x01},
	{0xff, 0x01},
	{0xff, 0xff, 0xff, 0xff, 0x07},
	{0xff, 0xff, 0xff, 0xff, 0x0f},
	{0x80, 0x80, 0x80, 0x80, 0x08},
}

func TestVarInt_WriteTo(t *testing.T) {
	var buf bytes.Buffer
	for i, v := range VarInts {
		buf.Reset()
		if n, err := v.WriteTo(&buf); err != nil {
			t.Fatalf("Write to bytes.Buffer should never fail: %v", err)
		} else if n != int64(buf.Len()) {
			t.Errorf("Number of byte returned by WriteTo should equal to buffer.Len()")
		}
		if p := buf.Bytes(); !bytes.Equal(p, PackedVarInts[i]) {
			t.Errorf("pack int %d should be \"% x\", get \"% x\"", v, PackedVarInts[i], p)
		}
	}
}

func TestVarInt_ReadFrom(t *testing.T) {
	for i, v := range PackedVarInts {
		var vi pk.VarInt
		if _, err := vi.ReadFrom(bytes.NewReader(v)); err != nil {
			t.Errorf("unpack \"% x\" error: %v", v, err)
		}
		if vi != VarInts[i] {
			t.Errorf("unpack \"% x\" should be %d, get %d", v, VarInts[i], vi)
		}
	}
}

func TestVarInt_ReadFrom_tooLongData(t *testing.T) {
	var vi pk.VarInt
	data := []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}
	if _, err := vi.ReadFrom(bytes.NewReader(data)); err != nil {
		t.Logf("unpack \"% x\" error: %v", data, err)
	} else {
		t.Errorf("unpack \"% x\" should be error, get %d", data, vi)
	}
}

func FuzzVarInt_Len(f *testing.F) {
	for _, v := range VarInts {
		f.Add(int32(v))
	}
	var buf bytes.Buffer
	f.Fuzz(func(t *testing.T, v int32) {
		defer buf.Reset()
		if _, err := pk.VarInt(v).WriteTo(&buf); err != nil {
			t.Fatal(err)
		}
		if a, b := buf.Len(), pk.VarInt(v).Len(); a != b {
			t.Errorf("VarInt(%d) Length calculation error: calculated to be %d, actually %d", v, b, a)
		}
	})
}

var VarLongs = []pk.VarLong{0, 1, 2, 127, 128, 255, 2147483647, 9223372036854775807, -1, -2147483648, -9223372036854775808}

var PackedVarLongs = [][]byte{
	{0x00},
	{0x01},
	{0x02},
	{0x7f},
	{0x80, 0x01},
	{0xff, 0x01},
	{0xff, 0xff, 0xff, 0xff, 0x07},
	{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
	{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01},
	{0x80, 0x80, 0x80, 0x80, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x01},
	{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01},
}

func TestVarLong_WriteTo(t *testing.T) {
	var buf bytes.Buffer
	for i, v := range VarLongs {
		buf.Reset()
		if _, err := v.WriteTo(&buf); err != nil {
			t.Error(err)
		}
		if !bytes.Equal(buf.Bytes(), PackedVarLongs[i]) {
			t.Errorf("pack long %d should be \"% x\", get \"% x\"", v, PackedVarLongs[i], buf.Bytes())
		}
	}
}

func TestVarLong_ReadFrom(t *testing.T) {
	for i, v := range PackedVarLongs {
		var vi pk.VarLong
		if _, err := vi.ReadFrom(bytes.NewReader(v)); err != nil {
			t.Errorf("unpack \"% x\" error: %v", v, err)
		}
		if vi != VarLongs[i] {
			t.Errorf("unpack \"% x\" should be %d, get %d", v, VarLongs[i], vi)
		}
	}
}

func FuzzVarLong_Len(f *testing.F) {
	for _, v := range VarLongs {
		f.Add(int64(v))
	}
	var buf bytes.Buffer
	f.Fuzz(func(t *testing.T, v int64) {
		defer buf.Reset()
		if _, err := pk.VarLong(v).WriteTo(&buf); err != nil {
			t.Fatal(err)
		}
		if a, b := buf.Len(), pk.VarLong(v).Len(); a != b {
			t.Errorf("VarLong(%d) Length calculation error: calculated to be %d, actually %d", v, b, a)
		}
	})
}

func ExampleNBT() {
	type T struct {
		Name string `nbt:"name"`
	}
	var send, recv T
	send.Name = "Tnze"

	// Example of how pk.NBT() is used in pk.Marshal()
	p := pk.Marshal(
		packetid.ServerboundPacketID(0),
		//...
		pk.NBT(send),           // without tag name
		pk.NBT(send, "player"), // with tag name
		//...
	)
	fmt.Println("Marshal:")
	fmt.Println(hex.Dump(p.Data))

	// Example of how pk.NBT() is used in pk.Packet.Scan()
	_ = p.Scan(
		//...
		pk.NBT(&recv),
		// pk.NBT(&recv) // The tag name are going to be ignored. To receive the tag name, pk.NBTField has to be used.
		//...
	)
	fmt.Println("Scan:", recv.Name)

	// Output:
	// Marshal:
	// 00000000  0a 00 00 08 00 04 6e 61  6d 65 00 04 54 6e 7a 65  |......name..Tnze|
	// 00000010  00 0a 00 06 70 6c 61 79  65 72 08 00 04 6e 61 6d  |....player...nam|
	// 00000020  65 00 04 54 6e 7a 65 00                           |e..Tnze.|
	//
	// Scan: Tnze
}

func TestNBTField_ReadFrom(t *testing.T) {
	var send struct {
		KnownField   int32
		UnknownField int32
	}
	var recv struct {
		KnownField int32
	}

	p := pk.Marshal(
		packetid.ServerboundPacketID(0),
		pk.NBTField{V: send},
	)

	err := p.Scan(pk.NBTField{V: &recv})
	if err == nil {
		t.Errorf("should be a unknown field error here")
	}

	err = p.Scan(pk.NBT(&recv))
	if err == nil {
		t.Errorf("disallow unknown field by default")
	}

	err = p.Scan(pk.NBTField{V: &recv, AllowUnknownFields: true})
	if err != nil {
		t.Errorf("should allow the unknown field here: %v", err)
	}
}

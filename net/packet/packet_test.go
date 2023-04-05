package packet_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"testing"

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

//go:embed joingame_test.bin
var testJoinGameData []byte

func ExamplePacket_Scan_joinGame() {
	p := pk.Packet{ID: 0x24, Data: testJoinGameData}
	var (
		EID            pk.Int
		Hardcore       pk.Boolean
		Gamemode       pk.UnsignedByte
		PreGamemode    pk.Byte
		WorldNames     = []pk.Identifier{} // This cannot replace with "var DimensionNames []pk.Identifier" because "nil" has no type information
		DimensionCodec struct {
			DimensionType any `nbt:"minecraft:dimension_type"`
			WorldgenBiome any `nbt:"minecraft:worldgen/biome"`
		}
		Dimension                 any
		WorldName                 pk.Identifier
		HashedSeed                pk.Long
		MaxPlayers                pk.VarInt
		ViewDistance              pk.VarInt
		RDI, ERS, IsDebug, IsFlat pk.Boolean
	)
	err := p.Scan(
		&EID,
		&Hardcore,
		&Gamemode,
		&PreGamemode,
		pk.Array(&WorldNames),
		pk.NBT(&DimensionCodec),
		pk.NBT(&Dimension),
		&WorldName,
		&HashedSeed,
		&MaxPlayers,
		&ViewDistance,
		&RDI, &ERS, &IsDebug, &IsFlat,
	)
	fmt.Print(err)
	// Output: <nil>
}

func ExampleMarshal_setSlot() {
	for _, pf := range []struct {
		WindowID  byte
		Slot      int16
		Present   bool
		ItemID    int
		ItemCount byte
		NBT       any
	}{
		{WindowID: 0, Slot: 5, Present: false},
		{WindowID: 0, Slot: 5, Present: true, ItemID: 0x01, ItemCount: 1, NBT: pk.Byte(0)},
		{WindowID: 0, Slot: 5, Present: true, ItemID: 0x01, ItemCount: 1, NBT: pk.NBT(int32(0x12345678))},
	} {
		p := pk.Marshal(0x15,
			pk.Byte(pf.WindowID),
			pk.Short(pf.Slot),
			pk.Boolean(pf.Present),
			pk.Opt{Has: pf.Present, Field: pk.Tuple{
				pk.VarInt(pf.ItemID),
				pk.Byte(pf.ItemCount),
				pf.NBT,
			}},
		)
		fmt.Printf("%02X % 02X\n", p.ID, p.Data)
	}
	// Output:
	// 15 00 00 05 00
	// 15 00 00 05 01 01 01 00
	// 15 00 00 05 01 01 01 03 00 00 12 34 56 78
}

func BenchmarkPacket_Pack(b *testing.B) {
	p := pk.Packet{ID: 0, Data: make([]byte, 64)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := p.Pack(io.Discard, -1); err != nil {
			b.Fatal(err)
		}
	}
}

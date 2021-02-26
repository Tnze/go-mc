package packet_test

import (
	"bytes"
	_ "embed"
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

func TestPackVarInt(t *testing.T) {
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
func TestUnpackVarInt(t *testing.T) {
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

func TestUnpackVarInt_TooLongData(t *testing.T) {
	var vi pk.VarInt
	var data = []byte{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01}
	if _, err := vi.ReadFrom(bytes.NewReader(data)); err != nil {
		t.Logf("unpack \"% x\" error: %v", data, err)
	} else {
		t.Errorf("unpack \"% x\" should be error, get %d", data, vi)
	}
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

func TestPackVarLong(t *testing.T) {
	for i, v := range VarLongs {
		p := v.Encode()
		if !bytes.Equal(p, PackedVarLongs[i]) {
			t.Errorf("pack long %d should be \"% x\", get \"% x\"", v, PackedVarLongs[i], p)
		}
	}
}
func TestUnpackVarLong(t *testing.T) {
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

//go:embed joingame_test.bin
var joingame []byte

func TestJoinGamePacket(t *testing.T) {
	p := pk.Packet{ID: 0x24, Data: joingame}
	var (
		EID            pk.Int
		Hardcore       pk.Boolean
		Gamemode       pk.UnsignedByte
		PreGamemode    pk.Byte
		WorldCount     pk.VarInt
		WorldNames     = pk.Ary{Len: &WorldCount, Ary: &[]pk.String{}}
		DimensionCodec struct {
			DimensionType interface{} `nbt:"minecraft:dimension_type"`
			WorldgenBiome interface{} `nbt:"minecraft:worldgen/biome"`
		}
		Dimension                 interface{}
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
		&WorldCount,
		WorldNames,
		&pk.NBT{V: &DimensionCodec},
		&pk.NBT{V: &Dimension},
		&WorldName,
		&HashedSeed,
		&MaxPlayers,
		&ViewDistance,
		&RDI, &ERS, &IsDebug, &IsFlat,
	)
	if err != nil {
		t.Error(err)
	}
}

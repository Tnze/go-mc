package packet_test

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/Tnze/go-mc/nbt"
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

//go:embed joingame_test.bin
var testJoinGameData []byte

func ExamplePacket_Scan_joinGame() {
	p := pk.Packet{ID: 0x24, Data: testJoinGameData}
	var (
		EID            pk.Int
		Hardcore       pk.Boolean
		Gamemode       pk.UnsignedByte
		PreGamemode    pk.Byte
		WorldCount     pk.VarInt
		WorldNames     = make([]pk.Identifier, 0) // This cannot replace with "var WorldNames []pk.Identifier" because "nil" has no type information
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
		pk.Ary{
			Len: &WorldCount,
			Ary: &WorldNames,
		},
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
	var buf bytes.Buffer
	type Enchantment struct {
		ID  int16 `nbt:"id"`
		Lvl int16 `nbt:"lvl"`
	}
	type SlotNBT struct {
		StoredEnchantments []Enchantment
		Unbreakable        int32
	}
	for _, pf := range [][]pk.FieldEncoder{
		{
			pk.Byte(0),
			pk.Short(5),
			pk.Boolean(false),
			pk.Opt{Has: false},
		},
		{
			pk.Byte(0),
			pk.Short(5),
			pk.Boolean(true),
			pk.Opt{Has: true, Field: pk.Tuple{
				pk.VarInt(0x01),     // ItemID
				pk.Byte(1),          // ItemCount
				pk.Byte(nbt.TagEnd), // NBT, 0 when this is no data
			}},
		},
		{
			pk.Byte(0),
			pk.Short(5),
			pk.Boolean(true),
			pk.Opt{Has: true, Field: pk.Tuple{
				pk.VarInt(0x01), // ItemID
				pk.Byte(1),      // ItemCount
				pk.NBT(SlotNBT{
					StoredEnchantments: []Enchantment{
						{ID: 01, Lvl: 02},
						{ID: 03, Lvl: 04},
					},
					Unbreakable: 1, // true
				}), // NBT
			}},
		},
	} {
		buf.Reset()
		p := pk.Marshal(0x15, pf...)
		fmt.Printf("%02X % 02X\n", p.ID, p.Data)
	}
	// Output:
	// 15 00 00 05 00
	// 15 00 00 05 01 01 01 00
	// 15 00 00 05 01 01 01 0A 00 00 09 00 12 53 74 6F 72 65 64 45 6E 63 68 61 6E 74 6D 65 6E 74 73 0A 00 00 00 02 02 00 02 69 64 00 01 02 00 03 6C 76 6C 00 02 00 02 00 02 69 64 00 03 02 00 03 6C 76 6C 00 04 00 03 00 0B 55 6E 62 72 65 61 6B 61 62 6C 65 00 00 00 01 00
}

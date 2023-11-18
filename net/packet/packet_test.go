package packet_test

import (
	_ "embed"
	"fmt"
	"io"
	"testing"

	pk "github.com/Tnze/go-mc/net/packet"
)

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
	// 15 00 00 05 01 01 01 03 12 34 56 78
}

func BenchmarkPacket_Pack_packWithoutCompression(b *testing.B) {
	p := pk.Packet{ID: 0, Data: make([]byte, 64)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := p.Pack(io.Discard, -1); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPacket_Pack_packWithCompression(b *testing.B) {
	p := pk.Packet{ID: 0, Data: make([]byte, 64)}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if err := p.Pack(io.Discard, 32); err != nil {
			b.Fatal(err)
		}
	}
}

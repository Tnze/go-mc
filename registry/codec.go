package registry

import "github.com/Tnze/go-mc/nbt"

type NetworkCodec struct {
	ChatType      Registry[ChatType]       `nbt:"minecraft:chat_type"`
	DimensionType Registry[Dimension]      `nbt:"minecraft:dimension_type"`
	WorldGenBiome Registry[nbt.RawMessage] `nbt:"minecraft:worldgen/biome"`
}

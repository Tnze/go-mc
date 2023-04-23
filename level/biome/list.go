package biome

import (
	"errors"
	"hash/maphash"
	"math/bits"
)

type Type int

var hashSeed = maphash.MakeSeed()

func (t Type) MarshalText() (text []byte, err error) {
	if t >= 0 && int(t) < len(biomesNames) {
		return biomesNames[t], nil
	}
	return nil, errors.New("invalid type")
}

func (t *Type) UnmarshalText(text []byte) error {
	var ok bool
	*t, ok = biomesIDs[maphash.Bytes(hashSeed, text)]
	if ok {
		return nil
	}
	return errors.New("invalid type")
}

// String returns the biome id. Debugging purposes only.
func (t Type) String() string {
	if t >= 0 && int(t) < len(biomesNames) {
		return string(biomesNames[t])
	}
	return "<invalid biome type>"
}

var (
	// BitsPerBiome reports how many bits are required to represent all possible biomes.
	BitsPerBiome int
	biomesIDs    map[uint64]Type
	biomesNames  = [][]byte{
		[]byte("minecraft:the_void"),
		[]byte("minecraft:plains"),
		[]byte("minecraft:sunflower_plains"),
		[]byte("minecraft:snowy_plains"),
		[]byte("minecraft:ice_spikes"),
		[]byte("minecraft:desert"),
		[]byte("minecraft:swamp"),
		[]byte("minecraft:mangrove_swamp"),
		[]byte("minecraft:forest"),
		[]byte("minecraft:flower_forest"),
		[]byte("minecraft:birch_forest"),
		[]byte("minecraft:dark_forest"),
		[]byte("minecraft:old_growth_birch_forest"),
		[]byte("minecraft:old_growth_pine_taiga"),
		[]byte("minecraft:old_growth_spruce_taiga"),
		[]byte("minecraft:taiga"),
		[]byte("minecraft:snowy_taiga"),
		[]byte("minecraft:savanna"),
		[]byte("minecraft:savanna_plateau"),
		[]byte("minecraft:windswept_hills"),
		[]byte("minecraft:windswept_gravelly_hills"),
		[]byte("minecraft:windswept_forest"),
		[]byte("minecraft:windswept_savanna"),
		[]byte("minecraft:jungle"),
		[]byte("minecraft:sparse_jungle"),
		[]byte("minecraft:bamboo_jungle"),
		[]byte("minecraft:badlands"),
		[]byte("minecraft:eroded_badlands"),
		[]byte("minecraft:wooded_badlands"),
		[]byte("minecraft:meadow"),
		[]byte("minecraft:grove"),
		[]byte("minecraft:snowy_slopes"),
		[]byte("minecraft:frozen_peaks"),
		[]byte("minecraft:jagged_peaks"),
		[]byte("minecraft:stony_peaks"),
		[]byte("minecraft:river"),
		[]byte("minecraft:frozen_river"),
		[]byte("minecraft:beach"),
		[]byte("minecraft:snowy_beach"),
		[]byte("minecraft:stony_shore"),
		[]byte("minecraft:warm_ocean"),
		[]byte("minecraft:lukewarm_ocean"),
		[]byte("minecraft:deep_lukewarm_ocean"),
		[]byte("minecraft:ocean"),
		[]byte("minecraft:deep_ocean"),
		[]byte("minecraft:cold_ocean"),
		[]byte("minecraft:deep_cold_ocean"),
		[]byte("minecraft:frozen_ocean"),
		[]byte("minecraft:deep_frozen_ocean"),
		[]byte("minecraft:mushroom_fields"),
		[]byte("minecraft:dripstone_caves"),
		[]byte("minecraft:lush_caves"),
		[]byte("minecraft:deep_dark"),
		[]byte("minecraft:nether_wastes"),
		[]byte("minecraft:warped_forest"),
		[]byte("minecraft:crimson_forest"),
		[]byte("minecraft:soul_sand_valley"),
		[]byte("minecraft:basalt_deltas"),
		[]byte("minecraft:the_end"),
		[]byte("minecraft:end_highlands"),
		[]byte("minecraft:end_midlands"),
		[]byte("minecraft:small_end_islands"),
		[]byte("minecraft:end_barrens"),
	}
)

func init() {
	BitsPerBiome = bits.Len(uint(len(biomesNames)))
	biomesIDs = make(map[uint64]Type, len(biomesNames))
	for i, v := range biomesNames {
		h := maphash.Bytes(hashSeed, v)
		biomesIDs[h] = Type(i)
	}
}

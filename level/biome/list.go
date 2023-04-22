package biome

import (
	"errors"
	"math/bits"
)

type Type int

func (t Type) MarshalText() (text []byte, err error) {
	if t >= 0 && int(t) < len(biomesNames) {
		return []byte(biomesNames[t]), nil
	}
	return nil, errors.New("invalid type")
}

func (t *Type) UnmarshalText(text []byte) error {
	var ok bool
	*t, ok = biomesIDs[string(text)]
	if ok {
		return nil
	}
	return errors.New("unknown type")
}

var (
	BitsPerBiome int
	biomesIDs    map[string]Type
	biomesNames  = []string{
		"minecraft:the_void",
		"minecraft:plains",
		"minecraft:sunflower_plains",
		"minecraft:snowy_plains",
		"minecraft:ice_spikes",
		"minecraft:desert",
		"minecraft:swamp",
		"minecraft:mangrove_swamp",
		"minecraft:forest",
		"minecraft:flower_forest",
		"minecraft:birch_forest",
		"minecraft:dark_forest",
		"minecraft:old_growth_birch_forest",
		"minecraft:old_growth_pine_taiga",
		"minecraft:old_growth_spruce_taiga",
		"minecraft:taiga",
		"minecraft:snowy_taiga",
		"minecraft:savanna",
		"minecraft:savanna_plateau",
		"minecraft:windswept_hills",
		"minecraft:windswept_gravelly_hills",
		"minecraft:windswept_forest",
		"minecraft:windswept_savanna",
		"minecraft:jungle",
		"minecraft:sparse_jungle",
		"minecraft:bamboo_jungle",
		"minecraft:badlands",
		"minecraft:eroded_badlands",
		"minecraft:wooded_badlands",
		"minecraft:meadow",
		"minecraft:grove",
		"minecraft:snowy_slopes",
		"minecraft:frozen_peaks",
		"minecraft:jagged_peaks",
		"minecraft:stony_peaks",
		"minecraft:river",
		"minecraft:frozen_river",
		"minecraft:beach",
		"minecraft:snowy_beach",
		"minecraft:stony_shore",
		"minecraft:warm_ocean",
		"minecraft:lukewarm_ocean",
		"minecraft:deep_lukewarm_ocean",
		"minecraft:ocean",
		"minecraft:deep_ocean",
		"minecraft:cold_ocean",
		"minecraft:deep_cold_ocean",
		"minecraft:frozen_ocean",
		"minecraft:deep_frozen_ocean",
		"minecraft:mushroom_fields",
		"minecraft:dripstone_caves",
		"minecraft:lush_caves",
		"minecraft:deep_dark",
		"minecraft:nether_wastes",
		"minecraft:warped_forest",
		"minecraft:crimson_forest",
		"minecraft:soul_sand_valley",
		"minecraft:basalt_deltas",
		"minecraft:the_end",
		"minecraft:end_highlands",
		"minecraft:end_midlands",
		"minecraft:small_end_islands",
		"minecraft:end_barrens",
	}
)

func init() {
	BitsPerBiome = bits.Len(uint(len(biomesNames)))
	biomesIDs = make(map[string]Type, len(biomesNames))
	for i, v := range biomesNames {
		biomesIDs[v] = Type(i)
	}
}

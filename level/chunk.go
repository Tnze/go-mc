package level

import (
	"bytes"
	"io"
	"math/bits"
	"strings"
	"sync"
	"unsafe"

	"github.com/Tnze/go-mc/data/block"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/save"
)

type ChunkPos struct{ X, Z int }
type Chunk struct {
	sync.Mutex
	Sections   []Section
	HeightMaps HeightMaps
}
type HeightMaps struct {
	MotionBlocking *BitStorage
	WorldSurface   *BitStorage
}

func EmptyChunk(secs int) *Chunk {
	sections := make([]Section, secs)
	for i := range sections {
		sections[i] = Section{
			blockCount: 0,
			States:     NewStatesPaletteContainer(16*16*16, 0),
			Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
		}
	}
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			MotionBlocking: NewBitStorage(bits.Len(uint(secs)*16), 16*16, nil),
		},
	}
}

var biomesIDs = map[string]int{
	"the_void":                 0,
	"plains":                   1,
	"sunflower_plains":         2,
	"snowy_plains":             3,
	"ice_spikes":               4,
	"desert":                   5,
	"swamp":                    6,
	"forest":                   7,
	"flower_forest":            8,
	"birch_forest":             9,
	"dark_forest":              10,
	"old_growth_birch_forest":  11,
	"old_growth_pine_taiga":    12,
	"old_growth_spruce_taiga":  13,
	"taiga":                    14,
	"snowy_taiga":              15,
	"savanna":                  16,
	"savanna_plateau":          17,
	"windswept_hills":          18,
	"windswept_gravelly_hills": 19,
	"windswept_forest":         20,
	"windswept_savanna":        21,
	"jungle":                   22,
	"sparse_jungle":            23,
	"bamboo_jungle":            24,
	"badlands":                 25,
	"eroded_badlands":          26,
	"wooded_badlands":          27,
	"meadow":                   28,
	"grove":                    29,
	"snowy_slopes":             30,
	"frozen_peaks":             31,
	"jagged_peaks":             32,
	"stony_peaks":              33,
	"river":                    34,
	"frozen_river":             35,
	"beach":                    36,
	"snowy_beach":              37,
	"stony_shore":              38,
	"warm_ocean":               39,
	"lukewarm_ocean":           40,
	"deep_lukewarm_ocean":      41,
	"ocean":                    42,
	"deep_ocean":               43,
	"cold_ocean":               44,
	"deep_cold_ocean":          45,
	"frozen_ocean":             46,
	"deep_frozen_ocean":        47,
	"mushroom_fields":          48,
	"dripstone_caves":          49,
	"lush_caves":               50,
	"nether_wastes":            51,
	"warped_forest":            52,
	"crimson_forest":           53,
	"soul_sand_valley":         54,
	"basalt_deltas":            55,
	"the_end":                  56,
	"end_highlands":            57,
	"end_midlands":             58,
	"small_end_islands":        59,
	"end_barrens":              60,
}

func ChunkFromSave(c *save.Chunk, secs int) *Chunk {
	sections := make([]Section, secs)
	for _, v := range c.Sections {
		var blockCount int16
		stateData := *(*[]uint64)((unsafe.Pointer)(&v.BlockStates.Data))
		statePalette := v.BlockStates.Palette
		stateRawPalette := make([]int, len(statePalette))
		for i, v := range statePalette {
			// TODO: Consider the properties of block, not only index the block name
			stateRawPalette[i] = int(stateIDs[strings.TrimPrefix(v.Name, "minecraft:")])
			if v.Name != "minecraft:air" {
				blockCount++
			}
		}

		biomesData := *(*[]uint64)((unsafe.Pointer)(&v.BlockStates.Data))
		biomesPalette := v.Biomes.Palette
		biomesRawPalette := make([]int, len(biomesPalette))
		for i, v := range biomesPalette {
			biomesRawPalette[i] = biomesIDs[strings.TrimPrefix(v, "minecraft:")]
		}

		i := int32(int8(v.Y)) - c.YPos
		sections[i].blockCount = blockCount
		sections[i].States = NewStatesPaletteContainerWithData(16*16*16, stateData, stateRawPalette)
		sections[i].Biomes = NewBiomesPaletteContainerWithData(16*16*16*2, biomesData, biomesRawPalette)
	}
	for i := range sections {
		if sections[i].States == nil {
			sections[i] = Section{
				blockCount: 0,
				States:     NewStatesPaletteContainer(16*16*16, 0),
				Biomes:     NewBiomesPaletteContainer(4*4*4, 0),
			}
		}
	}

	motionBlocking := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.MotionBlocking))
	worldSurface := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.WorldSurface))
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			MotionBlocking: NewBitStorage(bits.Len(uint(secs)), 16*16, motionBlocking),
			WorldSurface:   NewBitStorage(bits.Len(uint(secs)), 16*16, worldSurface),
		},
	}
}

// TODO: This map should be moved to data/block.
var stateIDs = make(map[string]uint32)

func init() {
	for i, v := range block.StateID {
		name := block.ByID[v].Name
		if _, ok := stateIDs[name]; !ok {
			stateIDs[name] = i
		}
	}
}

func (c *Chunk) WriteTo(w io.Writer) (int64, error) {
	data, err := c.Data()
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		// Heightmaps
		pk.NBT(struct {
			MotionBlocking []uint64 `nbt:"MOTION_BLOCKING"`
			WorldSurface   []uint64 `nbt:"WORLD_SURFACE"`
		}{
			MotionBlocking: c.HeightMaps.MotionBlocking.Raw(),
			WorldSurface:   c.HeightMaps.MotionBlocking.Raw(),
		}),
		pk.ByteArray(data),
		pk.VarInt(0), // TODO: Block Entity
		&lightData{
			SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
			BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
			SkyLight:       []pk.ByteArray{},
			BlockLight:     []pk.ByteArray{},
		},
	}.WriteTo(w)
}

func (c *Chunk) Data() ([]byte, error) {
	var buff bytes.Buffer
	for _, section := range c.Sections {
		_, err := section.WriteTo(&buff)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

type Section struct {
	blockCount int16
	States     *PaletteContainer
	Biomes     *PaletteContainer
}

func (s *Section) GetBlock(i int) int {
	return s.States.Get(i)
}
func (s *Section) SetBlock(i int, v int) {
	// TODO: Handle cave air and void air
	if s.States.Get(i) != 0 {
		s.blockCount--
	}
	if v != 0 {
		s.blockCount++
	}
	s.States.Set(i, v)
}

func (s *Section) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Short(s.blockCount),
		s.States,
		s.Biomes,
	}.WriteTo(w)
}

func (s *Section) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		pk.Short(s.blockCount),
		s.States,
		s.Biomes,
	}.ReadFrom(r)
}

type lightData struct {
	SkyLightMask   pk.BitSet
	BlockLightMask pk.BitSet
	SkyLight       []pk.ByteArray
	BlockLight     []pk.ByteArray
}

func bitSetRev(set pk.BitSet) pk.BitSet {
	rev := make(pk.BitSet, len(set))
	for i := range rev {
		rev[i] = ^set[i]
	}
	return rev
}

func (l *lightData) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Boolean(true), // Trust Edges
		l.SkyLightMask,
		l.BlockLightMask,
		bitSetRev(l.SkyLightMask),
		bitSetRev(l.BlockLightMask),
		pk.Array(l.SkyLight),
		pk.Array(l.BlockLight),
	}.WriteTo(w)
}

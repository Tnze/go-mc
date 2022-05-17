package level

import (
	"bytes"
	"fmt"
	"io"
	"math/bits"
	"strings"
	"unsafe"

	"github.com/Tnze/go-mc/level/block"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/save"
)

type ChunkPos struct{ X, Z int }

func (c ChunkPos) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.Int(c.X).WriteTo(w)
	if err != nil {
		return
	}
	n1, err := pk.Int(c.Z).WriteTo(w)
	return n + n1, err
}

func (c *ChunkPos) ReadFrom(r io.Reader) (n int64, err error) {
	var x, z pk.Int
	if n, err = x.ReadFrom(r); err != nil {
		return n, err
	}
	var n1 int64
	if n1, err = z.ReadFrom(r); err != nil {
		return n + n1, err
	}
	*c = ChunkPos{int(x), int(z)}
	return n + n1, nil
}

type Chunk struct {
	Sections    []Section
	HeightMaps  HeightMaps
	BlockEntity []BlockEntity
}

func EmptyChunk(secs int) *Chunk {
	sections := make([]Section, secs)
	for i := range sections {
		sections[i] = Section{
			BlockCount: 0,
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
var biomesNames = []string{
	0:  "the_void",
	1:  "plains",
	2:  "sunflower_plains",
	3:  "snowy_plains",
	4:  "ice_spikes",
	5:  "desert",
	6:  "swamp",
	7:  "forest",
	8:  "flower_forest",
	9:  "birch_forest",
	10: "dark_forest",
	11: "old_growth_birch_forest",
	12: "old_growth_pine_taiga",
	13: "old_growth_spruce_taiga",
	14: "taiga",
	15: "snowy_taiga",
	16: "savanna",
	17: "savanna_plateau",
	18: "windswept_hills",
	19: "windswept_gravelly_hills",
	20: "windswept_forest",
	21: "windswept_savanna",
	22: "jungle",
	23: "sparse_jungle",
	24: "bamboo_jungle",
	25: "badlands",
	26: "eroded_badlands",
	27: "wooded_badlands",
	28: "meadow",
	29: "grove",
	30: "snowy_slopes",
	31: "frozen_peaks",
	32: "jagged_peaks",
	33: "stony_peaks",
	34: "river",
	35: "frozen_river",
	36: "beach",
	37: "snowy_beach",
	38: "stony_shore",
	39: "warm_ocean",
	40: "lukewarm_ocean",
	41: "deep_lukewarm_ocean",
	42: "ocean",
	43: "deep_ocean",
	44: "cold_ocean",
	45: "deep_cold_ocean",
	46: "frozen_ocean",
	47: "deep_frozen_ocean",
	48: "mushroom_fields",
	49: "dripstone_caves",
	50: "lush_caves",
	51: "nether_wastes",
	52: "warped_forest",
	53: "crimson_forest",
	54: "soul_sand_valley",
	55: "basalt_deltas",
	56: "the_end",
	57: "end_highlands",
	58: "end_midlands",
	59: "small_end_islands",
	60: "end_barrens",
}

// ChunkFromSave convert save.Chunk to level.Chunk.
func ChunkFromSave(c *save.Chunk) *Chunk {
	secs := len(c.Sections)
	sections := make([]Section, secs)
	for _, v := range c.Sections {
		i := int32(v.Y) - c.YPos
		// TODO: the error is ignored
		sections[i].BlockCount, sections[i].States, _ = readStatesPalette(v.BlockStates.Palette, v.BlockStates.Data)
		sections[i].Biomes, _ = readBiomesPalette(v.Biomes.Palette, v.Biomes.Data)
	}

	motionBlocking := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.MotionBlocking))
	motionBlockingNoLeaves := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.MotionBlockingNoLeaves))
	oceanFloor := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.OceanFloor))
	worldSurface := *(*[]uint64)(unsafe.Pointer(&c.Heightmaps.WorldSurface))

	bitsForHeight := bits.Len( /* chunk height in blocks */ uint(secs) * 16)
	return &Chunk{
		Sections: sections,
		HeightMaps: HeightMaps{
			MotionBlocking:         NewBitStorage(bitsForHeight, 16*16, motionBlocking),
			MotionBlockingNoLeaves: NewBitStorage(bitsForHeight, 16*16, motionBlockingNoLeaves),
			OceanFloor:             NewBitStorage(bitsForHeight, 16*16, oceanFloor),
			WorldSurface:           NewBitStorage(bitsForHeight, 16*16, worldSurface),
		},
	}
}

func readStatesPalette(palette []save.BlockState, data []int64) (blockCount int16, paletteData *PaletteContainer, err error) {
	stateData := *(*[]uint64)((unsafe.Pointer)(&data))
	statePalette := make([]int, len(palette))
	for i, v := range palette {
		b, ok := block.FromID[v.Name]
		if !ok {
			return 0, nil, fmt.Errorf("unknown block id: %v", v.Name)
		}
		if v.Properties.Data != nil {
			if err := v.Properties.Unmarshal(&b); err != nil {
				return 0, nil, fmt.Errorf("unmarshal block properties fail: %v", err)
			}
		}
		s, ok := block.ToStateID[b]
		if !ok {
			return 0, nil, fmt.Errorf("unknown block: %v", b)
		}
		if !block.IsAir(s) {
			blockCount++
		}
		statePalette[i] = s
	}
	paletteData = NewStatesPaletteContainerWithData(16*16*16, stateData, statePalette)
	return
}

func readBiomesPalette(palette []string, data []int64) (*PaletteContainer, error) {
	biomesData := *(*[]uint64)((unsafe.Pointer)(&data))
	biomesRawPalette := make([]int, len(palette))
	var ok bool
	for i, v := range palette {
		biomesRawPalette[i], ok = biomesIDs[strings.TrimPrefix(v, "minecraft:")]
		if !ok {
			return nil, fmt.Errorf("unknown biomes: %s", v)
		}
	}
	return NewBiomesPaletteContainerWithData(4*4*4, biomesData, biomesRawPalette), nil
}

// ChunkToSave convert level.Chunk to save.Chunk
func ChunkToSave(c *Chunk, dst *save.Chunk) {
	secs := len(c.Sections)
	sections := make([]save.Section, secs)
	for i, v := range c.Sections {
		statePalette, stateData := writeStatesPalette(v.States)
		biomePalette, biomeData := writeBiomesPalette(v.Biomes)
		sections[i] = save.Section{
			Y: int8(int32(i) + dst.YPos),
			BlockStates: struct {
				Palette []save.BlockState `nbt:"palette"`
				Data    []int64           `nbt:"data"`
			}{
				Palette: statePalette, Data: stateData,
			},
			Biomes: struct {
				Palette []string `nbt:"palette"`
				Data    []int64  `nbt:"data"`
			}{
				Palette: biomePalette, Data: biomeData,
			},
			SkyLight:   nil,
			BlockLight: nil,
		}
	}
	dst.Sections = sections
}

func writeStatesPalette(paletteData *PaletteContainer) (palette []save.BlockState, data []int64) {
	rawPalette := paletteData.palette.export()
	palette = make([]save.BlockState, len(rawPalette))
	var buffer bytes.Buffer
	for i, v := range rawPalette {
		b := block.StateList[v]
		palette[i].Name = b.ID()

		buffer.Reset()
		if err := nbt.NewEncoder(&buffer).Encode(b, ""); err != nil {
			panic(err)
		}
		if _, err := nbt.NewDecoder(&buffer).Decode(&palette[i].Properties); err != nil {
			panic(err)
		}
	}

	rawData := paletteData.data.Raw()
	copy(data, *(*[]int64)(unsafe.Pointer(&rawData)))

	return
}

func writeBiomesPalette(paletteData *PaletteContainer) (palette []string, data []int64) {
	rawPalette := paletteData.palette.export()
	palette = make([]string, len(rawPalette))
	for i, v := range rawPalette {
		palette[i] = biomesNames[v]
	}

	rawData := paletteData.data.Raw()
	copy(data, *(*[]int64)(unsafe.Pointer(&rawData)))

	return
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
		pk.Array(c.BlockEntity),
		&lightData{
			SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
			BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
			SkyLight:       []pk.ByteArray{},
			BlockLight:     []pk.ByteArray{},
		},
	}.WriteTo(w)
}

func (c *Chunk) ReadFrom(r io.Reader) (int64, error) {
	var (
		heightmaps struct {
			MotionBlocking []uint64 `nbt:"MOTION_BLOCKING"`
			WorldSurface   []uint64 `nbt:"WORLD_SURFACE"`
		}
		data pk.ByteArray
	)

	n, err := pk.Tuple{
		pk.NBT(&heightmaps),
		&data,
		pk.Array(&c.BlockEntity),
		&lightData{
			SkyLightMask:   make(pk.BitSet, (16*16*16-1)>>6+1),
			BlockLightMask: make(pk.BitSet, (16*16*16-1)>>6+1),
			SkyLight:       []pk.ByteArray{},
			BlockLight:     []pk.ByteArray{},
		},
	}.ReadFrom(r)
	if err != nil {
		return n, err
	}

	err = c.PutData(data)
	return n, err
}

func (c *Chunk) Data() ([]byte, error) {
	var buff bytes.Buffer
	for i := range c.Sections {
		_, err := c.Sections[i].WriteTo(&buff)
		if err != nil {
			return nil, err
		}
	}
	return buff.Bytes(), nil
}

func (c *Chunk) PutData(data []byte) error {
	r := bytes.NewReader(data)
	for i := range c.Sections {
		_, err := c.Sections[i].ReadFrom(r)
		if err != nil {
			return err
		}
	}
	return nil
}

type HeightMaps struct {
	MotionBlocking         *BitStorage
	MotionBlockingNoLeaves *BitStorage
	OceanFloor             *BitStorage
	WorldSurface           *BitStorage
}

type BlockEntity struct {
	XZ   int8
	Y    int16
	Type int32
	Data nbt.RawMessage
}

func (b BlockEntity) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.Byte(b.XZ),
		pk.Short(b.Y),
		pk.VarInt(b.Type),
		pk.NBT(b.Data),
	}.WriteTo(w)
}

func (b *BlockEntity) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		(*pk.Byte)(&b.XZ),
		(*pk.Short)(&b.Y),
		(*pk.VarInt)(&b.Type),
		pk.NBT(&b.Data),
	}.ReadFrom(r)
}

type Section struct {
	BlockCount int16
	States     *PaletteContainer
	Biomes     *PaletteContainer
}

func (s *Section) GetBlock(i int) int {
	return s.States.Get(i)
}
func (s *Section) SetBlock(i int, v int) {
	if block.IsAir(s.States.Get(i)) {
		s.BlockCount--
	}
	if v != 0 {
		s.BlockCount++
	}
	s.States.Set(i, v)
}

func (s *Section) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Short(s.BlockCount),
		s.States,
		s.Biomes,
	}.WriteTo(w)
}

func (s *Section) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		(*pk.Short)(&s.BlockCount),
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

func (l *lightData) ReadFrom(r io.Reader) (int64, error) {
	var TrustEdges pk.Boolean
	var RevSkyLightMask, RevBlockLightMask pk.BitSet
	return pk.Tuple{
		&TrustEdges, // Trust Edges
		&l.SkyLightMask,
		&l.BlockLightMask,
		&RevSkyLightMask,
		&RevBlockLightMask,
		pk.Array(&l.SkyLight),
		pk.Array(&l.BlockLight),
	}.ReadFrom(r)
}

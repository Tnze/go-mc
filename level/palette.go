package level

import (
	"io"
	"strconv"

	"github.com/Tnze/go-mc/level/biome"
	"github.com/Tnze/go-mc/level/block"
	pk "github.com/Tnze/go-mc/net/packet"
)

type State interface {
	~int
}
type (
	BlocksState = block.StateID
	BiomesState = biome.Type
)

type PaletteContainer[T State] struct {
	bits    int
	config  paletteCfg[T]
	palette palette[T]
	data    *BitStorage
}

func NewStatesPaletteContainer(length int, defaultValue BlocksState) *PaletteContainer[BlocksState] {
	return &PaletteContainer[BlocksState]{
		bits:    0,
		config:  statesCfg{},
		palette: &singleValuePalette[BlocksState]{v: defaultValue},
		data:    NewBitStorage(0, length, nil),
	}
}

func NewStatesPaletteContainerWithData(length int, data []uint64, pat []BlocksState) *PaletteContainer[BlocksState] {
	var p palette[BlocksState]
	n := calcBitsPerValue(length, len(data))
	switch n {
	case 0:
		p = &singleValuePalette[BlocksState]{pat[0]}
	case 1, 2, 3, 4:
		n = 4
		p = &linearPalette[BlocksState]{
			values: pat,
			bits:   n,
		}
	case 5, 6, 7, 8:
		ids := make(map[BlocksState]int)
		for i, v := range pat {
			ids[v] = i
		}
		p = &hashPalette[BlocksState]{
			ids:    ids,
			values: pat,
			bits:   n,
		}
	default:
		p = &globalPalette[BlocksState]{}
	}
	return &PaletteContainer[BlocksState]{
		bits:    n,
		config:  statesCfg{},
		palette: p,
		data:    NewBitStorage(n, length, data),
	}
}

func NewBiomesPaletteContainer(length int, defaultValue BiomesState) *PaletteContainer[BiomesState] {
	return &PaletteContainer[BiomesState]{
		bits:    0,
		config:  biomesCfg{},
		palette: &singleValuePalette[BiomesState]{v: defaultValue},
		data:    NewBitStorage(0, length, nil),
	}
}

func NewBiomesPaletteContainerWithData(length int, data []uint64, pat []BiomesState) *PaletteContainer[BiomesState] {
	var p palette[BiomesState]
	n := calcBitsPerValue(length, len(data))
	switch n {
	case 0:
		p = &singleValuePalette[BiomesState]{pat[0]}
	case 1, 2, 3:
		p = &linearPalette[BiomesState]{
			values: pat,
			bits:   n,
		}
	default:
		p = &globalPalette[BiomesState]{}
	}
	return &PaletteContainer[BiomesState]{
		bits:    n,
		config:  biomesCfg{},
		palette: p,
		data:    NewBitStorage(n, length, data),
	}
}

func (p *PaletteContainer[T]) Get(i int) T {
	return p.palette.value(p.data.Get(i))
}

func (p *PaletteContainer[T]) Set(i int, v T) {
	if vv, ok := p.palette.id(v); ok {
		p.data.Set(i, vv)
	} else {
		length := p.data.Len()
		// resize
		newContainer := PaletteContainer[T]{
			bits:    vv,
			config:  p.config,
			palette: p.config.create(vv),
			data:    NewBitStorage(p.config.bits(vv), length, nil),
		}
		// copy
		for i := 0; i < length; i++ {
			newContainer.Set(i, p.Get(i))
		}

		if vv, ok := newContainer.palette.id(v); !ok {
			panic("not reachable")
		} else {
			newContainer.data.Set(i, vv)
		}
		*p = newContainer
	}
}

func (p *PaletteContainer[T]) ReadFrom(r io.Reader) (n int64, err error) {
	var nBits pk.UnsignedByte
	n, err = nBits.ReadFrom(r)
	if err != nil {
		return
	}
	p.bits = p.config.bits(int(nBits))
	p.palette = p.config.create(int(nBits))

	nn, err := p.palette.ReadFrom(r)
	n += nn
	if err != nil {
		return n, err
	}

	nn, err = p.data.ReadFrom(r)
	n += nn
	if err != nil {
		return n, err
	}
	return n, p.data.Fix(p.bits)
}

func (p *PaletteContainer[T]) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.UnsignedByte(p.bits),
		p.palette,
		p.data,
	}.WriteTo(w)
}

// Palette export the raw palette values for @maxsupermanhd.
// Others shouldn't call this because this might be removed
// after max doesn't need it anymore.
func (p *PaletteContainer[T]) Palette() []T {
	return p.palette.export()
}

type paletteCfg[T State] interface {
	bits(int) int
	create(bits int) palette[T]
}

type statesCfg struct{}

func (s statesCfg) bits(bits int) int {
	switch bits {
	case 0:
		return 0
	case 1, 2, 3, 4:
		return 4
	case 5, 6, 7, 8:
		return bits
	default:
		return block.BitsPerBlock
	}
}

func (s statesCfg) create(bits int) palette[BlocksState] {
	switch bits {
	case 0:
		return &singleValuePalette[BlocksState]{v: -1}
	case 1, 2, 3, 4:
		return &linearPalette[BlocksState]{bits: 4, values: make([]BlocksState, 0, 1<<4)}
	case 5, 6, 7, 8:
		return &hashPalette[BlocksState]{
			bits:   bits,
			ids:    make(map[BlocksState]int),
			values: make([]BlocksState, 0, 1<<bits),
		}
	default:
		return &globalPalette[BlocksState]{}
	}
}

type biomesCfg struct{}

func (b biomesCfg) bits(bits int) int {
	switch bits {
	case 0:
		return 0
	case 1, 2, 3:
		return bits
	default:
		return biome.BitsPerBiome
	}
}

func (b biomesCfg) create(bits int) palette[BiomesState] {
	switch bits {
	case 0:
		return &singleValuePalette[BiomesState]{v: -1}
	case 1, 2, 3:
		return &linearPalette[BiomesState]{bits: bits, values: make([]BiomesState, 0, 1<<bits)}
	default:
		return &globalPalette[BiomesState]{}
	}
}

type palette[T State] interface {
	pk.Field
	// id return the index of state v in the palette and true if existed.
	// otherwise return the new bits for resize and false.
	id(v T) (int, bool)
	value(i int) T
	export() []T
}

type singleValuePalette[T State] struct {
	v T
}

func (s *singleValuePalette[T]) id(v T) (int, bool) {
	if s.v == v {
		return 0, true
	}
	// We have 2 values now. At least 1 bit is required.
	return 1, false
}

func (s *singleValuePalette[T]) value(i int) T {
	if i == 0 {
		return s.v
	}
	panic("singleValuePalette: " + strconv.Itoa(i) + " out of bounds")
}

func (s *singleValuePalette[T]) export() []T {
	return []T{s.v}
}

func (s *singleValuePalette[T]) ReadFrom(r io.Reader) (n int64, err error) {
	var i pk.VarInt
	n, err = i.ReadFrom(r)
	if err != nil {
		return
	}
	s.v = T(i)
	return
}

func (s *singleValuePalette[T]) WriteTo(w io.Writer) (n int64, err error) {
	return pk.VarInt(s.v).WriteTo(w)
}

type linearPalette[T State] struct {
	values []T
	bits   int
}

func (l *linearPalette[T]) id(v T) (int, bool) {
	for i, t := range l.values {
		if t == v {
			return i, true
		}
	}
	if cap(l.values)-len(l.values) > 0 {
		l.values = append(l.values, v)
		return len(l.values) - 1, true
	}
	return l.bits + 1, false
}

func (l *linearPalette[T]) value(i int) T {
	if i >= 0 && i < len(l.values) {
		return l.values[i]
	}
	panic("linearPalette: " + strconv.Itoa(i) + " out of bounds")
}

func (l *linearPalette[T]) export() []T {
	return l.values
}

func (l *linearPalette[T]) ReadFrom(r io.Reader) (n int64, err error) {
	var size, value pk.VarInt
	if n, err = size.ReadFrom(r); err != nil {
		return
	}
	if int(size) > cap(l.values) {
		l.values = make([]T, size)
	} else {
		l.values = l.values[:size]
	}
	for i := 0; i < int(size); i++ {
		if nn, err := value.ReadFrom(r); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
		l.values[i] = T(value)
	}
	return
}

func (l *linearPalette[T]) WriteTo(w io.Writer) (n int64, err error) {
	if n, err = pk.VarInt(len(l.values)).WriteTo(w); err != nil {
		return
	}
	for _, v := range l.values {
		if nn, err := pk.VarInt(v).WriteTo(w); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
	}
	return
}

type hashPalette[T State] struct {
	ids    map[T]int
	values []T
	bits   int
}

func (h *hashPalette[T]) id(v T) (int, bool) {
	if i, ok := h.ids[v]; ok {
		return i, true
	}
	if cap(h.values)-len(h.values) > 0 {
		h.ids[v] = len(h.values)
		h.values = append(h.values, v)
		return len(h.values) - 1, true
	}
	return h.bits + 1, false
}

func (h *hashPalette[T]) value(i int) T {
	if i >= 0 && i < len(h.values) {
		return h.values[i]
	}
	panic("hashPalette: " + strconv.Itoa(i) + " out of bounds")
}

func (h *hashPalette[T]) export() []T {
	return h.values
}

func (h *hashPalette[T]) ReadFrom(r io.Reader) (n int64, err error) {
	var size, value pk.VarInt
	if n, err = size.ReadFrom(r); err != nil {
		return
	}
	if int(size) > cap(h.values) {
		h.values = make([]T, size)
	} else {
		h.values = h.values[:size]
	}
	for i := 0; i < int(size); i++ {
		if nn, err := value.ReadFrom(r); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
		h.values[i] = T(value)
		h.ids[T(value)] = i
	}
	return
}

func (h *hashPalette[T]) WriteTo(w io.Writer) (n int64, err error) {
	if n, err = pk.VarInt(len(h.values)).WriteTo(w); err != nil {
		return
	}
	for _, v := range h.values {
		if nn, err := pk.VarInt(v).WriteTo(w); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
	}
	return
}

type globalPalette[T State] struct{}

func (g *globalPalette[T]) id(v T) (int, bool) {
	return int(v), true
}

func (g *globalPalette[T]) value(i int) T {
	return T(i)
}

func (g *globalPalette[T]) export() []T {
	return []T{}
}

func (g *globalPalette[T]) ReadFrom(_ io.Reader) (int64, error) {
	return 0, nil
}

func (g *globalPalette[T]) WriteTo(_ io.Writer) (int64, error) {
	return 0, nil
}

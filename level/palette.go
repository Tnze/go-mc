package level

import (
	"io"
	"math/bits"
	"strconv"

	pk "github.com/Tnze/go-mc/net/packet"
)

type state = int

type PaletteContainer struct {
	bits    int
	config  paletteCfg
	palette palette
	data    *BitStorage
}

func NewStatesPaletteContainer(length int, defaultValue state) *PaletteContainer {
	return &PaletteContainer{
		bits:    0,
		config:  statesCfg{},
		palette: &singleValuePalette{v: defaultValue},
		data:    NewBitStorage(0, length, nil),
	}
}

func NewStatesPaletteContainerWithData(length int, data []uint64, pat []int) *PaletteContainer {
	var p palette
	n := bits.Len(uint(len(pat) - 1))
	switch n {
	case 0:
		p = &singleValuePalette{pat[0]}
	case 1, 2, 3, 4:
		n = 4
		p = &linearPalette{
			values: pat,
			bits:   n,
		}
	case 5, 6, 7, 8:
		ids := make(map[state]int)
		for i, v := range pat {
			ids[v] = i
		}
		p = &hashPalette{
			ids:    ids,
			values: pat,
			bits:   n,
		}
	default:
		p = &globalPalette{}
	}
	return &PaletteContainer{
		bits:    n,
		config:  statesCfg{},
		palette: p,
		data:    NewBitStorage(n, length, data),
	}
}

func NewBiomesPaletteContainer(length int, defaultValue state) *PaletteContainer {
	return &PaletteContainer{
		bits:    0,
		config:  biomesCfg{},
		palette: &singleValuePalette{v: defaultValue},
		data:    NewBitStorage(0, length, nil),
	}
}

func NewBiomesPaletteContainerWithData(length int, data []uint64, pat []int) *PaletteContainer {
	var p palette
	n := bits.Len(uint(len(pat) - 1))
	switch n {
	case 0:
		p = &singleValuePalette{pat[0]}
	case 1, 2, 3:
		p = &linearPalette{
			values: pat,
			bits:   n,
		}
	default:
		p = &globalPalette{}
	}
	return &PaletteContainer{
		bits:    n,
		config:  biomesCfg{},
		palette: p,
		data:    NewBitStorage(n, length, data),
	}
}

func (p *PaletteContainer) Get(i int) state {
	return p.palette.value(p.data.Get(i))
}

func (p *PaletteContainer) Set(i int, v state) {
	if vv, ok := p.palette.id(v); ok {
		p.data.Set(i, vv)
	} else {
		// resize
		oldLen := p.data.Len()
		newPalette := PaletteContainer{
			bits:    vv,
			config:  p.config,
			palette: p.config.create(vv),
			data:    NewBitStorage(vv, oldLen+1, nil),
		}
		// copy
		for i := 0; i < oldLen; i++ {
			raw := p.data.Get(i)
			if vv, ok := newPalette.palette.id(raw); !ok {
				panic("not reachable")
			} else {
				newPalette.data.Set(i, vv)
			}
		}

		if vv, ok := newPalette.palette.id(v); !ok {
			panic("not reachable")
		} else {
			newPalette.data.Set(oldLen, vv)
		}
		*p = newPalette
	}
}

func (p *PaletteContainer) ReadFrom(r io.Reader) (n int64, err error) {
	var bits pk.UnsignedByte
	n, err = bits.ReadFrom(r)
	if err != nil {
		return
	}
	p.bits = p.config.bits(int(bits))
	p.palette = p.config.create(int(bits))

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
	return n, nil
}

func (p *PaletteContainer) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.UnsignedByte(p.bits),
		p.palette,
		p.data,
	}.WriteTo(w)
}

type paletteCfg interface {
	bits(int) int
	create(bits int) palette
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
		return bits
	}
}

func (s statesCfg) create(bits int) palette {
	switch bits {
	case 0:
		return &singleValuePalette{v: -1}
	case 1, 2, 3, 4:
		return &linearPalette{bits: 4, values: make([]state, 0, 1<<4)}
	case 5, 6, 7, 8:
		// TODO: HashMapPalette
		return &linearPalette{bits: bits, values: make([]state, 0, 1<<bits)}
	default:
		return &globalPalette{}
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
		return bits
	}
}
func (b biomesCfg) create(bits int) palette {
	switch bits {
	case 0:
		return &singleValuePalette{v: -1}
	case 1, 2, 3:
		return &linearPalette{bits: bits, values: make([]state, 0, 1<<bits)}
	default:
		return &globalPalette{}
	}
}

type palette interface {
	pk.Field
	// id return the index of state v in the palette and true if existed.
	// otherwise return the new bits for resize and false.
	id(v state) (int, bool)
	value(i int) state
	export() []state
}

type singleValuePalette struct {
	v state
}

func (s *singleValuePalette) id(v state) (int, bool) {
	if s.v == v {
		return 0, true
	}
	// We have 2 values now. At least 1 bit is required.
	return 1, false
}

func (s *singleValuePalette) value(i int) state {
	if i == 0 {
		return s.v
	}
	panic("singleValuePalette: " + strconv.Itoa(i) + " out of bounds")
}

func (s *singleValuePalette) export() []state {
	return []int{s.v}
}

func (s *singleValuePalette) ReadFrom(r io.Reader) (n int64, err error) {
	var i pk.VarInt
	n, err = i.ReadFrom(r)
	if err != nil {
		return
	}
	s.v = state(i)
	return
}

func (s *singleValuePalette) WriteTo(w io.Writer) (n int64, err error) {
	return pk.VarInt(s.v).WriteTo(w)
}

type linearPalette struct {
	values []state
	bits   int
}

func (l *linearPalette) id(v state) (int, bool) {
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

func (l *linearPalette) value(i int) state {
	if i >= 0 && i < len(l.values) {
		return l.values[i]
	}
	panic("linearPalette: " + strconv.Itoa(i) + " out of bounds")
}

func (l *linearPalette) export() []state {
	return l.values
}

func (l *linearPalette) ReadFrom(r io.Reader) (n int64, err error) {
	var size, value pk.VarInt
	if n, err = size.ReadFrom(r); err != nil {
		return
	}
	if int(size) > cap(l.values) {
		l.values = make([]state, size)
	} else {
		l.values = l.values[:size]
	}
	for i := 0; i < int(size); i++ {
		if nn, err := value.ReadFrom(r); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
		l.values[i] = state(value)
	}
	return
}

func (l *linearPalette) WriteTo(w io.Writer) (n int64, err error) {
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

type hashPalette struct {
	ids    map[state]int
	values []state
	bits   int
}

func (h *hashPalette) id(v state) (int, bool) {
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

func (h *hashPalette) value(i int) state {
	if i >= 0 && i < len(h.values) {
		return h.values[i]
	}
	panic("hashPalette: " + strconv.Itoa(i) + " out of bounds")
}

func (h *hashPalette) export() []state {
	return h.values
}

func (h *hashPalette) ReadFrom(r io.Reader) (n int64, err error) {
	var size, value pk.VarInt
	if n, err = size.ReadFrom(r); err != nil {
		return
	}
	if int(size) > cap(h.values) {
		h.values = make([]state, size)
	} else {
		h.values = h.values[:size]
	}
	for i := 0; i < int(size); i++ {
		if nn, err := value.ReadFrom(r); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
		h.values[i] = state(value)
		h.ids[state(value)] = i
	}
	return
}

func (h *hashPalette) WriteTo(w io.Writer) (n int64, err error) {
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

type globalPalette struct{}

func (g *globalPalette) id(v state) (int, bool) {
	return v, true
}

func (g *globalPalette) value(i int) state {
	return i
}

func (g *globalPalette) export() []state {
	return []state{}
}

func (g *globalPalette) ReadFrom(_ io.Reader) (int64, error) {
	return 0, nil
}

func (g *globalPalette) WriteTo(_ io.Writer) (int64, error) {
	return 0, nil
}

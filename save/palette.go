package save

import (
	"io"
	"strconv"

	pk "github.com/Tnze/go-mc/net/packet"
)

type BlockState interface {
}

type PaletteContainer struct {
	bits    int
	maps    IdMaps
	config  func(maps IdMaps, bits int) palette
	palette palette
	data    *BitStorage
}

func NewStatesPaletteContainer(maps IdMaps, length int) *PaletteContainer {
	return &PaletteContainer{
		bits:    0,
		maps:    maps,
		config:  createStatesPalette,
		palette: createStatesPalette(maps, 0),
		data:    NewBitStorage(0, length, nil),
	}
}

func NewBiomesPaletteContainer(maps IdMaps, length int) *PaletteContainer {
	return &PaletteContainer{
		bits:    0,
		maps:    maps,
		config:  createBiomesPalette,
		palette: createBiomesPalette(maps, 0),
		data:    NewBitStorage(0, length, nil),
	}
}

func (p *PaletteContainer) Get(i int) BlockState {
	return p.palette.value(p.data.Get(i))
}

func (p *PaletteContainer) Set(i int, v BlockState) {
	if vv, ok := p.palette.id(v); ok {
		p.data.Set(i, vv)
	} else {
		// resize
		oldLen := p.data.Len()
		newPalette := PaletteContainer{
			bits:    vv,
			maps:    p.maps,
			config:  p.config,
			palette: p.config(p.maps, vv),
			data:    NewBitStorage(vv, oldLen+1, nil),
		}
		// copy
		for i := 0; i < oldLen; i++ {
			raw := p.palette.value(i)
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
	p.palette = p.config(p.maps, int(bits))

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

func createStatesPalette(maps IdMaps, bits int) palette {
	switch bits {
	case 0:
		return &singleValuePalette{
			maps: maps,
			v:    nil,
		}
	case 1, 2, 3, 4:
		return &linearPalette{
			maps: maps,
			bits: 4,
		}
	case 5, 6, 7, 8:
		// TODO: HashMapPalette
		return &linearPalette{
			maps: maps,
			bits: bits,
		}
	default:
		return &globalPalette{
			maps: maps,
		}
	}
}

func createBiomesPalette(maps IdMaps, bits int) palette {
	switch bits {
	case 0:
		return &singleValuePalette{
			maps: maps,
			v:    nil,
		}
	case 1, 2, 3:
		return &linearPalette{
			maps: maps,
			bits: bits,
		}
	default:
		return &globalPalette{
			maps: maps,
		}
	}
}

func (p *PaletteContainer) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.UnsignedByte(p.bits),
		p.palette,
		p.data,
	}.WriteTo(w)
}

type palette interface {
	pk.FieldEncoder
	pk.FieldDecoder
	id(v BlockState) (int, bool)
	value(i int) BlockState
}

type IdMaps interface {
	getID(state BlockState) (id int)
	getValue(id int) (state BlockState)
}

type singleValuePalette struct {
	maps IdMaps
	v    BlockState
}

func (s *singleValuePalette) id(v BlockState) (int, bool) {
	if s.v == nil {
		s.v = v
		return 0, true
	}
	if s.v == v {
		return 0, true
	}
	// We have 2 values now. At least 1 bit is required.
	return 1, false
}

func (s *singleValuePalette) value(i int) BlockState {
	if s.v != nil && i == 0 {
		return s.v
	}
	panic("singleValuePalette: " + strconv.Itoa(i) + " out of bounds")
}

func (s *singleValuePalette) ReadFrom(r io.Reader) (n int64, err error) {
	var i pk.VarInt
	n, err = i.ReadFrom(r)
	if err != nil {
		return
	}
	s.v = s.maps.getValue(int(i))
	return
}

func (s *singleValuePalette) WriteTo(w io.Writer) (n int64, err error) {
	return pk.VarInt(s.maps.getID(s.v)).WriteTo(w)
}

type linearPalette struct {
	maps   IdMaps
	values []BlockState
	bits   int
}

func (l *linearPalette) id(v BlockState) (int, bool) {
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

func (l *linearPalette) value(i int) BlockState {
	if i >= 0 && i < len(l.values) {
		return l.values[i]
	}
	return nil
}

func (l *linearPalette) ReadFrom(r io.Reader) (n int64, err error) {
	var size, value pk.VarInt
	if n, err = size.ReadFrom(r); err != nil {
		return
	}
	for i := 0; i < int(size); i++ {
		if nn, err := value.ReadFrom(r); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
		l.values[i] = l.maps.getValue(int(value))
	}
	return
}

func (l *linearPalette) WriteTo(w io.Writer) (n int64, err error) {
	if n, err = pk.VarInt(len(l.values)).WriteTo(w); err != nil {
		return
	}
	for _, v := range l.values {
		if nn, err := pk.VarInt(l.maps.getID(v)).WriteTo(w); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
	}
	return
}

type globalPalette struct {
	maps IdMaps
}

func (g *globalPalette) id(v BlockState) (int, bool) {
	return g.maps.getID(v), true
}

func (g *globalPalette) value(i int) BlockState {
	return g.value(i)
}

func (g *globalPalette) ReadFrom(_ io.Reader) (int64, error) {
	return 0, nil
}

func (g *globalPalette) WriteTo(_ io.Writer) (int64, error) {
	return 0, nil
}

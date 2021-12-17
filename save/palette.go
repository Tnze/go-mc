package save

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

type BlockState interface {
}

type PaletteContainer struct {
	maps blockMaps
	palette
	BitStorage
}

func (p *PaletteContainer) ReadFrom(r io.Reader) (n int64, err error) {
	var bits pk.UnsignedByte
	n, err = bits.ReadFrom(r)
	if err != nil {
		return
	}
	switch bits {
	case 0:
		// TODO: SingleValuePalette
	case 1, 2, 3, 4:
		p.palette = &linearPalette{
			onResize: nil,
			maps:     p.maps,
			bits:     4,
		}
	case 5, 6, 7, 8:
		// TODO: HashMapPalette
	default:
		// TODO: GlobalPalette
	}

	nn, err := p.palette.ReadFrom(r)
	n += nn
	if err != nil {
		return n, err
	}

	nn, err = p.BitStorage.ReadFrom(r)
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
		p.BitStorage,
	}.WriteTo(w)
}

type palette interface {
	pk.FieldEncoder
	pk.FieldDecoder
	id(v BlockState) int
	value(i int) BlockState
}

type blockMaps interface {
	getID(state BlockState) (id int)
	getValue(id int) (state BlockState)
}

type linearPalette struct {
	onResize func(n int, v BlockState) int
	maps     blockMaps
	values   []BlockState
	bits     int
}

func (l *linearPalette) id(v BlockState) int {
	for i, t := range l.values {
		if t == v {
			return i
		}
	}
	if cap(l.values)-len(l.values) > 0 {
		l.values = append(l.values, v)
		return len(l.values) - 1
	}
	return l.onResize(l.bits+1, v)
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

type hashMapPalette struct {
	maps   blockMaps
	values map[int]BlockState
	bits   int
}

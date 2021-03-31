package save

import (
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"io"
)

type BlockState interface {
}

type palette interface {
	id(v BlockState) int
	value(i int) BlockState
	io.ReaderFrom
	io.WriterTo
	read(r nbt.DecoderReader) (int, error)
}

type linearPalette struct {
	onResize func(n int, v BlockState) int
	sToID    map[BlockState]int
	idTos    map[int]BlockState
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
		l.values[i] = l.idTos[int(value)]
	}
	return
}

func (l *linearPalette) WriteTo(w io.Writer) (n int64, err error) {
	if n, err = pk.VarInt(len(l.values)).WriteTo(w); err != nil {
		return
	}
	for _, v := range l.values {
		if nn, err := pk.VarInt(l.sToID[v]).WriteTo(w); err != nil {
			return n + nn, err
		} else {
			n += nn
		}
	}
	return
}

func (l *linearPalette) read(r nbt.DecoderReader) (int, error) {
	panic("not implemented yet")
}

package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*Instrument)(nil)

// TODO
type Instrument struct {
	Type        pk.VarInt
	SoundEvent  SoundEvent
	UseDuration pk.Float
	Range       pk.Float
}

// ID implements DataComponent.
func (Instrument) ID() string {
	return "minecraft:instrument"
}

// ReadFrom implements DataComponent.
func (i *Instrument) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		&i.Type,
		pk.Opt{
			Has: func() bool { return i.Type == 0 },
			Field: pk.Tuple{
				&i.SoundEvent,
				&i.UseDuration,
				&i.Range,
			},
		},
	}.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (i *Instrument) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		&i.Type,
		pk.Opt{
			Has: func() bool { return i.Type == 0 },
			Field: pk.Tuple{
				&i.SoundEvent,
				&i.UseDuration,
				&i.Range,
			},
		},
	}.WriteTo(w)
}

// TODO
type SoundEvent struct {
	Type       pk.VarInt
	SoundName  pk.Identifier
	FixedRange pk.Option[pk.Float, *pk.Float]
}

func (s *SoundEvent) ReadFrom(r io.Reader) (int64, error) {
	return pk.Tuple{
		&s.Type,
		pk.Opt{
			Has: func() bool { return s.Type == 0 },
			Field: pk.Tuple{
				&s.SoundName,
				&s.FixedRange,
			},
		},
	}.ReadFrom(r)
}

func (s SoundEvent) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		&s.Type,
		pk.Opt{
			Has: func() bool { return s.Type == 0 },
			Field: pk.Tuple{
				&s.SoundName,
				&s.FixedRange,
			},
		},
	}.WriteTo(w)
}

package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*Food)(nil)

type Food struct {
	Nutrition    pk.VarInt
	Saturation   pk.Float
	CanAlwaysEat pk.Boolean
	EatSeconds   pk.Float
	// TODO: using_converts_to
	// TODO: effects
}

// ID implements DataComponent.
func (Food) ID() string {
	return "minecraft:food"
}

// ReadFrom implements DataComponent.
func (f *Food) ReadFrom(r io.Reader) (n int64, err error) {
	pk.Tuple{
		&f.Nutrition,
		&f.Saturation,
		&f.CanAlwaysEat,
		&f.EatSeconds,
		// TODO
	}.ReadFrom(r)
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (f *Food) WriteTo(w io.Writer) (n int64, err error) {
	pk.Tuple{
		&f.Nutrition,
		&f.Saturation,
		&f.CanAlwaysEat,
		&f.EatSeconds,
		// TODO
	}.WriteTo(w)
	panic("unimplemented")
}

package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*PotionContents)(nil)

type PotionContents struct {
	PotionID    pk.Option[pk.VarInt, *pk.VarInt]
	CustomColor pk.Option[pk.Int, *pk.Int]
	PotionEffects []any
}

// ID implements DataComponent.
func (PotionContents) ID() string {
	return "minecraft:potion_contents"
}

// ReadFrom implements DataComponent.
func (p *PotionContents) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (p *PotionContents) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}

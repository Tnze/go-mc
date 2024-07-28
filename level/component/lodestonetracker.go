package component

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*LodestoneTracker)(nil)

type LodestoneTracker struct {
	HasGlobalPosition pk.Boolean
	Dimension         pk.Identifier
	Position          pk.Position
	Tracked           pk.Boolean
}

// ID implements DataComponent.
func (LodestoneTracker) ID() string {
	return "minecraft:lodestone_tracker"
}

// ReadFrom implements DataComponent.
func (l *LodestoneTracker) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		&l.HasGlobalPosition,
		&l.Dimension,
		&l.Position,
		&l.Tracked,
	}.ReadFrom(r)
}

// WriteTo implements DataComponent.
func (l *LodestoneTracker) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		&l.HasGlobalPosition,
		&l.Dimension,
		&l.Position,
		&l.Tracked,
	}.WriteTo(w)
}

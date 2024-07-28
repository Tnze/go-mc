package component

import (
	"io"

	"github.com/Tnze/go-mc/level/block"
	pk "github.com/Tnze/go-mc/net/packet"
)

var _ DataComponent = (*DebugStickState)(nil)

type DebugStickState struct {
	Data block.State
}

// ID implements DataComponent.
func (DebugStickState) ID() string {
	return "minecraft:debug_stick_state"
}

// ReadFrom implements DataComponent.
func (d *DebugStickState) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.NBT(&d.Data).ReadFrom(r)
}

// WriteTo implements DataComponent.
func (d *DebugStickState) WriteTo(w io.Writer) (n int64, err error) {
	return pk.NBT(&d.Data).WriteTo(w)
}

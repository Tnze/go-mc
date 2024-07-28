package component

import (
	"io"
)

var _ DataComponent = (*JukeboxPlayable)(nil)

type JukeboxPlayable struct {
	// TODO
}

// ID implements DataComponent.
func (JukeboxPlayable) ID() string {
	return "minecraft:jukebox_playable"
}

// ReadFrom implements DataComponent.
func (j *JukeboxPlayable) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (j *JukeboxPlayable) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}

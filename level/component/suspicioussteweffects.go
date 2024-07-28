package component

import (
	"io"
)

var _ DataComponent = (*SuspiciousStewEffects)(nil)

type SuspiciousStewEffects struct {
	Effects []any
}

// ID implements DataComponent.
func (SuspiciousStewEffects) ID() string {
	return "minecraft:suspicious_stew_effects"
}

// ReadFrom implements DataComponent.
func (s *SuspiciousStewEffects) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (s *SuspiciousStewEffects) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}

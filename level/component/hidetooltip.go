package component

import "io"

var _ DataComponent = (*HideTooptip)(nil)

type HideTooptip struct{}

// ID implements DataComponent.
func (HideTooptip) ID() string {
	return "minecraft:hide_tooltip"
}

// ReadFrom implements DataComponent.
func (h *HideTooptip) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, nil
}

// WriteTo implements DataComponent.
func (h *HideTooptip) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

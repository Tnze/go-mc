package component

import "io"

var _ DataComponent = (*HideAdditionalTooptip)(nil)

type HideAdditionalTooptip struct{}

// ID implements DataComponent.
func (HideAdditionalTooptip) ID() string {
	return "minecraft:hide_additional_tooltip"
}

// ReadFrom implements DataComponent.
func (h *HideAdditionalTooptip) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, nil
}

// WriteTo implements DataComponent.
func (h *HideAdditionalTooptip) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

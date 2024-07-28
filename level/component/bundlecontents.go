package component

import "io"

var _ DataComponent = (*BundleContents)(nil)

type BundleContents struct {
	// TODO
}

// ID implements DataComponent.
func (BundleContents) ID() string {
	return "minecraft:bundle_contents"
}

// ReadFrom implements DataComponent.
func (b *BundleContents) ReadFrom(r io.Reader) (n int64, err error) {
	panic("unimplemented")
}

// WriteTo implements DataComponent.
func (b *BundleContents) WriteTo(w io.Writer) (n int64, err error) {
	panic("unimplemented")
}

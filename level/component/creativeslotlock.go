package component

import "io"

var _ DataComponent = (*CreativeSlotLock)(nil)

type CreativeSlotLock struct{}

// ID implements DataComponent.
func (c *CreativeSlotLock) ID() string {
	return "minecraft:creative_slot_lock"
}

// ReadFrom implements DataComponent.
func (c *CreativeSlotLock) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, nil
}

// WriteTo implements DataComponent.
func (c *CreativeSlotLock) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

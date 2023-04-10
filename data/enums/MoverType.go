package enums

// MoverType is the type of the entity's movement.
type MoverType int

const (
	MoverTypeSelf MoverType = iota
	MoverTypePlayer
	MoverTypePiston
	MoverTypeShulkerBox
	MoverTypeShulker
)

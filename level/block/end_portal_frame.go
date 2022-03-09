package block

type EndPortalFrame struct {
	Eye    string
	Facing string
}

func (EndPortalFrame) ID() string {
	return "minecraft:end_portal_frame"
}

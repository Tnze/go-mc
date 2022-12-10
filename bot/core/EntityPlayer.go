package core

type EntityPlayer struct {
	*EntityLiving
	DisplayName string
	expBar      float32
	TotalExp    int32
	Level       int32
}

func (e *EntityPlayer) SetExp(bar float32, exp, total int32) {
	e.expBar, e.TotalExp, e.Level = bar, total, exp
}

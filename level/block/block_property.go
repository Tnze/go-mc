package block

type BlockProperty struct {
	//material Material
	HasCollision bool `nbt:"HasCollision"`
	//soundType SoundType
	ExplosionResistance        float64 `nbt:"ExplosionResistance"`
	DestroyTime                float64 `nbt:"DestroyTime"`
	RequiresCorrectToolForDrop bool    `nbt:"RequiresCorrectToolForDrop"`
	Friction                   float64 `nbt:"Friction"`
	SpeedFactor                float64 `nbt:"SpeedFactor"`
	JumpFactor                 float64 `nbt:"JumpFaction"`
	CanOcclude                 bool    `nbt:"CanOcclude"`
	IsAir                      bool    `nbt:"IsAir"`
	DynamicShape               bool    `nbt:"DynamicShape"`
}

func NewBlockProperty(other *BlockProperty) *BlockProperty {
	b := &BlockProperty{}
	if other != nil {
		b.copy(other)
	}
	return b
}

func (b *BlockProperty) copy(other *BlockProperty) *BlockProperty {
	if other == nil {
		return b
	}
	b.HasCollision = other.HasCollision
	b.ExplosionResistance = other.ExplosionResistance
	b.DestroyTime = other.DestroyTime
	b.RequiresCorrectToolForDrop = other.RequiresCorrectToolForDrop
	b.Friction = other.Friction
	b.SpeedFactor = other.SpeedFactor
	b.JumpFactor = other.JumpFactor
	b.CanOcclude = other.CanOcclude
	b.IsAir = other.IsAir
	b.DynamicShape = other.DynamicShape
	return b
}

func (b *BlockProperty) setHasCollision(hasCollision bool) *BlockProperty {
	b.HasCollision = hasCollision
	return b
}

func (b *BlockProperty) setExplosionResistance(explosionResistance float64) *BlockProperty {
	b.ExplosionResistance = explosionResistance
	return b
}

func (b *BlockProperty) setDestroyTime(destroyTime float64) *BlockProperty {
	b.DestroyTime = destroyTime
	return b
}

func (b *BlockProperty) setRequiresCorrectTool(requiresCorrectTool bool) *BlockProperty {
	b.RequiresCorrectToolForDrop = requiresCorrectTool
	return b
}

func (b *BlockProperty) setFriction(friction float64) *BlockProperty {
	b.Friction = friction
	return b
}

func (b *BlockProperty) setSpeedFactor(speedFactor float64) *BlockProperty {
	b.SpeedFactor = speedFactor
	return b
}

func (b *BlockProperty) setJumpFactor(jumpFactor float64) *BlockProperty {
	b.JumpFactor = jumpFactor
	return b
}

func (b *BlockProperty) setCanOcclude(canOcculde bool) *BlockProperty {
	b.CanOcclude = canOcculde
	return b
}

func (b *BlockProperty) setIsAir(isAir bool) *BlockProperty {
	b.IsAir = isAir
	return b
}

func (b *BlockProperty) setDynamicShape(dynamicShape bool) *BlockProperty {
	b.DynamicShape = dynamicShape
	return b
}

// func (b *BlockProperty) GetMaterial() Material {
// 	return b.material
// }

// func (b *BlockProperty) SetMaterial(material Material) *BlockProperty {
// 	b.material = material
// 	return b
// }

// func (b *BlockProperty) GetSoundType() SoundType {
// 	return b.soundType
// }

// func (b *BlockProperty) SetSoundType(soundType SoundType) *BlockProperty {
// 	b.soundType = soundType
// 	return b
// }

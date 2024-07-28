package component

import pk "github.com/Tnze/go-mc/net/packet"

type DataComponent interface {
	pk.Field
	ID() string
}

func NewComponent(id int32) DataComponent {
	switch id {
	case 0:
		return new(CustomData)
	case 1:
		return new(MaxStackSize)
	case 2:
		return new(MaxDamage)
	case 3:
		return new(Damage)
	case 4:
		return new(Unbreakable)
	case 5:
		return new(CustomName)
	case 6:
		return new(ItemName)
	case 7:
		return new(Lore)
	case 8:
		return new(Rarity)
	case 9:
		return new(Enchantments)
	case 10:
		return new(CanPlaceOn)
	case 11:
		return new(CanBreak)
	case 12:
		return new(AttributeModifiers)
	case 13:
		return new(CustomModelData)
	case 14:
		return new(HideAdditionalTooptip)
	case 15:
		return new(HideTooptip)
	case 16:
		return new(RepairCost)
	case 17:
		return new(CreativeSlotLock)
	case 18:
		return new(EnchantmentGlintOverride)
	case 19:
		return new(IntangibleProjectile)
	case 20:
		return new(Food)
	case 21:
		return new(FireResistant)
	case 22:
		return new(Tool)
	case 23:
		return new(StoredEnchantments)
	case 24:
		return new(DyedColor)
	case 25:
		return new(MapColor)
	case 26:
		return new(MapID)
	case 27:
		return new(MapDecorations)
	case 28:
	case 29:
	case 30:
	case 31:
	case 32:
	case 33:
	case 34:
	case 35:
	case 36:
	case 37:
	case 38:
	case 39:
	case 40:
	case 41:
	case 42:
	case 43:
		return new(Recipes)
	case 44:
	case 45:
	case 46:
	case 47:
	case 48:
	case 49:
	case 50:
	case 51:
	case 52:
	case 53:
	case 54:
	case 55:
	case 56:
	}
	return nil
}

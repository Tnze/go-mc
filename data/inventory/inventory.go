package inventory

type InventoryID = int32

const (
	Generic9x1 InventoryID = iota
	Generic9x2
	Generic9x3
	Generic9x4
	Generic9x5
	Generic9x6
	Generic3x3
	Crafter3x3
	Anvil
	Beacon
	BlastFurnace
	BrewingStand
	Crafting
	Enchantment
	Furnace
	Grindstone
	Hopper
	Lectern
	Loom
	Merchant
	ShulkerBox
	Smithing
	Smoker
	Cartography
	Stonecutter
)

func IDToName(t InventoryID) string {
	switch t {
	case Generic9x1:
		return "generic_9x1"
	case Generic9x2:
		return "generic_9x2"
	case Generic9x3:
		return "generic_9x3"
	case Generic9x4:
		return "generic_9x4"
	case Generic9x5:
		return "generic_9x5"
	case Generic9x6:
		return "generic_9x6"
	case Generic3x3:
		return "generic_3x3"
	case Crafter3x3:
		return "crafter_3x3"
	case Anvil:
		return "anvil"
	case Beacon:
		return "beacon"
	case BlastFurnace:
		return "blast_furnace"
	case BrewingStand:
		return "brewing_stand"
	case Crafting:
		return "crafting"
	case Enchantment:
		return "enchantment"
	case Furnace:
		return "furnace"
	case Grindstone:
		return "grindstone"
	case Hopper:
		return "hopper"
	case Lectern:
		return "lectern"
	case Loom:
		return "loom"
	case Merchant:
		return "merchant"
	case ShulkerBox:
		return "shulker_box"
	case Smithing:
		return "smithing"
	case Smoker:
		return "smoker"
	case Cartography:
		return "cartography"
	case Stonecutter:
		return "stonecutter"
	}
	return ""
}

func NameToID(name string) InventoryID {
	switch name {
	case "generic_9x1":
		return Generic9x1
	case "generic_9x2":
		return Generic9x2
	case "generic_9x3":
		return Generic9x3
	case "generic_9x4":
		return Generic9x4
	case "generic_9x5":
		return Generic9x5
	case "generic_9x6":
		return Generic9x6
	case "generic_3x3":
		return Generic3x3
	case "crafter_3x3":
		return Crafter3x3
	case "anvil":
		return Anvil
	case "beacon":
		return Beacon
	case "blast_furnace":
		return BlastFurnace
	case "brewing_stand":
		return BrewingStand
	case "crafting":
		return Crafting
	case "enchantment":
		return Enchantment
	case "furnace":
		return Furnace
	case "grindstone":
		return Grindstone
	case "hopper":
		return Hopper
	case "lectern":
		return Lectern
	case "loom":
		return Loom
	case "merchant":
		return Merchant
	case "shulker_box":
		return ShulkerBox
	case "smithing":
		return Smithing
	case "smoker":
		return Smoker
	case "cartography":
		return Cartography
	case "stonecutter":
		return Stonecutter
	}
	return -1
}

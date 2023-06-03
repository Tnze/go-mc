package screen

import (
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/data/grids"
	"github.com/Tnze/go-mc/data/slots"
)

type Manager struct {
	Screens       map[int]Container
	CurrentScreen *Container
	Inventory     *grids.GenericInventory
	Cursor        *slots.Slot
	HeldItem      *slots.Slot
	// The last state received from the server
	StateID int32
}

func NewManager() *Manager {
	inventory := new(grids.GenericInventory)
	return &Manager{
		Screens:   fillContainers(inventory),
		Inventory: inventory,
	}
}

func fillContainers(inventory *grids.GenericInventory) map[int]Container {
	return map[int]Container{
		0:  grids.NewGeneric9x1(inventory),
		1:  grids.NewGeneric9x2(inventory),
		2:  grids.NewGeneric9x3(inventory),
		3:  grids.NewGeneric9x4(inventory),
		4:  grids.NewGeneric9x5(inventory),
		5:  grids.NewGeneric9x6(inventory),
		6:  grids.NewGeneric3x3(inventory),
		7:  grids.NewAnvil(inventory),
		8:  grids.NewBeacon(inventory),
		9:  grids.NewBlastFurnace(inventory),
		10: grids.NewBrewingStand(inventory),
		11: grids.NewCraftingTable(inventory),
		12: grids.NewEnchantmentTable(inventory),
		13: grids.NewFurnace(inventory),
		14: grids.NewGrindstone(inventory),
		15: grids.NewHopper(inventory),
		16: grids.InitGenericContainer("nil", 0, 0, inventory), // TODO: This is the only one that is not a container, I don't know why mojang did this.
		17: grids.NewLoom(inventory),
		18: grids.NewMerchant(inventory),
		19: grids.NewShulkerBox(inventory),
		20: grids.NewSmithingTable(inventory),
		21: grids.NewSmoker(inventory),
		22: grids.NewCartographyTable(inventory),
		23: grids.NewStonecutter(inventory),
	}
}

type Container interface {
	GetSlot(int) *slots.Slot
	SetSlot(int, slots.Slot) basic.Error
	ApplyData([]slots.Slot)
	OnClose() basic.Error
}

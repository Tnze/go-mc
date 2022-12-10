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
	return &Manager{
		Screens:   fillContainers(),
		Inventory: new(grids.GenericInventory),
	}
}

func fillContainers() map[int]Container {
	return map[int]Container{
		0:  grids.NewGeneric_9x1(),
		1:  grids.NewGeneric_9x2(),
		2:  grids.NewGeneric_9x3(),
		3:  grids.NewGeneric_9x4(),
		4:  grids.NewGeneric_9x5(),
		5:  grids.NewGeneric_9x6(),
		6:  grids.NewGeneric_3x3(),
		7:  grids.NewAnvil(),
		8:  grids.NewBeacon(),
		9:  grids.NewBlastFurnace(),
		10: grids.NewBrewingStand(),
		11: grids.NewCraftingTable(),
		12: grids.NewEnchantmentTable(),
		13: grids.NewFurnace(),
		14: grids.NewGrindstone(),
		15: grids.NewHopper(),
		//16: grids.NewLectern(), // TODO: This is the only one that is not a container
		17: grids.NewLoom(),
		18: grids.NewMerchant(),
		19: grids.NewShulkerBox(),
		20: grids.NewSmithingTable(),
		21: grids.NewSmoker(),
		22: grids.NewCartographyTable(),
		23: grids.NewStonercutter(),
	}
}

type Container interface {
	GetSlot(int) *slots.Slot
	GetInventorySlot(int) *slots.Slot
	GetHotbarSlot(int) *slots.Slot
	SetSlot(int, slots.Slot) basic.Error
	SwitchSlot(int, int) basic.Error
	OnClose() basic.Error
}

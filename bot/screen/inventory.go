package screen

import (
	"fmt"

	"github.com/Tnze/go-mc/data/item"
	"github.com/Tnze/go-mc/nbt"
)

type Inventory struct {
}

func (inv Inventory) SetSlot(i int, id int32, count byte, NBT nbt.RawMessage) {
	// TODO: accept inv data
	fmt.Printf("Inventory[%d] = minecraft:%v * %d\n", i, item.ByID[item.ID(id)].Name, count)
}

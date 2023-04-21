package block

import (
	"fmt"
	"testing"
)

func TestBlock(t *testing.T) {
	var block = NewBlock("air", NewBlockProperty(nil).setHasCollision(false).setIsAir(true))
	fmt.Println(block)
}

func TestBlockStateHolder(t *testing.T) {
	var block = NewBlock("air", NewBlockProperty(nil).setHasCollision(false).setIsAir(true))
	var stateHolder = NewBlockState(&block, map[uint64]any{
		FacingProperty.HashCode():   North,
		Age15Property.HashCode():    0,
		MoistureProperty.HashCode(): 0,
		BedPartProperty.HashCode():  BedPartFoot,
	})

	// Initial print
	fmt.Println(stateHolder.GetValue(Age15Property), stateHolder.GetValue(MoistureProperty), stateHolder.GetValue(BedPartProperty))

	// Simulate a block update
	stateHolder.SetValue(Age15Property, 1)
	stateHolder.SetValue(MoistureProperty, 5)
	stateHolder.SetValue(BedPartProperty, BedPartHead)

	// Updated print
	fmt.Println(stateHolder.GetValue(Age15Property), stateHolder.GetValue(MoistureProperty), stateHolder.GetValue(BedPartProperty))
}

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
	var block = TurtleEgg

	fmt.Println("Default state id", block.GetDefaultValue())

	// Initial print
	fmt.Println(block.GetValue(HatchProperty), block.GetValue(EggsProperty), ToStateID[block])

	// Simulate a block update
	block.SetValue(HatchProperty, 1)
	block.SetValue(EggsProperty, 3)

	// Updated print
	fmt.Println(block.GetValue(HatchProperty), block.GetValue(EggsProperty), ToStateID[block])
}

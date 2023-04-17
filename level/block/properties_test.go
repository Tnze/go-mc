package block

import (
	"fmt"
	"testing"
)

var (
	chestType = NewPropertyEnum[ChestType]("type", map[string]ChestType{
		"single": ChestTypeSingle,
		"left":   ChestTypeLeft,
		"right":  ChestTypeRight,
	})
	moisture          = NewPropertyInteger("moisture", 0, 7)
	stabilityDistance = NewPropertyInteger("distance", 0, 7)
	stairsShape       = NewPropertyEnum[StairsShape]("shape", map[string]StairsShape{
		"straight":    StairsShapeStraight,
		"inner_left":  StairsShapeInnerLeft,
		"inner_right": StairsShapeInnerRight,
		"outer_left":  StairsShapeOuterLeft,
		"outer_right": StairsShapeOuterRight,
	})
)

func TestNewPropertyEnum(t *testing.T) {
	fmt.Println(chestType, moisture, stabilityDistance, stairsShape)
}

// gen_shape.go generates block shape information.

//+build ignore

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
)

const (
	infoURL = "https://raw.githubusercontent.com/PrismarineJS/minecraft-data/master/data/pc/1.16.1/blockCollisionShapes.json"
)

func downloadInfo() (map[string]interface{}, error) {
	resp, err := http.Get(infoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func main() {
	info, err := downloadInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(`// Package shape stores information about the shapes of blocks in Minecraft.
package shape

import (
  "github.com/Tnze/go-mc/bot/world"
  "github.com/Tnze/go-mc/data/block"
)

// ID describes a numeric shape ID.
type ID uint32

// Shape describes how collisions should be tested for an object.
type Shape struct {
	ID    ID
  Boxes []BoundingBox
}

type BoundingTriplet struct {
  X, Y, Z float64
}

type BoundingBox struct {
  Min,Max BoundingTriplet
}

// CollisionBoxes returns the set of bounding boxes for that block state ID.
func CollisionBoxes(bStateID world.BlockStatus) ([]BoundingBox, error) {
  bID := block.StateID[uint32(bStateID)]
  if bID == 0 {
    return nil, fmt.Errorf("unknown state ID: %v", bStateID)
  }
  b, ok := block.ByID[bID]
  if !ok {
    return nil, fmt.Errorf("unknown block ID: %v", bID)
  }
  shapes, ok := ByBlockID[bID]
  if !ok {
    return nil, fmt.Errorf("unknown shape for block ID: %v", bID)
  }
  shapeIdx := (uint32(bStateID) - b.MinStateID) % uint32(len(shapes))
  if int(shapeIdx) > len(shapes) {
    return nil, fmt.Errorf("shape index out of bounds: %v >= %v", shapeIdx, len(shapes))
  }
  return Dimensions[shapes[shapeIdx]].Boxes, nil
}

`)

	fmt.Println()
	fmt.Println()
	fmt.Println("// ByBlockID is an index of shapes for each minecraft block variant.")
	fmt.Println("var ByBlockID = map[block.ID][]ID{")
	blocks := info["blocks"].(map[string]interface{})
	for name, shapes := range blocks {
		switch s := shapes.(type) {
		case []interface{}:
			set := make([]string, len(s))
			for i := range s {
				set[i] = fmt.Sprint(s[i])
			}
			fmt.Printf("  block.%s.ID: []ID{%s},\n", strcase.ToCamel(name), strings.Join(set, ", "))
		default:
			fmt.Printf("  block.%s.ID: []ID{%s},\n", strcase.ToCamel(name), fmt.Sprint(s))
		}
	}
	fmt.Println("}")

	fmt.Println()
	fmt.Println()
	fmt.Println("// Dimensions describes the bounding boxes of a shape ID.")
	fmt.Println("var Dimensions = map[ID]Shape{")
	shapes := info["shapes"].(map[string]interface{})
	for id, boxes := range shapes {
		fmt.Printf("  %s: Shape{\n", id)
		fmt.Printf("    ID: %s,\n", id)
		fmt.Printf("    Boxes: []BoundingBox{\n")
		for _, box := range boxes.([]interface{}) {
			elements := box.([]interface{})
			if len(elements) != 6 {
				panic("expected 6 elements")
			}
			fmt.Printf("      {\n")
			fmt.Printf("        Min: BoundingTriplet{X: %v, Y: %v, Z: %v},\n", elements[0], elements[1], elements[2])
			fmt.Printf("        Max: BoundingTriplet{X: %v, Y: %v, Z: %v},\n", elements[3], elements[4], elements[5])
			fmt.Printf("      },\n")
		}
		fmt.Printf("    },\n")
		fmt.Println("  },")
	}
	fmt.Println("}")
}

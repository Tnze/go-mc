//go:build generate
// +build generate

package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"github.com/Tnze/go-mc/level/block"
	"github.com/Tnze/go-mc/nbt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/Tnze/go-mc/internal/generateutils"
)

//go:embed blocks.go.tmpl
var tempSource string

var temp = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst": generateutils.UpperTheFirst,
		"ToGoTypeName":  generateutils.ToGoTypeName,
		"ToStructLiteral": func(s interface{}) string {
			return fmt.Sprintf("%#v", s)[6:]
		},
		"GetDefaultValues": GetDefaultValues,
		"Generator":        func() string { return "generator/blocks/main.go" },
	}).
	Parse(tempSource),
)

//go:embed states.go.tmpl
var tempSource2 string

var temp2 = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst": generateutils.UpperTheFirst,
		"ToGoTypeName":  generateutils.ToGoTypeName,
		"ToStructLiteral": func(s interface{}) string {
			return fmt.Sprintf("%#v", s)[6:]
		},
		"GetProperty": GetProperty,
		"CanContinue": CanContinue,
		"GetRealName": GetRealName,
		"Generator":   func() string { return "generator/blocks/main.go" },
	}).
	Parse(tempSource2),
)

type State struct {
	Name       string
	Properties block.BlockProperty
	Default    map[string]any
}

//go:generate go run $GOFILE
//go:generate go fmt blocks.go
//go:generate go fmt states_properties.go
func main() {
	fmt.Println("Generating source file...")
	var states []State
	readBlockStates(&states)

	// generate go source file
	genSourceFile(states)
}

func readBlockStates(states *[]State) {
	// open block_states data file
	f, err := os.Open("blocks.nbt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		log.Panic(err)
	}

	// parse the nbt format
	if _, err := nbt.NewDecoder(r).Decode(states); err != nil {
		log.Panic(err)
	}
}

func genSourceFile(states []State) {
	var source bytes.Buffer
	if err := temp.Execute(&source, states); err != nil {
		log.Panic(err)
	}

	err := os.WriteFile("blocks.go", source.Bytes(), 0666)
	if err != nil {
		log.Panic(err)
	}

	source.Reset()

	if err = temp2.Execute(&source, states); err != nil {
		log.Panic(err)
	}

	formatted, err := format.Source(source.Bytes())
	if err != nil {
		log.Panic(err)
	}

	err = os.WriteFile("states_properties.go", formatted, 0666)
	if err != nil {
		panic(err)
	}
}

var exist = map[string]bool{}

func GetProperty(key string, value any) string {
	if key == "Age" {
		exist[key+fmt.Sprintf("%v", value)] = true
	} else {
		exist[key] = true
	}
	switch value.(type) {
	case int8:
		return "NewPropertyBoolean(\"" + uncapitalize(key) + "\")"
	case int32:
		return "NewPropertyInteger(\"" + uncapitalize(key) + "\", " + fmt.Sprint(separateIntOr(toInt(value))[0]) + ", " + fmt.Sprint(separateIntOr(toInt(value))[1]) + ")"
	default:
		return ""
	}
}

func GetDefaultValues(mapped map[string]any) string {
	if len(mapped) == 0 {
		return ""
	}
	var result string
	for key, value := range mapped {
		switch value.(type) {
		case int8:
			result += fmt.Sprintf(".registerState(%s, %v)", GetRealName(key, value)+"Property", value != 0)
		case int32:
			// Here we can assume that the default state is the first state
			result += fmt.Sprintf(".registerState(%s, %v)", GetRealName(key, value)+"Property", separateIntOr(toInt(value))[0])
		case string:
			result += fmt.Sprintf(".registerState(states.%s, states.%v)", GetRealName(key, value)+"Property", GetTrueValue(key, upperTheFirst(value.(string))))
		}
	}
	return result
}

func CanContinue(key string, value any) bool {
	if key == "Age" {
		return !exist[key+fmt.Sprintf("%v", value)] && GetProperty(key, value) != ""
	} else {
		return !exist[key] && GetProperty(key, value) != ""
	}
}

func GetRealName(key string, value any) string {
	if key == "Age" {
		value = separateIntOr(toInt(value))[1]
		return upperTheFirst("Age" + fmt.Sprintf("%v", value))
	}
	return upperTheFirst(GetTrueType(key, value))
}

func GetTrueType(str string, class any) string {
	switch str {
	case "Shape":
		if strings.Contains(class.(string), "RailShape") {
			return "RailShape"
		} else if strings.Contains(class.(string), "StairsShape") {
			return "StairsShape"
		}
	case "Face":
		if strings.Contains(class.(string), "AttachFace") {
			return "AttachFace"
		}
	case "Type":
		if strings.Contains(class.(string), "PistonType") {
			return "PistonType"
		} else if strings.Contains(class.(string), "SlabType") {
			return "SlabType"
		} else if strings.Contains(class.(string), "ChestType") {
			return "ChestType"
		}
	case "Part":
		if strings.Contains(class.(string), "BedPart") {
			return "BedPart"
		}
	case "South", "West", "North", "East":
		switch class.(type) {
		case string:
			if strings.Contains(class.(string), "Redstone") {
				return str + "Redstone"
			}
			if strings.Contains(class.(string), "WallSide") {
				return str + "Wall"
			}
		}
	case "Hinge":
		if strings.Contains(class.(string), "DoorHinge") {
			return "DoorHinge"
		}
	case "Mode":
		if strings.Contains(class.(string), "ComparatorMode") {
			return "ComparatorMode"
		}
		if strings.Contains(class.(string), "StructureMode") {
			return "StructureMode"
		}
	case "Leaves":
		if strings.Contains(class.(string), "Leaves") {
			return "BambooLeaves"
		}
	case "Attachment":
		if strings.Contains(class.(string), "BellAttach") {
			return "BellAttachment"
		}
	case "Thickness":
		if strings.Contains(class.(string), "DripstoneThickness") {
			return "DripstoneThickness"
		}
	}
	return str
}

func GetTrueValue(str string, class any) any {
	switch str {
	case "Orientation":
		switch class.(type) {
		case string:
			if strings.Contains(class.(string), "FrontAndTop") {
				return strings.TrimPrefix(class.(string), "FrontAndTop")
			}
		}
	}

	return class
}

func upperTheFirst(word string) string {
	var sb strings.Builder
	for _, word := range strings.Split(word, "_") {
		runes := []rune(word)
		if len(runes) > 0 {
			runes[0] = unicode.ToUpper(runes[0])
		}
		sb.WriteString(string(runes))
	}
	return sb.String()
}

func uncapitalize(s string) string {
	return strings.ToLower(s)
}

func separateIntOr(value int32) [2]int32 {
	a := value >> 16
	b := value & 0xffff
	if b == 0 {
		b++
	}
	return [2]int32{a, b}
}

func toInt(s any) int32 {
	switch s.(type) {
	case int32:
		return s.(int32)
	default:
		panic("not a valid int")
	}
}

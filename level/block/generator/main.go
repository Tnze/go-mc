package main

import (
	_ "embed"
	"log"
	"os"
	"strings"
	"text/template"
	"unicode"

	"github.com/Tnze/go-mc/nbt"
)

//go:embed blocks.go.tmpl
var tempSource string

var temp = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst": UpperTheFirst,
		"ToGoTypeName":  ToGoTypeName,
		"GetType":       GetType,
		"Generator":     func() string { return "generator/main.go" },
	}).
	Parse(tempSource))

type State struct {
	Name string
	Meta map[string]string
}

func main() {
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

	// parse the nbt format
	if _, err := nbt.NewDecoder(f).Decode(states); err != nil {
		log.Panic(err)
	}
}

func genSourceFile(states []State) {
	file, err := os.Create("blocks.go")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	// clean up the file
	if err := file.Truncate(0); err != nil {
		return
	}

	if err := temp.Execute(file, states); err != nil {
		log.Panic(err)
	}
}

func ToGoTypeName(name string) string {
	name = strings.TrimPrefix(name, "minecraft:")
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = UpperTheFirst(words[i])
	}
	return strings.Join(words, "")
}

var typeMaps = map[string]string{
	"BooleanProperty":   "Boolean",
	"DirectionProperty": "Direction",
	"EnumProperty":      "string",
	"IntegerProperty":   "Integer",
}

func GetType(v string) string {
	if mapped, ok := typeMaps[v]; ok {
		return mapped
	}
	return v
}

func UpperTheFirst(word string) string {
	runes := []rune(word)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

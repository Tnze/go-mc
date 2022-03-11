package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"go/format"
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
		"Generator":     func() string { return "generator/blocks/main.go" },
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

	formattedSource, err := format.Source(source.Bytes())
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("blocks.go", formattedSource, 0666)
	if err != nil {
		panic(err)
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

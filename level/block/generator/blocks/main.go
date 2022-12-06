package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/Tnze/go-mc/internal/generateutils"
	"github.com/Tnze/go-mc/nbt"
)

//go:embed blocks.go.tmpl
var tempSource string

var temp = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst": generateutils.UpperTheFirst,
		"ToGoTypeName":  generateutils.ToGoTypeName,
		"GetType":       GetType,
		"Generator":     func() string { return "generator/blocks/main.go" },
	}).
	Parse(tempSource),
)

type State struct {
	Name string
	Meta map[string]string
}

func main() {
	fmt.Println("Generating source file...")
	var states []State
	readBlockStates(&states)

	// generate go source file
	genSourceFile(states)
}

func readBlockStates(states *[]State) {
	// open block_states data file
	f, err := os.Open("./level/block/blocks.nbt")
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

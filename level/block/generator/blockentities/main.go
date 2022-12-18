package main

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/Tnze/go-mc/internal/generateutils"
	"github.com/Tnze/go-mc/nbt"
)

//go:embed blockentities.go.tmpl
var tempSource string

var temp = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst":      generateutils.UpperTheFirst,
		"ToGoTypeName":       generateutils.ToGoTypeName,
		"ToFuncReceiverName": generateutils.ToFuncReceiverName,
		"Generator":          func() string { return "generator/blockentities/main.go" },
	}).
	Parse(tempSource),
)

type BlockEntity struct {
	Name        string
	ValidBlocks []string
}

func main() {
	var states []BlockEntity
	readBlockEntities(&states)

	// generate go source file
	genSourceFile(states)
}

func readBlockEntities(states *[]BlockEntity) {
	f, err := os.Open("block_entities.nbt")
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

func genSourceFile(states []BlockEntity) {
	var source bytes.Buffer
	if err := temp.Execute(&source, states); err != nil {
		log.Panic(err)
	}

	formattedSource, err := format.Source(source.Bytes())
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("blockentities.go", formattedSource, 0o666)
	if err != nil {
		panic(err)
	}
	log.Print("Generated blockentities.go")
}

package main

import (
	"bytes"
	_ "embed"
	"encoding/gob"
	"fmt"
	"github.com/Tnze/go-mc/data/block"
	"image"
	"image/png"
	"log"
	"os"
)

func savePng(img image.Image, name string) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

//go:embed colors.gob
var colorsBin []byte // gob([]color.RGBA64)

func init() {
	if err := gob.NewDecoder(bytes.NewReader(colorsBin)).Decode(&colors); err != nil {
		panic(err)
	}
}

func usage() {
	_, _ = fmt.Fprintf(os.Stderr, "usage: %s [-region <region path>] \n", os.Args[0])
	os.Exit(1)
}

func mkmax(c, n *int) {
	if *c < *n {
		*c = *n
	}
}
func mkmin(c, n *int) {
	if *c > *n {
		*c = *n
	}
}

var idByName = make(map[string]uint32, len(block.ByID))

func init() {
	for _, v := range block.ByID {
		idByName["minecraft:"+v.Name] = uint32(v.ID)
	}
}

package main

import (
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
	"unicode"

	"github.com/Tnze/go-mc/nbt"
)

const tempSource = `package block

type {{.Name | ToGoTypeName}} struct { {{range $key, $elem := .Properties}}
	{{$key | UpperTheFirst}}	{{$elem | GetType}} {{ end }}
}

func ({{.Name | ToGoTypeName}}) ID() string {
	return {{.Name | printf "%q"}}
}
`

var temp = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst": UpperTheFirst,
		"ToGoTypeName":  ToGoTypeName,
		"GetType":       GetType,
	}).
	Parse(tempSource))

type State struct {
	Name       string
	Properties map[string]interface{}
}

func main() {
	var states []State
	readBlockStates(&states)

	// generate go sources for each blocks
	for _, state := range states {
		genSourceFile(state)
	}
}

func readBlockStates(states *[]State) {
	// open block_states data file
	f, err := os.Open("testdata/block_states.nbt")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	// parse the nbt format
	if _, err := nbt.NewDecoder(f).Decode(states); err != nil {
		log.Panic(err)
	}
}

func genSourceFile(state State) {
	filename := strings.TrimPrefix(state.Name, "minecraft:") + ".go"
	file, err := os.Create(filename)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	// clean up the file
	if err := file.Truncate(0); err != nil {
		return
	}

	if err := temp.Execute(file, state); err != nil {
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

func GetType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func UpperTheFirst(word string) string {
	runes := []rune(word)
	if len(runes) > 0 {
		runes[0] = unicode.ToUpper(runes[0])
	}
	return string(runes)
}

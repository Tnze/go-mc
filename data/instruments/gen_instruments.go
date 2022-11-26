//go:build generate
// +build generate

package main

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"

	"github.com/iancoleman/strcase"
)

const (
	version = "1.19"
	infoURL = "https://raw.githubusercontent.com/PrismarineJS/minecraft-data/master/data/pc/" + version + "/instruments.json"
	//language=gohtml
	tmpl = `// Code generated by gen_instruments.go DO NOT EDIT.
// Package instruments stores information about instruments in Minecraft.
package instruments
// ID describes the numeric ID of an instrument.
type ID uint32

// Instrument describes information about a type of instrument.
type Instrument struct {
	ID          ID
	Name        string
}

var (
	{{- range .}}
	{{.CamelName}} = Instrument{
		ID: {{.ID}},
		Name: "{{.Name}}",
	}{{end}}
)

// ByID is an index of minecraft instruments by their ID.
var ByID = map[ID]*Instrument{ {{range .}}
	{{.ID}}: &{{.CamelName}},{{end}}
}

// ByName is an index of minecraft instruments by their name.
var ByName = map[string]*Instrument{ {{range .}}
	"{{.Name}}": &{{.CamelName}},{{end}}
}`
)

type Instrument struct {
	ID        uint32 `json:"id"`
	CamelName string `json:"-"`
	Name      string `json:"name"`
}

func downloadInfo() ([]*Instrument, error) {
	resp, err := http.Get(infoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []*Instrument
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	for _, d := range data {
		d.CamelName = strcase.ToCamel(d.Name)
	}
	return data, nil
}

//go:generate go run $GOFILE
//go:generate go fmt instruments.go
func main() {
	instruments, err := downloadInfo()
	if err != nil {
		panic(err)
	}

	f, err := os.Create("instruments.go")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := template.Must(template.New("").Parse(tmpl)).Execute(f, instruments); err != nil {
		panic(err)
	}
}

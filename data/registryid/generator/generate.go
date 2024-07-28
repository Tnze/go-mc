package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Tnze/go-mc/internal/generateutils"
)

type registry struct {
	Default string `json:"default"`
	Entries map[string]struct {
		ProtocolID int `json:"protocol_id"`
	}
	ProtocolID int `json:"protocol_id"`
}

// This file is generated with following command
//
//	java -DbundlerMainClass="net.minecraft.data.Main" -jar server.jar --all
//
// And you can found it at the generated\reports\ folder.
//
//go:embed registries.json
var registersJson []byte

//go:embed template.go.tmpl
var tempSource string

type tempData struct {
	PackageName string
	Default     string
	Entries     []string
	TypeName    string
}

var temp = template.Must(template.
	New("block_template").
	Funcs(template.FuncMap{
		"UpperTheFirst": generateutils.UpperTheFirst,
		"ToGoTypeName":  generateutils.ToGoTypeName,
		"Generator":     func() string { return "data/registry/generate.go" },
	}).
	Parse(tempSource),
)

func main() {
	var registries map[string]registry
	if err := json.Unmarshal(registersJson, &registries); err != nil {
		log.Fatal(err)
	}

	for key, reg := range registries {
		registryName := strings.TrimPrefix(key, "minecraft:")
		typeName := generateutils.ToGoTypeName(strings.ReplaceAll(registryName, "/", "_"))
		filename := strings.NewReplacer("_", "", "/", "_").Replace(registryName)
		generateRegistry(reg, typeName, filename)
	}
}

func generateRegistry(r registry, typeName, filename string) {
	entries := make([]string, len(r.Entries))
	for name, v := range r.Entries {
		entries[v.ProtocolID] = name
	}

	var buff bytes.Buffer
	err := temp.Execute(&buff, tempData{
		PackageName: filename,
		Default:     r.Default,
		Entries:     entries,
		TypeName:    typeName,
	})
	if err != nil {
		log.Fatal(err)
	}

	formattedSource, err := format.Source(buff.Bytes())
	if err != nil {
		log.Print(filename, err)
		formattedSource = buff.Bytes()
	}

	err = os.WriteFile(filepath.Join("..", filename+".go"), formattedSource, 0o666)
	if err != nil {
		log.Fatal(err)
	}
}

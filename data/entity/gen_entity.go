// gen_entity.go generates entity information.

//+build ignore

package main

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/iancoleman/strcase"
)

const (
	infoURL = "https://raw.githubusercontent.com/PrismarineJS/minecraft-data/master/data/pc/1.16.2/entities.json"
)

type Entity struct {
	ID          uint32 `json:"id"`
	InternalID  uint32 `json:"internalId"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`

	Width  float64 `json:"width"`
	Height float64 `json:"height"`

	Type     string `json:"type"`
	Category string `json:"category"`
}

func downloadInfo() ([]Entity, error) {
	resp, err := http.Get(infoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []Entity
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func makeEntityDeclaration(entities []Entity) *ast.DeclStmt {
	out := &ast.DeclStmt{Decl: &ast.GenDecl{Tok: token.VAR}}

	for _, e := range entities {
		t := reflect.TypeOf(e)
		fields := make([]ast.Expr, t.NumField())

		for i := 0; i < t.NumField(); i++ {
			ft := t.Field(i)

			if ft.Name == "Category" {
				val := &ast.BasicLit{Kind: token.IDENT, Value: "Unknown"}
				switch e.Category {
				case "Passive mobs":
					val.Value = "PassiveMob"
				case "Hostile mobs":
					val.Value = "HostileMob"
				case "UNKNOWN":
				default:
					val.Value = e.Category
				}
				fields[i] = &ast.KeyValueExpr{
					Key:   &ast.Ident{Name: ft.Name},
					Value: val,
				}
				continue
			}

			var val ast.Expr
			switch ft.Type.Kind() {
			case reflect.Uint32, reflect.Int:
				val = &ast.BasicLit{Kind: token.INT, Value: fmt.Sprint(reflect.ValueOf(e).Field(i))}
			case reflect.Float64:
				val = &ast.BasicLit{Kind: token.FLOAT, Value: fmt.Sprint(reflect.ValueOf(e).Field(i))}
			case reflect.String:
				val = &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(reflect.ValueOf(e).Field(i).String())}
			case reflect.Bool:
				val = &ast.BasicLit{Kind: token.IDENT, Value: fmt.Sprint(reflect.ValueOf(e).Field(i).Bool())}

			case reflect.Slice:
				val = &ast.CompositeLit{
					Type: &ast.ArrayType{
						Elt: &ast.BasicLit{Kind: token.IDENT, Value: ft.Type.Elem().Name()},
					},
				}
				v := reflect.ValueOf(e).Field(i)
				switch ft.Type.Elem().Kind() {
				case reflect.Uint32, reflect.Int:
					for x := 0; x < v.Len(); x++ {
						val.(*ast.CompositeLit).Elts = append(val.(*ast.CompositeLit).Elts, &ast.BasicLit{
							Kind:  token.INT,
							Value: fmt.Sprint(v.Index(x)),
						})
					}
				}
			}

			fields[i] = &ast.KeyValueExpr{
				Key:   &ast.Ident{Name: ft.Name},
				Value: val,
			}
		}

		out.Decl.(*ast.GenDecl).Specs = append(out.Decl.(*ast.GenDecl).Specs, &ast.ValueSpec{
			Names: []*ast.Ident{{Name: strcase.ToCamel(e.Name)}},
			Values: []ast.Expr{
				&ast.CompositeLit{
					Type: &ast.Ident{Name: reflect.TypeOf(e).Name()},
					Elts: fields,
				},
			},
		})
	}

	return out
}

func main() {
	entities, err := downloadInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(`// Package entity stores information about entities in Minecraft.
package entity

// ID describes the numeric ID of an entity.
type ID uint32

// Category groups like entities.
type Category uint8

// Valid entity categories.
const (
	Unknown Category = iota
	Blocks
	Immobile
	Vehicles
	Drops
	Projectiles
	PassiveMob
	HostileMob
)

// Entity describes information about a type of entity.
type Entity struct {
	ID          ID
	InternalID  uint32
	DisplayName string
	Name        string

	Width  float64
	Height float64

	Type     string
	Category Category
}

`)
	format.Node(os.Stdout, token.NewFileSet(), makeEntityDeclaration(entities))

	fmt.Println()
	fmt.Println()
	fmt.Println("// ByID is an index of minecraft entities by their ID.")
	fmt.Println("var ByID = map[ID]*Entity{")
	for _, e := range entities {
		fmt.Printf("  %d: &%s,\n", e.ID, strcase.ToCamel(e.Name))
	}
	fmt.Println("}")

	fmt.Println()
}

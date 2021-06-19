// gen_blocks.go generates block information.

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
	infoURL = "https://raw.githubusercontent.com/PrismarineJS/minecraft-data/master/data/pc/1.17/items.json"
)

type Item struct {
	ID          uint32 `json:"id"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	StackSize   uint   `json:"stackSize"`
}

func downloadInfo() ([]Item, error) {
	resp, err := http.Get(infoURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []Item
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}
	return data, nil
}

func makeItemDeclaration(blocks []Item) *ast.DeclStmt {
	out := &ast.DeclStmt{Decl: &ast.GenDecl{Tok: token.VAR}}

	for _, b := range blocks {
		t := reflect.TypeOf(b)
		fields := make([]ast.Expr, t.NumField())

		for i := 0; i < t.NumField(); i++ {
			ft := t.Field(i)

			var val ast.Expr
			switch ft.Type.Kind() {
			case reflect.Uint32, reflect.Int, reflect.Uint:
				val = &ast.BasicLit{Kind: token.INT, Value: fmt.Sprint(reflect.ValueOf(b).Field(i))}
			case reflect.Float64:
				val = &ast.BasicLit{Kind: token.FLOAT, Value: fmt.Sprint(reflect.ValueOf(b).Field(i))}
			case reflect.String:
				val = &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(reflect.ValueOf(b).Field(i).String())}
			case reflect.Bool:
				val = &ast.BasicLit{Kind: token.IDENT, Value: fmt.Sprint(reflect.ValueOf(b).Field(i).Bool())}

			case reflect.Slice:
				val = &ast.CompositeLit{
					Type: &ast.ArrayType{
						Elt: &ast.BasicLit{Kind: token.IDENT, Value: ft.Type.Elem().Name()},
					},
				}
				v := reflect.ValueOf(b).Field(i)
				switch ft.Type.Elem().Kind() {
				case reflect.Uint32, reflect.Int:
					for x := 0; x < v.Len(); x++ {
						val.(*ast.CompositeLit).Elts = append(val.(*ast.CompositeLit).Elts, &ast.BasicLit{
							Kind:  token.INT,
							Value: fmt.Sprint(v.Index(x)),
						})
					}
				}

			case reflect.Map:
				// Must be the NeedsTools map of type map[uint32]bool.
				m := &ast.CompositeLit{
					Type: &ast.MapType{
						Key:   &ast.BasicLit{Kind: token.IDENT, Value: ft.Type.Key().Name()},
						Value: &ast.BasicLit{Kind: token.IDENT, Value: ft.Type.Elem().Name()},
					},
				}
				iter := reflect.ValueOf(b).Field(i).MapRange()
				for iter.Next() {
					m.Elts = append(m.Elts, &ast.KeyValueExpr{
						Key:   &ast.BasicLit{Kind: token.INT, Value: fmt.Sprint(iter.Key().Uint())},
						Value: &ast.BasicLit{Kind: token.IDENT, Value: fmt.Sprint(iter.Value().Bool())},
					})
				}

				val = m
			}

			fields[i] = &ast.KeyValueExpr{
				Key:   &ast.Ident{Name: ft.Name},
				Value: val,
			}
		}

		out.Decl.(*ast.GenDecl).Specs = append(out.Decl.(*ast.GenDecl).Specs, &ast.ValueSpec{
			Names: []*ast.Ident{{Name: strcase.ToCamel(b.Name)}},
			Values: []ast.Expr{
				&ast.CompositeLit{
					Type: &ast.Ident{Name: reflect.TypeOf(b).Name()},
					Elts: fields,
				},
			},
		})
	}

	return out
}

func main() {
	items, err := downloadInfo()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(`// Package item stores information about items in Minecraft.
package item

import (
	"math"
)

// ID describes the numeric ID of an item.
type ID uint32

// Item describes information about a type of item.
type Item struct {
	ID          ID
	DisplayName string
	Name        string
	StackSize   uint
}

`)
	format.Node(os.Stdout, token.NewFileSet(), makeItemDeclaration(items))

	fmt.Println()
	fmt.Println()
	fmt.Println("// ByID is an index of minecraft items by their ID.")
	fmt.Println("var ByID = map[ID]*Item{")
	for _, i := range items {
		fmt.Printf("  %d: &%s,\n", i.ID, strcase.ToCamel(i.Name))
	}
	fmt.Println("}")
	fmt.Println()

	fmt.Println("// ByName is an index of minecraft items by their name.")
	fmt.Println("var ByName = map[string]*Item{")
	for _, i := range items {
		fmt.Printf("  %q: &%s,\n", i.Name, strcase.ToCamel(i.Name))
	}
	fmt.Println("}")
	fmt.Println()
}

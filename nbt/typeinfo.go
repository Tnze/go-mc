package nbt

import (
	"reflect"
	"strings"
	"sync"
)

type typeInfo struct {
	fields      []structField
	nameToIndex map[string]int // index of the field in struct, not previous slice
}

type structField struct {
	name  string
	index int

	omitEmpty bool
	list      bool
}

var tInfoMap sync.Map

func typeFields(typ reflect.Type) *typeInfo {
	if ti, ok := tInfoMap.Load(typ); ok {
		return ti.(*typeInfo)
	}

	tInfo := new(typeInfo)
	tInfo.nameToIndex = make(map[string]int)
	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		tInfo.fields = make([]structField, 0, n)
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			tag := f.Tag.Get("nbt")
			if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
				continue // Private field
			}

			// parse tags
			var field structField
			name, opts, _ := strings.Cut(tag, ",")
			if keytag := f.Tag.Get("nbtkey"); keytag != "" {
				name = keytag
			} else if name == "" {
				name = f.Name
			}
			field.name = name
			field.index = i

			// parse options
			for opts != "" {
				var name string
				name, opts, _ = strings.Cut(opts, ",")
				switch name {
				case "omitempty":
					field.omitEmpty = true
				case "list":
					field.list = true
				}
			}
			if f.Tag.Get("nbt_type") == "list" {
				field.list = true
			}
			tInfo.fields = append(tInfo.fields, field)

			tInfo.nameToIndex[field.name] = i
			if _, ok := tInfo.nameToIndex[f.Name]; !ok {
				tInfo.nameToIndex[f.Name] = i
			}
		}
	}

	ti, _ := tInfoMap.LoadOrStore(typ, tInfo)
	return ti.(*typeInfo)
}

func (t *typeInfo) findIndexByName(name string) int {
	i, ok := t.nameToIndex[name]
	if !ok {
		return -1
	}
	return i
}

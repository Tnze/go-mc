package nbt

import (
	"reflect"
)

type typeInfo struct {
	tagName     string
	nameToIndex map[string]int
}

func getTypeInfo(typ reflect.Type) *typeInfo {
	tinfo := new(typeInfo)
	tinfo.nameToIndex = make(map[string]int)
	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			tag := f.Tag.Get("nbt")
			if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
				continue // Private field
			}

			tinfo.nameToIndex[tag] = i
			if _, ok := tinfo.nameToIndex[f.Name]; !ok {
				tinfo.nameToIndex[f.Name] = i
			}
		}
	}
	return tinfo
}

func (t *typeInfo) findIndexByName(name string) int {
	i, ok := t.nameToIndex[name]
	if !ok {
		return -1
	}
	return i
}

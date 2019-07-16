package nbt

import (
	"reflect"
	"sync"
)

type typeInfo struct {
	tagName     string
	nameToIndex map[string]int
}

var tInfoMap sync.Map

func getTypeInfo(typ reflect.Type) *typeInfo {
	if ti, ok := tInfoMap.Load(typ); ok {
		return ti.(*typeInfo)
	}

	tInfo := new(typeInfo)
	tInfo.nameToIndex = make(map[string]int)
	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			tag := f.Tag.Get("nbt")
			if (f.PkgPath != "" && !f.Anonymous) || tag == "-" {
				continue // Private field
			}

			tInfo.nameToIndex[tag] = i
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

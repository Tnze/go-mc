package nbt

import (
	"reflect"
)

type typeInfo struct {
}

func getTypeInfo(typ reflect.Type) (*typeInfo, error) {
	tinfo := new(typeInfo)
	if typ.Kind() == reflect.Struct {
		n := typ.NumField()
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if (f.PkgPath != "" && !f.Anonymous) || f.Tag.Get("nbt") == "-" {
				continue // Private field
			}
		}
	}
	return tinfo, nil
}

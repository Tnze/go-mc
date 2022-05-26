package ecs

import (
	"reflect"
)

type System interface {
	Update(w *World)
}

type funcsystem struct {
	update func(w *World)
}

func FuncSystem(F any) System {
	type Storage interface {
		BitSetLike
		GetValueAny(eid Index) any
	}
	f := reflect.ValueOf(F)
	in := f.Type().NumIn()
	argTypes := make([]reflect.Type, in)
	needCopy := make([]bool, in)
	for i := 0; i < in; i++ {
		if t := f.Type().In(i); t.Kind() == reflect.Pointer {
			argTypes[i] = t.Elem()
		} else {
			argTypes[i] = t
			needCopy[i] = true
		}
	}
	return &funcsystem{
		update: func(w *World) {
			storages := make([]Storage, in)
			for i := 0; i < in; i++ {
				storages[i] = w.GetResourceRaw(argTypes[i]).(Storage)
			}
			args := make([]reflect.Value, len(storages))
			if len(storages) > 0 {
				set := BitSetLike(storages[0])
				for _, v := range storages[1:] {
					set = set.And(v)
				}
				set.Range(func(eid Index) {
					for i := range args {
						arg := storages[i].GetValueAny(eid)
						if arg == nil {
							args[i] = reflect.Zero(argTypes[i])
						} else if needCopy[i] {
							args[i] = reflect.ValueOf(arg).Elem()
						} else {
							args[i] = reflect.ValueOf(arg)
						}
					}
					f.Call(args)
				})
			} else {
				f.Call(args)
			}
		},
	}
}

func (f *funcsystem) Update(w *World) {
	f.update(w)
}

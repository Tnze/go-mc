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
		GetValueAny(eid Index) any
		And(*BitSet) *BitSet
		Range(f func(eid Index))
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
			storages := make([]reflect.Value, in)
			for i := 0; i < in; i++ {
				storages[i] = reflect.ValueOf(w.GetResourceRaw(argTypes[i]))
			}
			args := make([]reflect.Value, len(storages))
			if len(storages) > 0 {
				set := reflect.Indirect(storages[0]).FieldByName("BitSet").Addr()
				for _, v := range storages[1:] {
					p := reflect.Indirect(v).FieldByName("BitSet").Addr()
					set = set.MethodByName("And").Call([]reflect.Value{p})[0]
				}
				set.MethodByName("Range").Call([]reflect.Value{
					reflect.ValueOf(func(eid Index) {
						for i := range args {
							arg := storages[i].MethodByName("GetValue").Call([]reflect.Value{reflect.ValueOf(eid)})[0]
							if arg.IsNil() {
								args[i] = reflect.Zero(argTypes[i])
							} else if needCopy[i] {
								args[i] = arg.Elem()
							} else {
								args[i] = arg
							}
						}
						f.Call(args)
					}),
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

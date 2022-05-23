package ecs

import "reflect"

type System interface {
	Update(w *World)
}

type funcsystem struct {
	update reflect.Value
	args   func(w *World) []Storage
}

func FuncSystem(F any) System {
	f := reflect.ValueOf(F)
	in := f.Type().NumIn()
	argTypes := make([]reflect.Type, in)
	for i := 0; i < in; i++ {
		argTypes[i] = f.Type().In(i)
	}
	return &funcsystem{
		update: f,
		args: func(w *World) (args []Storage) {
			args = make([]Storage, in)
			for i := 0; i < in; i++ {
				args[i] = w.GetResourceRaw(argTypes[i]).(Storage)
			}
			return
		},
	}
}

func (f *funcsystem) Update(w *World) {
	storages := f.args(w)
	if len(storages) == 0 {
		return
	}
	eids := storages[0].BitSet()
	for _, v := range storages[1:] {
		eids = eids.And(v.BitSet())
	}
	args := make([]reflect.Value, len(storages))
	eids.Range(func(eid Index) {
		for i := range args {
			args[i] = reflect.ValueOf(storages[i].Get(eid))
		}
		f.update.Call(args)
	})
}

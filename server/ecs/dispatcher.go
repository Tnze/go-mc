package ecs

import "sync"

type Dispatcher struct {
	waiters map[string][]*sync.WaitGroup
	tasks   []func(w *World, wg *sync.WaitGroup)
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		waiters: make(map[string][]*sync.WaitGroup),
		tasks:   nil,
	}
}

func (d *Dispatcher) Add(s System, name string, deps []string) {
	var start sync.WaitGroup
	start.Add(len(deps))
	for _, dep := range deps {
		if wg, ok := d.waiters[dep]; ok {
			d.waiters[dep] = append(wg, &start)
		} else {
			panic("Unknown deps: " + dep)
		}
	}
	d.tasks = append(d.tasks, func(w *World, done *sync.WaitGroup) {
		start.Wait()
		defer done.Done()
		s.Update(w)
		for _, wg := range d.waiters[name] {
			wg.Done()
		}
	})
}

func (d *Dispatcher) Run(w *World) {
	var wg sync.WaitGroup
	wg.Add(len(d.tasks))
	for _, f := range d.tasks {
		go f(w, &wg)
	}
	wg.Wait()
}

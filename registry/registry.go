package registry

type Registry[E any] struct {
	keys    map[string]int32
	values  []E
	indices map[*E]int32
}

func NewRegistry[E any]() Registry[E] {
	return Registry[E]{
		keys:    make(map[string]int32),
		values:  make([]E, 0, 256),
		indices: make(map[*E]int32),
	}
}

func (r *Registry[E]) Get(key string) (int32, *E) {
	id, ok := r.keys[key]
	if !ok {
		return -1, nil
	}
	return id, &r.values[id]
}

func (r *Registry[E]) GetByID(id int32) *E {
	if id >= 0 && id < int32(len(r.values)) {
		return &r.values[id]
	}
	return nil
}

func (r *Registry[E]) Put(name string, data E) (id int32, val *E) {
	id = int32(len(r.values))
	r.keys[name] = id
	r.values = append(r.values, data)
	val = &r.values[id]
	r.indices[val] = id
	return
}

func (r *Registry[E]) Clear() {
	r.keys = make(map[string]int32)
	r.values = r.values[:0]
	r.indices = make(map[*E]int32)
}

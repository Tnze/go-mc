package registry

type Registry[E any] struct {
	Type  string `nbt:"type"`
	Value []struct {
		Name    string `nbt:"name"`
		ID      int32  `nbt:"id"`
		Element E      `nbt:"element"`
	} `nbt:"value"`
}

func (r *Registry[E]) Find(name string) (int32, *E) {
	for i := range r.Value {
		if r.Value[i].Name == name {
			return int32(i), &r.Value[i].Element
		}
	}
	return -1, nil
}

func (r *Registry[E]) FindByID(id int32) *E {
	if id >= 0 && id < int32(len(r.Value)) && r.Value[id].ID == id {
		return &r.Value[id].Element
	}
	for i := range r.Value {
		if r.Value[i].ID == id {
			return &r.Value[i].Element
		}
	}
	return nil
}

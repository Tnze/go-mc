package registry

import "github.com/Tnze/go-mc/nbt"

type Registry[E any] struct {
	Type  string     `nbt:"type"`
	Value []Entry[E] `nbt:"value"`
}

type Entry[E any] struct {
	Name    string `nbt:"name"`
	ID      int32  `nbt:"id"`
	Element E      `nbt:"element"`
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

func (r *Registry[E]) Insert(name string, data E) {
	r.Value = append(r.Value, Entry[E]{Name: name, Element: data})
}

func (r *Registry[E]) InsertNBT(name string, data nbt.RawMessage) error {
	entry := Entry[E]{Name: name, ID: int32(len(r.Value))}
	if data.Type != 0 {
		if err := data.UnmarshalDisallowUnknownField(&entry.Element); err != nil {
			return err
		}
	}
	r.Value = append(r.Value, entry)
	return nil
}

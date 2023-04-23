package fastnbt

import "github.com/Tnze/go-mc/nbt"

func (v *Value) Set(key string, val *Value) {
	if v.tag != nbt.TagCompound {
		panic("cannot set non-Compound Tag")
	}
	v.comp.Set(key, val)
}

func (v *Value) Get(keys ...string) *Value {
	for _, key := range keys {
		if v.tag == nbt.TagCompound {
			v = v.comp.Get(key)
			if v == nil {
				return nil
			}
		} else {
			return nil
		}
	}
	return v
}

func (c *Compound) Set(key string, val *Value) {
	for i := range c.kvs {
		if c.kvs[i].tag == key {
			c.kvs[i].v = val
			return
		}
	}
	c.kvs = append(c.kvs, kv{key, val})
}

func (c *Compound) Get(key string) *Value {
	for _, tag := range c.kvs {
		if tag.tag == key {
			return tag.v
		}
	}
	return nil
}

func (c *Compound) Len() int {
	return len(c.kvs)
}

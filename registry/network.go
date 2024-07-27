package registry

import (
	"errors"
	"io"
	"strconv"

	pk "github.com/Tnze/go-mc/net/packet"
)

func (reg *Registry[E]) ReadFrom(r io.Reader) (int64, error) {
	var length pk.VarInt
	n, err := length.ReadFrom(r)
	if err != nil {
		return n, err
	}

	reg.Clear()

	var key pk.Identifier
	var hasData pk.Boolean
	for i := 0; i < int(length); i++ {
		var data E
		var n1, n2, n3 int64

		n1, err = key.ReadFrom(r)
		if err != nil {
			return n + n1, err
		}

		n2, err = hasData.ReadFrom(r)
		if err != nil {
			return n + n1 + n2, err
		}

		if hasData {
			n3, err = pk.NBTField{V: &data, AllowUnknownFields: true}.ReadFrom(r)
			if err != nil {
				return n + n1 + n2 + n3, err
			}
			reg.Put(string(key), data)
		}

		n += n1 + n2 + n3
	}
	return n, nil
}

func (reg *Registry[E]) ReadTagsFrom(r io.Reader) (int64, error) {
	var count pk.VarInt
	n, err := count.ReadFrom(r)
	if err != nil {
		return n, err
	}

	var tag pk.Identifier
	var length pk.VarInt
	for i := 0; i < int(count); i++ {
		var n1, n2, n3 int64

		n1, err = tag.ReadFrom(r)
		if err != nil {
			return n + n1, err
		}

		n2, err = length.ReadFrom(r)
		if err != nil {
			return n + n1 + n2, err
		}

		n += n1 + n2
		values := make([]*E, length)

		var id pk.VarInt
		for i := 0; i < int(length); i++ {
			n3, err = id.ReadFrom(r)
			if err != nil {
				return n + n3, err
			}

			if id < 0 || int(id) >= len(reg.values) {
				err = errors.New("invalid id: " + strconv.Itoa(int(id)))
				return n + n3, err
			}

			values[i] = &reg.values[id]
			n += n3
		}

		reg.tags[string(tag)] = values
	}
	return n, nil
}

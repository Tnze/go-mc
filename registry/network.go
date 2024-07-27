package registry

import (
	"io"

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
			n3, err = pk.NBTField{V:&data, AllowUnknownFields: true}.ReadFrom(r)
			if err != nil {
				return n + n1 + n2 + n3, err
			}
			reg.Put(string(key), data)
		}

		n += n1 + n2 + n3
	}
	return n, nil
}

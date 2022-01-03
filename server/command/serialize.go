package command

import (
	"io"
	"unsafe"

	pk "github.com/Tnze/go-mc/net/packet"
)

const (
	isExecutable = 1 << (iota + 2)
	hasRedirect
	hasSuggestionsType
)

func (g *Graph) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Array(g.nodes),
		pk.VarInt(0),
	}.WriteTo(w)
}

func (n Node) WriteTo(w io.Writer) (int64, error) {
	var flag byte
	flag |= n.kind & 0x03
	if n.Run != nil {
		flag |= isExecutable
	}
	return pk.Tuple{
		pk.Byte(flag),
		pk.Array((*[]pk.VarInt)(unsafe.Pointer(&n.Children))),
		pk.Opt{
			Has:   func() bool { return n.kind&hasRedirect != 0 },
			Field: nil, // TODO: send redirect node
		},
		pk.Opt{
			Has:   func() bool { return n.kind == ArgumentNode || n.kind == LiteralNode },
			Field: pk.String(n.Name),
		},
		pk.Opt{
			Has:   func() bool { return n.kind == ArgumentNode },
			Field: n.Parser, // Parser identifier and Properties
		},
		pk.Opt{
			Has:   func() bool { return flag&hasSuggestionsType != 0 },
			Field: nil, // TODO: send Suggestions type
		},
	}.WriteTo(w)
}

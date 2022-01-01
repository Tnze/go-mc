package command

import (
	"io"
	"unsafe"

	pk "github.com/Tnze/go-mc/net/packet"
)

type CmdSet struct {
	root  int32
	nodes []Node
}

func NewCmdSet() *CmdSet {
	return &CmdSet{
		root: 0,
		nodes: []Node{
			{
				flags:  TypeRoot,
				parser: nil,
			},
		},
	}
}

func (c *CmdSet) AddNode(n Node) int32 {
	c.nodes = append(c.nodes, n)
	return int32(len(c.nodes) - 1)
}

func (c *CmdSet) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Array(c.nodes),
		pk.VarInt(c.root),
	}.WriteTo(w)
}

func (c *CmdSet) Execute(cmd string) error {
	for _, child := range c.nodes[c.root].children {
		if node := c.nodes[child]; node.flags&0x03 == TypeLiteral {

		}
	}
}

type Node struct {
	flags       byte
	children    []int32
	redirect    int32
	name        string
	parser      Parser
	suggestions string
}

func (n *Node) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Byte(n.flags),
		pk.Array(*(*[]pk.VarInt)(unsafe.Pointer(&n.children))),
		pk.Opt{ // Redirect
			Has:   n.flags&HasRedirect != 0,
			Field: pk.VarInt(n.redirect),
		},
		pk.Opt{ // Name
			Has:   n.flags&0x03 == TypeLiteral || n.flags&0x03 == TypeArgument,
			Field: pk.String(n.name),
		},
		pk.Opt{ // Parser & Properties
			Has:   n.flags&0x03 == TypeArgument,
			Field: n.parser,
		},
		pk.Opt{
			Has:   n.flags&HasSuggestionsType != 0,
			Field: pk.Identifier(n.suggestions),
		},
	}.WriteTo(w)
}

type Parser interface {
	pk.FieldEncoder
	Parse(cmd string) (token string, value interface{}, err error)
}

const (
	TypeRoot byte = iota
	TypeLiteral
	TypeArgument
	TypeUnknown
)

const (
	IsExecutable = 1 << (iota + 2)
	HasRedirect
	HasSuggestionsType
)

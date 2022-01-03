package command

import (
	"errors"
	"strings"
)

const (
	RootNode = iota
	LiteralNode
	ArgumentNode
)

type Graph struct {
	nodes []Node
}

func NewGraph() *Graph {
	g := new(Graph)
	root := Node{
		g:    g,
		kind: RootNode,
	}
	g.nodes = append(g.nodes, root)
	return g
}

func (g *Graph) Run(cmd string) error {
	var args []ParsedData
	node := &g.nodes[0] // root
	for {
		left, value, next, err := node.parse(cmd)
		if err != nil {
			return err
		}
		args = append(args, value)
		left = strings.TrimSpace(left)
		if len(left) == 0 {
			err := node.Run(args)
			if err != nil {
				return err
			}
			return nil
		}
		if next == 0 {
			return errors.New("command contains extra text: " + left)
		}

		cmd = left
		node = &g.nodes[next]
	}
}

type ParsedData interface{}

type HandlerFunc func(args []ParsedData) error

type Node struct {
	g               *Graph
	index           int32
	kind            byte
	Name            string
	Children        []int32
	SuggestionsType string
	Parser          Parser
	Run             HandlerFunc
}
type Literal Node
type Argument Node

func (n *Node) parse(cmd string) (left string, value ParsedData, next int32, err error) {
	switch n.kind & 0x03 {
	case RootNode:
		left = cmd
		value = nil
		err = nil
	case LiteralNode:
		if !strings.HasPrefix(cmd, n.Name) {
			panic("expect " + cmd + " prefixed with " + n.Name)
		}
		left = strings.TrimPrefix(cmd, n.Name)
		value = LiteralData(n.Name)
	case ArgumentNode:
		left, value, err = n.Parser.Parse(cmd)
		if err != nil {
			return "", nil, 0, err
		}
	default:
		panic("unreachable")
	}
	// find next
	if len(n.Children) > 0 {
		// look up the first child's type
		switch n.g.nodes[n.Children[0]].kind & 0x03 {
		case RootNode:
			panic("root node can't be child")
		default:
			panic("unreachable")
		case LiteralNode:
			_, value, err := StringParser(0).Parse(strings.TrimSpace(left))
			if err != nil {
				return "", nil, 0, err
			}
			literal := value.(string)
			for _, i := range n.Children {
				if n.g.nodes[i].Name == literal {
					next = i
					break
				}
			}
		case ArgumentNode:
			next = n.Children[0]
		}
	}
	return
}

func unhandledCmd([]ParsedData) error {
	return errors.New("unhandled function")
}

type LiteralData string

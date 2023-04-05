package command

import (
	"context"
	"errors"
	"strings"
)

const (
	RootNode = iota
	LiteralNode
	ArgumentNode
)

// Graph is a directed graph with a root node, representing all commands and how they are parsed.
type Graph struct {
	// List of all nodes. The first element is the root node
	nodes []*Node
}

func NewGraph() *Graph {
	var g Graph
	g.nodes = append(g.nodes,
		&Node{g: &g, kind: RootNode},
	)
	return &g
}

func (g *Graph) Execute(ctx context.Context, cmd string) error {
	var args []ParsedData
	node := g.nodes[0] // root
	for {
		// parser command
		left, value, err := node.parse(cmd)
		if err != nil {
			return err
		}
		args = append(args, value)
		left = strings.TrimSpace(left)
		if len(left) == 0 {
			return node.Run(ctx, args)
		}
		// find next node
		next, err := node.next(left)
		if err != nil {
			return err
		}
		if next == 0 {
			return errors.New("command contains extra text: " + left)
		}

		cmd = left
		node = g.nodes[next]
	}
}

type ParsedData any

type HandlerFunc func(ctx context.Context, args []ParsedData) error

// Node is the node of the Graph. There are 3 kinds of node: Root, Literal and Argument.
type Node struct {
	g     *Graph
	index int32
	kind  byte

	Name            string
	Children        []int32
	SuggestionsType string
	Parser          Parser
	Run             HandlerFunc
}
type (
	Literal  Node
	Argument Node
)

func (n *Node) parse(cmd string) (left string, value ParsedData, err error) {
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

	default:
		panic("unreachable")
	}
	return
}

func (n *Node) next(left string) (next int32, err error) {
	if len(n.Children) == 0 {
		return 0, nil
	}
	// look up the first child's type
	switch n.g.nodes[n.Children[0]].kind & 0x03 {
	case RootNode:
		panic("root node can't be child")
	default:
		panic("unreachable")
	case LiteralNode:
		_, value, err := StringParser(0).Parse(strings.TrimSpace(left))
		if err != nil {
			return 0, err
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
	return
}

func unhandledCmd(context.Context, []ParsedData) error {
	return errors.New("unhandled function")
}

type LiteralData string

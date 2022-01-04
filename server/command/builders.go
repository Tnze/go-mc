package command

func (g *Graph) AppendLiteral(child *Literal) *Graph {
	g.nodes[0].Children = append(g.nodes[0].Children, child.index)
	return g
}

// Literal create a new LiteralNode in the Graph.
func (g *Graph) Literal(str string) LiteralBuilder {
	index := int32(len(g.nodes))
	g.nodes = append(g.nodes, &Node{
		g:     g,
		index: index,
		kind:  LiteralNode,
		Name:  str,
	})
	return LiteralBuilder{current: g.nodes[index]}
}

// Argument create a new ArgumentNode in the Graph.
func (g *Graph) Argument(name string, parser Parser) ArgumentBuilder {
	index := int32(len(g.nodes))
	g.nodes = append(g.nodes, &Node{
		g:      g,
		index:  index,
		kind:   ArgumentNode,
		Name:   name,
		Parser: parser,
	})
	return ArgumentBuilder{current: g.nodes[index]}
}

type LiteralBuilder struct {
	current *Node
}

func (n LiteralBuilder) AppendLiteral(node *Literal) LiteralBuilderWithLiteral {
	n.current.Children = append(n.current.Children, node.index)
	return LiteralBuilderWithLiteral{n: n}
}

func (n LiteralBuilder) AppendArgument(node *Argument) LiteralBuilderWithArgument {
	n.current.Children = append(n.current.Children, node.index)
	return LiteralBuilderWithArgument{n: n}
}

func (n LiteralBuilder) HandleFunc(f HandlerFunc) *Literal {
	n.current.Run = f
	return (*Literal)(n.current)
}

func (n LiteralBuilder) Unhandle() *Literal {
	return n.HandleFunc(unhandledCmd)
}

type ArgumentBuilder struct {
	current *Node
}

func (n ArgumentBuilder) AppendLiteral(node *Literal) ArgumentBuilderWithLiteral {
	n.current.Children = append(n.current.Children, node.index)
	return ArgumentBuilderWithLiteral{n: n}
}

func (n ArgumentBuilder) AppendArgument(node *Argument) ArgumentBuilderWithArgument {
	n.current.Children = append(n.current.Children, node.index)
	return ArgumentBuilderWithArgument{n: n}
}

func (n ArgumentBuilder) HandleFunc(f HandlerFunc) *Argument {
	n.current.Run = f
	return (*Argument)(n.current)
}

func (n ArgumentBuilder) Unhandle() *Argument {
	return n.HandleFunc(unhandledCmd)
}

type LiteralBuilderWithLiteral struct {
	n LiteralBuilder
}

func (n LiteralBuilderWithLiteral) AppendLiteral(node *Literal) LiteralBuilderWithLiteral {
	return n.n.AppendLiteral(node)
}

func (n LiteralBuilderWithLiteral) HandleFunc(f HandlerFunc) *Literal {
	return n.n.HandleFunc(f)
}

func (n LiteralBuilderWithLiteral) Unhandle() *Literal {
	return n.n.Unhandle()
}

type LiteralBuilderWithArgument struct {
	n LiteralBuilder
}

func (n LiteralBuilderWithArgument) HandleFunc(f HandlerFunc) *Literal {
	return n.n.HandleFunc(f)
}

func (n LiteralBuilderWithArgument) Unhandle() *Literal {
	return n.n.Unhandle()
}

type ArgumentBuilderWithLiteral struct {
	n ArgumentBuilder
}

func (n ArgumentBuilderWithLiteral) AppendLiteral(node *Literal) ArgumentBuilderWithLiteral {
	return n.n.AppendLiteral(node)
}

func (n ArgumentBuilderWithLiteral) HandleFunc(f HandlerFunc) *Argument {
	return n.n.HandleFunc(f)
}

func (n ArgumentBuilderWithLiteral) Unhandle() *Argument {
	return n.n.Unhandle()
}

type ArgumentBuilderWithArgument struct {
	n ArgumentBuilder
}

func (n ArgumentBuilderWithArgument) HandleFunc(f HandlerFunc) *Argument {
	return n.n.HandleFunc(f)
}

func (n ArgumentBuilderWithArgument) Unhandle() *Argument {
	return n.n.Unhandle()
}

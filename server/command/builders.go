package command

func (g *Graph) AppendLiteral(child *Literal) *Graph {
	g.nodes[0].Children = append(g.nodes[0].Children, child.index)
	return g
}

func (g *Graph) Literal(str string) LiteralBuilder {
	index := len(g.nodes)
	g.nodes = append(g.nodes, Node{
		g:     g,
		index: index,
		flags: LiteralNode,
		Name:  str,
	})
	return LiteralBuilder{g: g, current: index}
}

func (g *Graph) Argument(name string, parser Parser) ArgumentBuilder {
	index := len(g.nodes)
	g.nodes = append(g.nodes, Node{
		g:      g,
		index:  index,
		flags:  ArgumentNode,
		Name:   name,
		Parser: parser,
	})
	return ArgumentBuilder{g: g, current: index}
}

type LiteralBuilder struct {
	g       *Graph
	current int
}

func (n LiteralBuilder) AppendLiteral(node *Literal) LiteralBuilderWithLiteral {
	current := &n.g.nodes[n.current]
	current.Children = append(current.Children, node.index)
	return LiteralBuilderWithLiteral{n: n}
}

func (n LiteralBuilder) AppendArgument(node *Argument) LiteralBuilderWithArgument {
	current := &n.g.nodes[n.current]
	current.Children = append(current.Children, node.index)
	return LiteralBuilderWithArgument{n: n}
}

func (n LiteralBuilder) HandleFunc(f HandlerFunc) *Literal {
	current := &n.g.nodes[n.current]
	current.Run = f
	return (*Literal)(current)
}

func (n LiteralBuilder) Unhandle() *Literal {
	return n.HandleFunc(unhandledCmd)
}

type ArgumentBuilder struct {
	g       *Graph
	current int
}

func (n ArgumentBuilder) AppendLiteral(node *Literal) ArgumentBuilderWithLiteral {
	current := &n.g.nodes[n.current]
	current.Children = append(current.Children, node.index)
	return ArgumentBuilderWithLiteral{n: n}
}

func (n ArgumentBuilder) AppendArgument(node *Argument) ArgumentBuilderWithArgument {
	current := &n.g.nodes[n.current]
	current.Children = append(current.Children, node.index)
	return ArgumentBuilderWithArgument{n: n}
}

func (n ArgumentBuilder) HandleFunc(f HandlerFunc) *Argument {
	current := &n.g.nodes[n.current]
	current.Run = f
	return (*Argument)(current)
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
	return (*Literal)((*Argument)(n.n.HandleFunc(f)))
}

func (n LiteralBuilderWithArgument) Unhandle() *Literal {
	return (*Literal)((*Argument)(n.n.Unhandle()))
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

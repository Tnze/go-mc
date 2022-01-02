package command

import (
	"github.com/Tnze/go-mc/chat"
	"strings"
)

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

type Node interface {
	Then(nodes ...Node) Node
	Runnable() HandleFunc
	Parse(cmd string, store map[string]interface{}) (token string, next Node, err error)
}

type Root struct {
	children map[string]Node
}

type HandleFunc func(args map[string]interface{})

func NewGraph() *Root {
	return &Root{children: make(map[string]Node)}
}

func (r *Root) Then(nodes ...Node) Node {
	for _, node := range nodes {
		l, ok := node.(*Literal)
		if !ok {
			panic("you could only add Literal node to Root")
		}
		if _, ok := r.children[l.Name]; ok {
			panic("the Literal(" + l.Name + ") is already exist")
		}
		r.children[l.Name] = node
	}
	return r
}

func Lite(literal string) *Literal {
	return &Literal{Name: literal}
}

func Arg(name string, parser Parser) *Argument {
	return &Argument{Name: name, Parser: parser}
}

func (r *Root) Runnable() HandleFunc {
	return nil
}

func (r *Root) Parse(cmd string, _ map[string]interface{}) (token string, next Node, err error) {
	token, value, err := StringParser(0).Parse(cmd)
	if err != nil {
		return "", nil, err
	}
	child, ok := r.children[value.(string)]
	if !ok {
		return "", nil, Error{chat.TranslateMsg("command.unknown.command")}
	}
	return "", child, nil
}

func (r *Root) Run(cmd string) error {
	var node Node = r
	args := make(map[string]interface{})
	for {
		token, next, err := node.Parse(cmd, args)
		if err != nil {
			return err
		}
		cmd = strings.TrimSpace(strings.TrimPrefix(cmd, token))
		if next == nil || len(cmd) == 0 {
			f := node.Runnable()
			if f == nil {
				return Error{chat.TranslateMsg("command.unknown.command")}
			}
			f(args)
			return nil
		}
		node = next
	}
}

type Children struct {
	Lite map[string]*Literal
	Args *Argument
}

func (c *Children) add(nodes []Node) {
	for _, node := range nodes {
		switch node.(type) {
		case *Literal:
			if c.Args != nil {
				panic("this Node already has an Argument Node children, cannot Add Literal Node")
			}
			if c.Lite == nil {
				c.Lite = make(map[string]*Literal)
			}
			child := node.(*Literal)
			c.Lite[child.Name] = child
		case *Argument:
			if c.Lite != nil {
				panic("this Node already has Literal Node children, cannot Add Argument Node")
			}
			if c.Args != nil {
				panic("this Node already set an Argument Node child, cannot be duplicate")
			}
			c.Args = node.(*Argument)
		}
	}
	return
}

type Handler struct {
	Handler HandleFunc
}

func (h *Handler) Handle(f HandleFunc) {
	if h.Handler != nil {
		panic("handler func set twice")
	}
	h.Handler = f
}

func (h *Handler) Runnable() HandleFunc {
	return h.Handler
}

type Literal struct {
	Name string
	Children
	Handler
}

func (l *Literal) Then(nodes ...Node) Node {
	l.add(nodes)
	return l
}

func (l *Literal) Parse(cmd string, store map[string]interface{}) (token string, next Node, err error) {
	cmd = strings.TrimPrefix(cmd, l.Name)
	store[l.Name] = nil
	if l.Children.Lite != nil {
		token, value, err := StringParser(0).Parse(strings.TrimSpace(cmd))
		if err != nil {
			return token, nil, err
		}
		if next, ok := l.Children.Lite[value.(string)]; ok {
			return l.Name, next, nil
		}
		return "", nil, Error{chat.TranslateMsg("command.unknown.command")}
	} else if l.Children.Args != nil {
		return l.Name, l.Children.Args, nil
	} else {
		return l.Name, nil, nil
	}
}

func (l *Literal) Handle(f HandleFunc) Node {
	l.Handler.Handle(f)
	return l
}

type Argument struct {
	Name string
	Children
	Parser Parser
	Handler
}

func (a *Argument) Parse(cmd string, store map[string]interface{}) (token string, next Node, err error) {
	token, value, err := a.Parser.Parse(cmd)
	if err != nil {
		return "", nil, err
	}
	store[a.Name] = value
	cmd = strings.TrimSpace(strings.TrimPrefix(cmd, token))

	if a.Children.Lite != nil {
		_, value, err := StringParser(0).Parse(strings.TrimSpace(cmd))
		if err != nil {
			return token, nil, err
		}
		if next, ok := a.Children.Lite[value.(string)]; ok {
			return token, next, nil
		}
		return "", nil, Error{chat.TranslateMsg("command.unknown.command")}
	} else if a.Children.Args != nil {
		return token, a.Children.Args, nil
	} else {
		return token, nil, nil
	}
}

func (a *Argument) Then(nodes ...Node) Node {
	a.add(nodes)
	return a
}

func (a *Argument) Handle(f HandleFunc) Node {
	a.Handler.Handle(f)
	return a
}

type Error struct {
	msg chat.Message
}

func (e Error) Error() string {
	return e.msg.String()
}

package command

import (
	"context"
	"log"
	"testing"
)

func TestRoot_Run(t *testing.T) {
	handleFunc := func(ctx context.Context, args []ParsedData) error {
		log.Printf("Command: args: %v", args)
		return nil
	}
	g := NewGraph()
	g.AppendLiteral(g.Literal("me").
		AppendArgument(g.Argument("action", StringParser(2)).
			HandleFunc(handleFunc)).
		Unhandle(),
	).AppendLiteral(g.Literal("help").
		AppendArgument(g.Argument("command", StringParser(0)).
			HandleFunc(handleFunc)).
		HandleFunc(handleFunc),
	).AppendLiteral(g.Literal("list").
		AppendLiteral(g.Literal("uuids").
			HandleFunc(handleFunc)).
		HandleFunc(handleFunc),
	)

	err := g.Execute(context.TODO(), "me Tnze Xi_Xi_Mi")
	if err != nil {
		t.Fatal(err)
	}
}

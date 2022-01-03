package command

import "testing"

func TestRoot_Run(t *testing.T) {
	g := NewGraph()
	g.AppendLiteral(g.Literal("whitelist").
		AppendLiteral(g.Literal("add").
			AppendArgument(g.Argument("targets", StringParser(0)).
				HandleFunc(func(args []ParsedData) error {
					t.Logf("whitelist add: %v", args)
					return nil
				})).Unhandle()).
		AppendLiteral(g.Literal("remove").
			AppendArgument(g.Argument("targets", StringParser(0)).
				HandleFunc(func(args []ParsedData) error {
					t.Logf("whitelist remove: %v", args)
					return nil
				})).Unhandle()).
		AppendLiteral(g.Literal("on").
			HandleFunc(func(args []ParsedData) error {
				t.Logf("whitelist on: %v", args)
				return nil
			})).
		AppendLiteral(g.Literal("off").
			HandleFunc(func(args []ParsedData) error {
				t.Logf("whitelist off: %v", args)
				return nil
			})).
		Unhandle(),
	)

	targetB := g.Argument("targetB", StringParser(0)).
		HandleFunc(func(args []ParsedData) error {
			t.Logf("tp A <from/to> B parsed: %v", args)
			return nil
		})
	g.AppendLiteral(g.Literal("tp").AppendArgument(
		g.Argument("targetA", StringParser(0)).
			AppendLiteral(g.Literal("from").
				AppendArgument(targetB).
				Unhandle()).
			AppendLiteral(g.Literal("to").
				AppendArgument(targetB).
				Unhandle()).
			Unhandle(),
	).Unhandle())

	err := g.Run("tp Tnze to Xi_Xi_Mi")
	if err != nil {
		t.Fatal(err)
	}

	//g2 := NewGraph()
	//g2.Then(
	//	g.Lite("whitelist").Then(
	//		// using reflect to generate all arguments nodes
	//		g.Lite("add").Then(g.Func(func(player GameProfile) {
	//
	//		})),
	//	),
	//)
}

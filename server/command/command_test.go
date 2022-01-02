package command

import "testing"

func TestRoot_Run(t *testing.T) {
	g := NewGraph().Then(
		Lite("whitelist").Then(
			Lite("add").Then(
				Arg("targets", StringParser(0)).
					Handle(func(args map[string]interface{}) {
						t.Logf("whitelist add: %v", args)
					})),
			Lite("remove").Then(
				Arg("targets", StringParser(0)).
					Handle(func(args map[string]interface{}) {
						t.Logf("whitelist remove: %v", args)
					})),
			Lite("on").Handle(func(args map[string]interface{}) {
				t.Logf("whitelist on: %v", args)
			}),
			Lite("off").Handle(func(args map[string]interface{}) {
				t.Logf("whitelist off: %v", args)
			}),
		),
	).(*Root)

	targetA := Arg("targetA", StringParser(0))
	targetB := Arg("targetB", StringParser(0))
	targetB.Handle(func(args map[string]interface{}) {
		t.Logf("tp A <from/to> B parsed: %v", args)
	})
	tp := Lite("tp").Then(targetA.Then(
		Lite("form").Then(targetB),
		Lite("to").Then(targetB),
	))
	g.Then(tp)

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

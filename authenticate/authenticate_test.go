package authenticate

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEncodingPayload(t *testing.T) {
	j, err := json.Marshal(Payload{
		Agent: Agent{
			Name:    "Minecraft",
			Version: 1,
		},
		UserName:    "mojang account name",
		Password:    "mojang account password",
		ClientToken: "client identifier",
		RequestUser: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(j))
}

func ExampleAuthenticate() {
	resp, err := Authenticate("", "", "")
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}

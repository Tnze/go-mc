package yggdrasil

import (
	"fmt"
)

func ExampleAuthenticate() {
	resp, err := Authenticate("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.SelectedProfile())
	fmt.Println(resp.AccessToken())
}

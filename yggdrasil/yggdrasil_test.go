package yggdrasil

import (
	"fmt"
	"os"
)

func ExampleAuthenticate() {
	resp, err := Authenticate("", "")
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.SelectedProfile())
	fmt.Println(resp.AccessToken())
}

func Example() {
	var user, password string // set your proof

	// Sign in
	resp, err := Authenticate(user, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	id, name := resp.SelectedProfile()
	fmt.Println("user:", name)
	fmt.Println("uuid:", id)
	fmt.Println("astk:", resp.AccessToken())

	// Refresh access token
	if err := resp.Refresh(nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	id, name = resp.SelectedProfile()
	fmt.Println("user:", name)
	fmt.Println("uuid:", id)
	fmt.Println("astk:", resp.AccessToken())

	// Check access token
	ok, err := resp.Validate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("at status: ", ok)

	// Invalidate access token
	err = resp.Invalidate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Check access token
	ok, err = resp.Validate()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("at status: ", ok)

	// Sign out
	err = SignOut(user, password)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

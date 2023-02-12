// This example used to act as a launcher, log in and obtain the access token.
// The Yggdrasil Authentication is no longer available. This example doesn't work now.
//
// For now, you should use Microsoft Authentication. The description and example code can be found here:
// https://wiki.vg/Microsoft_Authentication_Scheme
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Tnze/go-mc/yggdrasil"
)

var (
	user = flag.String("user", "", "Can be an email address or player name for unmigrated accounts")
	pswd = flag.String("password", "", "Your password")
)

func main() {
	flag.Parse()

	resp, err := yggdrasil.Authenticate(*user, *pswd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	id, name := resp.SelectedProfile()
	fmt.Println("user:", name)
	fmt.Println("uuid:", id)
	fmt.Println("astk:", resp.AccessToken())
}

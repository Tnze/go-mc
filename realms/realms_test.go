package realms

import (
	"fmt"
	"time"
)

func ExampleRealms() {
	var r *Realms

	r = New(
		"1.14.4",
		"Name",
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		"xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	)
	fmt.Println(r.Available())
	fmt.Println(r.Compatible())

	servers, err := r.Worlds()
	if err != nil {
		panic(err)
	}

	for _, v := range servers {
		fmt.Println(v.Name, v.ID)
	}

	time.Sleep(time.Second * 5)
	if err := r.TOS(); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 5)
	fmt.Println(r.Address(servers[0]))
}

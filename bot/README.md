# bot

Bot provides some tools to join servers (offline or online).

For examples, see the [examples](../examples) directory.

## Offline

```go
package main

import (
	"log"
	
	"github.com/Tnze/go-mc/bot"
)

func main() {
	c := bot.NewClient()
	if err := c.JoinServer("localhost"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")
}
```

## Online

Similar to offline, however we need to authenticate with Yggdrasil.

```go
package main

import (
	"log"
	
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/yggdrasil"
)

func main() {
	c := bot.NewClient()

	// Login to Mojang account to get AccessToken
	auth, err := yggdrasil.Authenticate("Your E-mail", "Your Password")
	if err != nil {
		panic(err)
	}

	c.Auth.UUID, c.Auth.Name = auth.SelectedProfile()
	c.Auth.AsTk = auth.AccessToken()

	// Login
	if err := c.JoinServer("localhost"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")
}
```

## Player

The above methods will work, however on servers they will not respond to heartbeats/keepalive packets.  
In order to respond to heartbeats, etc. you can wrap the client in a `basic.Player`.

```go
package main

import (
	"log"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/yggdrasil"
)

func main() {
	c := bot.NewClient()

	// Login to Mojang account to get AccessToken
	auth, err := yggdrasil.Authenticate("Your E-mail", "Your Password")
	if err != nil {
		panic(err)
	}

	c.Auth.UUID, c.Auth.Name = auth.SelectedProfile()
	c.Auth.AsTk = auth.AccessToken()
	
	// Wrap with player
	// NOTE: basic.DefaultSettings has zh_CN as default locale
	// NOTE: This also returns a `Player` that can be used to respawn, etc. however for this example we just wrap for heartbeat keepalive
	basic.NewPlayer(c, basic.DefaultSettings)

	// Login
	if err := c.JoinServer("localhost"); err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")
}
```

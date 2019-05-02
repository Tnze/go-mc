# GO-MC
![](https://img.shields.io/badge/Minecraft-1.14-blue.svg)
![](https://img.shields.io/badge/Protocol-477-blue.svg)
[![GoDoc](https://godoc.org/github.com/Tnze/go-mc?status.svg)](https://godoc.org/github.com/Tnze/go-mc)
[![Go Report Card](https://goreportcard.com/badge/github.com/Tnze/go-mc)](https://goreportcard.com/report/github.com/Tnze/go-mc)

There's some library in Go support you to create your Minecraft client or server. 
- [x] Mojang authenticate
- [x] Minecraft network protocal
- [x] Parse chat message
- [x] Simple MC robot lib
- [ ] Parse NBT

Some examples are at `/cmd` folder.

> `1.13.2` version is at [gomcbot](https://github.com/Tnze/gomcbot).

# Getting start
After you install golang tools:
- run `go run cmd/ping/ping.go` to ping and list the Miaoscraft mc-server.  
- run `go run cmd/daze/daze.go` to join local server at *localhost:25565* as Steve on offline mode.

See `/bot` folder to get more infomation about how to create your own robot.
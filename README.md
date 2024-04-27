# Go-MC

![Version](https://img.shields.io/badge/Minecraft-1.20.2-blue.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/Tnze/go-mc.svg)](https://pkg.go.dev/github.com/Tnze/go-mc)
[![Go Report Card](https://goreportcard.com/badge/github.com/Tnze/go-mc)](https://goreportcard.com/report/github.com/Tnze/go-mc)
[![Discord](https://img.shields.io/discord/915805561138860063?label=Discord)](https://discord.gg/A4qh8BT8Ue)

### [教程 · Tutorial](https://go-mc.github.io/tutorial/)
### [文档 · Documents](https://pkg.go.dev/github.com/Tnze/go-mc)

Require Go version: 1.20

There's some library in Go support you to create your Minecraft client or server.  
这是一些Golang库，用于帮助你编写自己的Minecraft客户端或服务器。

- [x] 👍 Minecraft network protocol
- [x] 👍 Robot framework
- [x] 👍 Server framework
- [x] 👍 Dual role RCON protocol (Server & Client)
- [x] 👍 Chat Message (Support both Json and old `§` format)
- [x] 👍 NBT (Based on reflection)
- [x] 👌 SNBT ⇋ NBT
- [x] 👍 Regions & Chunks & Blocks
- [x] ⌛ Yggdrasil (Mojang login)
- [x] ⌛ Realms Server

> We don't promise that API is 100% backward compatible.

## Getting start

Go-MC tag the old version after new version released. For example,
if *1.19.4* is the latest Minecraft version, the newest go-mc tag will be *v1.19.3*.
To get the latest Go-MC that support *1.19.4*, usually you must use `go get -u github.com/Tnze/go-mc@master`.
Special cases are version like *1.19*, the Go-MC support it is tagged `v1.19.0` to avoid automatically upgrade. 

Examples:  
To get the latest version: `go get github.com/Tnze/go-mc@master`  
To get old versions (e.g. 1.18.2): `go get github.com/Tnze/go-mc@v1.18.2`
To get the first of each primary version: `go get github.com/Tnze/go-mc@v1.19.0`

### Run Examples

- Run `go run github.com/Tnze/go-mc/examples/mcping localhost` to ping and list the localhost mc server.
- Run `go run github.com/Tnze/go-mc/examples/daze` to join the local server at *localhost:25565* as player named Daze on the offline mode.

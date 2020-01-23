# Go-MC
![Version](https://img.shields.io/badge/Minecraft-1.15.2-blue.svg)
![Protocol](https://img.shields.io/badge/Protocol-578-blue.svg)
[![GoDoc](https://godoc.org/github.com/Tnze/go-mc?status.svg)](https://godoc.org/github.com/Tnze/go-mc)
[![Go Report Card](https://goreportcard.com/badge/github.com/Tnze/go-mc)](https://goreportcard.com/report/github.com/Tnze/go-mc)
[![Build Status](https://travis-ci.org/Tnze/go-mc.svg?branch=master)](https://travis-ci.org/Tnze/go-mc)

There's some library in Go support you to create your Minecraft client or server.  
这是一些Golang库，用于帮助你编写自己的Minecraft客户端或服务器，
- [x] Chat
- [x] NBT
- [x] Yggdrasil
- [x] Realms Server
- [x] RCON protocol
- [x] Saves decoding /encoding
- [x] Minecraft network protocol
- [x] Simple MC robot lib

bot:  
- [x] Swing arm
- [x] Get inventory
- [x] Pick item
- [x] Drop item
- [x] Swap item in hands
- [x] Use item
- [x] Use entity
- [x] Attack entity
- [x] Use/Place block
- [x] Mine block
- [x] Custom packets
- [ ] Record entities


> 由于仍在开发中，部分API在未来版本中可能会变动

Some examples are at `/cmd` folder.  
有一些例子在cmd目录下

> `1.13.2` version is at [gomcbot](https://github.com/Tnze/gomcbot).

# Getting start
After you install golang tools:  
To get latest version: `go get github.com/Tnze/go-mc@master`  
To get old versions (eg. 1.14.3): `go get github.com/Tnze/go-mc@v1.14.3`
- Run `go run github.com/Tnze/go-mc/cmd/mcping localhost` to ping and list the localhost mc server.  
- Run `go run github.com/Tnze/go-mc/cmd/daze` to join local server at *localhost:25565* as Steve on offline mode.

See `/bot` folder to get more information about how to create your own robot.

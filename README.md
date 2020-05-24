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

## Getting start
After you install golang tools:  
To get latest version: `go get github.com/Tnze/go-mc@master`  
To get old versions (eg. 1.14.3): `go get github.com/Tnze/go-mc@v1.14.3`

First of all, you might have a try of the simple examples. It's a good start.

### Run Examples

- Run `go run github.com/Tnze/go-mc/cmd/mcping localhost` to ping and list the localhost mc server.  
- Run `go run github.com/Tnze/go-mc/cmd/daze` to join local server at *localhost:25565* as Steve on offline mode.

### Basic Useage

One of the most useful functions of this lib is that it implements the network communication protocol of minecraft. It allows you to construct, send, receive, and parse network packets. All of them are encapsulated in `go-mc/net` and `go-mc/net/packet`.

```go
import "github.com/Tnze/go-mc/net"
import pk "github.com/Tnze/go-mc/net/packet"
```

It's very easy to create a packet. For example, after any client connected the server, it sends a [Handshake Packet](https://wiki.vg/Protocol#Handshake). You can create this package with the following code:

```go
p := pk.Marshal(
    0x00,                       // Handshake packet ID
    pk.VarInt(ProtocolVersion), // Protocol version
    pk.String("localhost"),     // Server's address
    pk.UnsignedShort(25565),    // Server's port
    pk.Byte(1),                 // 1 for status ping, 2 for login
)
```

Then you can send it to server using `conn.WritePacket(p)`. The `conn` is a `net.Conn` which is returned by `net.Dial()`. And don't forget to handle the error.^_^

Receiving packet is quite easy too. To read a packet, call `p.Scan()` like this:

```go
var (
    x, y, z    pk.Double
    yaw, pitch pk.Float
    flags      pk.Byte
    TeleportID pk.VarInt
)

err := p.Scan(&x, &y, &z, &yaw, &pitch, &flags, &TeleportID)
if err != nil {
    return err
}
```

As the net package implements the minecraft network protocol, there is no update between the versions at this level. So net package actually supports any version. It's just that the ID and content of the package are different between different versions.

Originally it's all right to write a bot with only `go-mc/net` package. But considering that the process of handshake, login and encryption is not difficult but complicated, I have implemented it in `go-mc/bot` package. You may use it directly or as a reference for your own implementation.

Now, go and have a look at the example! 
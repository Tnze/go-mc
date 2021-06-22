# NBT [![Go Reference](https://pkg.go.dev/badge/github.com/Tnze/go-mc/nbt.svg)](https://pkg.go.dev/github.com/Tnze/go-mc/nbt)

This package implement the [Named Binary Tag](https://wiki.vg/NBT) format of Minecraft.

The API is very similar to the standard library `encoding/json`. If you (high probability) have used that, it is easy to
use this.

## Basic Usage

For the following NBT tag:

```
TAG_Compound("hello world") {
    TAG_String("name"): "Bananrama"
}   
```

To read and write would look like:

```go
package main

import "github.com/Tnze/go-mc/nbt"

type Compound struct {
    Name string `nbt:"name"` // The field must be started with the capital letter
}

func main() {
    banana := Compound{Name: "Bananrama"}
    data, _ := nbt.Marshal(banana)

    var rama Compound
    _ = nbt.Unmarshal(data, &rama)
}
```

## Struct field tags

There are two tags supported:

- nbt
- nbt_type

The `nbt` tag is used to change the name of the NBT Tag field, whereas the `nbt_type`
tag is used to enforce a certain NBT Tag type when it is ambiguous.

For example:

```go
type Compound struct {
    LongArray []int64
    LongList []int64 `nbt_type:"list"` // forces a long list instead of a long array
}
```

## Stringified NBT

You can transform SNBT string into binary NBT by using go-mc, which could be useful when parsing command, but the
reverse conversion is not supported at the moment.

```go
package main

import (
    "bytes"
    "github.com/Tnze/go-mc/nbt"
)

func main() {
    var buf bytes.Buffer
    e := nbt.NewEncoder(&buf)
    err := e.WriteSNBT(`{Tnze: 1, 'sp ace': [I;1,2,3]}`)
    if err != nil {
        panic(err)
    }
}
```
# NBT
This package implement the Named Binary Tag format of Minecraft. It supports all formats
of NBT including compact arrays for longs.

# Usage
For the following NBT tag:

```
TAG_Compound("hello world") {
    TAG_String('name'): 'Bananrama'
}   
```

To read and write would look like:

```go
package main

import "bytes"
import "github.com/Tnze/go-mc/nbt"

type Compound struct {
    Name string `nbt:"name"` // Note that if name is private (name), the field will not be used
}

func main() {
    var out bytes.Buffer
    banana := Compound{Name: "Bananrama"}
    _ = nbt.Marshal(&out, banana)

    var rama Compound
    _ = nbt.Unmarshal(out.Bytes(), &rama)
}
```

# Struct field tags
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

# Docs
[![GoDoc](https://godoc.org/github.com/Tnze/go-mc/nbt?status.svg)](https://godoc.org/github.com/Tnze/go-mc/nbt)
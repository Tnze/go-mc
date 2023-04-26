# NBT [![Go Reference](https://pkg.go.dev/badge/github.com/Tnze/go-mc/nbt.svg)](https://pkg.go.dev/github.com/Tnze/go-mc/nbt)

This package implements the [Named Binary Tag](https://wiki.vg/NBT) format of Minecraft.

The API is very similar to the standard library `encoding/json`.
(But fix some its problem)
If you (high probability) have used that, it is easy to use this.

## Supported Struct Tags and Options

- `nbt` - The primary tag name. See below.
- `nbtkey` - The key name of the field (Used to support commas `,` in tag names)

### The `nbt` tag

In most cases, you only need this one to specify the name of the tag. 

The format of `nbt` struct tag: `<nbt tag>[,opt]`.

It's a comma-separated list of options.
The first item is the name of the tag, and the rest are options.

Like this:
```go
type MyStruct struct {
    Name string `nbt:"name"`
}
```

To tell the encoder not to encode a field, use `-`:
```go
type MyStruct struct {
    Internal string `nbt:"-"`
}
```

To tell the encoder to skip the field if it is zero value, use `omitempty`:
```go
type MyStruct struct {
    Name string `nbt:"name,omitempty"`
}
```

Fields typed `[]byte`, `[]int32` and `[]int64` will be encoded as `TagByteArray`, `TagIntArray` and `TagLongArray` respectively by default.
You can override this behavior by specifying encode them as`TagList` by using `list`:
```go
type MyStruct struct {
    Data []byte `nbt:"data,list"`
}
```

### The `nbtkey` tag

Common issue with JSON standard libraries: inability to specify keys containing commas for structures.
(e.g `{"a,b" : "c"}`)

So this is a workaround for that:

```go
type MyStruct struct {
    AB string `nbt:",omitempty" nbtkey:"a,b"`
}
```
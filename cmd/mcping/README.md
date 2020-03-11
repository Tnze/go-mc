# mcping

Ping tool for Minecraft: Java Edition.
Just for example. Not recommended for daily use. Use [github.com/go-mc/mcping](github.com/go-mc/mcping) instead, which including SRV parse.

适用于Minecraft: Java Edition的ping工具。
只起示例作用，日常使用建议使用完整版[github.com/go-mc/mcping](github.com/go-mc/mcping)，包含SRV解析等功能。

Install with go tools:  
    ```go get -u github.com/Tnze/go-mc/cmd/mcping```
    `$GOPATH/bin` should in your `$PATH`.

Install with Homebrew:  
    ```brew tap Tnze/tap && brew install mcping```

Useage:  
    ```mcping <hostname>[:port]```

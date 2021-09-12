# Go Build

## Build

Go 支授多種 CPU 與作業系統平台，也可以編譯成 Android, iOS, 及 WebAssembly。執行指令：`env GOOS=target-OS GOARCH=target-architecture go build package-import-path`。 eg: 編繹成 linux 64bit 版本，`env GOOS=linux GOARCH=amd64 go build .`

在 Go 1.11 版本，支援 Web Assembly, 可以將 Go 的程式，編繹了 WebAssembly，讓瀏覽器執行。eg: `env GOOS=js GOARCH=wasm go build .`

要查支援的環境，可以執行 `go tool dist list` 查詢。


Go 編譯成 Anodroid/iOS 說明：[Gomobile](https://github.com/golang/go/wiki/Mobile)
Go WebAssembly 說明：[Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)



// TODO: Build Constrain
// TODO: Build CFlags to modify global variables.
// TODO: with Makefile
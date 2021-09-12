# Go Build

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Go Build](#go-build)
  - [1. 跨平台](#1-跨平台)
  - [2. Build Constraints](#2-build-constraints)
    - [2.1 目錄結構](#21-目錄結構)
    - [2.3 主程式說明](#23-主程式說明)
      - [mybuild/main.go](#mybuildmaingo)
    - [2.4 使用檔名控制](#24-使用檔名控制)
      - [mybuild/hello_darwin.go](#mybuildhello_darwingo)
      - [mybuild/hello_linux.go](#mybuildhello_linuxgo)
      - [mybuild/hello_windows.go](#mybuildhello_windowsgo)
    - [2.5 在檔案加入 Constraint 設定](#25-在檔案加入-constraint-設定)
      - [mybuild/start1.go](#mybuildstart1go)
      - [mybuild/start2.go](#mybuildstart2go)
      - [mybuild/start3.go](#mybuildstart3go)
    - [2.6 使用 tag 方式指定](#26-使用-tag-方式指定)
      - [mybuild/hi_gopher.go](#mybuildhi_gophergo)
      - [mybuild/hi_everyone.go](#mybuildhi_everyonego)
    - [2.7 Makefile](#27-makefile)
    - [2.8 測試](#28-測試)

<!-- /code_chunk_output -->

## 1. 跨平台

Go 支授多種 CPU 與作業系統平台，也可以編譯成 Android, iOS, 及 WebAssembly。執行指令：`env GOOS=target-OS GOARCH=target-architecture go build package-import-path`。 eg: 編繹成 linux 64bit 版本，`env GOOS=linux GOARCH=amd64 go build .`

在 Go 1.11 版本，支援 Web Assembly, 可以將 Go 的程式，編繹了 WebAssembly，讓瀏覽器執行。eg: `env GOOS=js GOARCH=wasm go build .`

要查支援的環境，可以執行 `go tool dist list` 查詢。

- Go 編譯成 Anodroid/iOS 說明：[Gomobile](https://github.com/golang/go/wiki/Mobile)
- Go WebAssembly 說明：[Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)

## 2. Build Constraints

Go 可以透過內建的 build 工具，來指定不同平台，要編譯那些檔案。詳細的說明，請見 [Go Build Constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)

此範例預計要完成：

1. 使用 Mailefile 管理
1. 可以跨平台編譯
1. 指定編譯檔案
1. 在編譯時期，修改程式變數值。

### 2.1 目錄結構

```text
mybuild
├── Makefile
├── go.mod
├── hello_darwin.go
├── hello_linux.go
├── hello_windows.go
├── hi_everyone.go
├── hi_gopher.go
├── main.go
├── say1.go
├── say2.go
├── start1.go
├── start2.go
└── start3.go
```

### 2.3 主程式說明

#### mybuild/main.go

@import "mybuild/main.go" {class=line-numbers}

1. 有一個 Global 變數 `version`。預設值是 __"debug"__。會在編譯時，修改值。
1. `Hello()`, `Start()`, `Say()`，透過 Build Constraints，指定編譯的檔案，在不同平台會輸出對應平台的訊息。
1. `Hi()` 透過 Build Tags，指令編譯的檔案，會輸出不同結果。

### 2.4 使用檔名控制

Go 的編譯器，可以依據檔名，自動認定為指定平台需要加入編譯的檔案。如：

- __hello_darwin.go__: MacOS
- __hello_linux.go__: Linux
- __hello_windows.go__: Windows

#### mybuild/hello_darwin.go

@import "mybuild/hello_darwin.go" {class=line-numbers}

#### mybuild/hello_linux.go

@import "mybuild/hello_linux.go" {class=line-numbers}

#### mybuild/hello_windows.go

@import "mybuild/hello_windows.go" {class=line-numbers}

### 2.5 在檔案加入 Constraint 設定

可以在程式碼檔案的 __第一行__ ，加入 `//go:build 平台`，如：`//go:build darwin`。請見 __start1.go__, __start2.go__, __start3.go__

#### mybuild/start1.go

@import "mybuild/start1.go" {class=line-numbers, highlight=1}

#### mybuild/start2.go

@import "mybuild/start2.go" {class=line-numbers, highlight=1}

#### mybuild/start3.go

@import "mybuild/start3.go" {class=line-numbers, highlight=1}

### 2.6 使用 tag 方式指定

使用上述 Build Constraints 的方式，加入 tag。如 __hi_gopher.go__, __hi_everyone.go__。

#### mybuild/hi_gopher.go

__正向表列__

@import "mybuild/hi_gopher.go" {class=line-numbers, highlight=1}

#### mybuild/hi_everyone.go

__負向表列__

@import "mybuild/hi_everyone.go" {class=line-numbers, highlight=1}

### 2.7 Makefile

@import "mybuild/Makefile" {as="makefile", class="line-numbers"}

1. 利用 __-ldfags__ 來修改 __main.go__ 的變數名稱。如：`-X main.version=$(VERSION)`。可以用在程式執行檔的版控。
1. 指定平台編譯，可以用 `env GOOS=OS GOARCH=CPU ...` 的方式指定。
1. 透過 `tags` 的方式，指定編譯的檔案。

### 2.8 測試

切換至 __mybuild__ 目錄。以下是執行對應的指令，與輸出的結果。

- MacOS: `make clean && make darwin GOOS=darwin`

    ```text
    version is 1.0.0
    hello darwin
    start darwin
    I am not Windows
    hi, everyone!
    ```

- MacOS + Gopher: `make clean && make gopher GOOS=darwin && make darwin`

    ```text
    version is 1.0.0
    hello darwin
    start darwin
    I am not Windows
    hi, gopher!
    ```

- Linux: `make clean && make linux GOOS=linux`

    ```text
    version is 1.0.0
    hello linux
    start linux
    I am not Windows
    hi, everyone!
    ```

- Linux + Gopher: `make clean && make gopher GOOS=linux && make linux`

    ```text
    version is 1.0.0
    hello linux
    start linux
    I am not Windows
    hi, gopher!
    ```

- Windows: `make clean && make windows GOOS=windows`

    ```text
    version is 1.0.0
    hello windows
    start windows
    I am Windows
    hi, everyone!
    ```

- Windows + Gopher: `make clean && make windows GOOS=windows`

    ```text
    version is 1.0.0
    hello windows
    start windows
    I am Windows
    hi, gopher!
    ```
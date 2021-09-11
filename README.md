# Golang 學習筆記

個人多年來學習與實作上的心得筆記，本文件適合已有一種程式語言經驗的同好閱讀，內容如有錯誤或建議，可以隨時與我連絡。
## 開發環境

- GO 版本: Go 1.17
- 開發環境: Mac OS (amd64)
- 開發工具: [VSCode](https://code.visualstudio.com/)
- 文件使用 Markdown 撰寫，建議使用 [Markdown Preview Enhanced](https://github.com/shd101wyy/markdown-preview-enhanced) 閱讀
- [Source on Github](https://github.com/kigichang/go_course)

## 主要資料來源

1. 官網：[https://golang.org/](https://golang.org/)
1. 線上學習：[A tour of Go](https://tour.golang.org/list)
1. [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS)
1. [Go Web Programming](https://www.manning.com/books/go-web-programming)
1. [Go System Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming)
1. [Go Mobile](https://github.com/golang/go/wiki/Mobile)
1. [Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)

## 其他資料

1. [Go Dev](https://go.dev/)
1. [Go Wiki](https://github.com/golang/go/wiki)
1. [Awesome Go](https://awesome-go.com/)
1. [Effective Go](https://golang.org/doc/effective_go) (必讀)

## Summary

- [本文件](README.md)
### 一、Go 基礎說明

- [01 開發環境與語言簡介](01_introduction)
  - 參考文件
  - IDE 設定
  - 與 C/Java/PHP 簡單比較
- [02 程式結構與語法](02_syntax)
  - 關鍵字
  - 基本語法
  - `iota`
  - `type`
  - 指標
  - Package
- [03 Data Types - Basic Types](03_basic_types)
  - Number
  - Boolean
  - String
  - Zero Value
- [04 Data Types - Aggregate Types](04_aggregate_types)
  - Array
  - Struct
  - JSON
- [05 Data Types - Reference Types](05_reference_types)
  - Slice
    - Append Slice
  - Map
- [06 Functions](06_functions)
  - 語法
  - Pass By Value in Value and Reference Types.
  - First Class
- [07 Methods](07_methods)
  - Methods in Value and Pointer
- [08 Interface](08_interface)
  - Interface in Struct and Pointer
  - Interface value
  - Interface puzzle (interface nil problem)
- [09 Go and OOP](09_go_and_oop)
  - Encapsulation
  - Inheritance (fake)
  - Polymorphism
- [10 Defer and Error Handling](10_defer_and_error_handling)
  - Defer
    - Closure Binding
  - Panic and Recover
  - Errors (new feature in Go 1.13)
- [14 Testing](14_testing)
- [Build and Dependency](class_build_dependency)
  - Build cross platform
  - Go Module `go mod`
### 二、多執行緒
- [11 Concurrency - Goroutine](11_goroutine)
  - Keyword `go`
  - `sync.WaitGroup`
- [12 Concurrency - Channel](12_channel)
  - Buffered channel
  - Producer and Consumer Pattern
  - Actor Pattern
  - `select` to monitor channels
- [13 Context](13_context)

### 三、實作應用
- [15 flag and spf13 Cobra/Viper](15_flag_cobra_viper)
- [16 MySQL](16_mysql)
- [17 Web](class17/)
  - Go Template 語法
  - Context in Request (Request-Scoped)
  - Cookie
- [Gorilla](gorilla/)
  - Mux
  - Middleware
  - Shema
  - Secure Cookie
  - CSRF
- [Echo](echo/)
  - Hello world
- [18 RESTful, Protobuf and gRPC](class18/)
  - RESTful using Gorilla
  - Protobuf
    - Protoc (Protobuf Compliler)
  - gRPC (Client and Service)

### 四、Go 進階功能

- [19 Reflection and Struct Tag](class19/)
  - Type and Value
  - Strut Tag
  - Check Interface implementation
  - Zero Value
  - Make Slice
  - Make Map
  - Make Function
- [20 Cowork with C/C++](class20/)
  - Go `unsafe` Package
  - Go String and *C.char
  - Go call C
    - Swig
    - DIY
  - C Call Go with Static Library

### 五、實驗中功能

- [21 Go WebAssembly](class21/)
  - WebAssembly Introduction
  - DOM in Go WASM
    - Selector
    - Property
    - Method
    - Event
  - Create a Javascript Object
  - Go call Javascript
  - Javascript Call Go
  - File and FileReader
  - Conversion Javascript Uint8Array and Go Byte Slice

- [x Type Parameters and Go Generic]

## 新增

- Go 1.13 Error 功能
- Go 1.16 embed 功能
- Go WebAssembly
- Go Generic & Go2Go (Go 1.8)

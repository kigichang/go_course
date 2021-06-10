# Golang 學習筆記

- GO 版本: Go 1.16
- 開發環境: Mac OS
- [Source on Github](https://github.com/kigichang/go_course)
- 文件使用 Markdown 撰寫，建議使用 [Markdown Preview Enhanced](https://github.com/shd101wyy/markdown-preview-enhanced) 閱讀

資料主要來自：

1. 官網：[https://golang.org/](https://golang.org/)
1. 線上學習：[A tour of Go](https://tour.golang.org/list)
1. [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS)
1. [Go Web Programming](https://www.manning.com/books/go-web-programming)
1. [Go System Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming)
1. [Go Mobile](https://github.com/golang/go/wiki/Mobile)
1. [Go WebAssembly](https://github.com/golang/go/wiki/WebAssembly)

其他資源：

1. [Go Dev](https://go.dev/)
1. [Go Wiki](https://github.com/golang/go/wiki)
1. [Awesome Go](https://awesome-go.com/)
1. [Effective Go](https://golang.org/doc/effective_go)

## Summary

- [Summary](README.md)
- [01 開發環境與語言簡介](class01/)
  - 參考文件
  - IDE 設定
  - 與 C/Java/PHP 簡單比較
- [02 程式結構與語法](class02/)
  - 關鍵字
  - 基本語法
  - `iota`
  - `type`
  - 指標
  - Package
- [03 Data Types - Basic Types](class03/)
  - Number
  - Boolean
  - String
  - Zero Value
- [04 Data Types - Aggregate Types](class04/)
  - Array
  - Struct
  - JSON
- [05 Data Types - Reference Types](class05/)
  - Slice
    - Append Slice
  - Map
- [06 Functions](class06/)
  - 語法
  - Pass By Value in Value and Reference Types.
  - First Class
- [07 Methods](class07/)
  - Methods in Value and Pointer
- [Build and Dependency](class_build_dependency)
  - Build cross platform
  - Go Module `go mod`
- [08 Interface](class08/)
  - Interface in Struct and Pointer
  - Interface value
  - Interface puzzle (interface nil problem)
- [09 OOP in Go](class09/)
  - Encapsulation
  - Inheritance (fake)
  - Polymorphism
- [10 Defer and Error Handling](class10/)
  - Defer
    - Closure Binding
  - Panic and Recover
  - Errors (new feature in Go 1.13)
- [11 Concurrency - Goroutine](class11/)
  - Keyword `go`
  - `sync.WaitGroup`
- [12 Concurrency - Channel](class12/)
  - Buffered channel
  - Producer and Consumer Pattern
  - Actor Pattern
  - `select` to monitor channels
- [13 Context](class13/)
- [14 Testing](class14/)
- [15 flag and spf13 Cobra/Viper](class15/)
- [16 MySQL](class16/)
- [17 Web (Gorilla web toolkit)](class17/)
  - Go Template 語法
  - Context in Request (Request-Scoped)
  - Cookie
  - Gorilla
    - Mux
      - Middleware
    - Shema
    - Secure Cookie
    - CSRF
- [18 RESTful, Protobuf and gRPC](class18/)
  - RESTful using Gorilla
  - Protobuf
    - Protoc (Protobuf Compliler)
  - gRPC (Client and Service)
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

## 新增

- Go 1.13 Error 新功能
- Go WebAssembly

# 01 開發環境與語言簡介

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [01 開發環境與語言簡介](#01-開發環境與語言簡介)
  - [0. 前言](#0-前言)
  - [1. 相關資源](#1-相關資源)
    - [1.1 我的參考書籍](#11-我的參考書籍)
  - [2. 環境設定](#2-環境設定)
    - [2.1 IDE 建議](#21-ide-建議)
    - [2.2 GOPATH 目錄說明](#22-gopath-目錄說明)
  - [3. 語言特性](#3-語言特性)
    - [3.1 編譯式語言](#31-編譯式語言)
    - [3.2 Strong Type](#32-strong-type)
    - [3.3 可在不同平台執行](#33-可在不同平台執行)
    - [3.4 有 Garbage Collection](#34-有-garbage-collection)
    - [3.5 有指標 (Pointer)](#35-有指標-pointer)
    - [3.6 沒有物件導向 (OO, Object Oriented)](#36-沒有物件導向-oo-object-oriented)
    - [3.7 變數 Visibility](#37-變數-visibility)
    - [3.8 Summary](#38-summary)
  - [4. Soruce Code 放法](#4-soruce-code-放法)
  - [5. Hello World (ex01-01)](#5-hello-world-ex01-01)
    - [5.1 說明](#51-說明)
  - [6. Arguemnts (ex01-02)](#6-arguemnts-ex01-02)

<!-- /code_chunk_output -->

## 0. 前言

寫 Go 建議用 Git 當版控工具。基本上，Go 相關的套件，絕大部分放在 Github 上。  
本文件是使用 Go 1.13 以上的版本，會使用 go module 來管理 package。

## 1. 相關資源

自己的練習過程：

1. 官網：[https://golang.org/](https://golang.org/)
1. 線上學習：[A tour of Go](https://tour.golang.org/welcome/1)。請務必要上網練習，玩過一輪後，差不多就學完最基本的語法，再撘配 [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS) 書，效果會比較好。
1. [Effective Go](https://golang.org/doc/effective_go.html): Go 上手後，一定要看。

### 1.1 我的參考書籍

筆記中有關的說明與範例是出自下以書籍：

1. [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS)
1. [Go Web Programming](https://www.manning.com/books/go-web-programming)
1. [Go System Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming)

第一本看完就差不多了；第二本主要是講 Web，包含 DB, 版型；第三本主要是討論跟作業系統互動，對 goroutine, channel 有比較深入的說明。

## 2. 環境設定

1. [下載 golang](https://golang.org/dl/)。
1. 設定環境變數 __`$GOPATH`__，如：`YOUR_HOME/go`，所有 Go 相關的工具，會一併裝在這個目錄。並將 `$GOPATH/bin` 也加入 `$PATH`。
1. Go 1.11 之後，有官方有支援 module  的功能，撰寫程式，不一定要放在 `$GOPATH` 下，可以依專案有各自的目錄，也不需要將目錄設定在 `$GOPATH` 內。

### 2.1 IDE 建議

建議用 Visual Studio Code，再安裝 Go Plugin。因為使用到 [PlantUML](https://plantuml.com/)，需要安裝 JDK。

[Go by Microsoft](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

其他相關的 plugin (非必要，但為了開發方便，還是裝一下)

1. [Markdown Preview Enhanced](https://marketplace.visualstudio.com/items?itemName=shd101wyy.markdown-preview-enhanced): Markdown 工具，支援數學公式，流程圖，心智圖，匯出 PDF 等多種功能。
1. [TabNine Autocomplete AI](https://marketplace.visualstudio.com/items?itemName=TabNine.tabnine-vscode): 很熱的 AI 自動補碼工具。
1. [GitLens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens): Git 工具。
1. [hexdump for VSCode](https://marketplace.visualstudio.com/items?itemName=slevesque.vscode-hexdump): 查看二進位檔案。

### 2.2 GOPATH 目錄說明

```text
.
├── bin
├── pkg
└── src
```

- bin: 主要放 Go 相關的工具程式，VSCode 的 Go plugin 用到的工具，都會被在這邊。~~及專案的執行檔~~
- pkg: ~~編譯過程會產生的中間檔~~，存放 Go module 下載來的相依 package。
~~- src: 放 source code.~~
- 使用 Go Module 的話，則專案不需要放在 GOPATH 下。

## 3. 語言特性

Go 是編譯式，強型別 (Strong Type)，且有垃圾回收 (Garbage Collection) 機制的程式語言。
### 3.1 編譯式語言

需要有編譯器 (Compiler)，將程式碼編譯成機器碼 (Machine Code) 或 Byte Code。Go 是編譯成 __Machine Code__。

- 與 PHP 直譯式不同，程式碼需要經過編譯成執行檔，才可以用。
- 與 Java 不同，直接編譯成對應的 OS 與 CPU 的 Machine Code。

### 3.2 Strong Type

變數宣告後，它的資料型別也就固定，不能更改。不像 PHP/Pyhone 可以隨意更改變數的資料型別。

### 3.3 可在不同平台執行

Go 可以實現跨平台，前題是沒使用到平台特有的函式庫。Go 的跟平台與 Java 不同，需要重新 compile 成相對應的平台的 Machine Code。

- Go 可以在 MacOS 環境，編譯出 Windows, 或 Linux 的執行檔。
- Go 可以在編譯過程，指定需要(或不需要)編譯的程式碼。

### 3.4 有 Garbage Collection

與 Java 類似，有 GC，可以不用自己管理記憶體，但也要注意，免得浪費記憶體還不知道。

### 3.5 有指標 (Pointer)

- 程式撰寫的觀念與 C 類似，有 pointer，但 pointer 不能做 pointer 位移，pointer 也不可以做 cast。
- 可以使用內建的 `unsafe` package 來轉成 C 的 pointer。

### 3.6 沒有物件導向 (OO, Object Oriented)

OOP 有三個基本特性:
1. 封裝
1. 繼承
1. 多型

Go __沒有繼承__。但 Go 可以透 Anonymous Embbed 與 Interface 達到繼承的效果。
### 3.7 變數 Visibility

只有 public, private, 使用變數的第一個字母 __大__、 __小__ 寫來區分。

- 大寫：__Public__
- 小寫：__Private__

### 3.8 Summary

- 寫 Go 與寫 C 類似，但有 GC，可以省去記憶體管理工作.
- 因為沒有 OO，只有封裝，沒有繼承等功能
- 沒有泛型 (Generic) 目前已計劃實作。可參考：
	1. [A Proposal for Adding Generics to Go](https://blog.golang.org/generics-proposal)
	1. [Go Type Parameters](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)
- 檔案的編碼，一定要是 **UTF-8**。
- 參數傳遞 (Paramenter Passing) 是 **Pass by Value** (Call by Value)。
- 官方有工具會自動校正 Coding style.

## 4. Soruce Code 放法

1. 每個專案自己開一個目錄。自己有各自的 git repo。
1. 專案的主目錄下，每個 package 開一個目錄。
1. 一個目錄只能有一個 package 及測試的 package。package 名稱建議 __要與目錄名稱相同__；如此比較好維護程式碼，想找某個 package 時，只要去找相對應的目錄即可。

## 5. Hello World (ex01-01)

1. 為專案產生一個目錄。
1. 在上述 1 的目錄下，執行 `go mod init ex01-01`，產生 `go.mod` 檔案。此檔案是用來設定依賴的 package。
1. 產生一個檔案 `main.go` 內容如下：

	```go {.line-numbers}
	package main

	import "fmt"

	func main() {
		fmt.Println("Hello, 世界")
	}
	```

1. 在目錄下，執行 `go run .`，可以看到結果。
1. 在目錄下，執行 `go build`，編繹成執行檔。

### 5.1 說明

1. 寫執行檔的程式，檔名不一定要命名成 `main.go`，但建議使用 `main.go`，程式碼的 package 宣告一定要是 **main**。
1. 經過 build 之後，產生的執行檔名，會是當初 go module 初始化的名稱。eg: `go mod init mytest`，則編譯的檔名就會是 `mytest`。
1. 進到目錄下，執行 `go run .` 的方式來執行。
1. `import` 是將會用到的 package 加入，跟 Java 一樣，有用到的 package 用 import 加入。Go 的工具，會幫忙找內建的 package ，自動加入到程式碼中，很方便。如果是第三方套件，就要修改 `go.mod`，通常 IDE 工具都會自動編輯這個檔案。如果沒有的話，則自己修改後，在該目錄下執行 `go mod tidy` 則會自動更新依賴的 package。
1. 程式的進入點 (Entry point): `func main()`，跟大多數的程式語言一樣，寫執行檔都會需要有一個主函式 **main**

## 6. Arguemnts (ex01-02)

重覆上述的動作，sample code 如下：

```go {.line-numbers}
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	if len(os.Args) > 1 {
		fmt.Println("hi, ", os.Args[1])
	}
	fmt.Println("hello world")
}
```

1. 執行方式：`go run .` or `go run . Gopher`
利用 `os.Args` 來接 command line 傳進來的參數。`os.Args[0]` 是執行檔的完整檔名，所以傳入的參數值要從 `os.Args[1]` 開始。Golang 有內建 `flag` 套件來管理參數，但如果要寫較複雜的 command line 程式，建議用 [spf13/cobra](https://github.com/spf13/cobra)。

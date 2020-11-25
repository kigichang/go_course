# 01 開發環境與語言簡介

寫 Go 建議用 Git 當 source control。基本上，Go 相關的套件，絕大部分放在 Github 上。  
本文件是使用 Go 1.13 以上的版本，會使用 go module 來管理 package。

資源：

1. 官網：[https://golang.org/](https://golang.org/)
1. 線上學習：[A tour of Go](https://tour.golang.org/welcome/1)。請務必要上網練習，玩過一輪後，差不多就學完最基本的語法，再撘配 [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS) 書，效果會比較好。
1. [Effective Go](https://golang.org/doc/effective_go.html): Go 上手後，一定要看。

Books:

1. [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS)
1. [Go Web Programming](https://www.manning.com/books/go-web-programming)
1. [Go System Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming)

第一本看完就差不多了；第二本主要是講 Web，包含 DB, 版型；第三本主要是討論跟作業系統互動，對 routine, channel 有比較深入的說明。

## 環境設定

1. 下載 golang。[~下載連結~](https://golang.org/dl/)
1. 設定環境變數 **`$GOPATH`**: `$GOPATH` 是專門放 Go 開發專案的目錄，所有 Go 相關的工具，也會一併裝在這個目錄。
1. Go 1.11 之後，有官方有支援 module  的功能，撰寫程式，不一定要放在 `$GOPATH` 下，可以依專案有各自的目錄，也不需要將目錄設定在 `$GOPATH` 內。

## IDE 建議

建議用 Visual Studio Code，再安裝 Go Plugin

[Go by Microsoft](https://marketplace.visualstudio.com/items?itemName=ms-vscode.Go)

其他相關的 plugin (非必要，但為了開發方便，還是裝一下)

1. [gotemplate-syntax](https://marketplace.visualstudio.com/items?itemName=casualjim.gotemplate)
1. [vscode-proto3](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3)
1. [Docker by Microsoft](https://marketplace.visualstudio.com/items?itemName=PeterJausovec.vscode-docker)
1. [Markdown Preview Enhanced](https://marketplace.visualstudio.com/items?itemName=shd101wyy.markdown-preview-enhanced)
1. [markdownlint](https://marketplace.visualstudio.com/items?itemName=DavidAnson.vscode-markdownlint)
1. [hexdump for VSCode](https://marketplace.visualstudio.com/items?itemName=slevesque.vscode-hexdump)
1. [Compareit](https://marketplace.visualstudio.com/items?itemName=in4margaret.compareit)
1. [TabNine Autocomplete AI](https://marketplace.visualstudio.com/items?itemName=TabNine.tabnine-vscode)

## GOPATH 目錄說明

```text
.
├── bin
├── pkg
└── src
```

- bin: 主要放 Go 相關的工具程式，及專案的執行檔
- pkg: 編譯過程會產生的中間檔
- src: 放 source code.
- 使用 Go Module 的話，則專案不需要放在 GOPATH 下。

### Soruce code 放法

1. 每個專案自己開一個目錄。自己有各自的 git repo。
1. 專案的主目錄下，每個 package 開一個目錄。
1. 一個目錄只能有一個 package 及測試的 package。package 名稱建議要與目錄名稱相同；如此比較好維護程式碼，想找某個 package 時，只要去找相對應的目錄即可。

## 語言特性

### 編譯式語言

- 與 PHP 直譯式不同，程式碼需要經過編譯成執行檔，才可以用。
- 與 Java 不同，直接編譯成 os 平台的 machine code。

### Strong Type

變數宣告後，它的資料型別也就固定，不能更改。不像 PHP 可以隨意更改變數的資料型別。

### 可在不同平台執行

與 Java 不同，需要重新 compile 成相對應的平台。

### 有 Garbage Collection

與 Java 類似，有 GC，可以不用自己管理記憶體，但也要注意，免得浪費記憶體還不知道。

### 有 pointers

- 程式撰寫的觀念與 C 類似，有 pointer，但 pointer 不能做 pointer 位移，pointer 也不可以做 cast。
- 可以使用內建的 `unsafe` package 來轉成 C 的 pointer。

### 沒有 OO (Object Oriented)

OOP 有三個基本特性: 封裝，繼承，多型。而 Go 沒有繼承。在 Go 要有繼承的效果，需要用 interface 的方式來達成。

### 變數 Visibility

只有 public, private, 使用變數的第一個字母大、小寫來區分。

- 大寫：public
- 小寫: private

### Summary

- 寫 Go 與寫 C 類似，但有 GC，可以省去記憶體管理工作.
- 因為沒有 OO，只有封裝，沒有繼承等功能
- 沒有泛型 (Generic) 所以有關型別方面的寫作，就沒這個彈性。
- 有工具會自動校正 coding style.
- 檔案的編碼，一定要是 **UTF-8**。
- Function paramenter **pass by value** (call by value)

## Hello World (ex01-01)

1. 在 src 下開一個目錄。
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

### 說明

1. 寫執行檔的程式，檔名不一定要命名成 `main.go`，但程式碼的 package 宣告一定要是 **main**。
1. 經過 build 之後，產生的執行檔名，會是當初 go module 初始化的名稱。eg: `go mod init mytest`，則編譯的檔名就會是 `mytest`。
1. 可以使用 `go run main.go` 直接執行程式，如果程式是拆分成多個 .go 的檔案，則需要將每個檔名也加入。eg: `go run main.go a.go b.go`，也可進到目錄下，執行 `go run .` 的方式來執行。
1. `import` 是將會用到的 package 加入，跟 Java 一樣，有用到的 package 用 import 加入。Go 的工具，會幫忙找內建的 package ，自動加入到程式碼中，很方便。如果是第三方套件，就要修改 `go.mod`，通當 IDE 工具都會自動編輯這個檔案。如果沒有的話，則自己修改後，在該目錄下執行 `go mod tidy` 則會自動更新依賴的 package。
1. 程式的進入點 (Entry point): `func main()`，跟大多數的程式語言一樣，寫執行檔都會需要有一個主函式 **main**

## Arguemnts (ex01-02)

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

利用 `os.Args` 來接 command line 傳進來的參數。`os.Args[0]` 是執行檔的完整檔名，所以傳入的參數值要從 `os.Args[1]` 開始。Golang 有內建 `flag` 套件來管理參數，但如果要寫較複雜的 command line 程式，建議用 [spf13/cobra](https://github.com/spf13/cobra)。

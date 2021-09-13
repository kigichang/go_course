# Generic


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [Generic](#generic)
  - [0. 前言](#0-前言)
  - [1. 什麼是泛型 (Generic)](#1-什麼是泛型-generic)
  - [2. Type Parameters](#2-type-parameters)
  - [3. Go 1.17 支援](#3-go-117-支援)
      - [generic_1.17/main.go](#generic_117maingo)
      - [generic_1.17/Makefile](#generic_117makefile)
  - [4. Go2Go](#4-go2go)
    - [4.1 自行編譯 Go2Go 工具](#41-自行編譯-go2go-工具)
    - [4.2 Go2Go 實作](#42-go2go-實作)
      - [generic_go2go/main.go2](#generic_go2gomaingo2)
      - [generic_go2go/Makefile](#generic_go2gomakefile)
      - [generic_go2go/main.go](#generic_go2gomaingo)
  - [5. Go Generic](#5-go-generic)
    - [5.1 自定義限制條件](#51-自定義限制條件)
    - [5.2 實作](#52-實作)
  - [6. Summary](#6-summary)

<!-- /code_chunk_output -->

## 0. 前言

目前 Go 泛型 (Generic) 正在實作中，預計在 2022 年的 Go 1.18 版推出。相關的說明，可以看：

1. [A Proposal for Adding Generics to Go](https://go.dev/blog/generics-proposal)
1. [Type Parameters Proposal](https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md)


像 Java / Scala 都有支援泛型，因 GO 本身沒有 OOP，在彈性上，不如 Java / Scala。

## 1. 什麼是泛型 (Generic)

因為 Go 是強型別語言，當某些資料型別(如：int, float64)，有共同的行為(如: Add)時，我們必須為每種資料型別，實作相同行為。如：

```go {.line-numbers}
func AddInt(a, b int) int {
	return a + b
}

func AddFloat(a, b float64) float64 {
	return a + b
}
```

上述實作中，`AddInt` 與 `AddFloat` 中，除了資料型別外，程式碼都相同 `return a + b`。

如果想達到共用的目的，可以使用 __interface{}__ 與 Type Assertion 機制實作。但實作上會比較麻煩，且只能在執行時期確認資料型別是否正確，也會造成系統的不穩定。

因此泛型 (Generic) 的機制就是想要達到：

1. 在編譯時期檢查資料型別是否符合。
1. 達到程式碼共用。

## 2. Type Parameters

泛型有一個很重要觀念，就是 __Type Parameters__ 。一般我們很憝悉 Function 的參數與回傳值，都有各自的資料型別，如 `func AddInt(a, b int) int`。 __Type Parameters__ 的概念是 __資料型別__ 也是 Function 的一種參數。在 `func AddInt(a, b int) int` 裏，當 __int__ 也是 `AddInt` 的參數時，我們就有更多實作上的彈性。

## 3. Go 1.17 支援

目前 Go 1.17 版本，有支援編譯 Generic 語法。在編譯時，加入 __-gcflags=-G=3__。

#### generic_1.17/main.go

@import "generic_1.17/main.go" {class="line-numbers"}

#### generic_1.17/Makefile

@import "generic_1.17/Makefile" {as="makefile" class="line-numbers"}

由於還在實驗中，目前 1.17 版本不允許將 Generic 相關的實作公開，也因此上述程式，都只能用 private 形式。如果改成 public (改大寫)，則會 compile 錯誤。
## 4. Go2Go

Go 官方，有提供練習 Generic 語法網站：[The go2go Playground](https://go2goplay.golang.org/)。原理是先將有泛型的程式碼，轉成沒有泛型的程式碼，再編譯執行。

也可以自行編譯 Go2Go 的工具。

### 4.1 自行編譯 Go2Go 工具

我自己環境的設定方式如下：

1. 切換至 __HOME__ 目錄。
1. 執行 `git clone -b dev.go2go  git@github.com:golang/go.git goroot`。
1. 進到 __goroot/src__ 執行 __all.bash__ 進行編譯。
1. 執行 `~/goroot/bin/go version` 確認版本。
1. 如果要開始實作大型專案，則需要設定 __GO2PATH__ 環境變數。
    - 在 __HOME__ 目錄下，建立 __go2__ 目錄，並設定環境變數 __$GO2PATH__。

可以參考 __goroot/src/cmd/go2go/testdata__ Generic 與 Monoid 範例可學習。

### 4.2 Go2Go 實作

要使用 Go2GO 必須將程式碼的副檔案，改成 __*.go2__。使用 `~/goroot/bin/go tool go2go build` 方式編譯程式。編譯完成後，會看到 go2go 產生的程式碼 __main.go__ 。

#### generic_go2go/main.go2

@import "generic_go2go/main.go2" {as="go" class="line-numbers"}

#### generic_go2go/Makefile

@import "generic_go2go/Makefile" {as="makefile" class="line-numbers"}

#### generic_go2go/main.go

@import "generic_go2go/main.go" {class="line-numbers"}

## 5. Go Generic

在泛型依然要定義可支援的資料型別，如：`Add[T Number]` 中的 __T__ 就被限制 (Constraint) 是 __Number__。

目前已知 Go 內建的限制條件：

- __any__: 沒有限制，可以是任何資料型別。
- __comparable__: 支援 `==` 與 `!=` 操作的資料型別。
  - map 內的 key 的資料型別，一定要是 __comparable__。

### 5.1 自定義限制條件

可以自定義限制條件，如：

```go {.line-numbers}
type Number interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64
}
```

1. 宣告新的限制條件為 __Number__: `type Number interface`
1. 設定此條件下，支援的資料型別：

    ```go {.line-numbers}
    type int, int8, int16, int32, int64,
            uint, uint8, uint16, uint32, uint64,
            float32, float64
    ```

### 5.2 實作

使用自定義的限制條件。

```go {.line-numbers}
func Add[T Number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(1, 2), reflect.TypeOf(Add(1, 2)))
	fmt.Println(Add(1.0, 2.0), reflect.TypeOf(Add(1.0, 2.0)))
}
```

1. 泛型宣告：`func Add[T Number](a, b T) T`
    1. 定義資料型別的限制條件：`[T Number]`
    1. 將原本 Function 中資料型別宣告，改成 `T`。如：`func AddInt(a, b int) int` => `func Add[T Number](a, b T) T` 中 __int__ => __T__。

## 6. Summary

Generic 即將在 2022 年推出，之後 Go 相關的 eco-system 又會是一場大改版。可以在這段時間多多預習。
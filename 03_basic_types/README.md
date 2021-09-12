# 03 Data Types - Basic Types


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [03 Data Types - Basic Types](#03-data-types-basic-types)
  - [0. 前言](#0-前言)
  - [1. Zero Value (ex03_01)](#1-zero-value-ex03_01)
  - [2. Numbers](#2-numbers)
    - [2.1 Integers](#21-integers)
    - [2.2 Float-Point Numbers](#22-float-point-numbers)
    - [2.3 Complex Numbers (複數)](#23-complex-numbers-複數)
  - [3. Booleans](#3-booleans)
  - [4. Strings](#4-strings)
    - [4.1 宣告](#41-宣告)
    - [4.2 取得字串長度](#42-取得字串長度)
    - [4.3 Substring](#43-substring)
    - [4.4 字串連接 (Concate)](#44-字串連接-concate)
    - [4.5 ==strings== Package](#45-strings-package)
    - [4.6 Rune (ex03_02)](#46-rune-ex03_02)
  - [5. 字串與數字轉換 (ex03_03)](#5-字串與數字轉換-ex03_03)

<!-- /code_chunk_output -->

## 0. 前言

Go 的 Data Type 分成四個類別：

- Basic Type
  - numbers
  - strings
  - booleans
- Aggregate Types
  - arrays
  - structs
- Reference Types
  - pointers
  - slices
  - maps
  - functions
  - channels
- Interface Types
  - interface

Basic 與 Aggregate Type 對應到 Java 是 Value Type，Reference Type 對應至 Java 的 Reference Type。

## 1. Zero Value (ex03_01)

每一種資料型別在宣告時，沒有給定值的話，則 Go 會給予一個初始值，這個初始值則稱為該型別的 **zero value**。

- int: __`0`__
- float: __`0.0`__
- string: __`""`__
- boolean: __`false`__
- struct: 依每一個 field 的資料型別，給定對應的 zero value.
- array: 指定長度，內含 zero value.
- reference type: __`nil`__

```go {.line-numbers}
package main

import "fmt"

// Runable ...
type Runable interface {
    Run()
}

func main() {
    // value type
    var a int
    var b float64
    var c string
    var d bool
    var e [3]int
    var f struct {
        A int
        B string
    }

    // reference type
    var g *int
    var h []int
    var i map[string]int
    var j func(int) int
    var k chan int
    var m Runable

    // value type
    fmt.Printf("%v\n", a) // 0
    fmt.Printf("%v\n", b) // 0
    fmt.Printf("%q\n", c) // ""
    fmt.Printf("%v\n", d) // false
    fmt.Printf("%v\n", e) // [0 0 0]
    fmt.Printf("%v\n", f) // {0 }

    // reference type
    fmt.Printf("%v, nil? %v\n", g, g == nil) // <nil>, nil? true
    fmt.Printf("%v, nil? %v\n", h, h == nil) // [], nil? true
    fmt.Printf("%v, nil? %v\n", i, i == nil) // map[], nil? true
    fmt.Printf("%v, nil? %v\n", j, j == nil) // <nil>, nil? true
    fmt.Printf("%v, nil? %v\n", k, k == nil) // <nil>, nil? true
    fmt.Printf("%v, nil? %v\n", m, m == nil) // <nil>, nil? true
}
```

## 2. Numbers

### 2.1 Integers

Go 數字型別跟 C 一樣，有分 signed 及 unsigned，可指定 bit 數。

- int8, uint8
- int16, uint16
- int32, uint32
- int64, uint64
- int, uint: 會依作業系統(32bit, or 64bit)，變成 int32/int64 or uint32/uint64

```go {.line-numbers}
var a int32   // zero value: 0
b := 10       // type: int
```

### 2.2 Float-Point Numbers

- float32: 32 bit 浮點數，C 的 float。
- float64: 64 bit 浮點數，C 的 double。

eg:

```go {.line-numbers}
var f float32 // zero value: 0.0
d := 0.0      // type: float64
```

### 2.3 Complex Numbers (複數)

GO complex 是數學上的複數

- complex64: 由兩個 float32 組成
- complex128: 由兩個 float64 組成

1. 複數宣告

    ```go {.line-numbers}
    x := 1 + 2i // complex128
    y := 3 + 4i // complex128
    ```

1. 使用 `complex` function 宣告， `real` 與 `imag` function

    ```go {.line-numbers}
    var x complex128 = complex(1, 2)    // 1+2i
    var y complex128 = complex(3, 4)    // 3+4i
    fmt.Println(x*y)                    // "(-5+10i)"
    fmt.Println(real(x*y))              // "-5"
    fmt.Println(imag(x*y))              // "10"
    ```

## 3. Booleans

只有 `true` 及 `false`，且不能像 C 用 integers 來當 boolean 使用

```go {.line-numbers}
var b boolean   // zero value: false
ok := true
```

## 4. Strings

Go 的字串處理方式與 Swift 同，但與 Java 不同，在 Go 是屬於 Value type, 而在 Java 是 Class (Reference type)。
Go 的 `string` 有以下特性：

- **Immutable** sequence of bytes: 不可變動的連續 byte 陣列。
- **UTF-8** encoded: 預設是 UTF-8 編碼。

### 4.1 宣告

```go {.line-numbers}
var str string    // zero value: "" (empty string)
str2 := "hello world"
str3 := `hello world`
```

如果字串內有 __`""`__ 時，需配合 __`\`__ 使用。如: `"My name is \"Kigi\"."`，程式碼比較不好閱讀。新的程式語言，大都有各自的設計來解決這個問題。在 Go 可以使用 __`__，如：

```go {.line-numbers}
str := `My name is "Kigi".`
```

### 4.2 取得字串長度

使用內建的 `len` 函數，來取得字串長度。

```go {.line-numbers}
len := len(str2)    // use len() to get length of bytes in string
```

### 4.3 Substring

使用 `str[i:j]` 取得 substring. 會從第 i 個開始，取到第 j-1 個為止。可以省略 i 及 j。

```go {.line-numbers}
s := "hello world"

substr2 := s[0:5]    // hello
substr3 := s[1:]     // 從第 1 個開始，取到最後
substr4 := s[:5]     // 從 0 開始，取到第 4 個
substr5 := s[:]      // 全取
```

### 4.4 字串連接 (Concate)

可以直接使用 `+` 連接字串。

```go {.line-numbers}
a := "hello"
b := " world"
c := a + b    // "hello world"
```

### 4.5 ==strings== Package

Go 內建 `strings` package，內含常用的字串操作。常用 functions:

```go {.line-numbers}
// s 是否有包含 substr
func Contains(s, substr string) bool

// s 是否為 prefix 開頭
func HasPrefix(s, prefix string) bool

// substr 在 s 的位置
func Index(s, substr string) int

// 將多組字串，使用 sep 分隔，組合成一個字串。
func Join(a []string, sep string) string
```

### 4.6 Rune (ex03_02)

Go string 操作單位是 **byte**，但計算字串內有多少個字元時，非常不方便。Go 的 `rune` 是 Unicode 字元為單位，它的底層型別是 **int32** 也就是 4 bytes，可以以 Unicode 字元為單位來操作。

```go {.line-numbers}
package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    const nihongo = "日本語"
    for i := 0; i < len(nihongo); i++ {
        fmt.Printf("%d: %x\n", i, nihongo[i])
    }

    for index, runeValue := range nihongo {
        fmt.Printf("%U starts at byte position %d\n", runeValue, index)
    }

    fmt.Println(utf8.RuneCountInString(nihongo)) // 取 utf8 長度

    bytes1 := []byte(nihongo) // convert to byte slice.
    fmt.Println("bytes: ", bytes1)
    fmt.Println(string(bytes1)) // convert to string from byte slice.
}
```

## 5. 字串與數字轉換 (ex03_03)

使用 `fmt.Sprintf()` 與 `strconv` 這個套件。

1. 數字轉字串
    1. **fmt.Strintf**
    1. **strconv.Itoa**

    ```go {.line-numbers}
    x := 123
    y := fmt.Sprintf("%d", x)
    fmt.Println(y, strconv.Itoa(x)) // "123 123"
    ```

1. 轉換基底

    ```go {.line-numbers}
    fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
    fmt.Println(fmt.Sprintf("%b", x))           // "1111011"
    ```

1. 字串轉數字

    ```go {.line-numbers}
    a, _ := strconv.Atoi("123")             // a is an int
    b, _ := strconv.ParseInt("123", 10, 64) // b is base 10, up to 64 bits
    ```

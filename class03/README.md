# 03 Data Types - Basic Types

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

## Zero Value (ex03_01)

每一種資料型別在宣告時，沒有給定值的話，則 Go 會給予一個初始值，這個初始值則稱為該型別的 **zero value**。

- int: 0
- float: 0.0
- string: ""
- boolean: false
- struct: struct that field with zero value
- array: 指定長度，內含 zero value.
- reference type: **nil**

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

## Numbers

### Integers

跟 C 一樣，有分 signed 及 unsigned，可以指定變數的 bit 數。

- int8, uint8
- int16, uint16
- int32, uint32
- int64, uint64
- int, uint: 會依作業系統(32bit, or 64bit)，變成 int32/int64 or uint32/uint64

```go {.line-numbers}
var a int32   // zero value: 0
b := 10       // type: int
```

### Float-Point Numbers

- float32: C 的 float
- float64: C 的 double

eg:

```go {.line-numbers}
var f float32 // zero value: 0.0
d := 0.0      // type: float64
```

### Complex Numbers

數學上的複數

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

## Booleans

只有 `true` 及 `false`，不用能 integers 來當 boolean 使用

```go {.line-numbers}
var b boolean   // zero value: false
ok := true
```

## Strings

Go 的字串處理方式與 Swift 同，但與 Java 不同，在 Go 是屬於 Value type, 而在 Java 是 Class (Reference type)。
Go 的 `string` 有以下特性：

- **Immutable** sequence of bytes
- **UTF-8** encoded

1. 宣告

    ```go {.line-numbers}
    var str string    // zero value: "" (empty string)
    str2 := "hello world"
    ```

1. get length

    ```go {.line-numbers}
    len := len(str2)    // use len() to get length of bytes in string
    ```

1. substring

    使用 `str[i:j]` 取得 substring. 會從第 i 個開始，取到第 j-1 個為止。可以省略 i 及 j。

    ```go {.line-numbers}
    substr2 := s[0:5]    // hello
    substr3 := s[1:]     // 從第 1 個開始，取到最後
    substr4 := s[:5]     // 從 0 開始，取到第 4 個
    substr5 := s[:]      // 全取
    ```

1. concate

  ```go {.line-numbers}
  a := "hello"
  b := " world"
  c := a + b    // "hello world"
  ```

### ==strings== Package

常用 functions:

```go {.line-numbers}
func Contains(s, substr string) bool
func Count(s, sep string) int
func Fields(s string) []string
func HasPrefix(s, prefix string) bool
func Index(s, sep string) int
func Join(a []string, sep string) string
```

### Rune (ex03_02)

`rune` 是 unicode chacter 的概念，它的底層型別是 **int32** 也就是 4 bytes. 一般 string 操作單位是 **byte**。

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

## Conversions between Strings and Numbers (ex03_03)

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

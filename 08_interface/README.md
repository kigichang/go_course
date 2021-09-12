# 08 Interface


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [08 Interface](#08-interface)
  - [0. 我對 Interface 的心得](#0-我對-interface-的心得)
  - [1. Interface 宣告](#1-interface-宣告)
  - [2. Differentence from Value and Pointer](#2-differentence-from-value-and-pointer)
    - [2.1 Value](#21-value)
    - [2.2 Pointer](#22-pointer)
    - [2.3 Summary](#23-summary)
  - [3 Stringer interface](#3-stringer-interface)
  - [4. Type Assert and Interface Puzzler](#4-type-assert-and-interface-puzzler)
    - [Interface Puzzler](#interface-puzzler)
    - [Type Assertion](#type-assertion)
  - [5. Interface Value (interface{})](#5-interface-value-interface)
  - [6. Summary](#6-summary)

<!-- /code_chunk_output -->

## 0. 我對 Interface 的心得

>Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here.

[From Effective Go](https://golang.org/doc/effective_go.html#interfaces)

Interface 對於初學程式的人，會是個很抽象的概念。Interface 有著以下概念：

1. 限制 (Constraint)
1. 建立關聯
1. 去耦合(Decouple)

可以想像一種情境，進到公園，寵物需要繫上牽繩。繫上牽繩這種行為，是就是一種 Interface。

1. 限制 (Constraint)：凡是要進公園的寵物，都要繫上牽繩。
1. 建立關聯：只有繫上牽繩的寵物都可以進公園。
1. 去耦合(Decouple): 不管是那種寵物(貓，狗等)，有繫上牽繩即可進公園。

對程式來說，只要符合 Interface 定義的行為(Method)，即可建立與模組間的關係，但又不限定是那種資料型別。

當發現程式內的多個 Struct 都有相同的行為時，可以透過設計 Interface，來讓程式更有彈性，也更好維護。

比如，`DB`, `Net` 與 `File` 都有共同的行為 `Read`，礙於 Go 是強型別的程式語言，主程式需會要設計不同 Functions (`ReadDB`, `ReadNet`, `ReadFile`)，來讀取不同資料來源的資料。

```go {.line-numbers}
type DB struct{}

func (db *DB) Read() {
}

type Net struct{}

func (net *Net) Read() {
}

type File struct{}

func (f *File) Read() {
}

func ReadDB(db *DB) {
	db.Read()
}

func ReadNet(net *Net) {
	net.Read()
}

func ReadFile(f *File) {
	f.Read()
}
```

我們可以透過 Interface 的機制，定義 `Read` 這種行為。因為`DB`, `Net` 與 `File` 都有 `Read` 行為，就符合 `Reader` 這個 interface。主程式只要設計一個 `func Read(r Reader)` 的 function 即可。

```go {.line-numbers}
type Reader interface {
	Read()
}

type DB struct{}

func (db *DB) Read() {
}

type Net struct{}

func (net *Net) Read() {
}

type File struct{}

func (f *File) Read() {
}

type Integer int

func (i Integer) Read() {
}

func Read(r Reader) {
	r.Read()
}
```

由上的行為，可以再反思一開始概念：

1. 限制 (Constraint): 想要符合 `Reader`，就需要實作 `func Read()`。
1. 建立關聯: 需要變成 `Reader` 才能被主程式使用。
1. 去耦合(Decouple): 主程式可以不用理會是何種資料型別，只要是 `Reader`，主程式就可以使用。

## 1. Interface 宣告

```go {.line-numbers}
type Name interface {
    FuncName(ParameterName DataType) DataType
}
```

```go {.line-numbers}
type Chaincode interface {
    Init(stub ChaincodeStubInterface) pb.Response
    Invoke(stub ChaincodeStubInterface) pb.Response
}
```

在 interface 內，也可以再宣告另一個 interface。

```go {.line-numbers}
type ReadCloser interface {
    Reader
    Closer
}
```

Go 的 `io.ReadCloser` 是由 `io.Reader` 與 `io.Closer` 組成。

## 2. Differentence from Value and Pointer

Interface 與 Java 類似，用 struct 的 method 來實作 interface 指定的 method。這邊需注意的是，在實作 interface 的 method 時，是用 Value 還是 Pointer.

### 2.1 Value

使用 Value 方式，實作 interface 的 function。

```go {.line-numbers}
package main

import "fmt"

// Scale ...
type Scale interface {
    ScaleBy(float64)
}

// Point ...
type Point struct {
    X float64
    Y float64
}

// ScaleBy ...
func (p Point) ScaleBy(a float64) {
    p.X *= a
    p.Y *= a
}

// CallScale ...
func CallScale(s Scale, a float64) {
    s.ScaleBy(a)
}

func main() {
    p := Point{100.0, 200.0}

    fmt.Println(p) // {100 200}
    p.ScaleBy(10)
    fmt.Println(p) // {100 200}
    CallScale(p, 10)
    fmt.Println(p) // {100 200}
    CallScale(&p, 10)
    fmt.Println(p) // {100 200}
}
```

### 2.2 Pointer

```go {.line-numbers}
package main

import "fmt"

// Scale ...
type Scale interface {
    ScaleBy(float64)
}

// Point ...
type Point struct {
    X float64
    Y float64
}

// ScaleBy ...
func (p *Point) ScaleBy(a float64) {
    p.X *= a
    p.Y *= a
}

// CallScale ...
func CallScale(s Scale, a float64) {
    s.ScaleBy(a)
}

func main() {
    p := Point{100.0, 200.0}

    fmt.Println(p) // {100 200}
    p.ScaleBy(10)
    fmt.Println(p) // {1000 2000}
    CallScale(&p, 10)
    fmt.Println(p) // {10000 20000}
    CallScale(p, 10) // cannot use p (type Point) as type Scale in argument to CallScale: Point does not implement Scale (ScaleBy method has pointer receiver)
}
```

### 2.3 Summary

如果是用 Value 的方式，來實作 Interface 的 Methods 時，可以使用 Value 或 Pointer 方式，當作 interface，傳入 Function. __2.1__ 中的 CallSlice 可以是 Value or Pointer，但 __2.2__ 只能是 Pointer.

>The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.
[From Effective Go](https://golang.org/doc/effective_go.html#pointers_vs_values)

## 3 Stringer interface

**Stringer** interface 有一個 `String()`，功能類似 Java Object 的 **toString**. 可以自定義輸出的字串格式。可用在 `fmt.Println`, `fmt.Printf("%v")` 或 `fmt.Printf("%s")` 等。在 debug 時，非常好用。

```go {.line-numbers}
type Stringer interface {
    String() string
}
```

```go {.line-numbers}
type Point struct {
    X float64
    Y float64
}

func (p Point) String() string {
    return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}

p := Point{ 100.0, 200.0 }
fmt.Println(p)              // (100.000000, 200.000000)
```

## 4. Type Assert and Interface Puzzler

Go 的 Interface 會存原生的資料型別(是由那種資料型別實作)，與值，Go 的 interface 如果是 nil, 必須兩者都沒有被設定。

>1. Interfaces are implemented as two elements, a type T and a value V.
>1. An interface value is nil only if the V and T are both unset, (T=nil, V is not set).

[From FAQ of Go](https://golang.org/doc/faq#nil_error)

```go {.line-numbers}
package main

import (
    "fmt"
    "io"
    "os"
)

// MyFile ...
type MyFile struct{}

// Close implements io.Closer interface
func (f *MyFile) Close() error {
    return nil
}

// BadNew1 ...
func BadNew1() *MyFile {
    return nil
}

// BadNew2 ...
func BadNew2() io.Closer {
    var f *MyFile
    return f
}

func main() {
    var f1 io.Closer = BadNew1()
    f2 := BadNew2()

    fmt.Printf("%T, %v, %v\n", f1, f1, f1 == nil) // *main.MyFile, <nil>, false
    fmt.Printf("%T, %v, %v\n", f2, f2, f2 == nil) // *main.MyFile, <nil>, false

    switch v := f1.(type) {
    case *MyFile:
        fmt.Printf("%T, %v, %v\n", v, v, v == nil) // *main.MyFile, <nil>, true
    default:
        fmt.Println("unknown type")
    }

    value, ok := f2.(*MyFile)
    fmt.Printf("%v, %T, %v, %v\n", ok, value, value, value == nil) // true, *main.MyFile, <nil>, true

    f3, ok := f2.(*os.File)
    fmt.Printf("%v, %T, %v, %v\n", ok, f3, f3, f3 == nil) // false, *os.File, <nil>, true
}
```

### Interface Puzzler

1. `BadNew1` 回傳值，有指定資料型別為 `*MyFile`，即使回傳值是 `nil`，但 `f1` 會記錄 `*MyFile`，因此 `f1` 會不是 `nil`。應改為 `func BadNew1() io.Closer`
1. `BadNew2` 在 Function 內的 `f` 已指定資料型別為 `*MyFile`，即使值是 `nil`，回傳成 `io.Closer` 後，會記錄 `*MyFile`，依然不會是 `nil`。應改為 `var f io.Closer`

### Type Assertion

1. 要判斷 interface 是否為某種資料型別，可以用：
    ```go {.line-numbers}
    switch v := f1.(type) {
        case *MyFile:
        // v 會是該資料型別的值。
    }
    ```
    或
    ```go {.line-numbers}
    switch f1.(type) {
        case *MyFile:
    }
    ```

1. 要從 interface 轉成原生的資料型別，建議使用 `value, ok := f2.(*MyFile)`，不要使用 `value := f2.(*MyFile)`，原因是：如果該 interface 不是此種資料型別時，會發生 panic，中斷程式。
1. Type assertion 更多的實例，可參考 [spf13/cast](https://github.com/spf13/cast)。

## 5. Interface Value (interface{})

Go 有設計一種特別的資料型別 Interface Value (`interface{}`)，可以包含所有 Go 的資料型別。有點像 php 或 python 不管變數的資料型別 (弱型別)。

Interface Value 也可以視作一種完全沒有任何條件限制的 Interface。Interface Value 通常會跟 Type Assertion 與 Reflection 使用。如：內建的 `json`。 

An interface value (**interface{}**) can hold arbitrarily large dynamic values

```go {.line-numbers}
var any interface{}
fmt.Printf("%T\n", any)     // <nil>

a := 10
b := 100.0
c := "string"
d := struct {A, B string}{"foo", "boo"}
e := []string{}
f := map[string]int{}

any = a
fmt.Printf("%T\n", any)     // int

any = &a
fmt.Printf("%T\n", any)     // *int

any = b
fmt.Printf("%T\n", any)     // float64

any = &b
fmt.Printf("%T\n", any)     // *float64

any = c
fmt.Printf("%T\n", any)     // string

any = &c
fmt.Printf("%T\n", any)     // *string

any = d
fmt.Printf("%T\n", any)     // struct { A string; B string }

any = &d
fmt.Printf("%T\n", any)     // *struct { A string; B string }

any = e
fmt.Printf("%T\n", any)     // []string

any = &e
fmt.Printf("%T\n", any)     // *[]string

any = f
fmt.Printf("%T\n", any)     // map[string]int

any = &f
fmt.Printf("%T\n", any)     // *map[string]int
```

## 6. Summary

Go interface 設計上是以 decouple 的概念，將 interface 與 struct 的關係脫勾，與 Java 不太一樣的地方。在 Java 中，class 需要指定 implements 某個 interface，但在 go 則不用，只要 struct 有符合某個 interface 的定義，即成為該 interface；也因此我們也可以事後再定義 interface 來綁定 struct 的關係。

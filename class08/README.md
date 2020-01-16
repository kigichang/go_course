# 08 Interface

>Interfaces in Go provide a way to specify the behavior of an object: if something can do this, then it can be used here.

[From Effective Go](https://golang.org/doc/effective_go.html#interfaces)

## Interface 宣告

```go {.line-numbers}
type Name interface {
    FuncName(ParameterName DataType) DataType
}
```

eg:

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

## Differentence from Struct Value and Pointer

Interface 與 Java 類似，用 struct 的 method 來實作 interface 指定的 method。這邊需注意的是，在實作 interface 的 method 時，是用 value 還是 pointer.

### Value Version

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

### Pointer Version

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

## Stringer interface

**Stringer** interface 有一個 `String()`，功能類似 Java Object 的 **toString**.

```go {.line-numbers}
type Stringer interface {
    String() string
}
```

eg:

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

## Interface Puzzler

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

## Interface Value (interface{})

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

## Summary

Go interface 設計上是以 decouple 的概念，將 interface 與 struct 的關係脫勾，與 Java 不太一樣。在 Java 中，class 需要指定 implements 某個 interface，但在 go 則不用，只要 struct 有符合某個 interface 的定義，即成為該 interface；也因此我們也可以事後再定義 interface 來綁定 struct 的關係。

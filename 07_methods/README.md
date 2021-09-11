# 07 Methods


<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [07 Methods](#07-methods)
  - [0. 前言](#0-前言)
  - [1. Struct Receiver](#1-struct-receiver)
  - [2. Struct Pointer Receiver](#2-struct-pointer-receiver)
  - [3. Methods in Value and Pointer](#3-methods-in-value-and-pointer)
    - [3.1 Struct](#31-struct)
    - [3.2 Pointer](#32-pointer)
    - [3.3 注意](#33-注意)
  - [4. Method Signature](#4-method-signature)
  - [5. Side Effects in Functional Language](#5-side-effects-in-functional-language)

<!-- /code_chunk_output -->

## 0. 前言

在 OOP 中，定義在 Class 內的 Function，稱為 Methods。Class 的 Method 來處理資料，達到封裝的功能。在 Go 也有一樣的功能，主要是針對 struct 來定義 method。Go 定義 Method 時，(Method 的 receiver) 可以使用 Struct 或者是 Struct Pointer。

## 1. Struct Receiver

```go {.line-numbers}
package main

import (
    "fmt"
    "math"
)

// Point ...
type Point struct{ X, Y float64 }

// Distance ...
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance ...
func (p Point) Distance(q Point) float64 {
    return Distance(p, q)
}

func main() {
    p := Point{1, 2}
    q := Point{4, 6}
    fmt.Println(Distance(p, q)) // "5"
    fmt.Println(p.Distance(q))  // "5"
}
```

## 2. Struct Pointer Receiver

在定義 method 也可以使用 struct pointer。

```go {.line-numbers}
package main

import (
    "fmt"
    "math"
)

// Point ...
type Point struct{ X, Y float64 }

// Distance ...
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance ...
func (p *Point) Distance(q Point) float64 {
    return Distance(*p, q)
}

func main() {
    p := Point{1, 2}
    q := Point{4, 6}
    fmt.Println(Distance(p, q)) // "5"
    fmt.Println(p.Distance(q))  // "5"

    x := &p
    fmt.Println(Distance(*x, q)) // "5"
    fmt.Println(x.Distance(q))   // "5"

}
```

## 3. Methods in Value and Pointer

由上面的例子來看，不論 methods 是被宣告在 struct 或者是 struct pointer，只要是該 struct 或者是 struct pointer，都可以呼叫。而這兩者差別是，用 struct pointer 定義 method 要特別注意會修改到原本的值。

### 3.1 Struct

使用 Method 時，不會修改到原本 Struct 值。

```go {.line-numbers}
package main

import "fmt"

// Point ...
type Point struct{ X, Y float64 }

// ScaleBy ...
func (p Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

func main() {
    p := Point{1, 2}
    p.ScaleBy(10)
    fmt.Println(p) // {1 2}

    q := &Point{1, 2}
    q.ScaleBy(10)  // implicit (*q)
    fmt.Println(q) // &{1 2}

    (&Point{3, 4}).ScaleBy(100) // an pointer can be the receiver for value method.
}
```

### 3.2 Pointer

使用 Method 時，會修改到原本 Struct 值。(原因：傳 Pointer 至 Method)

```go {.line-numbers}
package main

import "fmt"

// Point ...
type Point struct{ X, Y float64 }

// ScaleBy ...
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

func main() {
    p := Point{1, 2}
    p.ScaleBy(10)  // implicit (&p)
    fmt.Println(p) // {10 20}

    q := &Point{1, 2}
    q.ScaleBy(10)
    fmt.Println(q) // &{10 20}

    // value can not be the receiver of pointer method.
    Point{3, 4}.ScaleBy(20) // compile error: cannot call pointer method on Point literal.
}
```

### 3.3 注意

與 slice 類似，但因為是 method 很難查覺是否有修改原本的資料。因此在實作上，儘量 method 都用 pointer 的方式。

1. 避免 pass by value 的記憶體浪費
1. 避免 golang 在 struct pointer 語法上的 puzzle (因為 struct 與 struct pointer 在 call method 的語法都一樣，不像 C 有分 `.` 與 `->`).


## 4. Method Signature

Method 本身就是 funcation，因此也有 signature.

```go {.line-numbers}
package main

import "fmt"

// Point ...
type Point struct{ X, Y float64 }

// ScaleBy ...
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor
}

func main() {
    p := Point{1, 2}
    q := Point{3, 4}

    p.ScaleBy(10)
    fmt.Println(p) // {10 20}

    x := p.ScaleBy
    fmt.Printf("%T\n", x) // func(float64)

    x(100)
    fmt.Println(p) // {1000 2000}

    x = q.ScaleBy
    x(20)
    fmt.Println(p) // {1000 2000}
    fmt.Println(q) // {60 80}
}
```

## 5. Side Effects in Functional Language

程式的函式有以下的行為時，就會稱該函式有 **Side Effects**。

* Reassigning a variable (val v.s. var)
* Modify a data structure in place (mutable v.s. immutable)
* Setting a field on an object (change object state)
  * 這裏指的 object 是 OOP 的 Object 不是 Scala 的 object (singleton)
  * OOP 修改物件的值，在程式語言的術語：改變物件的狀態，如上說的 **changing-state**
* Throwing an exception or halting with error
* Printing to the console or reading user input (I/O)
* Reading from or write to a file (I/O)
* Drawing on the screen (I/O)

截自 [Functional Language in Scala](http://www.amazon.com/Functional-Programming-Scala-Paul-Chiusano/dp/1617290653)

~~就實際狀況來說，我們寫程式不可能不去碰 I/O。~~

# 07 Methods

在 OOP 中，會定義 Class 的 Method 來處理資料。在 Go 也有一樣的功能，主要是針對 struct 來定義 method.

## Declaring Method for Struct

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

## Declaring Method for Struct Pointer

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

## Methods in Value and Pointer

由上面的例子來看，不論 methods 是被宣告在 struct 或者是 struct pointer，只要是該 struct 或者是 struct pointer，都可以呼叫。而這兩者差別是，用 struct pointer 定義 method 要特別注意會修改到原本的值。

### Value

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

### Pointer

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

### 注意

與 slice 類似，但因為是 method 很難查覺是否有修改原本的資料。因此在實作上，儘量 method 都用 pointer 的方式。

1. 避免 pass by value 的記憶體浪費
1. 避免 golang 在 struct pointer 語法上的 puzzle (因為 struct 與 struct pointer 在 call method 的語法都一樣，不像 C 有分 `.` 與 `->`).

## Summary

>The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values, but pointer methods can only be invoked on pointers.

[From Effective Go](https://golang.org/doc/effective_go.html#pointers_vs_values)

|                | Pointer | Value |
|:--------------:|:-------:|:-----:|
| Pointer Method | O       | X     |
| Value Method   | O       | O     |

## Method Signature

method 本身就是 funcation，因此也有 signature.

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

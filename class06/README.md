# 06 Functions

## 宣告

```go {.line-numbers}
func name(parameter-list) (result-list) {
    body
}
```

```go {.line-numbers}
func hypot(x float64, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}

fmt.Println(hypot(3, 4)) // "5"
```

### grouping 相同型別

```go {.line-numbers}
func f(i int, j int, k int, s string, t string) { /* ... */ } // original
func f(i, j, k int, s, t string)                { /* ... */ } // simplify
```

### 回傳值

```go {.line-numbers}
func add(x int, y int) int { return x+y }
func sub(x, y int) (z int) { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }
```

#### Return Tuple

Go 的 function 可以一次回傳多個值 (tuple)

```go {.line-numbers}
func swap(x, y int) (int, int) {
    return y, x
}

a, b := 1, 2        // a = 1, b = 2
a, b = swap(a, b)   // a = 2, b = 1
```

### Variadic Functions

function 的參數個數是不固定的。eg:

1. 宣告

    ```go {.line-numbers}
    func sum(vals ...int) int {
        total := 0
        for _, val := range vals {
            total += val
        }
        return total
    }

    fmt.Println(sum())           //  "0"
    fmt.Println(sum(3))          //  "3"
    fmt.Println(sum(1, 2, 3, 4)) //  "10"
    ```

1. 如何將 slice 傳入:

    ```go {.line-numbers}
    values := []int{1, 2, 3, 4}
    fmt.Println(sum(values...)) // "10"
    ```

### 空白 Body

可以定義 function 但沒有 body。通常是用另一種程式語言來實作，比如 C。越是底層的工作越容易看到這樣子的做法。

## Recursion 遞迴

```go {.line-numbers}
func gcd(a, b int) int {
    if b == 0 {
        return a
    }

    return gcd(b, a % b)
}

fmt.Println(gcd(24, 128)) // 8
```

## Pass by Value (Call by Value)

Go 在傳遞參數時，是以 **by value** 的方式進行，也就是說在傳入 function 前，會產生一份新的資料，給 function 使用，也因此 function 修改時，也是修改此新的資料。

此時要特別注意傳入的資料型別：

- Aggregate Types (Array, Struct)，在 Java 的定義下，是屬於 Value Types，也就是說會產生一筆新的資料給 function，function 做任何修改，都**不會**異動到原本的資料，如果 array/struct 資料很龐大時，會造成記憶體的浪費。
- Reference Types (Pointer, Slice, Map, Function, Channel)，一樣在傳入 function 時，會複製新的值給 function，只是這新的值，只是 copy 原本的參照值(reference, 可以當作記憶體位址)，因此 function 做任何修改時，也都是透過原來的參照值在做資料異動，會修改到原本的資料，要特別小心。

### Pass by Value with Struct and Struct Pointer

```go {.line-numbers}
package main

import "fmt"

// Person ...
type Person struct {
    Age  int
    Name string
}

func test(p Person) {
    p.Age++
    p.Name += " by test"
}

func testByPtr(p *Person) {
    p.Age++
    p.Name += " by test"
}

func main() {
    p := Person{
        Age:  5,
        Name: "Test",
    }

    fmt.Println(p) // {0 Test}
    test(p)        // 用原本的 struct
    fmt.Println(p) // {0 Test}

    testByPtr(&p)  // 改用 pointer
    fmt.Println(p) // {1 Test by test}
}
```

### Pass by Value with Aray and Slice

```go {.line-numbers}
package main

import "fmt"

func arrTest(a [3]int) {
    for i, x := range a {
        a[i] = x + 1
    }
}

func arrTestBySlice(a []int) {
    for i, x := range a {
        a[i] = x + 1
    }
}

func main() {
    a := [3]int{1, 2, 3}

    fmt.Println(a) // [1 2 3]
    arrTest(a)     // 用原本的 array
    fmt.Println(a) // [1 2 3]

    arrTestBySlice(a[:]) // 改用 Slice
    fmt.Println(a)       // [2 3 4]
}
```

## Signature

一個 function 的型別，通常也稱做 **signature**。兩個 function 有相同的 signature，需滿足以下兩個條件：

1. 參數 (parameters) 資料型別與順序相同，與參數名稱無關。
1. 回傳的值的資料型別與順序相同

eg:

```go {.line-numbers}
func add(x int, y int) int { return x+y }
func sub(x, y int) (z int) { z= x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int { return 0 }

fmt.Printf("%T\n", add)   // "func(int, int) int"
fmt.Printf("%T\n", sub)   // "func(int, int) int"
fmt.Printf("%T\n", first) // "func(int, int) int"
fmt.Printf("%T\n", zero)  // "func(int, int) int"
```

在 Go 的 function 也可以當作參數與回傳值。也因此 Go 也算是一種 first-class lanaugage.

### First-Class

function 也是一種資料型別，可以當作變數，或當作另一個 function 的參數及回傳值。
以 Go 來說，**signature** 是 Function 的資料型別。當宣告 funcation 沒有指定 name 時，則稱為 **anonymous function**

#### Assignment

```go {.line-numbers}
func square(n int) int { return n * n }
func negative(n int) int { return -n }
func product(m, n int) int { return m * n }

var f func(int) int     // signature
fmt.Printf("%T\n", f)   // "func(int) int"

f = square
fmt.Println(f(3))       // "9"

f = negative
fmt.Println(f(3))       // "-3"

f = product // cannot use product (type func(int, int) int) as type func(int) int in assignment
```

#### As Parameter and Return

```go {.line-numbers}
func square(n int) int { return n * n }
func negative(n int) int { return -n }

func compose(f, g func(int) int) func(int) int {
    return func(a int) int {        // anonymous function
        return g(f(a))
    }
}

k1 := compose(square, negative)
fmt.Printf("%T\n", k1)              // func(int) int
fmt.Println(k1(10))                 // -100 negative(square(10))

k2 := compose(negative, square)
fmt.Printf("%T\n", k2)              // func(int) int
fmt.Println(k2(10))                 // 100 square(negative(10))
```

# 04 Data Types: Aggregate Types

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [04 Data Types: Aggregate Types](#04-data-types-aggregate-types)
  - [0. 前言](#0-前言)
  - [1. Arrays](#1-arrays)
    - [1.1 Array Declaration](#11-array-declaration)
    - [1.2 Array Travel](#12-array-travel)
    - [1.3 Array Compare](#13-array-compare)
  - [2. Struct](#2-struct)
    - [2.1 Struct Declaration](#21-struct-declaration)
    - [2.2 Struct Pointer](#22-struct-pointer)
    - [2.3 Struct Compare](#23-struct-compare)
    - [2.4 Slice of Stuct Puzzle](#24-slice-of-stuct-puzzle)
  - [3. JSON](#3-json)

<!-- /code_chunk_output -->

## 0. 前言

Aggregate types 主要有兩種：

- arrays
- structs

它們的記憶體管理方式，都是：

1. 連續的記憶體空間
1. 固定長度

原文：

>Arrays and structs are aggregate types; their values are **concatenations** of other values in memory. Arrays are homogeneous—their elements all have the same type—whereas structs are heterogeneous. Both arrays and structs are **fixed size**.[^book1]

[^book1]: from [The Go Programming Language](https://www.amazon.com/Programming-Language-Addison-Wesley-Professional-Computing-ebook/dp/B0184N7WWS)

注意：

1. pass by value, 會 copy 原本資料，因此要小心記憶體的問題。可改用 slice 及 pointer。

## 1. Arrays

### 1.1 Array Declaration

宣告方式：

1. 指定長度
1. 使用 `...`， Go compiler 會依據內含的元素個數，給予長度。

```go {.line-numbers}
var a [3]int // [0,0,0]
var b = [3]int{1, 2, 3}
var r = [3]int{1, 2} // [1, 2, 0]
q := [...]int{1, 2, 3} // [1, 2, 3]
x := [...]int{5: -1} // [0, 0, 0, 0, 0, -1]
```

### 1.2 Array Travel

可以使用 `range`，來依序取得 array 內的值。

```go {.line-numbers}
x := [...]int{1, 2, 3}
// Print elements via indices.
for i := range x {
    fmt.Println(x[i])
}

// Print the indices and elements.
for i, v := range x {
    fmt.Printf("%d %d\n", i, v)
}

// Print the elements only.
for _, v := range x {
    fmt.Printf("%d\n", v)
}
```

### 1.3 Array Compare

如果 array 的長度相同，且 array 內，值的資料型別是可以比較的 (如: int, bool)，array 才可以比較。

>If an array’s element type is **comparable** then the array type is comparable too, so we may directly compare two arrays of that type using the == operator, which reports whether all corresponding elements are equal. The != operator is its negation.

如：

```go {.line-numbers}
a := [2]int{1, 2}
b := [...]int{1, 2}
c := [2]int{1, 3}
fmt.Println(a == b, a == c, b == c) // "true false false"

d := [3]int{1, 2}
fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
```

## 2. Struct

跟 C 的 struct 用法一樣。struct 可以組合多個不同型別的資料，每一個資料欄位，稱作 field.

### 2.1 Struct Declaration

```go {.line-numbers}
type Employee struct {
    ID        int
    Name      string
    Address   string
    DoB       time.Time
    Position  string
    Salary    int
    ManagerID int
}

var empty Employee // empty Employy struct

dilbert := Employee{
    ID:       1,
    Name:     "Dilbert",
    Position: "Engineer",
    Salary:   5000,
}
dilbert.Salary -= 5000 // demoted, for writing too few lines of code
position := &dilbert.Position
*position = "Senior " + *position // promoted, for outsourcing to Elbonia
```

與 C 不同的是，一個空白的 struct，會依照每一個 field 的資料型別，自動帶入該型別的 **zero value**。如：int 則為 0, string 則為 ""

### 2.2 Struct Pointer

與 C 一樣，struct 通常會撘配 Pointer 來處理。與 C 不同是操作語法。

- 在 C 中，如果是實例 (Instance)，則用 **`.`** 來操作, eg: **`x.A`**，如果是指標 (pointer)，則用 **`->`**, eg: **`x->A`**。
- 在 Go 則都用 **`.`** 來操作，也因此要小心是在用 Instance 還是 Pointer

```go {.line-numbers}
alice := &Employee{
    ID:   2,
    Name: "Alice",
}
fmt.Println("alice:", alice) // alice: &{2 Alice  0001-01-01 00:00:00 +0000 UTC  0 0}

fmt.Println(alice.ID, alice.Name)
```

### 2.3 Struct Compare

兩個 struct 可以比較的前題是所有的 field 的資料型別都是可以比較的。

>If all the fields of a struct are **comparable**, the struct itself is comparable, so two expressions of that type may be compared using == or !=. The == operation compares the corresponding fields of the two structs in order, so the two printed expressions below are equivalent:

如：

```go {.line-numbers}
type Point struct{ X, Y int }
p := Point{1, 2}
q := Point{2, 1}
fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
fmt.Println(p == q)                   // "false"
```

### 2.4 Slice of Stuct Puzzle

當 Slice 內的元素是 Struct 時，在使用 `for-range` 操作資料時，要特別留意操作的 struct 物件，並非 Slice 內的物件。如下：

@import "ex04_05/main.go" {class=line-numbers}

結果：

```text {.line-numbers}
before: [{0 0} {0 0} {0 0} {0 0} {0 0}]
0: 0xc000118000, 0xc00010c0c0
1: 0xc000118010, 0xc00010c0c0
2: 0xc000118020, 0xc00010c0c0
3: 0xc000118030, 0xc00010c0c0
4: 0xc000118040, 0xc00010c0c0
after: [{0 0} {0 0} {0 0} {0 0} {0 0}]
```

1. 在 `for i, pt := range points`，會 clone slice 內的值，因此會看到 `pt` 的記憶體位址，並非 `points[i]` 的位址。
1. `pt` 等同是一個新的 struct，因此任何修改的操作，都不的影響到原本 slice 內的值。

## 3. JSON

Go 有內建處理 JSON 的套件 `encoding/json`，會搭配 Struct Tag 來使用 JSON.

eg:

```go {.line-numbers}
type Movie struct {
    Title  string
    Year   int  `json:"released"`
    Color  bool `json:"color,omitempty"`
    Actors []string
}

var movies = []Movie{
    {
        Title:  "Casablanca",
        Year:   1942,
        Color:  false,
        Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
    },
    {
        Title:  "Cool Hand Luke",
        Year:   1967,
        Color:  true,
        Actors: []string{"Paul Newman"},
    },
    {
        Title:  "Bullitt",
        Year:   1968,
        Color:  true,
        Actors: []string{"Steve McQueen", "Jacqueline Bisset"},
    },
}

data, err := json.Marshal(movies)
if err != nil {
    log.Fatalf("JSON marshaling failed: %s", err)
}

fmt.Println("movies: ", string(data))

var titles []struct{ Title string }
if err := json.Unmarshal(data, &titles); err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
}
fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"

var movie1 Movie
err = json.Unmarshal([]byte(
    `{
        "Title": "Casablanca",
        "released":1942,
        "Actors":["Humphrey Bogart","Ingrid Bergman"]}
    `), &movie1)

if err != nil {
    log.Fatalf("JSON unmarshaling failed: %s", err)
}

fmt.Println("movie1", movie1)
```

1. 上例中，`Movie` 內的  ``` `json:"released"` ``` 即是 Struct Tag.
1. 在宣告 struct 時，加入 Struct Tag，如: `json:"color,omitempty"`, `color` 是指 json 的欄位名稱，`omitempty` 是指如果該欄位值是 **zero value**，則不輸出 json。
1. 如果沒有寫 json 的 annotation, 則直接用變數名稱。建議還是都加 Struct Tag 重新定義。
1. 使用 json.Umarshal 來取得資料。
1. 使用 json.Marshal 來輸出 json 資料。

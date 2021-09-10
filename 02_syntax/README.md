# 02 程式結構與語法

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [02 程式結構與語法](#02-程式結構與語法)
  - [1. Build-in](#1-build-in)
    - [1.1 關鍵字](#11-關鍵字)
    - [1.2 內建常數](#12-內建常數)
    - [1.3 資料型別](#13-資料型別)
    - [1.4 內建 Function](#14-內建-function)
  - [2. 宣告 Declaraion](#2-宣告-declaraion)
    - [2.1 完整寫法](#21-完整寫法)
    - [2.2 簡寫](#22-簡寫)
    - [2.3 Default Type](#23-default-type)
    - [2.4 注意事項](#24-注意事項)
  - [3. 數值宣告](#3-數值宣告)
  - [4. Constants](#4-constants)
  - [5. iota (ex02_01)](#5-iota-ex02_01)
  - [6. 指標 (Pointer)](#6-指標-pointer)
  - [7. 元組 (Tuple)](#7-元組-tuple)
  - [8. ==type== Keyword](#8-type-keyword)
    - [8.1 Type Declaration (ex02_02)](#81-type-declaration-ex02_02)
      - [華氏、攝氏型別宣告與轉換](#華氏-攝氏型別宣告與轉換)
    - [8.2 Type Alias (ex02_03)](#82-type-alias-ex02_03)
    - [8.3 Type Declaration & Alias 差別](#83-type-declaration-alias-差別)
  - [9. Package and Imports](#9-package-and-imports)
    - [9.1 Package Initialization (ex02_04)](#91-package-initialization-ex02_04)
    - [9.2 目錄結構](#92-目錄結構)
      - [程式碼](#程式碼)
  - [10. Go Wildcard](#10-go-wildcard)
  - [11. 變數 Visibiility](#11-變數-visibiility)
  - [12. 變數 Scope](#12-變數-scope)

<!-- /code_chunk_output -->

## 1. Build-in
### 1.1 關鍵字

```go
break    default     func   interface select
case     defer       go     map       struct
chan     else        goto   package   switch
const    fallthrough if     range     type
continue for         import return    var
```

### 1.2 內建常數

```go
true  false  iota  nil
```

### 1.3 資料型別

```go
int  int8  int16  int32  int64 uint  uint8  uint16  uint32  uint64 uintptr
float32  float64  complex64  complex128
bool  byte  rune  string  error
```

### 1.4 內建 Function

```go
make  len  cap  new  append  copy delete close
complex  real  imag
panic  recover
```

## 2. 宣告 Declaraion

### 2.1 完整寫法

```go
var name type = expression
```

- 不給初始值時，可以省略 `= expression`

	```go
	var name type
	```

- 給初始值時，則可以省略 ==type==

	```go
	var name = expression
	```

- 使用方式如下：

	1. 宣告變數，不給初始值

		```go {.line-numbers}
		var s string
		fmt.Println(s) // ""
		```

	1. 一次宣告多個變數，並給初始值(可省略型別)

		```go {.line-numbers}
		var i, j, k int                 // int, int, int
		var b, f, s = true, 2.3, "four" // bool, float64, string
		```

	1. 接收 return 值，可 return 多組值。

		```go {.line-numbers}
		var f, err = os.Open(name) // os.Open returns a file and an error
		```

### 2.2 簡寫

- 宣告時，省略 `var`

	```go {.line-numbers}
	name := expression
	```

- 使用方式如下：

	1. 省略型別宣告

		```go {.line-numbers}
		i, j := 0, 1
		```

	1. 接收 return 值

		```go {.line-numbers}
		anim := gif.GIF{LoopCount: nframes}
		freq := rand.Float64() * 3.0
		t := 0.0

		f, err := os.Open(name)
		if err != nil {
			return err
		}

		// ...use f...

		f.Close()
		```

### 2.3 Default Type

當省略型別時，

- 整數型別預設是 __int__
- 浮點數預設是 __float64__

```go {.line-numbers}
i := 0      // int
f := 0.0    // float64
s := ""     // string
```

### 2.4 注意事項

使用 `:=` 時，左邊的變數名稱，至少要有一個是新的。

1. 至少要有一個是新的變數名稱

	```go {.line-numbers}
	in, err := os.Open(infile)
	// ...
	out, err := os.Create(outfile)
	```

	以上，雖然 `err` 重覆，但 `out` 是新的變數名稱，compile 會過。

1. 都是舊的

	```go {.line-numbers}
	f, err := os.Open(infile)
	// ...
	f, err := os.Create(outfile) // compile error: no new variables
	```

	以上，`f` 與 `err` 都是舊的變數，所以在第二次，還是使用 `:=` 時，compile 會錯。通常 compile 會報錯，都不是什麼大問題，修正就好了。

## 3. 數值宣告

以往宣告很大數值時，無法像 excel 每千分位，用 `,` 來區隔。現在 Go 也支援這項功能，可以使用 `_` 來區隔。

```go
var x int64 = 123_456_789
var y float64 = 12_345.678_9
```

## 4. Constants

與 C 相同，利用 `const` 這個關鍵字來宣告常數。

```go {.line-numbers}
const pi = 3.14159 // approximately; math.Pi is a better approximation
```

```go {.line-numbers}
const (
	e = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)
```

## 5. iota (ex02_01)

可以使用 `iota` 來實作 Enumeraion 的功能。

1. 使用在 `const`。
1. 從 0 開始。

[itoa 詳細說明](https://github.com/golang/go/wiki/Iota)

```go {.line-numbers}
package main

import "fmt"

// number
const (
	Zero  = iota    // 0
	One   = iota    // 1
	Two   = iota    // 2
	Three = iota    // 3
)

// file mode
const (
	X = 1 << iota   // 1
	W = 1 << iota   // 2
	R = 1 << iota   // 4
)

// size
const (
	_          = iota // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// weekday
const (
	Sunday = 1 + iota       // iota = 0, sunday     = 1
	_                       // iota = 1, skip
	// this is a comment    // iota = 1, skip
							// iota = 1, skip
	Monday                  // iota = 2, monday     = 3
	Tuesday                 // iota = 3, tuesday    = 4
	Wednesday               // iota = 4, wednesday  = 5
	Thursday                // iota = 5, thursday   = 6
	Friday                  // iota = 6, friday     = 7
	Saturday                // iota = 7, saturday   = 8
)
```

## 6. 指標 (Pointer)

使用方法及觀念與 C 相同，`&` 取變數的指標，`*` 取指標內的值；需注意的是，C 可以對指標做位移，但 Go 不行。[^unsafe]

[^unsafe]: Golang 有 `unsafe` 套件，專門用在與 C 串接。需轉成 unsafe 的 pointer 才能做指標位移。

1. 用法

	```go {.line-numbers}
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2          // equivalent to x = 2
	fmt.Println(x)  // "2"
	```

1. 不可做指標位移:

	```go {.line-numbers}
	a := 10
	b := &a
	b++     // invalid operation: b++ (non-numeric type *int)
	```

## 7. 元組 (Tuple)

近期的程式語言，大多有支援 tuple 功能，早期的則沒有。如果沒有支援 tuple 時，就需要用 class or struct 來封裝回傳。

1. Tuple Assignment (1)

	```go {.line-numbers}
	x, y = y, x
	a[i], a[j] = a[j], a[i]
	```

1. Tuple Assignment (2): GCD sample

	```go {.line-numbers}
	func gcd(x, y int) int {

		for y != 0 {
			x, y = y, x%y
		}

		return x
	}
	```

1. Return Tuple

	```go {.line-numbers}
	func swap(x, y int) (int, int) {
		return y, x
	}
	```

## 8. ==type== Keyword

在 Go 可以使用 `type` 來宣告一個新的 data type，或幫舊的 data type 取一個別名，來增加程式碼的可讀性。

### 8.1 Type Declaration (ex02_02)

可以使用 `type` 來宣告一個新的 data type，通常用在宣告 struct 或 interface。我們也可以使用 `type` 的擴充既有型別的功能。

#### 華氏、攝氏型別宣告與轉換

```go {.line-numbers}
// Celsius ...
type Celsius float64

// ToF convert Celsius to Fahrenheit
func (c Celsius) ToF() Fahrenheit {
	return CToF(c)
}

// Fahrenheit ...
type Fahrenheit float64

// ToC convert Celsius to Fahrenheit
func (f Fahrenheit) ToC() Celsius {
	return FToC(f)
}

// const variable
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
```

1. Strong Type Detection
	雖然 `Celsius` 與 `Fahrenheit` 都是 `float64`，但還是遵守 **Strong Type** 原則，兩者不能直接做運算。如下：

	```go {.line-numbers}
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := BoilingC.ToF()
	fmt.Printf("%g\n", boilingF-FreezingC.ToF()) // "180" °F
	fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
	```

1. 型別轉換(cast)

	```go {.line-numbers}
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // "true"
	fmt.Println(f >= 0)          // "true"
	fmt.Println(c == Celsius(f)) // "true"!
	fmt.Println(c == f)          // compile error: type mismatch
	```

	注意，雖然 `Celsius` 及 `Fahrenheit` 底層都是 `float64`，但都還是要視為不同的型別。

### 8.2 Type Alias (ex02_03)

Go 可以幫 type 取別名 (alias), 宣告的語法：

```go
type type1 = type2
```

如此一來，type1 就直接等於是 type2，可以不用轉型。

```go {.line-numbers}
package main

import "fmt"

// Celsius ...
type Celsius = float64

// ToF convert Celsius to Fahrenheit
func (c Celsius) ToF() Fahrenheit { //  compile error: cannot define new methods on non-local type float64
	return CToF(c)
}

// Fahrenheit ...
type Fahrenheit = float64

// ToC convert Celsius to Fahrenheit
func (f Fahrenheit) ToC() Celsius { //  compile error: cannot define new methods on non-local type float64
	return FToC(f)
}

// const variable
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return c*9/5 + 32 } // do not cast

// FToC convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return (f - 32) * 5 / 9 } // do not cast

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // 100
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // 180
	fmt.Printf("%g\n", boilingF-FreezingC)       // 212

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // true
	fmt.Println(f >= 0)          // true
	fmt.Println(c == Celsius(f)) // true
	fmt.Println(c == f)          // true
}
```

### 8.3 Type Declaration & Alias 差別

**Type Declaration** 與 **Type Alias** 主要的差別：

|                  | 需做轉型 | 擴展功能  |
| ---------------- |:-------:|:-------:|
| Type Declaration | yes     | yes     |
| Type Alias       | no      | no      |

## 9. Package and Imports

Go Package 概念跟 Java package, C/PHP namespace 類似。主要用意：

1. Modularity
1. Encapsulation
1. Reuse

package 命名時要注意，除了 `main` 之外，目錄的名稱建議與 package 相同。

~~import 的路徑，由 `$GOPATH/src` 以下目錄開始。~~

比如有一個 package 為 `gopl.io/ch1/helloworld`, 則開一個 `helloworld` 目錄。完整路徑 `YOUR_PROJECT_PATH/helloworld`，並在該目錄下，執行 `go mod init gopl.io/ch1/helloworld`

import 則用

```go {.line-numbers}
import "gopl.io/ch1/helloworld"
```

package 名稱會是在 `go mod init` 定義。

### 9.1 Package Initialization (ex02_04)

package 中，可以在某一個程式檔案，定義 `func init()`。當 package 被載入時，會先執行 `init` 的程式碼。

### 9.2 目錄結構

```text
.
├── main.go
├── go.mod
└── util
	└── util.go
```

#### 程式碼

- go.mod

	```go
	module ex02_04

	go 1.17
	```

- util.go

	```go {.line-numbers}
	package util

	import "fmt"

	func init() {
		fmt.Println("package util initialize")
	}

	// Hello returns string concates
	func Hello(name string) string {
		return fmt.Sprintf("hello %s", name)
	}
	```

- main.go

	```go {.line-numbers}
	package main

	import (
		"ex02_04/util"
		"fmt"
	)

	func main() {
		fmt.Println("start...")
		fmt.Println(util.Hello("Gopher"))
	}
	```

- 結果

	```text
	package util initialize
	start...
	hello Gopher
	```

## 10. Go Wildcard

Go 的 wildcard 是 `_`，可以用在以下情境：

1. 因為 Go compiler 會檢查沒有使用的變數，如果不想使用該數值時，可以使用 `_` 來取代。

```go {.line-numbers}
_ = test()
```

1. 在宣告函式時，有些參數確定不會被用到，可以在宣告時使用 `_`。

```go {.line-numbers}
func mytest(_ int, str string) {
	
}
```

## 11. 變數 Visibiility

Go 沒有 **private** and **public** 關鍵字，而是利用字母的**大**、**小**寫來區分 **public** 及 **private**。如果變數或 function 是**小寫**開頭，則為 **private**，反之，**大寫**就是 **public**。

注意：**在同 package 下，可以存取 struct 內的 private 變數 (Package Private)。**

## 12. 變數 Scope

與其他語言相同，與 java 不同點是有 global 變數。

變數找尋的順序：local block -> outside block -> global

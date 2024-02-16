# 10 Defer and Error Handling

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [10 Defer and Error Handling](#10-defer-and-error-handling)
  - [1. Deferred (ex10_01)](#1-deferred-ex10_01)
    - [1.1 Defer in Loop](#11-defer-in-loop)
    - [1.2 Defer 與 os.Exit (ex10_03)](#12-defer-與-osexit-ex10_03)
  - [2. Panic](#2-panic)
  - [3. Recover](#3-recover)
  - [4. Errors](#4-errors)
    - [4.1 Errors in Go 1.13](#41-errors-in-go-113)

<!-- /code_chunk_output -->

## 1. Deferred (ex10_01)

在 code block 或 function 結束後，如果有一定會執行的程式碼，可以使用 __defer__。__defer__ 的呼叫順序是 __stack__ 的 LIFO (Last In First Out)，並且利用當下的變數值來執行。

@import "ex10_01/main.go" {class=line-numbers}

結果：

```text {.line-numbers}
before return
double(4) = 8
double of 4 is 8
main end
defer end 2
defer end 1
```

在有關 I/O 處理時，一定會用到。

```go {.line-numbers}
func ReadFile(filename string) ([]byte, error) {
    f, err := os.Open(filename)

    if err != nil {
        return nil, err
    }

    defer f.Close()
    return ReadAll(f)
}
```

### 1.1 Defer in Loop

在 loop 使用 defer 時，要注意：

1. defer 宣告的 function 會在離開 loop 時，才會執行，
1. 變數的 binding 時機點。

@import "ex10_02/main.go" {class=line-numbers}

1. 在 main 的 loop 宣告的 defer 函式，會在 main 結束前執行。
1. 在 a1, a2, a3 的 loop 宣告的 defer function，會在各別 function 結束前執行。
1. a1: 使用當下迴圈 i 的變數值。因此會是 __2 1 0__
~~1. a2: 每次迴圈完成時，會記錄要執行一個 __anonymous function__，當迴圈結束後，則開始執行 defer 記錄的 function，此時 i 的值已經是 __3__。~~ 在 Go 1.22 修正這個問題，之後的結果會與 a1 相同。
1. 與 a2 類似，多傳入當下 i 的值，因此結果與會 a1 相同。

使用 __defer__ 要特別小心被呼叫的時機點與綁定的變數值。

### 1.2 Defer 與 os.Exit (ex10_03)

當使用 `os.Exit` 時，設定的 defer function 並__不會被執行__。因此在宣告 defer 之後，就不要再用 `os.Exit` 中斷程式。

@import "ex10_03/main.go" {class=line-numbers}

## 2. Panic

`panic` 會導致程式中斷。在 Go 的設計中，除非是很嚴重的錯誤，才會使用 __panic__，如像 I/O, 設定檔錯誤等。如是預期到，在撰寫程式時，則儘量檢查並用 __error__ 來處理。

Panic 現象：

@import "ex10_04/main.go" {class=line-numbers}

Output:

```text {.line-numbers}
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.f(0x0)
    /Users/kigi/Data/cyberon/go/src/go_course/class10/ex10-04/main.go:6 +0x16f
main.f(0x1)
    /Users/kigi/Data/cyberon/go/src/go_course/class10/ex10-04/main.go:8 +0x14a
main.f(0x2)
    /Users/kigi/Data/cyberon/go/src/go_course/class10/ex10-04/main.go:8 +0x14a
main.f(0x3)
    /Users/kigi/Data/cyberon/go/src/go_course/class10/ex10-04/main.go:8 +0x14a
main.main()
    /Users/kigi/Data/cyberon/go/src/go_course/class10/ex10-04/main.go:12 +0x2a
exit status 2
```

不應該使用 Panic 的案例，請回傳 __error__:

```go {.line-numbers}
func Reset(x *Buffer) {
    if x == nil {
        panic("x is nil") // unnecessary!
    }
    x.elements = nil
}
```

## 3. Recover

用在取得 panic 發生的原因，通常與 __defer__ 撘配使用，用在 debug 執行時期(runtime)的錯誤。

eg:

@import "ex10_05/main.go" {class=line-numbers}

output:

```text
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
internal error: runtime error: integer divide by zero
```

## 4. Errors

`error` 是 Go 內建的 data type，它是一個 interface, 定義如下：

```go {.line-numbers}
type error interface {
    Error() string
}
```

因此，要自定義 error，只要實作這個 interface 即可。

```go {.line-numbers}
type MyError struct {
    Code int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
```

在 Go 的 function 設計中，很多都會回傳包含 error 的 tuple。eg:

```go {.line-numbers}
resp, err := http.Get(url)

if err != nil {
    return nil, err
}
```

在 Go 的 Code style 規範中，如果 function 會回傳 tuple 且含有 error 時，請把 error 放在 tuple 最後一個欄位。

```go {.line-numbers}
func FindMember(id int) (*Member, error)
```

### 4.1 Errors in Go 1.13

在 Go 1.13 的版本，Error 新增了 Wrap 另一個 error 的功能。詳細可看：[Working with Errors in Go 1.13](https://blog.golang.org/go1.13-errors)

1. `fmt.Errorf` 可使用 `%w` 來包裝另一個 error
1. 新增 `errors.As`, `errors.Is`, `errors.Unwrap`
    - `errors.As`: 用來判斷是那種 data type。
    - `errors.Is`: 用來判斷是否是某個 error。
    - `errors.Unwrap`: 取得包裝的 error, 如果沒有，則回傳 nil。

@import "ex10_06/main.go" {class=line-numbers}

結果：

```text {.line-numbers}
err is myErr: true
err as MyError type: true
otherErr is myErr: false
otherErr as MyError type: false
100: error message
another error
no internal error
```

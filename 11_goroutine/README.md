# 11 Concurrency - Goroutine

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [11 Concurrency - Goroutine](#11-concurrency-goroutine)
  - [0. 前言](#0-前言)
  - [1. ==go== Keyword](#1-go-keyword)
  - [2. Wait for goroutine](#2-wait-for-goroutine)
    - [2.1 已知有幾個 goroutine 會被執行](#21-已知有幾個-goroutine-會被執行)
    - [2.2 每執行 goroutine 前，WaitGroup Counter + 1](#22-每執行-goroutine-前waitgroup-counter-1)

<!-- /code_chunk_output -->

## 0. 前言
此章節的資料，來自 [Go Systems Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming)

> A **goroutine** is the minimum Go entity that can be executed concurrently.
> goroutine is live in **Thread**, so it is not an autonomous entity.

1. Goroutine 不等同於 thread。goroutine 是需要利用 thread 來執行，但並不會為每個 goroutine 產生一個專屬的 thread.
1. 當主程式結束時，goroutine 也一併會結束，即使還沒執行完畢。

## 1. ==go== Keyword

使用 `go` 這個關鍵字來啟動一個 goroutine.

```go {.line-numbers}
package main

import (
    "log"
    "time"
)

func namedFunction() {
    time.Sleep(3 * time.Second)
    log.Println("Printing from namedFunction!")
}

func main() {
    go namedFunction()

    time.Sleep(5 * time.Second)
    log.Println("Exiting....")
}
```

1. 先定義一組 function，之後要用 goroutine 來執行。function 故意延遲 3 秒。

    ```go {.line-numbers}
    func namedFunction() {
        time.Sleep(3 * time.Second)
        log.Println("Printing from namedFunction!")
    }
    ```

1. `go namedFunction()`: 產生一個 goroutine 來執行 `namedFunction()`
1. 主程式故意延遲 5 秒。否則 goroutine 會來不及執行。
1. 也可用 anonymous function

    ```go {.line-numbers}
    package main

    import (
        "log"
        "time"
    )

    func main() {
        go func() {
            time.Sleep(3 * time.Second)
            log.Println("Printing from anonymous")
        }()

        time.Sleep(5 * time.Second)
        log.Println("Exiting....")
    }
    ```

## 2. Wait for goroutine

可以利用 `sync.WaitGroup` 來等待 goroutine 結束。

### 2.1 已知有幾個 goroutine 會被執行

```go {.line-numbers}
package main

import (
    "log"
    "sync"
    "time"
)

func main() {
    waitGroup := &sync.WaitGroup{}

    waitGroup.Add(10)

    for i := 0; i < 10; i++ {
        go func(x int) {
            defer waitGroup.Done()
            time.Sleep(100 * time.Millisecond)
            log.Printf("%d ", x)
        }(i)
    }

    waitGroup.Wait()
    log.Println("Exit....")
}
```

1. `waitGroup := &sync.WaitGroup{}`: 產生一個 wait group
1. `waitGroup.Add(10)`: 告知 wait group 要等幾個 goroutine。
1. 產生 10 個 goroutine，並 `defer waitGroup.Done()`，確保 function 結束後，會告知 wait group 有 goroutine 結束了。

    ```go {.line-numbers}
    for i := 0; i < 10; i++ {
        go func(x int) {
            defer waitGroup.Done()
            time.Sleep(100 * time.Millisecond)
            fmt.Printf("%d ", x)
        }(i)
    }
    ```

1. `waitGroup.Wait()`: 主程序 wait

### 2.2 每執行 goroutine 前，WaitGroup Counter + 1

也可以每需要一個 goroutine 時，wait group 就加 1。記得 `waitGroup.Add(1)` 要在 `go nameFunction()` 之前。

```go {.line-numbers}
package main

import (
    "log"
    "sync"
    "time"
)

func test(x int, wait *sync.WaitGroup) {
    log.Println(x, "start")
    defer wait.Done()

    time.Sleep(5 * time.Second)
    log.Println(x, "end")

}

func main() {

    log.Println("Start...")
    wg := &sync.WaitGroup{}

    wg.Add(1)
    go test(10, wg)

    wg.Add(1)
    go test(11, wg)

    wg.Add(1)
    go test(12, wg)

    wg.Wait()
    log.Println("Exit....")
}
```


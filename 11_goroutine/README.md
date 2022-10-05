# 11 Concurrency: Goroutine

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [11 Concurrency: Goroutine](#11-concurrency-goroutine)
  - [0. 前言](#0-前言)
  - [1. ==go== Keyword (ex11_01, ex11_02)](#1-go-keyword-ex11_01-ex11_02)
  - [2. Wait for goroutine (ex11_03, ex11_04)](#2-wait-for-goroutine-ex11_03-ex11_04)
    - [2.1 已知有幾個 goroutine 會被執行](#21-已知有幾個-goroutine-會被執行)
    - [2.2 每執行 goroutine 前，WaitGroup Counter+1](#22-每執行-goroutine-前waitgroup-counter1)
  - [3. Mutex (ex11_05, ex11_06)](#3-mutex-ex11_05-ex11_06)
  - [4. Atmoic (ex11_07)](#4-atmoic-ex11_07)

<!-- /code_chunk_output -->

## 0. 前言

此章節的資料，來自 [Go Systems Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming)

> A __goroutine__ is the minimum Go entity that can be executed concurrently.
> goroutine is live in __Thread__, so it is not an autonomous entity.

1. Goroutine 不等同於 thread。goroutine 是需要利用 thread 來執行，但並不會為每個 goroutine 產生一個專屬的 thread.
1. 當主程式結束時，goroutine 也一併會結束，即使還沒執行完畢。

## 1. ==go== Keyword (ex11_01, ex11_02)

使用 `go` 這個關鍵字來啟動一個 goroutine.

@import "ex11_01/main.go" {class=line-numbers}

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

@import "ex11_02/main.go" {class=line-numbers}

## 2. Wait for goroutine (ex11_03, ex11_04)

可以利用 `sync.WaitGroup` 來等待 goroutine 結束。

### 2.1 已知有幾個 goroutine 會被執行

@import "ex11_03/main.go" {class=line-numbers}

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

### 2.2 每執行 goroutine 前，WaitGroup Counter+1

也可以每需要一個 goroutine 時，wait group 就加 1。記得 `waitGroup.Add(1)` 要在 `go nameFunction()` 之前。

@import "ex11_04/main.go" {class=line-numbers}

## 3. Mutex (ex11_05, ex11_06)

當多個 goroutine 對同一份資料進行修改時，如果沒有管控的話，會造成資料不一致。如下程式，並不保證一定會得到一組長度為 __20__ 的 slice。

@import "ex11_05/main.go" {class="line-numbers" highlight="9,24,26,29"}

1. 兩個 goroutine 同時 append data。
1. 最後 data 並不保證長度一定為 __20__ 。

可以使用 [sync.Mutex](https://pkg.go.dev/sync#Mutex)，用上鎖的概念，來確保資料的完整與一致。行為說明如下：

1. 當要讀寫共用的資料時，透過 Mutex 取得存取的權限 (上鎖)
1. 當存取完畢後，需釋放權限 (解鎖)
1. 此機制使用時，要留意，以免發生 Dead-Lock 問題。

使用方法如下：

@import "ex11_06/main.go" {class="line-numbers" highlight="17-19"}

## 4. Atmoic (ex11_07)

在 Go 1.19 推出新的 package __atomic__，提供更底層的同步控制。

[Atom (programming language)](https://en.wikipedia.org/wiki/Atom_(programming_language))
[sync/atomic](https://pkg.go.dev/sync/atomic)

@import "ex11_07/main.go" {class=line-numbers}

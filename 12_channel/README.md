# 12 Concurrency - Channel

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

- [12 Concurrency - Channel](#12-concurrency-channel)
  - [0. 前言](#0-前言)
  - [1. Channel](#1-channel)
  - [2. Buffered Channel](#2-buffered-channel)
    - [2.1 Deadlock 1: Reading/Writing with Non-Buffered Channel](#21-deadlock-1-readingwriting-with-non-buffered-channel)
    - [2.2 Deadlock 2: Reading Before Writing with Buffered Channel](#22-deadlock-2-reading-before-writing-with-buffered-channel)
  - [3. Producer and Consumer Pattern (Pipeline)](#3-producer-and-consumer-pattern-pipeline)
    - [利用 goroutine 執行 1 個 producer 及 2 個 consumer](#利用-goroutine-執行-1-個-producer-及-2-個-consumer)
      - [Deadlock: Closing Channel in Main Instead of Producer](#deadlock-closing-channel-in-main-instead-of-producer)
  - [4. Actor Pattern (Pipeline)](#4-actor-pattern-pipeline)
  - [5. Select and Timeout](#5-select-and-timeout)
    - [Select and Timeout 說明](#select-and-timeout-說明)
      - [createNumber](#createnumber)
      - [readNumber](#readnumber)
      - [main](#main)
      - [執行結果](#執行結果)

<!-- /code_chunk_output -->

## 0. 前言

Channel 可以想像是一個資料的通道 (pipe)，一頭是 write，另一頭是 read，資料順序是 FIFO (First In First Out)。通常用在 goroutine 間資料交換。channel 是 thread-safe，因此可以同時讀寫 channel。

channel 的注意事項：

1. 用 `make` 與 `chan` 關鍵字來產生一個 channel，不用時，要用 `close` 關閉。
1. 一個 channel 只能包含一種 data type
1. channel 當作參數傳給 function 時，最好指定是要做 read or write。

## 1. Channel

```go {.line-numbers}
package main

import (
    "log"
    "sync"
)

var (
    waitGroup = &sync.WaitGroup{}
)

func writeChannel(c chan<- int, x int) {
    defer waitGroup.Done()

    log.Println("writing ", x)
    c <- x
    log.Println("written ", x)
}

func readChannel(c <-chan int) {
    defer waitGroup.Done()

    log.Println("reading from channel")
    x := <-c
    log.Println("read: ", x)
}

func main() {
    c := make(chan int)
    defer close(c)

    waitGroup.Add(1)
    go readChannel(c)

    waitGroup.Add(1)
    go writeChannel(c, 10)

    waitGroup.Wait()
    log.Println("exit...")
}
```

說明：

1. `c := make(chan int)`: 產生一個 channel 且 data type 是 `int`。並 `defer close(c)` 確保 channel 會被關閉。
1. `go readChannel(c)`: goroutine 執行 readChannel。

    ```go {.line-numbers}
    func readChannel(c <-chan int) {
        log.Println("reading from channel")
        defer waitGroup.Done()
        x := <-c
        log.Println("read: ", x)
    }
    ```

    注意: `c <-chan`，是 **read only** channel

1. `go writeChannel(c, 10)`: goroutine 執行 writeChannel。

    ```go {.line-numbers}
    func writeChannel(c chan<- int, x int) {
        defer waitGroup.Done()

        log.Println("writing ", x)
        c <- x
        log.Println("wrote ", x)
    }
    ```

    注意：`c chan<- int` 是 **write only** channel。

## 2. Buffered Channel

`c := make(chan int)` 宣告時，沒有指定 channel 的容量，因此在 read/write 時，會 block。在上例中，因為是用 goroutine 執行, 所以不會有問題。

### 2.1 Deadlock 1: Reading/Writing with Non-Buffered Channel

```go {.line-numbers}
package main

import (
    "log"
)

func main() {
    c := make(chan int)
    defer close(c)

    log.Println("writing...")

    c <- 10

    log.Println("written")

    log.Println("reading")

    x := <-c

    log.Println("read ", x)

    log.Println("exit...")
}
```

執行結果，發生 deadlock：

```text
2020/01/16 13:54:02 writing...
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        main.go:13 +0xdb
exit status 2
```

此時，可以設定 channel 的容量，eg: `c := make(chan int, 1)`。則結果如下：

```text
2020/01/16 13:57:29 writing...
2020/01/16 13:57:29 written
2020/01/16 13:57:29 reading
2020/01/16 13:57:29 read  10
2020/01/16 13:57:29 exit...
```

先執 write，資料放在 channel，供之後 read。

### 2.2 Deadlock 2: Reading Before Writing with Buffered Channel

但如果程式的順序，改成先 read 再 write 時，一樣會發生 deadlock。因為還沒寫資料，根本沒資料供 read。

```go {.line-numbers}
func main() {
    c := make(chan int, 1)
    defer close(c)

    log.Println("reading")

    x := <-c

    log.Println("read ", x)

    log.Println("writing...")

    c <- 10

    log.Println("written")

    log.Println("exit...")
}
```

結果：

```text
2020/01/16 13:58:46 reading
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        main.go:13 +0xe2
exit status 2
```

## 3. Producer and Consumer Pattern (Pipeline)

Producer/Consumer 是 channel 最常用的實作模型。概念是一端產出資料 (可能是從資料庫或大檔案讀取資料)，另一端運算資料。

### 3.1 利用 goroutine 執行 1 個 producer 及 2 個 consumer

@import "ex12_05/main.go" {.line-numbers}

與先前的範例最大不同是，這次關閉 channel 是在 `producer` 執行，而非主程序，也就是說在產生完資料後，就關閉 channel，之後就不能再寫入。而 `consumer` 端，在 channel 資料讀完後，就會跳出 for-range 的迴圈而執行完畢。

#### 3.2 Deadlock: Closing Channel in Main Instead of Producer

如果不在 `producer` 關閉 channel，而是在主程序，則會發生 deadlock。

@import "ex12_06/main.go" {class="line-numbers"}

結果：

```text
2020/01/16 14:02:31 start...
2020/01/16 14:02:31 comsumer  2  starting...
2020/01/16 14:02:31 comsumer  1  starting...
2020/01/16 14:02:31 producer start...
2020/01/16 14:02:31 1  got  2
...
2020/01/16 14:02:31 1  got  98
2020/01/16 14:02:31 producer end and close channel
2020/01/16 14:02:31 1  got  99
2020/01/16 14:02:31 2  got  97
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x1199b68)
        /usr/local/go/src/runtime/sema.go:56 +0x42
sync.(*WaitGroup).Wait(0x1199b60)
        /usr/local/go/src/sync/waitgroup.go:130 +0x64
main.main()
        main.go:48 +0x198

goroutine 19 [chan receive]:
main.consumer(0x1, 0xc000060060)
        main.go:27 +0x1fd
created by main.main
        main.go:43 +0x144

goroutine 20 [chan receive]:
main.consumer(0x2, 0xc000060060)
        main.go:27 +0x1fd
created by main.main
        main.go:46 +0x188
exit status 2
```

## 4. Actor Pattern (Pipeline)

Actor Pattern 與 Producer/Consumer Pattern 類似，概念是每一個 Actor 只負責固定的工作。Producer 必須將資料，傳到每個 Actor。以下的範例，是模擬訂單成立後，傳給兩個 Actor，一個負責計算每個分類的業績，另一個計算全站的業績。

@import "ex12_07/main.go" {class="line-numbers"}

說明：

1. `Producer`: 負責模擬產生 100 筆訂單後，往後送給 consumer actor 處理。最後再關閉 consumer actor 的 channel，讓程式可以執行完畢。
1. `CategorySum`: 負責主要統計每個分類的業績。
1. `SiteSum`: 負責統計全站業績

## 5. Select and Timeout

可以透過 `select` 來偵測 channel 是否可以被寫入及是否有資料可以讀取。`select` 可以撘配 `time.After` 來實作 timeout 的機制。

eg:

```go {.line-numbers}
package main

import (
    "log"
    "math/rand"
    "time"
)

func createNumber(max int, randomChannel chan<- int, finishChannel <-chan bool) {
    for {
        select {
        case randomChannel <- rand.Intn(max):
            time.Sleep(1 * time.Second)
        case x := <-finishChannel:
            log.Println("finish channel got ", x)
            if x {
                close(randomChannel)
                log.Println("createNumber end")
                return
            }
        }
    }

}

func readNumber(randomChannel <-chan int) {
    for {
        select {
        case x, ok := <-randomChannel:
            if !ok {
                log.Println("readNumber end")
                return
            }
            log.Println("random channel got ", x)
        case <-time.After(500 * time.Millisecond):
            log.Println("time out")
        }
    }
}

func main() {
    rand.Seed(time.Now().Unix())

    randomChannel := make(chan int)
    finishChannel := make(chan bool)

    go createNumber(100, randomChannel, finishChannel)
    go readNumber(randomChannel)

    time.Sleep(2 * time.Second)
    finishChannel <- false
    time.Sleep(3 * time.Second)
    finishChannel <- true
    time.Sleep(1 * time.Second)
    close(finishChannel)
    log.Println("end")
}
```

### Select and Timeout 說明

#### createNumber

1. `for { }`: 無窮迴圈
1. `select - case`: 使用 `select` 來偵測 channel 狀態。
1. `case randomChannel <- rand.Intn(max)`: 對 `randomChannel` 寫入資料
1. `x := <-finishChannel`: 從 `finishChannel` 讀取資料，如果為 `true` 則關閉 `randomChannel` 並結束 `select - case` 迴圈。

#### readNumber

1. `for { }`: 無窮迴圈
1. `select - case`: 使用 `select` 來偵測 channel 狀態。
1. `case x, ok := <-randomChannel`: 從 `randomChannel` 讀取資料，這邊與先前從 channel 讀資料不同，多了一個 `ok` 來判斷 channel 是否已經被關閉了。如果 `randomChannel` 已被關閉，則跳出迴圈。
1. `case <-time.After(500 * time.Millisecond)`: Timeout 機制，如果 500 ms  內，randomChannel 一直沒有資料寫入的話，則會觸發。

#### main

1. 初始化 channel 及 goroutine.
1. 先停 2 sec. 後，先對 `finishChannel` 寫入 `false`，此時不會中止所有活動，但 `finishChannel` 會得到一個 `false` 值。
1. 再停 3 sec. 後，再對 `finishChannel` 寫入 `true`，此時會中斷 `createNumber` 的迴圈，且 `randomChannel` 會被關閉。
1. `randomChannel` 被關閉後，`readNumber` 會偵測到 `randomChannel` 被關閉，而中斷 `readNumber` 迴圈。
1. 再停 1 sec. 關閉 `finishChannel`。

#### 執行結果

```text
2020/01/16 14:08:28 random channel got  98
2020/01/16 14:08:29 time out
2020/01/16 14:08:29 random channel got  65
2020/01/16 14:08:30 time out
2020/01/16 14:08:30 random channel got  33
2020/01/16 14:08:31 time out
2020/01/16 14:08:31 finish channel got  false
2020/01/16 14:08:31 random channel got  92
2020/01/16 14:08:32 time out
2020/01/16 14:08:32 random channel got  46
2020/01/16 14:08:33 time out
2020/01/16 14:08:33 random channel got  57
2020/01/16 14:08:34 time out
2020/01/16 14:08:34 random channel got  57
2020/01/16 14:08:35 time out
2020/01/16 14:08:35 finish channel got  true
2020/01/16 14:08:35 createNumber end
2020/01/16 14:08:35 readNumber end
2020/01/16 14:08:36 end
```

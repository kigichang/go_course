# 11 Concurrency - Goroutine

此章節的資料，來自 [Go Systems Programming](https://www.packtpub.com/networking-and-servers/go-systems-programming)

1. A **goroutine** is the minimum Go entity that can be executed concurrently.
1. goroutine is live in **Thread**, so it is not an autonomous entity.
1. 當主程式結束時，goroutine 也一併會結束，即使還沒執行完畢。

## ==go== Keyword

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

## Wait for goroutine

可以利用 `sync.WaitGroup` 來等待 goroutine 結束。

### 已知有幾個 goroutine 會被執行

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

### 每執行 goroutine 前，WaitGroup Counter + 1

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

## Go Routine Puzzlers

使用 go routine 時，要注意 closure 的 binding 時機。

### Puzzlers Example 1

```go {.line-numbers}
package main

import (
    "fmt"
    "sync"
)

// Job ...
type Job struct {
    ID int
}

func main() {
    var jobs []Job
    for i := 0; i < 2; i++ {
        jobs = append(jobs, Job{i})
    }

    wait := sync.WaitGroup{}

    for _, x := range jobs {
        wait.Add(1)
        go func(job *Job) {
            defer wait.Done()
            fmt.Println(job.ID)
        }(&x)
    }

    wait.Wait()
}
```

結果：

```text
1
1
```

修正:

### Puzzlers Example 2

```go {.line-numbers}
package main

import (
    "fmt"
    "sync"
)

// Test ...
type Test struct {
    ID int
}

func main() {
    var tests []Test
    for i := 0; i < 2; i++ {
        tests = append(tests, Test{i})
    }

    wait := sync.WaitGroup{}

    for i := range tests {
        wait.Add(1)
        go func(t *Test) {
            defer wait.Done()
            fmt.Println(t.ID)
        }(&tests[i])
    }

    wait.Wait()
}
```

結果：

```text
1
0
```

因為 goroutine 會需要時間做初始化，所以在 Loop 的宣告的 goroutine, 有很大的機會會在 Loop 結束後才執行。因此在 closure binding 時，會有很大的機會會 binding 到最後一個值。在 Puzzlers Example 1 中，最後 `x` binding 的值會是最後一個，並取記憶體位址傳入。建議有這種情形時，要明確指定值。

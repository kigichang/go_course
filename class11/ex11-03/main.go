package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(10)

	for i := 0; i < 10; i++ {
		go func(x int) {
			defer waitGroup.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%d ", x)
		}(i)
	}

	waitGroup.Wait()
	fmt.Println("Exit....")
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg = new(sync.WaitGroup)

type Queue struct {
	closed *atomic.Bool
}

func (q *Queue) Close(idx int) (err error) {
	defer wg.Done()
	if flag := q.closed.Load(); flag {
		fmt.Printf("%d early return: %t\n", idx, flag)
		return
	}
	fmt.Printf("%d set flag to true\n", idx)
	q.closed.Store(true)
	// free resources
	return
}

func main() {
	wg.Add(10)
	queue := &Queue{
		closed: new(atomic.Bool),
	}

	for i := 0; i < 10; i++ {
		go queue.Close(i)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

var (
	data  = []int{}
	wait  = &sync.WaitGroup{}
	mutex = &sync.Mutex{}
)

func add(start int) {
	defer wait.Done()
	for i := 0; i < 10; i++ {
		mutex.Lock()
		data = append(data, start+i)
		mutex.Unlock()
	}

}

func main() {

	wait.Add(1)
	go add(0)
	wait.Add(1)
	go add(100)

	wait.Wait()
	fmt.Println(len(data))
}

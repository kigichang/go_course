package main

import (
	"fmt"
	"sync"
)

var (
	data = []int{}
	wait = &sync.WaitGroup{}
)

func add(start int) {
	defer wait.Done()
	for i := 0; i < 10; i++ {
		data = append(data, start+i)
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

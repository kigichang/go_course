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

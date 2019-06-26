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

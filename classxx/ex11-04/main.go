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
	waitGroup := sync.WaitGroup{}

	waitGroup.Add(1)

	go test(10, &waitGroup)

	waitGroup.Add(1)

	go test(11, &waitGroup)

	waitGroup.Add(1)

	go test(12, &waitGroup)

	waitGroup.Wait()
	log.Println("Exit....")

}

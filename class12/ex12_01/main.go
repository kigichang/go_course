package main

import (
	"log"
	"sync"
)

var (
	waitGroup = sync.WaitGroup{}
)

func writeChannel(c chan<- int, x int) {
	defer waitGroup.Done()

	log.Println("writing ", x)
	c <- x
	log.Println("written ", x)
}

func readChannel(c <-chan int) {
	log.Println("reading from channel")
	defer waitGroup.Done()
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

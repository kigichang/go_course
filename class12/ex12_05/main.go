package main

import (
	"log"
	"sync"
)

var (
	waitGroup = sync.WaitGroup{}
)

func producer(min, max int, c chan<- int) {
	defer waitGroup.Done()
	log.Println("producer start...")
	for i := min; i < max; i++ {
		c <- i
	}
	close(c)
	log.Println("producer end and close channel")
}

func consumer(id int, c <-chan int) {
	defer waitGroup.Done()
	count := 0

	log.Println("comsumer ", id, " starting...")
	for a := range c {
		log.Println(id, " got ", a)
		count++
	}
	log.Printf("consumer %d got %d times and end\n", id, count)
}

func main() {
	log.Println("start...")
	c := make(chan int)

	waitGroup.Add(1)
	go producer(1, 100, c)

	waitGroup.Add(1)
	go consumer(1, c)

	waitGroup.Add(1)
	go consumer(2, c)

	waitGroup.Wait()
	log.Println("end")
}

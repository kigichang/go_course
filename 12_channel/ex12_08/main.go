package main

import (
	"log"
	"math/rand"
	"time"
)

func createNumber(max int, randomChannel chan<- int, finishChannel <-chan bool) {
	for {
		select {
		case randomChannel <- rand.Intn(max):
			time.Sleep(1 * time.Second)
		case x := <-finishChannel:
			log.Println("finish channel got ", x)
			if x {
				close(randomChannel)
				log.Println("createNumber end")
				return
			}
		}
	}

}

func readNumber(randomChannel <-chan int) {
	for {
		select {
		case x, ok := <-randomChannel:
			if !ok {
				log.Println("readNumber end")
				return
			}
			log.Println("random channel got ", x)
		case <-time.After(500 * time.Millisecond):
			log.Println("time out")
		}
	}
}

func main() {
	rand.Seed(time.Now().Unix())

	randomChannel := make(chan int)
	finishChannel := make(chan bool)

	go createNumber(100, randomChannel, finishChannel)
	go readNumber(randomChannel)

	time.Sleep(2 * time.Second)
	finishChannel <- false

	time.Sleep(3 * time.Second)
	finishChannel <- true

	time.Sleep(1 * time.Second)
	close(finishChannel)

	log.Println("end")
}

package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx1, cancel1 := context.WithCancel(context.Background())
	ctx2, cancel2 := context.WithTimeout(ctx1, 10*time.Second)
	defer func() {
		log.Println("cancel 2")
		cancel2()
	}()

	<-time.After(2 * time.Second)
	log.Println("cancel 1")
	cancel1()
	log.Println("ctx1:", ctx1.Err())
	log.Println("ctx2:", ctx2.Err())
	log.Println("end")
}

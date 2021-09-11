package main

import (
	"log"
)

func main() {
	c := make(chan int)
	defer close(c)

	log.Println("writing...")

	c <- 10

	log.Println("written")

	log.Println("reading")

	x := <-c

	log.Println("read ", x)

	log.Println("exit...")
}

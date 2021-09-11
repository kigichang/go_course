package main

import (
	"log"
	"time"
)

func namedFunction() {
	time.Sleep(3 * time.Second)
	log.Println("Printing from namedFunction!")
}

func main() {
	go namedFunction()

	time.Sleep(5 * time.Second)
	log.Println("Exiting....")
}

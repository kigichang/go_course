package main

import (
	"log"
	"time"
)

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("Printing from anonymous")
	}()

	time.Sleep(5 * time.Second)
	log.Println("Exiting....")
}

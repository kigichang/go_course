package main

import (
	"fmt"
	"time"
)

func namedFunction() {
	time.Sleep(3 * time.Second)
	fmt.Println("Printing from namedFunction!")
}

func main() {
	go namedFunction()

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting....")
}

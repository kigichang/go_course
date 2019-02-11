package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Printing from anonymous")
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting....")
}

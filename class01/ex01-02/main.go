package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])

	if len(os.Args) > 1 {
		fmt.Println("hi, ", os.Args[1])
	}
	fmt.Println("hello world")
}

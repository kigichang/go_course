package main

import (
	"fmt"
	"os"
)

func main() {

	defer fmt.Println("call defer")
	fmt.Println("main end")
	os.Exit(1)
}

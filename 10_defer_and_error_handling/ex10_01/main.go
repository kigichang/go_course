package main

import "fmt"

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	fmt.Println("before return")
	return x + x
}

func main() {
	defer func() {
		fmt.Println("defer end 1")
	}()

	defer func() {
		fmt.Println("defer end 2")
	}()

	fmt.Println("double of 4 is", double(4))
	fmt.Println("main end")
}

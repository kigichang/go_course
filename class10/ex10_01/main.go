package main

import "fmt"

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	fmt.Println("before return")
	return x + x
}

func main() {
	defer func() {
		fmt.Printf("defer end")
	}()
	double(4)
	fmt.Println("main end")
}

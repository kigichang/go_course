package main

import "fmt"

func a1() {
	fmt.Println("\na1 loop start")
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
	fmt.Println("a1 loop end")
}

func a2() {
	fmt.Println("\na2 loop start")
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Print(i, " ") // warn: [go-vet] loop variable i captured by func literal
		}()
	}
	fmt.Println("a2 loop end")
}

func a3() {
	fmt.Println("\na3 loop start")
	for i := 0; i < 3; i++ {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Println("a3 loop end")
}

func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Printf("\n%d main end", i)
	}

	a1() // 2 1 0
	a2() // 3 3 3
	a3() // 2 1 0
}

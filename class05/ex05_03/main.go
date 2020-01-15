package main

import "fmt"

func minus(s [6]int) {
	for i, x := range s {
		s[i] = x - 1
	}
}

func plus(s []int) {
	for i, x := range s {
		s[i] = x + 1
	}
}

func main() {
	s := [6]int{0, 1, 2, 3, 4, 5}

	fmt.Println(s) // [0 1 2 3 4 5]
	minus(s)
	fmt.Println(s) // [0 1 2 3 4 5]

	s1 := s[2:]

	fmt.Println(s) // [0 1 2 3 4 5]
	plus(s1)
	fmt.Println(s) // [0 1 3 4 5 6]
}

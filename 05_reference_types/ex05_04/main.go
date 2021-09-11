package main

import "fmt"

func main() {
	s := [6]int{0, 1, 2, 3, 4, 5}
	fmt.Println(len(s), cap(s)) // 6 6

	s1 := s[2:]
	fmt.Println(len(s1), cap(s1)) // 4 4

	s1 = append(s1, 100)

	fmt.Println(s)  // [0 1 2 3 4 5]
	fmt.Println(s1) // [2 3 4 5 100]

	s2 := s[1:3]
	fmt.Println(len(s2), cap(s2)) // 2 5

	s2 = append(s2, 30)

	fmt.Println(s)  // [0 1 2 30 4 5]
	fmt.Println(s2) // [1 2 30]
}

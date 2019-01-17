package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, s == nil, len(s), cap(s)) // [] true 0 0

	s = nil
	fmt.Println(s, s == nil, len(s), cap(s)) // [] true 0 0

	s = []int(nil)
	fmt.Println(s, s == nil, len(s), cap(s)) // [] true 0 0

	s = []int{}
	fmt.Println(s, s == nil, len(s), cap(s)) // [] false 0 0

	s = []int{1, 2, 3}
	fmt.Println(s, s == nil, len(s), cap(s)) // [1 2 3] false 3 3

	s = make([]int, 4)
	fmt.Println(s, s == nil, len(s), cap(s)) // [0 0 0 0] false 4 4

	s = make([]int, 5, 6)
	fmt.Println(s, s == nil, len(s), cap(s)) // [0 0 0 0 0] false 5 6
}

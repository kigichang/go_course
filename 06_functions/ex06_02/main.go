package main

import "fmt"

func arrTest(a [3]int) {
	for i, x := range a {
		a[i] = x + 1
	}
}

func arrTestBySlice(a []int) {
	for i, x := range a {
		a[i] = x + 1
	}
}

func main() {
	a := [3]int{1, 2, 3}

	fmt.Println(a) // [1 2 3]
	arrTest(a)     // 用 array
	fmt.Println(a) // [1 2 3]

	arrTestBySlice(a[:]) // 用 Slice
	fmt.Println(a)       // [2 3 4]
}

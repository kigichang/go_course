package main

import "fmt"

func main() {
	var a [3]int // [0,0,0]
	fmt.Printf("%T: %v\n", a, a)

	var b = [3]int{1, 2, 3}
	fmt.Printf("%T: %v\n", b, b)

	var r = [3]int{1, 2}
	fmt.Println(r[2]) // "0"

	q := [...]int{1, 2, 3}
	fmt.Printf("%T: %v\n", q, q) // "[3]int"

	x := [...]int{5: -1} // value of 0 ~ 4th is "0", and 5th is "-1", length of array is 6
	fmt.Println(x)
	fmt.Println(len(x))

	// Print the indices and elements.
	for i, v := range x {
		fmt.Printf("%d %d\n", i, v)
	}

	// // Print the elements only.
	for _, v := range x {
		fmt.Printf("%d\n", v)
	}
}

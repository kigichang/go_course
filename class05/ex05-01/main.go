package main

import "fmt"

func main() {
	a := 10
	b := &a
	*b = 20

	fmt.Println(a) // 20

	arr := [3]int{0, 1, 2}

	p := &arr
	//p++ // invalid operation: p++ (non-numeric type *[3]int)
	fmt.Printf("%p: %v, %v", p, p, *p)
}

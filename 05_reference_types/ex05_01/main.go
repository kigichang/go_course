package main

import "fmt"

func AddByValue(a int) {
	fmt.Printf("Point(%p), Value(%d) of Parameter a\n", &a, a)
	a += 1
}

func AddByPointer(a *int) {
	fmt.Printf("Pointer(%p), Value(%x) of Paramter a\n", &a, a)
	*a = *a + 1
}

func main() {
	a := 10
	b := &a
	*b = 20

	fmt.Println(a) // 20

	arr := [3]int{0, 1, 2}

	p := &arr
	//p++ // invalid operation: p++ (non-numeric type *[3]int)
	fmt.Printf("%p: %v, %v\n", p, p, *p)

	fmt.Printf("Point(%p), Value(%d) of a\n", &a, a)
	AddByValue(a)
	fmt.Printf("%d\n", a) // 20

	fmt.Printf("Pointer(%p), Value(%x) of b (Pointer of a)\n", &b, b)
	AddByPointer(b)
	fmt.Printf("%d\n", a) // 21
}

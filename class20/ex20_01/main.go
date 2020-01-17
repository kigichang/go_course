package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x struct {
		a bool
		b float64
		c int16
	}

	fmt.Println("Sizeof x:", unsafe.Sizeof(x))
	fmt.Println("Alignof x:", unsafe.Alignof(x))

	fmt.Println("Sizeof x.a:", unsafe.Sizeof(x.a), "AlignOf x.a:", unsafe.Alignof(x.a), "Offsetof x.a:", unsafe.Offsetof(x.a))
	fmt.Println("Sizeof x.b:", unsafe.Sizeof(x.b), "AlignOf x.b:", unsafe.Alignof(x.b), "Offsetof x.b:", unsafe.Offsetof(x.b))
	fmt.Println("Sizeof x.c:", unsafe.Sizeof(x.c), "AlignOf x.c:", unsafe.Alignof(x.c), "Offsetof x.c:", unsafe.Offsetof(x.c))

	var y struct {
		a float64
		b int16
		c bool
	}

	fmt.Println("Sizeof y:", unsafe.Sizeof(y))
	fmt.Println("Alignof y:", unsafe.Alignof(y))

	fmt.Println("Sizeof y.a:", unsafe.Sizeof(y.a), "AlignOf y.a:", unsafe.Alignof(y.a), "Offsetof y.a:", unsafe.Offsetof(y.a))
	fmt.Println("Sizeof y.b:", unsafe.Sizeof(y.b), "AlignOf y.b:", unsafe.Alignof(y.b), "Offsetof y.b:", unsafe.Offsetof(y.b))
	fmt.Println("Sizeof y.c:", unsafe.Sizeof(y.c), "AlignOf y.c:", unsafe.Alignof(y.c), "Offsetof y.c:", unsafe.Offsetof(y.c))

	var z struct {
		a bool
		b int16
		c float64
	}

	fmt.Println("Sizeof z:", unsafe.Sizeof(z))
	fmt.Println("Alignof z:", unsafe.Alignof(z))

	fmt.Println("Sizeof z.a:", unsafe.Sizeof(z.a), "AlignOf y.a:", unsafe.Alignof(z.a), "Offsetof z.a:", unsafe.Offsetof(z.a))
	fmt.Println("Sizeof z.b:", unsafe.Sizeof(z.b), "AlignOf y.b:", unsafe.Alignof(z.b), "Offsetof z.b:", unsafe.Offsetof(z.b))
	fmt.Println("Sizeof z.c:", unsafe.Sizeof(z.c), "AlignOf y.c:", unsafe.Alignof(z.c), "Offsetof z.c:", unsafe.Offsetof(z.c))
}

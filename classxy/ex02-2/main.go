package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 10

	ap1 := uintptr(unsafe.Pointer(&a))

	fmt.Printf("%x, %T\n", ap1, ap1)

	ap2 := (*int)(unsafe.Pointer(ap1)) // warning: possible misuse of unsafe.Pointer

	fmt.Printf("%x, %v, %T\n", ap2, *ap2, ap2)
}

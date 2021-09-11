package main

import (
	"fmt"
	"unsafe"
)

type test struct {
	name string
	age  int
	addr string
}

func main() {

	slice := []int{1, 2, 3, 4, 5}

	snd := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&slice[0])) + 1*unsafe.Sizeof(slice[0])))
	fmt.Printf("%v, %T\n", snd, snd)

	t := test{"A", 10, "ABC"}

	addr := *(*string)(unsafe.Pointer((uintptr(unsafe.Pointer(&t)) + unsafe.Offsetof(t.addr))))
	fmt.Printf("%v, %T\n", addr, addr)
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {
	tmp := []int{1, 2, 3}

	ptr := uintptr(unsafe.Pointer(&tmp[0]))

	ptr += unsafe.Sizeof(tmp[0])

	snd := (*int)(unsafe.Pointer(ptr))
	fmt.Println(*snd)
}

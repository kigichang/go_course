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

	// 取結束位址
	out := unsafe.Pointer(uintptr(unsafe.Pointer(&slice[0])) + uintptr(len(slice)))
	fmt.Println(out)

	// 暫存 uintptr, 再轉成 unsafe.Pointer()
	firstp := uintptr(unsafe.Pointer(&slice[0]))
	sndp := firstp + unsafe.Sizeof(slice[0])

	snd := *(*int)(unsafe.Pointer(sndp)) // warning: possible misuse of unsafe.Pointer
	fmt.Printf("%v, %T\n", snd, snd)

	t := test{"A", 10, "ABC"}

	// 取結束位址
	end := unsafe.Pointer(uintptr(unsafe.Pointer(&t)) + unsafe.Sizeof(t))
	fmt.Println(end)
}

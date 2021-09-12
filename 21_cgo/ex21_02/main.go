package main

import (
	"fmt"
	"unsafe"
)

func main() {

	f := 3.1415926
	t1 := *(*uint64)(unsafe.Pointer(&f))
	fmt.Printf("%v, %T\n", t1, t1)

	t2 := *(*uint32)(unsafe.Pointer(&f))
	fmt.Printf("%v, %T\n", t2, t2)

	var f1 float32 = 3.1415926
	t3 := (*uint64)(unsafe.Pointer(&f1))
	fmt.Printf("%v, %v, %T\n", t3, *t3, t3)
	fmt.Printf("%v, %v, %T\n", &f1, f1, f1)

	*t3 = 1<<63 - 1
	fmt.Printf("%v, %T\n", *t3, t3)
	fmt.Printf("%v, %T\n", f1, f1)
}

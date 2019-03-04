package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type test struct {
	name string
	age  int
	addr string
}

func main() {
	u := reflect.ValueOf(new(int)).Pointer()
	p := (*int)(unsafe.Pointer(u))
	fmt.Printf("%v, %v, %T\n", p, *p, p)
	*p = 200
	fmt.Printf("%v, %v, %T\n", p, *p, p)
}

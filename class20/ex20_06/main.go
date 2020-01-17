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
	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
	fmt.Printf("%v, %v, %T\n", p, *p, p)
	*p = 200
	fmt.Printf("%v, %v, %T\n", p, *p, p)
}

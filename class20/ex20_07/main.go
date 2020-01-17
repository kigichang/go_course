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
	str := "hello world"

	hdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("%x, %x, %v\n", &str, uintptr(unsafe.Pointer(hdr)), hdr)

	str = str + "DEF"
	hdr = (*reflect.StringHeader)(unsafe.Pointer(&str))
	fmt.Printf("%x, %x, %v\n", &str, uintptr(unsafe.Pointer(hdr)), hdr)

	str2 := "ok"
	hdr2 := (*reflect.StringHeader)(unsafe.Pointer(&str2))
	*hdr = *hdr2
	fmt.Printf("%x, %x, %v\n", &str, uintptr(unsafe.Pointer(hdr)), hdr)

	a, b := "A", "BB"

	hdra := (*reflect.StringHeader)(unsafe.Pointer(&a))
	fmt.Printf("%x, %x, %v\n", &a, uintptr(unsafe.Pointer(hdra)), hdra)

	a = b
	hdra = (*reflect.StringHeader)(unsafe.Pointer(&a))
	fmt.Printf("%x, %x, %v\n", &a, uintptr(unsafe.Pointer(hdra)), hdra)
}

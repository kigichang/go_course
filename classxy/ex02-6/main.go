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

/*
由上例中可知：
    1. `&string` 代表取出 `*StringHeader`。
    1. 做完 `str = str + "DEF"`，`StringHeader.Data` 的值，已經被修改了，也就代表產生一個新的字串資料。但原本 StringHeader 的位址不變
    1. `*hdr = *hdr2` 等同 `str = str2`，因為整個 StringHeader 的值都被更改了。
*/

package main

/*
#include "stdio.h"
#include "stdlib.h"

void hello(char* name) {
	printf("hello %s\r\n", name);
}
*/
import "C"

import (
	"unsafe"
)

func main() {
	cs := C.CString("cyberon")
	C.hello(cs)
	C.free(unsafe.Pointer(cs))
}

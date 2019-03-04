package main

//#cgo LDFLAGS: -L. -lfoo -lpthread
//#cgo CFLAGS: -framework CoreFoundation -framework Security
//#include "cfoo.h"
//#include "stdlib.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func conv(args []string) []*C.char {
	ret := make([]*C.char, len(args))

	for i, x := range args {
		ret[i] = C.CString(x)
	}

	return ret
}

// New ...
func New() C.Test {
	var args = []string{"A", "B", "C"}

	xx := conv(args)

	aa := C.int(100)
	a := &aa

	t := C.TestNew(3, &xx[0], a)
	fmt.Println(aa)

	for _, v := range xx {
		C.free(unsafe.Pointer(v))
	}

	return t

}

func main() {

	foo := C.FooInit()
	C.FooBar(foo)
	C.FooFree(foo)

	t := New()
	C.TestFree(t)

}

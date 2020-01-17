package main

/*
#include "stdlib.h"
#include "stdio.h"

void list(const char** str, const int size) {
    for(int i = 0; i < size; i++) {
        printf("%d:%s\n", i, str[i]);
    }
    fflush(stdout);
}
*/
import "C"

import (
    "unsafe"
)

func main() {
    strings := []string{"hello", "world"}

    tmp := make([]*C.char, len(strings))

    // Go string to C *char
    for i, str := range strings {
        tmp[i] = C.CString(str)
    }

    C.list(&tmp[0], C.int(len(strings)))

    // Free all C *char
    for _, x := range tmp {
        C.free(unsafe.Pointer(x))
    }
}
package main

import "fmt"
import "C"

// Hello ...
//export Hello
func Hello(name *C.char) {
	fmt.Println("Hello,", C.GoString(name))
}

func main() {

}

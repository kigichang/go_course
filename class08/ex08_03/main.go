package main

import (
	"fmt"
	"io"
	"os"
)

// MyFile ...
type MyFile struct{}

// Close implements io.Closer interface
func (f *MyFile) Close() error {
	return nil
}

// BadNew1 ...
func BadNew1() *MyFile {
	return nil
}

// BadNew2 ...
func BadNew2() io.Closer {
	var f *MyFile
	return f
}

func main() {
	var f1 io.Closer = BadNew1()
	f2 := BadNew2()

	fmt.Printf("%T, %v, %v\n", f1, f1, f1 == nil) // *main.MyFile, <nil>, false
	fmt.Printf("%T, %v, %v\n", f2, f2, f2 == nil) // *main.MyFile, <nil>, false

	switch v := f1.(type) {
	case *MyFile:
		fmt.Printf("%T, %v, %v\n", v, v, v == nil) // *main.MyFile, <nil>, true
	default:
		fmt.Println("unknown type")
	}

	value, ok := f2.(*MyFile)
	fmt.Printf("%v, %T, %v, %v\n", ok, value, value, value == nil) // true, *main.MyFile, <nil>, true

	f3, ok := f2.(*os.File)
	fmt.Printf("%v, %T, %v, %v\n", ok, f3, f3, f3 == nil) // false, *os.File, <nil>, true
}

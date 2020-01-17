package main

import (
	"fmt"
	"swig_test/foo"
)

func main() {
	f := foo.NewCxxFoo(10)
	fmt.Println(foo.Add(10.0))
	f.Bar()
	foo.DeleteCxxFoo(f)
	fmt.Println(foo.Gcd(12, 16))
	foo.SetFoo(100.0)
	fmt.Println(foo.GetFoo())
	fmt.Println("end")
}

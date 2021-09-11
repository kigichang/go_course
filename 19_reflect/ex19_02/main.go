package main

import "fmt"

type test struct{ name string }

func (t *test) Test() string {
	return fmt.Sprintf("%s:Test", t.name)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case test:
		fmt.Println("this is a test struct,", v.Test())
	case *test:
		fmt.Println("this is a pointer of test struct,", v.Test())
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
	do(&test{"pointer"})
	do(test{"struct"})
}

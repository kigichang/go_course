package main

import "fmt"

func main() {
	var any interface{}
	fmt.Printf("%T\n", any) // <nil>

	a := 10
	b := 100.0
	c := "string"
	d := struct{ A, B string }{"foo", "boo"}
	e := []string{}
	f := map[string]int{}

	any = a
	fmt.Printf("%T\n", any) // int

	any = &a
	fmt.Printf("%T\n", any) // *int

	any = b
	fmt.Printf("%T\n", any) // float64

	any = &b
	fmt.Printf("%T\n", any) // *float64

	any = c
	fmt.Printf("%T\n", any) // string

	any = &c
	fmt.Printf("%T\n", any) // *string

	any = d
	fmt.Printf("%T\n", any) // struct { A string; B string }

	any = &d
	fmt.Printf("%T\n", any) // *struct { A string; B string }

	any = e
	fmt.Printf("%T\n", any) // []string

	any = &e
	fmt.Printf("%T\n", any) // *[]string

	any = f
	fmt.Printf("%T\n", any) // map[string]int

	any = &f
	fmt.Printf("%T\n", any) // *map[string]int
}

package main

import "fmt"

// Person ...
type Person struct {
	Age  int
	Name string
}

func test(p Person) {
	p.Age++
	p.Name += " by test"
}

func testByPtr(p *Person) {
	p.Age++
	p.Name += " by test"
}

func main() {
	p := Person{
		Age:  5,
		Name: "Test",
	}

	fmt.Println(p) // {0 Test}
	test(p)        // 用原本的 struct
	fmt.Println(p) // {0 Test}

	testByPtr(&p)  // 改用 pointer
	fmt.Println(p) // {1 Test by test}
}

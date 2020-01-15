package main

import "fmt"

// Runable ...
type Runable interface {
	Run()
}

func main() {
	// value type
	var a int
	var b float64
	var c string
	var d bool
	var e [3]int
	var f struct {
		A int
		B string
	}

	// reference type
	var g *int
	var h []int
	var i map[string]int
	var j func(int) int
	var k chan int
	var m Runable

	// value type
	fmt.Printf("%v\n", a) // 0
	fmt.Printf("%v\n", b) // 0
	fmt.Printf("%q\n", c) // ""
	fmt.Printf("%v\n", d) // false
	fmt.Printf("%v\n", e) // [0 0 0]
	fmt.Printf("%v\n", f) // {0 }

	// reference type
	fmt.Printf("%v, nil? %v\n", g, g == nil) // <nil>, nil? true
	fmt.Printf("%v, nil? %v\n", h, h == nil) // [], nil? true
	fmt.Printf("%v, nil? %v\n", i, i == nil) // map[], nil? true
	fmt.Printf("%v, nil? %v\n", j, j == nil) // <nil>, nil? true
	fmt.Printf("%v, nil? %v\n", k, k == nil) // <nil>, nil? true
	fmt.Printf("%v, nil? %v\n", m, m == nil) // <nil>, nil? true
}

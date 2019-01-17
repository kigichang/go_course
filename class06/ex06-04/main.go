package main

import "fmt"

func add(x int, y int) int   { return x + y }
func sub(x, y int) (z int)   { z = x - y; return }
func first(x int, _ int) int { return x }
func zero(int, int) int      { return 0 }

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func compose(f, g func(int) int) func(int) int {
	return func(a int) int { // anonymous function
		return g(f(a))
	}
}

func main() {
	fmt.Printf("%T\n", add)   // "func(int, int) int"
	fmt.Printf("%T\n", sub)   // "func(int, int) int"
	fmt.Printf("%T\n", first) // "func(int, int) int"
	fmt.Printf("%T\n", zero)  // "func(int, int) int"

	var f func(int) int   // signature
	fmt.Printf("%T\n", f) // "func(int) int"

	f = square
	fmt.Println(f(3)) // "9"

	f = negative
	fmt.Println(f(3)) // "-3"

	//f = product // cannot use product (type func(int, int) int) as type func(int) int in assignment

	k1 := compose(square, negative)
	fmt.Printf("%T\n", k1) // func(int) int
	fmt.Println(k1(10))    // -100 negative(square(10))

	k2 := compose(negative, square)
	fmt.Printf("%T\n", k2) // func(int) int
	fmt.Println(k2(10))    // 100 square(negative(10))
}

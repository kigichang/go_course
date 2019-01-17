package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x)) // "123 123"

	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
	fmt.Println(fmt.Sprintf("%b", x))           // "1111011"

	a, _ := strconv.Atoi("123")             // a is an int
	b, _ := strconv.ParseInt("123", 10, 64) // b is base 10, up to 64 bits

	fmt.Println(a)
	fmt.Println(b)
}

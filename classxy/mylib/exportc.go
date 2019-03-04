package main

import "fmt"
import "C"

//export MySum
func MySum(x, y int) int {
	return x + y
}

//export Hello
func Hello(str string) {
	fmt.Printf("Hello %s\n", str)
}

func main() {

}

package main

// #include "stdlib.h"
// #include "stdio.h"
// #define b (5)
// int add(int a) {
//   return a + b;
// }
import "C"
import "fmt"

func main() {
	fmt.Println(C.add(100))
}

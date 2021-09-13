package main

import (
	"fmt"
	"reflect"
)

type number interface {
	type int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64
}

func add[T number](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2), reflect.TypeOf(add(1, 2)))
	fmt.Println(add(1.0, 2.0), reflect.TypeOf(add(1.0, 2.0)))
}
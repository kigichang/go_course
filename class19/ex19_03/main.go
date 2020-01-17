package main

import (
	"reflect"
)

func main() {
	var i interface{} = 10

	reflect.ValueOf(i).Elem() // panic: reflect: call of reflect.Value.Elem on int Value
}

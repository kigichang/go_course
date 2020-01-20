package main

import (
	"fmt"
	"reflect"
)

func println(v reflect.Value) {
	fmt.Printf("`%v` is a Zero Value of %T? %v\n", v.Interface(), v.Interface(), v.IsZero())
}

func isnil(v reflect.Value) {
	fmt.Printf("`%v` is nil? %v\n", v.Interface(), v.IsNil())
}

func zero(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("Zero Value of %T is `%v`\n", x, reflect.Zero(t).Interface())
}

func main() {

	zero(100)
	zero(3.14)
	zero("ABC")

	println(reflect.ValueOf(0))
	println(reflect.ValueOf(1))
	println(reflect.ValueOf(0.0))

	println(reflect.ValueOf(""))

	slice := []int{}
	println(reflect.ValueOf(slice))
	isnil(reflect.ValueOf(slice))

	slice = nil
	println(reflect.ValueOf(slice))
	isnil(reflect.ValueOf(slice))

	mymap := make(map[string]int)
	println(reflect.ValueOf(mymap))
	isnil(reflect.ValueOf(mymap))

	mymap = nil
	println(reflect.ValueOf(mymap))
	isnil(reflect.ValueOf(mymap))
}

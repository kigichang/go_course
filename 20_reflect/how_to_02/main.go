package main

import (
	"fmt"
	"reflect"
)

// makeSlice returns a reflect.Value of go slice.
func makeSlice(t reflect.Type, lenAndCap ...int) reflect.Value {
	switch len(lenAndCap) {
	case 1:
		return reflect.MakeSlice(reflect.SliceOf(t), lenAndCap[0], lenAndCap[0])
	case 2:
		return reflect.MakeSlice(reflect.SliceOf(t), lenAndCap[0], lenAndCap[1])
	default:
		return reflect.MakeSlice(reflect.SliceOf(t), 0, 0)
	}
}

func main() {
	rslice1 := makeSlice(reflect.TypeOf(0), 2, 3)
	fmt.Printf("%v\n", rslice1.Interface().([]int))

	rslice1.Index(0).Set(reflect.ValueOf(10))
	rslice1.Index(1).Set(reflect.ValueOf(20))
	fmt.Printf("%v\n", rslice1.Interface().([]int))

	rslice1 = reflect.Append(rslice1, reflect.ValueOf(-1))
	fmt.Printf("%v\n", rslice1.Interface().([]int))

	rslice2 := makeSlice(rslice1.Type().Elem(), 5, 5)
	for i := 0; i < 5; i++ {
		rslice2.Index(i).Set(reflect.ValueOf(i * 20))
	}
	fmt.Printf("%v\n", rslice2.Interface().([]int))

	rsliceTotal := reflect.AppendSlice(rslice1, rslice2)
	fmt.Printf("%v\n", rsliceTotal.Interface().([]int))

	rSubSlice := rsliceTotal.Slice(2, 5)
	fmt.Printf("%v\n", rSubSlice.Interface().([]int))

	fmt.Printf("length: %d, cap: %d\n", rSubSlice.Len(), rSubSlice.Cap())
}

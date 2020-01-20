package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func makeMap(k, v reflect.Type, size int) reflect.Value {
	if size < 0 {
		return reflect.MakeMap(reflect.MapOf(k, v))
	}

	return reflect.MakeMapWithSize(reflect.MapOf(k, v), size)
}

func main() {
	rmap := makeMap(reflect.TypeOf(""), reflect.TypeOf(0), 5)

	for i := 0; i < 5; i++ {
		rmap.SetMapIndex(reflect.ValueOf(strconv.Itoa(i)), reflect.ValueOf(i))
	}

	iter := rmap.MapRange()

	for iter.Next() {
		fmt.Printf("%s: %d\n", iter.Key().String(), iter.Value().Int())
	}

	for _, key := range rmap.MapKeys() {
		fmt.Printf("%s: %d\n", key.String(), rmap.MapIndex(key).Int())
	}

	fmt.Printf(`if "ABC" in map?: %v`, rmap.MapIndex(reflect.ValueOf("ABC")).IsValid())
	fmt.Println()
}

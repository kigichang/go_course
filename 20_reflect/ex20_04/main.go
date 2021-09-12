package main

import (
	"fmt"
	"reflect"
)

// Test ...
type Test struct {
	ID     int64  `json:"id"`
	Name   string `json:"name,omitempty"`
	Hidden string `json:"-"`
}

func dig(x interface{}) {
	xval := reflect.ValueOf(x)
	xval = reflect.Indirect(xval)

	if xval.Kind() != reflect.Struct {
		fmt.Println("not a struct")
		return
	}

	typ := xval.Type()

	count := typ.NumField()

	for i := 0; i < count; i++ {
		field := typ.Field(i)
		name := field.Name
		tag := field.Tag.Get("json")
		fmt.Printf("%02d: %s, tag: %s\n", i, name, tag)
	}
}

func main() {
	test := &Test{}
	dig(test)
}

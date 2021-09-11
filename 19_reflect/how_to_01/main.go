package main

import (
	"fmt"
	"reflect"
)

type myinterface interface {
	Hello() string
}

type mystruct struct {
	name string
}

func (my *mystruct) Hello() string {
	return fmt.Sprintf("Hello, %s", my.name)
}

var myinterfaceType = reflect.TypeOf((*myinterface)(nil)).Elem()

func main() {

	my1 := &mystruct{name: "test1"}
	my2 := mystruct{name: "test2"}

	fmt.Printf("Type %T implements myinterface? %v\n", my1, reflect.TypeOf(my1).Implements(myinterfaceType))
	fmt.Printf("Type %T implements myinterface? %v\n", my2, reflect.TypeOf(my2).Implements(myinterfaceType))
}

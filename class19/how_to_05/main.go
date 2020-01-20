package main

import (
	"fmt"
	"reflect"
)

func sum1(a, b, c int) int {
	return a + b + c
}

func sum2(a int, x ...int) (int, int) {
	sum := 0
	for _, xx := range x {
		sum += xx
	}

	return len(x), a + sum
}

func sum3(x ...int) (int, int) {
	sum := 0
	for _, xx := range x {
		sum += xx
	}

	return len(x), sum
}

func power(base, exp int) int64 {
	p := int64(1)

	for i := 1; i <= exp; i++ {
		p *= int64(base)
	}
	return p
}

var powerFuncType = reflect.FuncOf(
	[]reflect.Type{reflect.TypeOf(0)},        // in
	[]reflect.Type{reflect.TypeOf(int64(0))}, // out
	false,
)

func makePowerOf(exp int) reflect.Value {

	return reflect.MakeFunc(
		powerFuncType,
		func(args []reflect.Value) []reflect.Value {
			base := int(args[0].Int())
			result := power(base, exp)
			return []reflect.Value{reflect.ValueOf(result)}
		},
	)
}

func callPower(base int, f reflect.Value) {

	input := []reflect.Value{
		reflect.ValueOf(base),
	}
	output := f.Call(input)
	fmt.Printf("base %d, result: %d\n", base, output[0].Int())

	if f.Type().ConvertibleTo(powerFuncType) {
		funcConv := f.Interface().(func(int) int64)
		fmt.Printf("convert and call: base %d, result %d\n", base, funcConv(base))
	}
}

func funcInfo(x interface{}) {
	t := reflect.TypeOf(x)
	isFunc := t.Kind() == reflect.Func
	fmt.Printf("%v is a function? %v\n", x, isFunc)

	if !isFunc {
		return
	}

	fmt.Printf("%v is a variadic function? %v\n", x, t.IsVariadic())

	fmt.Printf("\tnumber of input: %d\n", t.NumIn())

	for i := 0; i < t.NumIn(); i++ {
		input := t.In(i)
		fmt.Printf("\t\t%d %q %v\n", i, input.Name(), input.String())
	}

	fmt.Printf("\tnumber of output: %d\n", t.NumOut())
	for i := 0; i < t.NumOut(); i++ {
		output := t.Out(i)
		fmt.Printf("\t\t%d %q %v\n", i, output.Name(), output.String())
	}
}

func outputInfo(out []reflect.Value) {
	for i, x := range out {
		fmt.Printf("%d: %T %v\n", i, x.Interface(), x.Interface())
	}
}

func main() {

	funcInfo(sum1)
	funcInfo(sum2)
	funcInfo(100)

	params1 := []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(3),
		reflect.ValueOf(5),
	}

	params2 := []reflect.Value{
		reflect.ValueOf(2),
		reflect.ValueOf([]int{
			4, 6, 8,
		}),
	}

	params3 := []reflect.Value{
		reflect.ValueOf([]int{
			2, 4, 6, 8,
		}),
	}

	v1 := reflect.ValueOf(sum1)
	fmt.Println("call sum1 with reflect.Call")
	outputInfo(v1.Call(params1))

	v2 := reflect.ValueOf(sum2)
	fmt.Println("call sum2 with reflect.Call")
	outputInfo(v2.Call(params1))
	fmt.Println("call sum2 with reflect.CallSlice")
	outputInfo(v2.CallSlice(params2))

	v3 := reflect.ValueOf(sum3)
	fmt.Println("call sum3 with reflect.CallSlice")
	outputInfo(v3.CallSlice(params3))

	pow2 := makePowerOf(2)
	funcInfo(pow2.Interface())
	callPower(-3, pow2)

	pow3 := makePowerOf(3)
	funcInfo(pow3.Interface())
	callPower(-3, pow3)
}

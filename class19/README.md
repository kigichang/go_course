# 19 Reflection and Struct Tag

## Type Assertion

Type Assertion 主要是用來取出 interface{} 的值，類似 cast 行為。使用方式：

1. `value := i.(Data Type)`, eg: `f := i.(float64)`
1. `value, ok := i.(Data Type)`, eg: `f, ok := i.(float64)`

兩者的差別，前者如果資料型別不相符時，會 panic, 而後者不會，會回傳 zero-value, false。

eg:

```go
package main

import "fmt"

func main() {
    var i interface{} = "hello"

    s := i.(string)
    fmt.Println(s) // hello

    s, ok := i.(string)
    fmt.Println(s, ok) // hello true

    f, ok := i.(float64)
    fmt.Println(f, ok) // 0 false

    //f = i.(float64) // panic
    //fmt.Println(f)

    i = int64(100)

    f, ok = i.(float64)
    fmt.Println(f, ok) // 0 false
}
```

switch-type 可以讓 type assertion 程式碼更簡潔，使用方式: `switch v := i.(type)`。

1. 原本的 `i.(DATA_TYPE)` 改成 `i.(type)`
1. `i.(type)` 一定要配合 `switch` 使用，否則 compile 會失敗
1. `v` 會是 `interface{}` 所代表的值，因此可以直接使用，如下的: `v*2` or `v.Test()`

eg:

```go {.line-numbers}
package main

import "fmt"

type test struct{ name string }

func (t *test) Test() string {
    return fmt.Sprintf("%s:Test", t.name)
}

func do(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is %v bytes long\n", v, len(v))
    case test:
        fmt.Println("this is a test struct,", v.Test())
    case *test:
        fmt.Println("this is a pointer of test struct,", v.Test())
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

func main() {
    do(21)
    do("hello")
    do(true)
    do(&test{"pointer"})
    do(test{"struct"})
}
```

## Relfection TypeOf and ValueOf

reflect 常用的機制，主要有兩個 `Type` 及 `Value`，分別可以透過 `reflect.Type` 及 `reflect.Value` 取得。

1. `Type` 代表 `interface{}` 值的資料型別
1. `Value` 代表 `interface{}` 值

eg:

```go {.line-numbers}
// Indirect indirect returns the value, after dereferencing as many times
// as necessary to reach the base type (or nil).
// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
func Indirect(a interface{}) interface{} {
    if a == nil {
        return nil
    }
    if t := reflect.TypeOf(a); t.Kind() != reflect.Ptr {
        // Avoid creating a reflect.Value if it's not a pointer.
        return a
    }
    v := reflect.ValueOf(a)
    for v.Kind() == reflect.Ptr && !v.IsNil() {
        v = v.Elem()
    }
    return v.Interface()
}
```

使用 reflect 時，要特別注意 reflect 的 function 都會針對資料型別，誤用會發生 panic.

eg:

```go {.line-numbers}
package main

import (
    "reflect"
)

func main() {
    var i interface{} = 10

    reflect.ValueOf(i).Elem() // panic: reflect: call of reflect.Value.Elem on int Value
}
```

## Struct Tag

在 json, protobuf 都有用到 struct tag, 用來描述要如何處理 struct field.

```go {.line-numbers}
type User struct {
    ID    int    `validate:"-"`
    Name  string `validate:"presence,min=2,max=32"`
    Email string `validate:"email,required"`
    Time  time.Time
}
```

如何取得:

```go {.line-numbers}
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
```

1. 要取出 struct tag，必要用 `reflect.Type`: `typ := reflect.TypeOf(xx)`
1. 要取 struct tag 前，一定要先判斷傳入的 interface{} 是否是 struct: `typ.Kind() != reflect.Struct`
1. 透過 `Type.Field(i int)` 取得 `StructField`: `field := typ.Field(i)`
1. 過 `StructField.Tag` 取得 `StructTag`
1. 可以透過的 `Get` 與 `Lookup` 取得當初的設定: `field.Tag.Get(tagValidate)` or `field.Tag.Lookup(tagDefault)`

## Casting (spf13 Cast)

專門用來做資料轉型用。裏面有用到很多 type assertion 與 reflect，可以當學習範本。

[spf13 Cast](https://github.com/spf13/cast)

## How Tos

### Implements Some Interface or Not

```go {.line-numbers}
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
```

1. `var myinterfaceType = reflect.TypeOf((*myinterface)(nil)).Elem()`
    1. 先取得 pointer of interface 的 `reflect.Type`。
    1. 再使用 `reflect.Type.Elem` 取得 interface 的 `reflect.Type`。
1. 使用 `reflect.Type.Implements` 來判斷是否有實作該 interface。

### About Slice

```go {.line-numbers}
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
```

1. `reflect.MakeSlice` to make a slice.
1. `reflect.Value.Index` to get value in slice。
1. `reflect.Len` to get length of slice。
1. `reflect.Cap` to get capacity of slice。
1. `reflect.Append` to append single value to a slice.
1. `reflect.AppendSlice` to append two slices.

### Map

```go {.line-numbers}
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
```

1. `reflect.MakeMap` or `reflect.MakeMapWithSize` to make a map.
1. `reflect.Value.MapIndex` to get value with a key.
1. `reflect.Value.IsValid` to check map has the key.
1. `reflect.Value.SetMapIndex` to set (key, value) pair.
1. `reflect.Value.MapKeys` to get all keys in map to travel map.
1. `reflect.Value.MapRange` to get map iterator to travel map.

### Zero Value

```go {.line-numbers}
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
```

1. `reflect.Zero` to get ==Zero Value== of some type.
1. `reflect.Value.IsZero` to check value is ==Zero Value== of some type.
1. `reflect.Value.IsNil` to check value is ==nil==. It is only for reference type or **panic**.

### Function

```go {.line-numbers}
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
```

1. `reflect.Type.NumIn` to get number of input parameters.
1. `reflect.Type.NumOut` to get number of results.
1. `reflect.Type.In` to get parameter type.
1. `reflect.Type.Out` to get result type.
1. `reflect.Value.Call` to invoke a function.
    1. value must be a function or ==panic==
    1. `reflect.Value.Call` to invoke function including variadic one.
    1. `reflect.Type.IsVariadic` to verify function is variadic.
1. `reflect.Value.CallSlice` to invoke a variadic function.
    1. the last element in input parameters `[]reflect.Value` must be a slice.
1. `reflect.FuncOf` to make a function type used by `reflect.MakeFunc`.
1. `reflect.MakeFunc` to make a new function.
1. `reflect.Type.ConvertibleTo` to test source type can convert to another type.

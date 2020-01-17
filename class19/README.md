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

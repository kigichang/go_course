# 20 Reflection and Struct Tag

<!-- @import "[TOC]" {cmd="toc" depthFrom=1 depthTo=3 orderedList=false} -->

<!-- code_chunk_output -->

- [20 Reflection and Struct Tag](#20-reflection-and-struct-tag)
  - [0. 前言](#0-前言)
  - [1. Type Assertion](#1-type-assertion)
  - [2. TypeOf and ValueOf](#2-typeof-and-valueof)
  - [3. Struct Tag](#3-struct-tag)
  - [4. Casting (spf13 Cast)](#4-casting-spf13-cast)
  - [5. How Tos](#5-how-tos)
    - [5.1 Implements Some Interface or Not](#51-implements-some-interface-or-not)
    - [5.2 Slice](#52-slice)
    - [5.3 Map](#53-map)
    - [5.4 Zero Value](#54-zero-value)
    - [5.5 Function](#55-function)

<!-- /code_chunk_output -->

## 0. 前言

Reflection 對於程式初學者來說，不太能理解用途。大多數高階語言，如: Java, PHP, C#, Go 等，都有提供 Reflection 機制。主要是用在程式執行時期，解析輸入的資料，並產生相對應的結果。

以 Go 來說，[encoding/json](https://pkg.go.dev/encoding/json) 就是個實例，它提供程式人員，可以傳入任意值 (__interface{}__)，最終輸出 JSON 字串。其中如何解析傳入的任意值，並終產生對應的 JSON 字串，就是利用 Go Reflect 功能。

## 1. Type Assertion

Type Assertion 主要是用來取出 `interface{}` 的值，類似 __cast__ 行為。使用方式：

1. `value := i.(Data Type)`, eg: `f := i.(float64)`
1. `value, ok := i.(Data Type)`, eg: `f, ok := i.(float64)`

兩者的差別，前者如果資料型別不相符時，前者會 panic, 而後者不會，會回傳 __Zero Value__ 與 __false__。

eg:

@import "ex20_01/main.go" {class="line-numbers"}

__switch-type__ 可以讓 type assertion 程式碼更簡潔，使用方式: `switch v := i.(type)`。

eg:

@import "ex20_02/main.go" {class="line-numbers" highlight="12,14,18"}

1. 原本的 `i.(DATA_TYPE)` 改成 `i.(type)`
1. `i.(type)` 一定要配合 `switch` 使用，否則 compile 會失敗。
1. `v` 會是 `interface{}` 所代表的值，因此可以直接使用，如下的: `v*2` or `v.Test()`

## 2. TypeOf and ValueOf

reflect 常用的機制，主要有兩個 `Type` 及 `Value`，分別可以透過 `reflect.TypeOf` 及 `reflect.ValueOf` 取得。

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

使用 reflect 時，要特別注意資料型別是否正確，誤用會發生 panic。如：不是 Pointer 卻呼叫 `Elem()`。

@import "ex20_03/main.go" {class="line-numbers" highlight="10"}

## 3. Struct Tag

在 JSON, Protobuf 都有用到 struct tag，用來描述要如何處理 struct field。如：

```go {.line-numbers}
type User struct {
    ID    int    `validate:"-"`
    Name  string `validate:"presence,min=2,max=32"`
    Email string `validate:"email,required"`
    Time  time.Time
}
```

如何解析 Struct 並取得 Struct Tag：

@import "ex20_04/main.go" {class="line-numbers" highlight="19,24,29,31"}

1. 要取 struct tag 前，一定要先判斷傳入的 interface{} 是否是 struct: `typ.Kind() != reflect.Struct`
1. 要取出 struct tag，必要用 `reflect.Type`: `typ := reflect.TypeOf(xx)`
1. 透過 `Type.Field(i int)` 取得 `StructField`: `field := typ.Field(i)`
1. 過 `StructField.Tag` 取得 `StructTag`
1. 可以透過的 `Get` 與 `Lookup` 取得當初的設定: `field.Tag.Get(tagValidate)` or `field.Tag.Lookup(tagDefault)`

## 4. Casting (spf13 Cast)

專門用來做資料轉型用。裏面有用到很多 type assertion 與 reflect，可以當學習範本。

[spf13 Cast](https://github.com/spf13/cast)

## 5. How Tos

以下舉例我之前寫過的程式。

### 5.1 Implements Some Interface or Not

如何判斷是否有實作某個 interface。

@import "how_to_01/main.go" {class="line-numbers" highlight="20,27,28"}

1. 先取得已定義 interface 的 `Type`。
    1. `var myinterfaceType = reflect.TypeOf((*myinterface)(nil)).Elem()`
    1. 先取得 Pointer of Interface 的 `reflect.Type`。
    1. 再使用 `reflect.Type.Elem` 取得 interface 的 `reflect.Type`。
1. 使用 `reflect.Type.Implements` 來判斷是否有實作該 interface。

### 5.2 Slice

如何透過 reflect 產生 slice。

@import "how_to_02/main.go" {class="line-numbers"}

1. `reflect.MakeSlice` to make a slice.
1. `reflect.Value.Index` to get value in slice。
1. `reflect.Len` to get length of slice。
1. `reflect.Cap` to get capacity of slice。
1. `reflect.Append` to append single value to a slice.
1. `reflect.AppendSlice` to append two slices.

### 5.3 Map

如何透過 reflect 產生 map。

@import "how_to_03/main.go" {class="line-numbers"}

1. `reflect.MakeMap` or `reflect.MakeMapWithSize` to make a map.
1. `reflect.Value.MapIndex` to get value with a key.
1. `reflect.Value.IsValid` to check map has the key.
1. `reflect.Value.SetMapIndex` to set (key, value) pair.
1. `reflect.Value.MapKeys` to get all keys in map to travel map.
1. `reflect.Value.MapRange` to get map iterator to travel map.

### 5.4 Zero Value

如何取得某型別的 Zero Value，以及判斷是否為 Zero Value 或 Nil。

@import "how_to_04/main.go" {class="line-numbers"}

1. `reflect.Zero` to get Zero Value of some type.
1. `reflect.Value.IsZero` to check value is Zero Value of some type.
1. `reflect.Value.IsNil` to check value is nil. It is only for reference type or __panic__.

### 5.5 Function

如何解析 Function，呼叫 Function，以及創造一個新的 Function。

@import "how_to_05/main.go" {class="line-numbers"}

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

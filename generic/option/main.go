package main

import (
	"fmt"
	"strconv"
	"reflect"
)

type Option[T any] interface {
	IsDefined() bool
	Get() T
}

type option[T any] struct {
	defined bool
	v T
}

// IsDefined 如果是 Some 則回傳 true, 否則回傳 false。
func (o *option[T]) IsDefined() bool {
	return o.defined
}

// Get 取值，如果是 None 會發生 panic。
func (o *option[T]) Get() T {
	if o.defined {
		return o.v
	}
	panic(fmt.Sprintf("can not get value from None[%s]", reflect.TypeOf(o.v)))
}

// Some 回傳 Some，包含指定的值。
func Some[T any](v T) Option[T] {
	return &option[T] {
		defined: true,
		v: v,
	}
}

// None 回傳 None，內含的值會是 Zero Value。
func None[T any]() Option[T] {
	return &option[T] {
		defined: false,
	}
}

func OptionMap[T, U any](opt Option[T], f func(T) U) Option[U] {
	if opt.IsDefined() {
		return Some[U](f(opt.Get()))
	}
	return None[U]()
}


// OptionEquals 比較兩個 Option 是否相同，T 必須是 comparable。
func OptionEquals[T comparable](opt Option[T], that Option[T]) bool {
	if opt.IsDefined() == that.IsDefined() {
		return !opt.IsDefined() || opt.Get() == that.Get()
	}
	return false
}

func main() {
	
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	opt := Some[int](0)
	fmt.Println("defined:", opt.IsDefined())
	fmt.Println("value:", opt.Get(), reflect.TypeOf(opt.Get()))

	optStr := OptionMap(opt, strconv.Itoa)
	fmt.Println("defined:", optStr.IsDefined())
	fmt.Println("value:", optStr.Get(), reflect.TypeOf(optStr.Get()))

	opt = None[int]()
	fmt.Println("defined:", opt.IsDefined())
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println("value:", opt.Get())
	}()
	
	optStr = OptionMap(opt, strconv.Itoa)
	fmt.Println("defined:", optStr.IsDefined())
	func () {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()
		fmt.Println("value:", optStr.Get())
	}()
	

	fmt.Println(OptionEquals(Some(1), Some(1)))
	fmt.Println(OptionEquals(Some(1), Some(0)))
	fmt.Println(OptionEquals(Some(1), None[int]()))
	fmt.Println(OptionEquals(None[int](), None[int]()))


	//fmt.Println(OptionEquals(None[int64](), None[int]())) // type Option[int] of None[int]() does not match inferred type Option[int64] for Option[T]
	//fmt.Println(OptionEquals(None[[]int64](), None[[]int64]())) // []int64 does not satisfy comparable
}
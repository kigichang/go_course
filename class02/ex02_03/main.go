package main

import "fmt"

// Celsius ...
type Celsius = float64

// ToF convert Celsius to Fahrenheit
//func (c Celsius) ToF() Fahrenheit { //  cannot define new methods on non-local type float64
//	return CToF(c)
//}

// Fahrenheit ...
type Fahrenheit = float64

// ToC convert Celsius to Fahrenheit
//func (f Fahrenheit) ToC() Celsius { //  cannot define new methods on non-local type float64
//	return FToC(f)
//}

// const variable
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return c*9/5 + 32 }

// FToC convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return (f - 32) * 5 / 9 }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // 100
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // 180
	fmt.Printf("%g\n", boilingF-FreezingC)       // 212

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // true
	fmt.Println(f >= 0)          // true
	fmt.Println(c == Celsius(f)) // true
	fmt.Println(c == f)          // true
}

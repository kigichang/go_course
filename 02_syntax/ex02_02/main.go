package main

import "fmt"

// Celsius ...
type Celsius float64

// ToF convert Celsius to Fahrenheit
func (c Celsius) ToF() Fahrenheit {
	return CToF(c)
}

// Fahrenheit ...
type Fahrenheit float64

// ToC convert Celsius to Fahrenheit
func (f Fahrenheit) ToC() Celsius {
	return FToC(f)
}

// const variable
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF convert Celsius to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC convert Fahrenheit to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func main() {
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := BoilingC.ToF()
	fmt.Printf("%g\n", boilingF-FreezingC.ToF()) // "180" °F
	//fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch

	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)          // "true"
	fmt.Println(f >= 0)          // "true"
	fmt.Println(c == Celsius(f)) // "true"!
	//fmt.Println(c == f)          // compile error: type mismatch
}

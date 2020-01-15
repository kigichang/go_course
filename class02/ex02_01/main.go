package main

import "fmt"

// number
const (
	Zero  = iota
	One   = iota
	Two   = iota
	Three = iota
)

// file mode
const (
	X = 1 << iota
	W = 1 << iota
	R = 1 << iota
)

// size
const (
	_          = iota // ignore first value by assigning to blank identifier
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// weekday
const (
	Sunday = 1 + iota
	_
	// this is a comment

	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	var x int64 = 123_456_789
	var y float64 = 12_345.678_9
	fmt.Println(x, y)
	fmt.Println(Zero, One, Two, Three)
	fmt.Println(X, W, R)
	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

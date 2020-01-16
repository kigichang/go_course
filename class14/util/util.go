package util

import "fmt"

func init() {
	fmt.Println("util init")
}

const (
	defaultSum = 0 // package util_test 無法存取
)

// Sum ...
func Sum(x ...int) int {
	s := 0

	for _, v := range x {
		s += v
	}

	return s
}

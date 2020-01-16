package util_test

import (
	"fmt"
	"os"
	"testing"

	. "util"
)

func TestSum(t *testing.T) {

	x := Sum(1, 2, 3, 4, 5)

	if x != 15 {
		t.Fatal("sum error")
	}
}

func TestMain(m *testing.M) {
	// initialize test resource

	exitCode := m.Run()

	// destroy test resource

	os.Exit(exitCode)
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func ExampleSum() {
	fmt.Println("hello world")

	fmt.Println(Sum(1, 2, 3))
	// Output:
	// hello world
	// 6
}

func ExampleHello() {
	fmt.Println("hello world")

	fmt.Println(Sum(1, 2, 3))
	// Unordered output:
	// 6
	// hello world
}

//go:build !windows

package main

import "fmt"

func Say() {
	fmt.Println("I am not Windows")
}

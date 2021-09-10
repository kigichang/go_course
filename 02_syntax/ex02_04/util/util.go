package util

import "fmt"

func init() {
	fmt.Println("package util initialize")
}

// Hello returns string concates
func Hello(name string) string {
	return fmt.Sprintf("hello %s", name)
}

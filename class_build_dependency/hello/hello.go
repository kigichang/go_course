package hello

import "fmt"

// Hello ....
func Hello(name string) string {
	if name != "" {
		return fmt.Sprintf("Hello, %s", name)
	}

	return "Who are you?"

}

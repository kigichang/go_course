package main

import "fmt"

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	//ages := make(map[string]int) // mapping from strings to ints

	ages["alice"] = 32 // alice = 32
	fmt.Println("alice:", ages["alice"])
	ages["alice"]++ // alice = 33
	fmt.Println("alice:", ages["alice"])
	delete(ages, "cat")

	fmt.Println(ages["bob"]) // 0 (zero-value)

	a, ok := ages["bob"]
	fmt.Println(a, ok) // 0, false
}

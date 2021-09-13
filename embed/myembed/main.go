package main

import (
	"embed"
	"fmt"
)

//go:embed hello.txt
var s string

//go:embed hello.txt
var b []byte

//go:embed assets/*
//go:embed hello.txt
var content embed.FS

func main() {
	fmt.Printf("s is [%s]\n", s)
	fmt.Printf("b is [%s]\n", string(b))

	entries, err := content.ReadDir("assets")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range entries {
		fmt.Println(entries[i].Name())
	}

	a, err := content.ReadFile("assets/a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("assets/a.txt:", string(a))

	b, err := content.ReadFile("assets/b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("assets/b.txt:", string(b))

	_, err = content.ReadFile("assets/c.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
}

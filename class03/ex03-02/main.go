package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const nihongo = "日本語"
	for i := 0; i < len(nihongo); i++ {
		fmt.Printf("%d: %x\n", i, nihongo[i])
	}

	for index, runeValue := range nihongo {
		fmt.Printf("%U starts at byte position %d\n", runeValue, index)
	}

	fmt.Println(utf8.RuneCountInString(nihongo)) // 取 utf8 長度

	bytes1 := []byte(nihongo) // convert to byte slice.
	fmt.Println("bytes: ", bytes1)
	fmt.Println(string(bytes1)) // convert to string from byte slice.
}

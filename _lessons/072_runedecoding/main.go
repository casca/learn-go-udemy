package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	text := "öykü"

	for i := 0; i < len(text); {
		r, size := utf8.DecodeRuneInString(text[i:])
		fmt.Printf("%c", r)

		i += size
	}
	fmt.Println()

	for _, r := range text {
		fmt.Printf("%c", r)
	}
	fmt.Println()

	word := []byte("öykü")

	fmt.Printf("%s = % [1]x\n", word)

	// var size int
	// for i := range string(word) {
	// 	if i > 0 {
	// 		size = i
	// 		break
	// 	}
	// }
	_,size:=utf8.DecodeRune(word)
	copy(word[:size], bytes.ToUpper(word[:size]))
	fmt.Printf("%s = % [1]x\n", word)
}

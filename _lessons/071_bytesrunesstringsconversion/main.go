package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

func init() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line
}

func main() {
	str := "Yugeñ ☯ 😌"

	// str[0] = 'N'
	// str[1] = 'O'

	bytes := []byte(str)
	// bytes[0] = 'N'
	// bytes[1] = 'O'

	str = string(bytes)

	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", len(str))
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str))

	fmt.Printf("% x\n", bytes)
	fmt.Printf("\t%d bytes\n", len(bytes))
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))

	fmt.Println()
	for i, r := range str {
		// fmt.Printf("str[%2d] = %-2x\n", i, str[i])
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}

	fmt.Println()
	fmt.Printf("1st byte   : %c\n", str[0])
	fmt.Printf("5th byte   : %c\n", str[4])
	fmt.Printf("5th rune   : %s\n", str[4:6])
	fmt.Printf("last rune  : %s\n", str[len(str)-4:])

	runes := []rune(str)

	fmt.Println()
	fmt.Printf("%s\n", string(runes))
	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes))
	fmt.Printf("\t%d runes\n", len(runes))
	fmt.Printf("1st rune   : %c\n", runes[0])
	fmt.Printf("5th rune   : %c\n", runes[4])
	fmt.Printf("first five : %c\n", runes[:5])

}

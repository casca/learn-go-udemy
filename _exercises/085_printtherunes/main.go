package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print the runes
//
//  1. Loop over the "console" word and print its runes one by one,
//     in decimals, hexadecimals and binary.
//
//  2. Manually put the runes of the "console" word to a byte slice, one by one.
//
//     As the elements of the byte slice use only the rune literals.
//
//     Print the byte slice.
//
//  3. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only decimal numbers.
//
//  4. Repeat the step 2 but this time, as the elements of the byte slice,
//     use only hexadecimal numbers.
//
//
// EXPECTED OUTPUT
//   Run the solution to see the expected output.
// ---------------------------------------------------------

func main() {
	const word = "console"

	for _, r := range word {
		fmt.Printf("%[1]d % [1]x %[1]b\n", r)
	}

	bs := []byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}
	fmt.Println("with runes", string(bs))

	bs = []byte{99, 111, 110, 115, 111, 108, 101}
	fmt.Println("with decimals", string(bs))

	bs = []byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}
	fmt.Println("with hexadecimals", string(bs))

}

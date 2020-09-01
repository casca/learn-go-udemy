package main

import (
	"bufio"
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Unique Words
//
//  Create a program that prints the total and unique words
//  from an input stream.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Scan the input using a new Scanner.
//
//  3. Configure the scanner to scan for the words.
//
//  4. Count the unique words using a map.
//
//  5. Print the total and unique words.
//
//
// EXPECTED OUTPUT
//
//  There are 99 words, 70 of them are unique.
//
// ---------------------------------------------------------

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	// var words []string
	// seen := make(map[string]bool)

	// for in.Scan() {
	// 	w := in.Text()
	// 	if seen[w] {
	// 		continue
	// 	}
	// 	words = append(words, w)
	// 	seen[w] = true
	// }

	// for _, w := range words {
	// 	fmt.Println(w)
	// }
	// fmt.Println("length:", len(words))

	total, words := 0, make(map[string]int)
	for in.Scan() {
		total++
		words[in.Text()]++
	}
	fmt.Printf("There are %d words, %d of them are unique.\n",
		total, len(words))
}

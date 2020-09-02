package main

import "fmt"

func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		// title string
		// words int

		// text text

		text  // anonymous field
		isbn  string
		title string
	}

	moby := book{
		text: text{title: "moby dick", words: 206502},
		isbn: "102030",
	}

	moby.text.words = 1000
	moby.words++

	fmt.Printf("%s has %d words (isbn: %s)\n", moby.text.title, moby.text.words, moby.isbn)
	fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

	fmt.Printf("%#v\n", moby)
}

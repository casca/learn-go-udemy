package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/11-if/exercises/01-age-seasons/main.go
	age := 10
	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}

	// https://github.com/inancgumus/learngo/blob/master/11-if/exercises/02-simplify-it/main.go
	isSphere, radius := true, 200
	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}

	// https://github.com/inancgumus/learngo/blob/master/11-if/exercises/03-arg-count/main.go
	var (
		args = os.Args
		l    = len(args) - 1
	)
	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is one: %q\n", args[1])
	} else if l == 2 {
		fmt.Printf("There are two: \"%s %s\"\n", args[1], args[2])
	} else {
		fmt.Printf("There are %d arguments\n", l)
	}

	// https://github.com/inancgumus/learngo/blob/master/11-if/exercises/04-vowel-or-cons/main.go
	if l == 0 || len(args[1]) != 1 {
		fmt.Println("Give me a letter")
	}

	c := args[1]
	if c == "y" || c == "w" {
		fmt.Printf("%q is sometimes a vowel, sometimes not.\n", c)
	} else if strings.IndexAny(c, "aeiou") != -1 {
		fmt.Printf("%q is is a vowel.\n", c)
	} else {
		fmt.Printf("%q is a consonant.\n", c)
	}

}

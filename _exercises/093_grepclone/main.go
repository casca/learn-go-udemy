package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Grep Clone
//
//  Create a grep clone. grep is a command-line utility for
//  searching plain-text data for lines that match a specific
//  pattern.
//
//  1. Feed the shakespeare.txt to your program.
//
//  2. Accept a command-line argument for the pattern
//
//  3. Only print the lines that contains that pattern
//
//  4. If no pattern is provided, print all the lines
//
//
// EXPECTED OUTPUT
//
//  go run main.go come < shakespeare.txt
//
//    come night come romeo come thou day in night
//    come gentle night come loving black-browed night
//
// ---------------------------------------------------------

func main() {
	// args := os.Args[1:]
	// var pattern *regexp.Regexp
	// if len(args) == 1 {
	// 	pattern = regexp.MustCompile(args[0])
	// }

	// in := bufio.NewScanner(os.Stdin)

	// for in.Scan() {
	// 	line := in.Text()
	// 	if pattern != nil && !pattern.MatchString(line) {
	// 		continue
	// 	}
	// 	fmt.Println(line)
	// }

	in := bufio.NewScanner(os.Stdin)

	var pattern string
	if args := os.Args[1:]; len(args) == 1 {
		pattern = args[0]
	}

	for in.Scan() {
		s := in.Text()
		if strings.Contains(s, pattern) {
			fmt.Println(s)
		}
	}
}

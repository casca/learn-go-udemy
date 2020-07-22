package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/01-windows-path/main.go
	path := `c:\program files\duper super\fun.txt
c:\program files\really\funny.png`
	fmt.Println(path)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/02-print-json/main.go
	json := `
{
        "Items": [{
                "Item": {
                        "name": "Teddy Bear"
                }
        }]
}`
	fmt.Println(json)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/03-raw-concat/main.go
	// uncomment the code below
	// name := "and get the name from the command-line"

	// replace and concatenate the `name` variable
	// after `hi ` below

	name := os.Args[1]
	msg := `hi ` + name + `!
how are you?`

	fmt.Println(msg)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/04-count-the-chars/main.go
	length := utf8.RuneCountInString(os.Args[1])
	fmt.Println(length)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/05-improved-banger/main.go
	msg = os.Args[1]
	s := msg + strings.Repeat("!", utf8.RuneCountInString(msg))
	fmt.Println(s)

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/06-tolowercase/main.go
	msg = os.Args[1]
	fmt.Println(strings.ToLower(msg))

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/07-trim-it/main.go
	msg = `
	
	The weather looks good.
I should go and play.
	`

	fmt.Println(strings.TrimSpace(msg))

	// https://github.com/inancgumus/learngo/blob/master/08-numbers-and-strings/02-strings/exercises/08-right-trim-it/main.go
	name = "inanç           "
	fmt.Println(utf8.RuneCountInString(strings.TrimRight(name, " ")))
}

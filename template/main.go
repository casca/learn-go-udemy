package main

import s "github.com/inancgumus/prettyslice"

func init() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line
}

func main() {
}

package main

import (
	"fmt"

	"github.com/casca/learn-go-udemy/_exercises/072_limitbackingarraysharing/api"
	s "github.com/inancgumus/prettyslice"
)

func init() {
	s.PrintBacking = true // prints the backing array
	s.MaxPerLine = 10     // prints 10 slice elements per line
	s.Width = 60          // prints 60 character per line
}

func main() {
	// DO NOT CHANGE ANYTHING IN THIS CODE.

	// get the first three elements from api.temps
	received := api.Read(0, 3)

	// append changes the api package's temps slice's
	// backing array as well.
	received = append(received, []int{1, 3}...)

	fmt.Println("api.temps     :", api.All())
	fmt.Println("main.received :", received)
}

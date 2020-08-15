package main

import (
	"fmt"
	"unsafe"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	var games []string
	s.Show("games", games)

	games = []string{}
	s.Show("games", games)
	s.Show("another empty", []int{}) // same backing array as games

	games = []string{"pacman", "mario", "tetris", "doom"}
	s.Show("games", games)
	s.Show("games", []string{"one", "more"})

	part := games
	s.Show("part", part)

	part = games[:0]
	s.Show("part[:0]", part)
	s.Show("part[:cap]", part[:cap(part)])

	for cap(part) != 0 {
		s.Show("part", part)
		part = part[1:cap(part)]
	}

	fmt.Println(unsafe.Sizeof(""))

	s.Show("games", games)
	s.Show("part", part)
}

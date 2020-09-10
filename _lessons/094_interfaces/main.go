package main

import "fmt"

func main() {
	var (
		mobydick  = book{title: "mobydick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
		rubik     = puzzle{title: "rubik's cube", price: 5}
	)

	var store list
	store = append(store, &minecraft, &tetris, mobydick, rubik)
	store.print()

	fmt.Println(store[0] == &minecraft)
	fmt.Println(store[3] == rubik)
}

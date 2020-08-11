package main

import "fmt"

func main() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	newBooks := [5]string{"ulysses", "fire"}
	books = newBooks
	// books[0] = "foo" // books gets a copy of the array

	// var games []string
	games := []string{"kokemon", "sims"}
	newGames := []string{"pacman", "doom", "pong"}

	newGames = games

	games = []string{}

	// if games == newGames {
	// }

	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	// if games == nil {
	if len(games) != len(newGames) {
		ok = "not "
	}

	fmt.Printf("games and newGames are %sequal\n\n", ok)

	fmt.Printf("books: %#v\n", books)
	fmt.Printf("newBooks: %#v\n", newBooks)
	fmt.Printf("games: %#v\n", games)
	fmt.Printf("newGames: %#v\n", newGames)
	fmt.Printf("games: %T\n", games)
	fmt.Printf("games: %d\n", len(games))
	fmt.Printf("nil?: %t\n", games == nil)
}

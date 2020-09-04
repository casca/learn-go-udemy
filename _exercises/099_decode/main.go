package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

type jsonGame struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	Price int    `json:"price"`
}

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func main() {
	var jsonGames []jsonGame
	if err := json.Unmarshal([]byte(data), &jsonGames); err != nil {
		fmt.Println(err)
		return
	}

	var games []game
	for _, jsonGame := range jsonGames {
		games = append(games, game{
			item: item{
				id:    jsonGame.ID,
				name:  jsonGame.Name,
				price: jsonGame.Price,
			},
			genre: jsonGame.Genre,
		})
	}
	// games := []game{
	// 	game{item{1, "god of war", 50}, "action adventure"},
	// 	game{item{2, "x-com 2", 30}, "strategy"},
	// 	game{item{3, "minecraft", 20}, "sandbox"},
	// }

	catalog := make(map[int]game)
	for _, g := range games {
		catalog[g.id] = g
	}

	fmt.Printf("Milo's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : export games to JSON
  > quit   : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.id, g.name, "("+g.genre+")", g.price)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := catalog[id]
			if !ok {
				fmt.Println("sorry. i don't have the game")
				continue
			}

			fmt.Printf("#%d: %-15q %-20s $%d\n",
				g.id, g.name, "("+g.genre+")", g.price)

		case "save":
			var jsonGames []jsonGame
			for _, game := range games {
				jsonGames = append(jsonGames, jsonGame{
					ID:    game.id,
					Name:  game.name,
					Genre: game.genre,
					Price: game.price,
				})
			}
			out, err := json.MarshalIndent(jsonGames, "", "\t")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(out))
			return
		}
	}
}

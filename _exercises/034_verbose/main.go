package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	maxTurns = 5
	usage    = `Welcome to the Lucky Number Game! 🍀

The program will pick %d random numbers.
Your mission is to guess one of those numbers.

The greater your number is, the harder it gets.

Wanna play?
`
)

func main() {
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var verbose bool

	if os.Args[1] == "-v" {
		verbose = true
	}

	guess, err := strconv.Atoi(os.Args[len(os.Args)-1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	minGuess := guess
	if guess < 10 {
		minGuess = 10
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(minGuess + 1)
		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			fmt.Println("🎉 YOU WIN!")
			return
		}
	}
	fmt.Println("☠ YOU LOST... Try again?")
}

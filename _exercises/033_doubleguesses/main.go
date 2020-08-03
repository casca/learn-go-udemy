package main

import (
	"fmt"
	"math"
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
	if len(args) != 2 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	guess2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess1 < 0 || guess2 < 0 {
		fmt.Println("Please pick a positive number.")
		return
	}

	max := int(math.Max(float64(guess1), float64(guess2)))
	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(max + 1)

		if n == guess1 || n == guess2 {
			fmt.Println("🎉 YOU WIN!")
			return
		}
	}
	fmt.Println("☠ YOU LOST... Try again?")
}

package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const usage = "[your name] [positive|negative]"

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println(usage)
		return
	}

	name, mood := args[0], args[1]

	if mood != "positive" && mood != "negative" {
		fmt.Println(usage)
		return
	}

	moods := [...][2]string{
		{"happy 😊", "broccoli 🥦"},
		{"sad 😞", "rain 🌧"},
	}

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(moods[0]))

	var mi int
	if mood != "positive" {
		mi = 1
	}

	fmt.Printf("%s feels %s\n", name, moods[mi][n])
}

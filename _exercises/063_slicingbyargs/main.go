package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	ships := []string{"Normandy", "Verrikan", "Nexus", "Warsaw"}

	// args := os.Args[1:]
	// if len(args) < 1 || len(args) > 2 {
	// 	fmt.Println("Provide only the [starting] and [stopping] positions")
	// 	return
	// }

	// from, _ := strconv.Atoi(args[0])
	// to := len(ships)
	// if len(args) == 2 {
	// 	to, _ = strconv.Atoi(args[1])
	// }

	from, to := 0, len(ships)
	switch poss := os.Args[1:]; len(poss) {
	default:
		fallthrough
	case 0:
		fmt.Println("Provide only the [starting] and [stopping] positions")
		return
	case 2:
		to, _ = strconv.Atoi(poss[1])
		fallthrough
	case 1:
		from, _ = strconv.Atoi(poss[0])

	}

	if l := len(ships); from < 0 || from > l || to > l || from > to {
		fmt.Println("Wrong positions")
		return
	}

	fmt.Printf("%q\n", ships)
	fmt.Printf("%q\n", ships[from:to])

}

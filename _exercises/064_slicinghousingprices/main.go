package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	const (
		data = `Location,Size,Beds,Baths,Price
New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	rows := strings.Split(data, "\n")
	header, records := strings.Split(rows[0], separator), rows[1:]

	from, to := 0, len(header)
	switch poss := os.Args[1:]; len(poss) {
	case 2:
		to = findIndex(poss[1], header, to) + 1
		fallthrough
	case 1:
		from = findIndex(poss[0], header, from)
	}

	if from > to {
		from = 0
	}

	for _, h := range header[from:to] {
		fmt.Printf("%-12s", h)
	}
	fmt.Println()

	for _, r := range records {
		for _, v := range strings.Split(r, separator)[from:to] {
			fmt.Printf("%-12v", v)
		}
		fmt.Println()
	}
}

func findIndex(input string, header []string, fallback int) int {
	for i, h := range header {
		if input == h {
			return i
		}
	}
	return fallback
}

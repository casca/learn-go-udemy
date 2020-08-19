package main

import (
	"fmt"
	"strconv"
	"strings"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	// spendings := [][]int{
	// 	{200, 100},
	// 	{500},
	// 	{50, 25, 75},
	// }

	// spendings := make([][]int, 0, 5)
	// spendings = append(spendings, []int{200, 100})
	// spendings = append(spendings, []int{25, 10, 45, 60})
	// spendings = append(spendings, []int{5, 15, 35})
	// spendings = append(spendings, []int{95, 10}, []int{50, 25})
	spendings := fetch()

	for i, daily := range spendings {
		var total int

		for _, spending := range daily {
			total += spending
		}
		fmt.Printf("Day %d: %d\n", i+1, total)
	}
}

func fetch() [][]int {
	content := `200 100
25 10 45 60
5 15 35
95 10
50 25`

	lines := strings.Split(content, "\n")

	spendings := make([][]int, len(lines))

	for i, line := range lines {
		// fmt.Printf("%d, %#v\n", i+1, line)

		fields := strings.Fields(line)

		spendings[i] = make([]int, len(fields))

		for j, field := range fields {
			// fmt.Printf("\t%d, %#v\n", j+1, field)

			spending, _ := strconv.Atoi(field)

			spendings[i][j] = spending
		}
	}

	return spendings
}

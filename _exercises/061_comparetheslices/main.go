package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	namesA := "Da Vinci, Wozniak, Carmack"
	namesC := strings.Split(namesA, ", ")
	sort.Strings(namesC)

	namesB := []string{"Wozniak", "Da Vinci", "Carmack"}
	sort.Strings(namesB)

	if len(namesC) != len(namesB) {
		fmt.Println("not equal")
	}

	for i := range namesC {
		if namesC[i] != namesB[i] {
			fmt.Println("not equal")
			return
		}
	}
	fmt.Println("equal")
}

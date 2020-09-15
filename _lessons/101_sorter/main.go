package main

import (
	"fmt"
	"sort"
)

func main() {
	l := list{
		{title: "mobydick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
	}

	// sort.Sort(l)
	// sort.Sort(sort.Reverse(l))
	// sort.Sort(byReleaseDate(l))
	sort.Sort(sort.Reverse(byReleaseDate(l)))

	fmt.Print(l)

}

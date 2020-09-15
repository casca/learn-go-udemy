package main

import (
	"sort"
	"strings"
)

type list []*product

func (l list) String() string {
	if len(l) == 0 {
		return "Sorry. We're waiting for delivery. 🚛\n"
	}

	var str strings.Builder
	for _, p := range l {
		// fmt.Printf("(%-10T) --> ", p)
		// p.print()
		// fmt.Printf("* %s\n", p)
		str.WriteString("* ")
		str.WriteString(p.String())
		str.WriteRune('\n')
	}
	return str.String()
}

func (l list) discount(ratio float64) {
	type discounter interface{ discount(float64) }

	// for _, it := range l {
	// 	g, ok := it.(discounter)
	// 	// fmt.Printf("%T game? %v\n", it, ok)
	// 	if !ok {
	// 		continue
	// 	}
	// 	g.discount(ratio)
	// }

	for _, it := range l {
		// if it, ok := it.(discounter); ok {
		it.discount(ratio)
		// }
	}
}

func (l list) Len() int {
	return len(l)
}

func (l list) Less(i, j int) bool {
	return l[i].Title < l[j].Title
}

func (l list) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

type byRelease struct {
	list
}

func (br byRelease) Less(i, j int) bool {
	return br.list[i].Released.Before(br.list[j].Released.Time)
}

func byReleaseDate(l list) sort.Interface {
	return &byRelease{l}
}

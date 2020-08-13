package main

import (
	fmt
	s "github.com/inancgumus/prettyslice"
)

func main() {
	items := []string{"foo", "bar", "baz", "qux", "flo", "par", "qe"}

	const pageSize = 4

	l := len(items)
	for from := 0; from < l; from += pageSize {
		to := from + pageSize
		if to > l {
			to = l
		}
		currentPage := items[from:to]
		head := fmt.Sprinf("Page #%d", (from/pageSize)+1)
		s.Show(head, currentPage)
	}
	// s.Show("0:2", items[0:2])
	// s.Show("0:2", items[2:4])
	// s.Show("0:2", items[4:6])
	// s.Show("0:2", items[6:7])

}

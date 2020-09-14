package main

func main() {
	store := list{
		&book{product{"mobydick", 10}, 118281600},
		&book{product{"odyssey", 15}, "733622400"},
		&book{product{"hobbit", 25}, nil},
		&puzzle{product{"rubik's cube", 5}},
		&game{product{"minecraft", 20}},
		&game{product{"tetris", 5}},
		&toy{product{"yoda", 50}},
	}

	store.discount(.5)
	store.print()

	// t := &toy{product{"yoda", 150}}
	// fmt.Printf("%#v\n", t)

	// b := &book{product{"mobydick", 10}, 118281600}
	// fmt.Printf("%#v\n", b)
}

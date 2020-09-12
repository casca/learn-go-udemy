package main

func main() {
	store := list{
		book{title: "mobydick", price: 10, published: 118281600},
		book{title: "odyssey", price: 15, published: "733622400"},
		book{title: "hobbit", price: 25},
		puzzle{title: "rubik's cube", price: 5},
		&game{title: "minecraft", price: 20},
		&game{title: "tetris", price: 5},
		&toy{title: "yoda", price: 50},
	}

	store.discount(.5)
	store.print()

	// var p printer

	// p = mobydick
	// p = rubik
	// p = &minecraft

	// p = &tetris
	// tetris.discount(.5)
	// // p.discount(.5)

	// p.print()
}

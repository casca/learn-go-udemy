package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	// game.discount(minecraft, .5)
	// game.print(minecraft)
	// game.print(tetris)
	// book.print(mobydick)

	minecraft.discount(.5)

	mobydick.print()
	minecraft.print()
	tetris.print()

	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}

	var items []*game
	items = append(items, &minecraft, &tetris)

	my := list(items)
	my = nil
	my.print()
}

package main

func main() {
	l := list{
		{title: "mobydick", price: 10, released: toTimestamp(118281600)},
		{title: "odyssey", price: 15, released: toTimestamp("733622400")},
		{title: "hobbit", price: 25},
		// {title: "rubik's cube", price: 5},
		// {title: "minecraft", price: 20},
		// {title: "tetris", price: 5},
		// {title: "yoda", price: 50},
	}

	l.discount(.5)
	l.print()

	// t := &toy{product{"yoda", 150}}
	// fmt.Printf("%#v\n", t)

	// b := &book{product{"mobydick", 10}, 118281600}
	// fmt.Printf("%#v\n", b)
}

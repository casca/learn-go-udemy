package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var null *computer

	// compare the pointer variable to a nil value, and say it's nil
	if null == nil {
		fmt.Println("null is nil")
	}

	// create an apple computer by putting its address to a pointer variable
	ac := &computer{"Apple"}

	// put the apple into a new pointer variable
	yaac := ac

	// compare the apples: if they are equal say so and print their addresses
	if yaac == ac {
		fmt.Printf("yaac=%p\tac=%p\n", yaac, ac)
	}

	// create a sony computer by putting its address to a new pointer variable
	sc := &computer{"Sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if yaac != sc {
		fmt.Printf("yaac=%p\tsc=%p\n", yaac, sc)
	}

	// put apple's value into a new ordinary variable
	apple := *yaac

	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Printf("&yaac=%p\tyaac=%p\t&apple=%p\n", &yaac, yaac, &apple)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so

	// print the values:
	// the value that is pointed by the apple and the new variable
	if *yaac == apple {
		fmt.Printf("*yaac=%+v\tapple=%+v\n", *yaac, apple)
	}

	// create a new function: change
	// the func can change the given computer's brand to another brand

	// change sony's brand to hp using the func — print sony's brand
	change(sc, "HP")
	fmt.Println("Sony's brand:", sc.brand)

	// write a func that returns the value that is pointed by the given *computer
	// print the returned value
	fmt.Printf("*yaac=%+v\n", valueOf(yaac))

	// write a new constructor func that returns a pointer to a computer
	// and call the func 3 times and print the returned values' addresses
	fmt.Printf("new *c=%p\n", newComputer("asus"))
	fmt.Printf("new *c=%p\n", newComputer("asus"))
	fmt.Printf("new *c=%p\n", newComputer("asus"))
}

func change(c *computer, brand string) {
	c.brand = brand
}

func valueOf(c *computer) computer {
	return *c
}

func newComputer(brand string) *computer {
	return &computer{brand: brand}
}

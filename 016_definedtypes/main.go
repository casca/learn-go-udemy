package main

import (
	"fmt"

	"github.com/casca/learn-go-udemy/016_definedtypes/weights"
)

type gram float64
type ounce float64

func main() {
	var g gram = 1000
	var o ounce

	o = ounce(g) * 0.035274

	fmt.Printf("%g grams is %.2f ounces\n", g, o)

	type (
		// Gram's underlying type is int
		Gram int

		// Kilogram's underlying type is int
		Kilogram Gram

		// Ton's underlying type is int
		Ton Kilogram
	)

	var (
		salt   Gram     = 100
		apples Kilogram = 5
		truck  Ton      = 10
	)

	fmt.Printf("salt: %d, apples: %d, truck: %d\n", salt, apples, truck)

	salt = Gram(apples)
	apples = Kilogram(truck)
	truck = Ton(Gram(int(apples)))

	salt = Gram(weights.Gram(100))

	fmt.Printf("Type of weights.Gram: %T\n", weights.Gram(1))
	fmt.Printf("Type of main.Gram: %T\n", Gram(1))

	type myGram weights.Gram
	var pepper myGram = 20
	pepper = myGram(salt)
	fmt.Printf("Type of pepper: %T\n", pepper)
}

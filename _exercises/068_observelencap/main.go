package main

import (
	"fmt"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Observe the length and capacity
//
//  Follow the instructions inside the code below to
//  gain more intuition about the length and capacity of a slice.
//
// ---------------------------------------------------------

func main() {
	// --- #1 ---
	// 1. create a new slice named: games
	//  games := []string{"foo", "bar"}

	// 2. print the length and capacity of the games slice
	// printLenCap("games", games)

	// 3. comment out the games slice
	//    then declare it as an empty slice
	var games []string

	// 4. print the length and capacity of the games slice
	printLenCap("games empty", games)

	// 5. append the elements: "pacman", "mario", "tetris", "doom"
	games = append(games, "pacman", "mario", "tetris", "doom")

	// 6. print the length and capacity of the games slice
	printLenCap("games after append", games)

	// 7. comment out everything
	//
	// 8. declare the games slice again using a slice literal
	//    (use the same elements from step 3)

	// --- #2 ---
	// 1. use a loop from 0 to 4 to slice the games slice, element by element.
	for i := 0; i <= 4; i++ {
		printLenCap("slice "+strconv.Itoa(i), games[:i])
	}

	//
	// 2. print its length and capacity along the way (in the loop).

	fmt.Println()
	// for ... {
	// 	fmt.Printf("games[:%d]'s len: %d cap: %d\n", ...)
	// }

	// --- #3 ---
	// 1. slice the games slice up to zero element
	//    (save the result to a new slice named: "zero")
	zero := games[:0]

	// 2. print the games and the new slice's len and cap
	printLenCap("zero", zero)

	// 3. append a new element to the new slice
	zero = append(zero, "luigi")

	// 4. print the new slice's lens and caps
	printLenCap("zero after append", zero)

	// 5. repeat the last two steps 5 times (use a loop)
	for i := 0; i < 5; i++ {
		zero = append(zero, "game "+strconv.Itoa(i))
		printLenCap("zero after game "+strconv.Itoa(i), zero)
	}

	// 6. notice the growth of the capacity after the 5th append
	//
	// Use this slice's elements to append to the new slice:
	// []string{"ultima", "dagger", "pong", "coldspot", "zetra"}
	zero = append(zero, []string{"ultima", "dagger", "pong", "coldspot", "zetra"}...)
	printLenCap("append more to zero", zero)
	fmt.Println()

	// zero := ...
	// fmt.Printf("games's     len: %d cap: %d\n", ...)
	// fmt.Printf("zero's      len: %d cap: %d\n", ...)

	// for ... {
	//   ...
	//   fmt.Printf("zero's      len: %d cap: %d\n", ...)
	// }

	// --- #4 ---
	// using a range loop, slice the zero slice element by element,
	// and print its length and capacity along the way.
	//
	// observe that, the range loop only loops for the length, not the cap.
	for i := 0; i <= len(zero); i++ {
		printLenCap("(len) zero at "+strconv.Itoa(i), zero[:i])
	}

	fmt.Println()

	// for ... {
	//   s := zero[:n]
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d\n", ...)
	// }

	// --- #5 ---
	// 1. do the 3rd step above again but this time, start by slicing
	//    the zero slice up to its capacity (use the cap function).
	//
	// 2. print the elements of the zero slice in the loop.
	for i := 0; i <= cap(zero); i++ {
		printLenCap("(cap) zero at "+strconv.Itoa(i), zero[:i])
		fmt.Println(zero[:i])
	}
	fmt.Println()

	// zero = ...
	// for ... {
	//   fmt.Printf("zero[:%d]'s  len: %d cap: %d - %q\n", ...)
	// }

	// --- #6 ---
	// 1. change the one of the elements of the zero slice
	zero[0] = "foo"
	//
	// 2. change the same element of the games slice
	games[0] = "bar"
	// 3. print the games and the zero slices
	fmt.Printf("%s %q\n", "zero", zero)
	fmt.Printf("%s %q\n", "games", games)

	// 4. observe that they don't have the same backing array
	fmt.Println()
	// ...
	// fmt.Printf("zero  : %q\n", zero)
	// fmt.Printf("games : %q\n", games)

	// --- #7 ---
	// try to slice the games slice beyond its capacity
	// _ = games[:1e3] // panic: runtime error: slice bounds out of range [:1000] with capacity 4
}

func printLenCap(msg string, s []string) {
	fmt.Printf("%-20s: len=%-5d cap=%-5d\n", msg, len(s), cap(s))
}

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Warm-up
//
//  Create and print the following maps.
//
//  1. Phone numbers by last name
//  2. Product availability by Product ID
//  3. Multiple phone numbers by last name
//  4. Shopping basket by Customer ID
//
//     Each item in the shopping basket has a Product ID and
//     quantity. Through the map, you can tell:
//     "Mr. X has bought Y bananas"
//
// ---------------------------------------------------------

func main() {
	// Hint: Store phone numbers as text

	// #1
	// Key        : Last name
	// Element    : Phone number
	phones := map[string]string{}
	phones["Doe"] = "1234"
	phones["Foo"] = "7890"
	fmt.Println(phones)

	// #2
	// Key        : Product ID
	// Element    : Available / Unavailable
	availability := map[string]bool{}
	availability["ABC"] = true
	availability["DEF"] = false
	fmt.Println(availability)

	// #3
	// Key        : Last name
	// Element    : Phone numbers
	multiplePhoneNumbers := map[string][]string{}
	multiplePhoneNumbers["John"] = []string{"123", "456"}
	multiplePhoneNumbers["Doe"] = []string{"789", "012"}
	fmt.Println(multiplePhoneNumbers)

	// #4
	// Key        : Customer ID
	// Element Key:
	//   Key: Product ID Element: Quantity
	baskets := map[string]map[int]int{}
	baskets["123"] = map[int]int{
		1: 4,
		2: 5,
	}
	fmt.Println(baskets)
}

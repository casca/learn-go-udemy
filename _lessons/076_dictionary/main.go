package main

import (
	"fmt"
)

func main() {
	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	fmt.Println("[english word] -> [turkish word]")
	// 	return
	// }
	// query := args[0]

	// english := []string{"good", "great", "perfect"}
	// turkish := []string{"iyi", "harika", "mükemmel"}

	// for i, w := range english {
	// 	if query == w {
	// 		fmt.Printf("%q means %q\n", w, turkish[i])
	// 		return
	// 	}
	// }
	// fmt.Printf("%q not found\n", query)

	// var broken map[[]int]int
	// var broken map[map[int]string]int
	var broken map[int][]int
	_ = broken

	var dict map[string]string
	// dict["up"] = "yukari"
	// dict["down"] = "asagi"

	fmt.Printf("Zero Value: %#v\n", dict)
	fmt.Printf("# of Keys : %d\n", len(dict))

	key := "good"
	value := dict[key]
	fmt.Printf("%q means %#v\n", key, value)
}

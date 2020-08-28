package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

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

	// var dict map[string]string
	dict := map[string]string{
		// "good":    "kötü",
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel",
		// 42: "forty two",
		// 42: 42,
	}
	// dict["good"] = "iyi"
	// dict["up"] = "yukari"
	// dict["down"] = "asagi"
	// dict["mistake"] = ""

	delete(dict, "awesome")
	delete(dict, "awesome")
	delete(dict, "nothing")
	// dict = nil
	// for k := range dict {
	// 	delete(dict, k)
	// }
	// mapclear()

	// turkish := dict
	turkish := make(map[string]string, len(dict))
	fmt.Println("len(turkish map)", len(turkish))
	for k, v := range dict {
		turkish[v] = k
	}

	turkish["good"] = "güzel"
	dict["great"] = "kusursuz"

	fmt.Printf("english: %q\nturkish:%q\n", dict, turkish)

	fmt.Printf("%#v\n", dict)

	for k, v := range dict {
		fmt.Printf("%q means %#v\n", k, v)
	}

	// fmt.Printf("Zero Value: %#v\n", dict)

	// key := "good"

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found\n", query)

	fmt.Printf("# of Keys : %d\n", len(dict))
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/11-if/exercises/05-movie-ratings/main.go
	// if len(os.Args) != 2 {
	// 	fmt.Println("Requires age")
	// 	return
	// }
	// age, err := strconv.Atoi(os.Args[1])
	// if err != nil {
	// 	fmt.Printf("Wrong age: %q\n", age)
	// 	return
	// }
	// if age <= 13 {
	// 	fmt.Println("PG-Rated")
	// } else if age > 13 && age < 17 {
	// 	fmt.Println("PG-13")
	// } else {
	// 	fmt.Println("R-Rated")
	// }
	if len(os.Args) != 2 {
		fmt.Println("Requires age")
	} else if age, err := strconv.Atoi(os.Args[1]); err != nil {
		fmt.Printf("Wrong age: %q\n", os.Args[1])
	} else if age <= 13 {
		fmt.Println("PG-Rated")
	} else if age > 13 && age < 17 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("R-Rated")
	}
}

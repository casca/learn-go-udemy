package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		// fmt.Println("Scanned Text:", in.Text())
		lines++
	}
	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR:", err)
	}

	// fmt.Println("Scanned Text:", in.Text())
	// fmt.Println("Scanned Text:", in.Bytes())
}

package main

import "fmt"

func main() {
	// [...]string{ "Kafka's Revenge", "Stay Golden" }
	books := [4]string{
		"Kafka's Revenge",
		"Stay Golden",
		"Everythingship",
		"Kafka's Revenge 2nd Edition",
	}
	fmt.Printf("%#v\n", books)
}

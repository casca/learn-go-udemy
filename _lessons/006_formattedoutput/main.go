package main

import "fmt"

func main() {
	var speed int
	var heat float64
	var off bool
	var brand string

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)
	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)

	fmt.Printf("%v is %v away. Think %[2]v kms! %[1]v OMG!\n", planet, distance)

	fmt.Println()
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
}

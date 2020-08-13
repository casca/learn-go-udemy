package main

import "fmt"

func main() {
	msg := []byte{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(msg[0:5])
	fmt.Println(msg[0:])
	fmt.Println(msg[:5])
	fmt.Println(msg[:])
	fmt.Println(append(msg[:4], '!'))
	fmt.Println(append(msg, '!'))
	fmt.Println(msg)
}

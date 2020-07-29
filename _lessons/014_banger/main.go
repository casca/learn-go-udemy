package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	msg := os.Args[1]
	l := len(msg)
	b := strings.Repeat("!", l)
	s := b + msg + b
	s = strings.ToUpper(s)
	fmt.Println(s)
}

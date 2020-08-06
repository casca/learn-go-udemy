package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const (
		EUR = iota
		GBP
	)

	rates := [...]float64{EUR: 1.1, GBP: 0.9}

	if len(os.Args) != 2 {
		fmt.Println("Please provide the amount to be converted.")
		return
	}

	amount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid amount. It should be a number.")
		return
	}

	fmt.Printf("%.2f USD is %.2f EUR\n", amount, amount*rates[EUR])
	fmt.Printf("%.2f USD is %.2f GBP\n", amount, amount*rates[GBP])

}

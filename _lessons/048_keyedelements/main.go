package main

import "fmt"

func main() {
	// r := [3]float64{1: 2.5, 0: 0.5, 2: 1.5}
	// r := [3]float64{2: 1.5}
	// r := [...]float64{2: 1.5}
	r := [...]float64{5: 1.5, 3: 1, 2.5, 0: 0.5}
	fmt.Printf("%#v\n", r)

	const (
		ETH = iota
		WAN
	)

	rates := [...]float64{
		ETH: 25.5,  // ethereum
		WAN: 120.5, // wanchain
	}

	fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	fmt.Printf("1 BTC is %g WAN\n", rates[WAN])

	fmt.Printf("%#v\n", rates)
}

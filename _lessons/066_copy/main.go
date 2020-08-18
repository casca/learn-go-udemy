package main

import (
	"fmt"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	data := []float64{10, 25, 30, 50}

	// newData := []float64{99, 100}
	// for i := range newData {
	// 	datadata[i] = newData[i]
	// }
	copy(data, []float64{99, 100})
	// n := copy(data, []float64{10, 5, 15, 0, 20})
	// fmt.Printf("%d probabilities copied.\n", n)

	// data = append(data[:0], []float64{10, 5, 15, 0, 20}...)

	// saved := make([]float64, len(data))
	// copy(saved, data)

	saved := append([]float64(nil), data...)

	data[0] = 0

	s.Show("Probabilities (saved)", saved)
	s.Show("Probabilities (data)", data)
	fmt.Printf("Is it gonna rain? %.f%% chance.\n", (data[0]+data[1]+data[2]+data[3])/float64(len(data)))
}

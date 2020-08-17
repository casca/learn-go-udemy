package main

import (
	"math/rand"
	"time"

	"github.com/inancgumus/prettyslice"
	"github.com/inancgumus/screen"
)

func main() {
	prettyslice.PrintBacking = true
	prettyslice.MaxPerLine = 20
	prettyslice.Width = 150

	var nums []int

	screen.Clear()
	for cap(nums) <= 128 {
		screen.MoveTopLeft()

		prettyslice.Show("nums", nums)
		nums = append(nums, rand.Intn(9)+1)

		time.Sleep(time.Second / 4)
	}
}

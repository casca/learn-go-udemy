package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	screen.Clear()

	for {
		// fmt.Println("\f") // for Go Playground
		screen.MoveTopLeft()

		now := time.Now()
		hour, min, sec, nano := now.Hour(), now.Minute(), now.Second(), now.Nanosecond()

		clock := [...]placeholder{
			digits[hour/10], digits[hour%10],
			colon,
			digits[min/10], digits[min%10],
			colon,
			digits[sec/10], digits[sec%10],
			dot,
			digits[nano/1e8],
		}

		alarmed := sec%10 == 0

		for line := range clock[0] {
			if alarmed {
				clock = alarm
			}
			for index, digit := range clock {
				next := clock[index][line]
				if (digit == colon || digit == dot) && sec%2 == 0 {
					next = "   "
				}
				fmt.Print(next, "  ")
			}
			fmt.Println()
		}

		fmt.Println()

		time.Sleep(100 * time.Millisecond)
	}
}

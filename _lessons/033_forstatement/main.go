package main

import "fmt"

func main() {
	var sum1 int
	for i := 1; i <= 1000; i++ {
		sum1 += i
	}
	fmt.Println(sum1)

	var sum2, i2 int
	for i2 <= 1000 {
		sum2 += i2
		i2++
	}
	fmt.Println(sum2)

	var sum3, i3 int
	i3 = 0
	for {
		if i3 > 1000 {
			break
		}
		if i3%2 != 0 {
			i3++
			continue
		}
		sum3 += i3
		i3++
	}
	fmt.Println(sum3)
}

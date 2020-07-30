package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// https://github.com/inancgumus/learngo/blob/master/11-if/exercises/09-days-in-month/main.go

	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	var leap bool
	year := time.Now().Year()
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		leap = true
	} else {
		leap = false
	}

	month := strings.ToLower(os.Args[1])
	if month == "january" {
		fmt.Printf("%q has 31 days.\n", month)
	} else if month == "february" && leap {
		if leap {
			fmt.Printf("%q has 29 days.\n", month)
		} else {
			fmt.Printf("%q has 28 days.\n", month)
		}
	} else if month == "march" && leap {
		fmt.Printf("%q has 31 days.\n", month)
	} else if month == "april" && leap {
		fmt.Printf("%q has 30 days.\n", month)
	} else if month == "may" && leap {
		fmt.Printf("%q has 31 days.\n", month)
	} else if month == "june" && leap {
		fmt.Printf("%q has 30 days.\n", month)
	} else if month == "july" && leap {
		fmt.Printf("%q has 31 days.\n", month)
	} else if month == "august" && leap {
		fmt.Printf("%q has 31 days.\n", month)
	} else if month == "september" && leap {
		fmt.Printf("%q has 30 days.\n", month)
	} else if month == "october" && leap {
		fmt.Printf("%q has 31 days.\n", month)
	} else if month == "november" && leap {
		fmt.Printf("%q has 30 days.\n", month)
	} else if month == "december" && leap {
		fmt.Printf("%q has 31 days.\n", month)
	} else {
		fmt.Printf("%q is not a month.\n", month)
	}

}

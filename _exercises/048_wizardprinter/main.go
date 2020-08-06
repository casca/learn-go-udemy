package main

import (
	"fmt"
	"strings"
)

const pad = 15

func main() {
	scientists := [...][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"},
	}

	fmt.Printf("%-*s%-*s%-*s\n", pad, "First Name", pad, "Last Name", pad, "Nickname")
	fmt.Println(strings.Repeat("=", pad*3))
	for _, s := range scientists {
		fmt.Printf("%-*s%-*s%-*s\n", pad, s[0], pad, s[1], pad, s[2])
	}

}

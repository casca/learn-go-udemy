package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int
		domains []string
		total   int
		lines   int
	)
	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v (line #%d)\n", fields, lines)
			return
		}
		// fmt.Printf("domain: %s - visits: %s\n",fields[0],fields[1])

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if err != nil || visits < 0 {
			fmt.Printf("Wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}

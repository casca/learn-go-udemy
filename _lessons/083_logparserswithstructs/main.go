package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	domain string
	visits int
}

type parser struct {
	sum     map[string]result
	domains []string
	total   int
	lines   int
}

func main() {
	p := parser{
		sum: make(map[string]result),
	}

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		p.lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("Wrong input: %v (line #%d)\n", fields, p.lines)
			return
		}
		// fmt.Printf("domain: %s - visits: %s\n",fields[0],fields[1])

		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if err != nil || visits < 0 {
			fmt.Printf("Wrong input: %q (line #%d)\n", fields[1], p.lines)
			return
		}

		if _, ok := p.sum[domain]; !ok {
			p.domains = append(p.domains, domain)
		}
		p.total += visits
		// r := result{
		// 	domain: domain,
		// 	visits: visits + p.sum[domain].visits,
		// }
		// // p.sum[domain].visits += visits
		// p.sum[domain] = r
		p.sum[domain] = result{
			domain: domain,
			visits: visits + p.sum[domain].visits,
		}
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(p.domains)
	for _, domain := range p.domains {
		// visits := p.sum[domain]
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}
}

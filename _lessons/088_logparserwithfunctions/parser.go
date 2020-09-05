package main

import (
	"fmt"
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

func newParser() parser {
	return parser{
		sum: make(map[string]result),
	}
}

func parse(p parser, line string) (parsed result, err error) {
	fields := strings.Fields(line)
	if len(fields) != 2 {
		err = fmt.Errorf("Wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	parsed.domain = fields[0]
	parsed.visits, err = strconv.Atoi(fields[1])
	if err != nil || parsed.visits < 0 {
		err = fmt.Errorf("Wrong input: %q (line #%d)", fields[1], p.lines)
		return
	}

	return
}

func update(p parser, parsed result) parser {
	domain, visits := parsed.domain, parsed.visits

	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}
	p.total += visits

	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}

	// fmt.Printf("update.parser after:\n\t%+v\n", p)

	return p
}

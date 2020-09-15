package main

import (
	"fmt"
)

type product struct {
	title    string
	price    money
	released timestamp
}

func (p *product) String() string {
	// 	// fmt.Printf("%s: %s (%s)\n", p.title, p.price.string(), format(p.released))
	// 	fmt.Printf("%s: %s (%s)\n", p.title, p.price.string(), p.released.string())
	// }

	// func (p *product) discount(ratio float64) {
	// 	p.price *= money(1 - ratio)
	// }

	// func format(v interface{}) string {
	// 	var t int
	// 	switch v := v.(type) {
	// 	case int:
	// 		t = v
	// 	case string:
	// 		t, _ = strconv.Atoi(v)
	// 	default:
	// 		return "unknown"
	// 	}

	// 	const layout = "2006-01-02"
	// 	u := time.Unix(int64(t), 0)
	// 	return u.Format(layout)
	// }

	// fmt.Printf("%s: %s (%s)\n", p.title, p.price.string(), format(p.released))
	// fmt.Printf("%s: %s (%s)\n", p.title, p.price.string(), p.released.string())
	return fmt.Sprintf("%s: %s (%s)", p.title, p.price, p.released)
}

func (p *product) discount(ratio float64) {
	p.price *= money(1 - ratio)
}

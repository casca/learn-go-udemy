package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme something' to mask!")
		return
	}

	const (
		link  = "http://"
		nlink = len(link)
		mask  = '*'
	)

	var (
		text = args[0]
		size = len(text)
		// stringBuf string
		buf = make([]byte, 0, size)
		// buf = []byte(text)[:0]

		in bool
	)

	// fmt.Println("text size: ", size)

	for i := 0; i < size; i++ {
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true
			// fmt.Printf(`text[%d : %[1]d+%d] = `, i, nlink)
			// fmt.Println(text[i : i+nlink])

			buf = append(buf, link...)
			i += nlink
		}

		// buf = append(buf, text[i])
		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}
		buf = append(buf, c)
		// stringBuf += string(text[i])
	}

	fmt.Println(text)
	fmt.Println(string(buf))
}

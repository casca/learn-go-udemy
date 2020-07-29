package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	var s string
	s = "how are you?"
	s = `how are you?`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n<html>"
	fmt.Println(s)

	s = `
<html>
	<body>"Hello"</body>
</html>`
	fmt.Println(s)

	fmt.Println("c:\\my\\dir\\file")
	fmt.Println(`c:\my\dir\file`)

	name, last := "carl", "sagan"
	name += " edward"
	fmt.Println(name + " " + last)

	fmt.Println("hello" + ", " + "how" + " " + "are you" + " " + "today?")
	fmt.Println(`hello` + `, ` + `how` + ` ` + `are you` + ` ` + "today?")

	eq := "1 + 2 = "
	sum := 1 + 2
	fmt.Println(eq + strconv.Itoa(sum))

	// eq = true + " " + false
	eq = strconv.FormatBool(true) + " " + strconv.FormatBool(false)
	fmt.Println(eq)

	name = "cärl"
	fmt.Println(len(name)) // number of bytes
	fmt.Println(utf8.RuneCountInString(name))
}

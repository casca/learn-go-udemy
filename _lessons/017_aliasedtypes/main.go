package main

func main() {
	type foo = byte
	// byte uint8
	// rune int32

	var (
		byteVal  byte
		uint8Val uint8
		intVal   int
	)

	uint8Val = byteVal
	_, _ = uint8Val, intVal

	var (
		runeVal  rune
		int32Val int32
	)
	runeVal = int32Val
	// runeVal = rune(runeVal) // unnecessary
	_ = runeVal
}

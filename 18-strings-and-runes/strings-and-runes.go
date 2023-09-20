package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"

	fmt.Println("s:", s)
	fmt.Println("len(s):", len(s))

	fmt.Println("For loop over s:")
	for i := 0; i < len(s); i++ {
		fmt.Printf(" %x", s[i])
	}
	fmt.Println()
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	fmt.Println("Ranging over s, getting its runes:")
	for idx, runeValue := range s {
		fmt.Printf(" rune %#U starts at index %d\n", runeValue, idx)
	}
	fmt.Println()

	fmt.Println("Iterating over s, decoding runes manually:")
	for i, w := 0, 0; i < len(s); i += w {
		var runeValue rune
		runeValue, w = utf8.DecodeRuneInString(s[i:])
		fmt.Printf(" rune %#U starts at index %d\n", runeValue, i)
		printIfSoSua(runeValue)
	}
}

func printIfSoSua(r rune) {
	if r == 'ส' {
		fmt.Println(" found so sua!")
	}
}

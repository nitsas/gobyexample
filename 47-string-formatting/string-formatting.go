package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("p (struct %%v): %v\n", p)
	fmt.Printf("p (struct-with-fields %%+v): %+v\n", p)
	fmt.Printf("p (struct-constructor %%#v): %#v\n", p)
	fmt.Printf("p (type-of %%T): %T\n", p)
	fmt.Println()

	fmt.Printf("true (boolean %%t): %t\n", true)
	fmt.Printf("123 (integer-base-10 %%d): %d\n", 123)
	fmt.Printf("123 (binary %%b): %b\n", 123)
	fmt.Printf("123 (hex %%x): %x\n", 123)
	fmt.Printf("33 (character %%c): %c\n", 33)
	fmt.Printf("&p (pointer %%p): %p\n", &p)

	flt := 1230000.45
	fmt.Printf("1230000.45 (float %%f): %f\n", flt)
	fmt.Printf("1230000.45 (scientific-e %%e): %e\n", flt)
	fmt.Printf("1230000.45 (scientific-E %%E): %E\n", flt)
	fmt.Println()

	s := "string"
	fmt.Printf("\"string\" (string %%s): %s\n", s)
	fmt.Printf("\"string\" (string-quoted %%q): %q\n", s)
	fmt.Printf("base 16 \"string\" (hex-base-16 %%x): %x\n", s)
	fmt.Println()

	fmt.Printf("123 (integer-fixed-width %%6d): %6d\n", 123)
	fmt.Printf("1.2 (float-fixed-width %%6.2f): %6.2f\n", 1.2)
	fmt.Printf("1.2, 1.3 (float-left-justify %%-6.2f): %-6.2f, %-6.2f\n", 1.2, 1.3)
	fmt.Printf("\"hi\" (string-fixed-width %%6s): %6s\n", "hi")
	fmt.Printf("\"ho\", \"ho\" (string-left-justify %%-6s): %-6s, %-6s\n", "ho", "ho")
	fmt.Println()

	fmt.Printf("s2 := fmt.Sprintf(\"a %%s\", \"string\")\n")
	s2 := fmt.Sprintf("a %s", "string")
	fmt.Printf("s2 (%%s): %s\n", s2)
	fmt.Println()

	fmt.Fprintf(os.Stderr, "print to stderr and other files: fmt.Fprintf(os.Stderr, \"...\", arg)")
}

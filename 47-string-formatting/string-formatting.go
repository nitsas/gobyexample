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

	fmt.Printf("p (%%v): %v\n", p)
	fmt.Printf("p (%%+v): %+v\n", p)
	fmt.Printf("p (%%#v): %#v\n", p)
	fmt.Printf("type of p (%%T): %T\n", p)
	fmt.Println()

	fmt.Printf("true (%%t): %t\n", true)
	fmt.Printf("123 (%%d): %d\n", 123)
	fmt.Printf("binary 123 (%%b): %b\n", 123)
	fmt.Printf("hex 123 (%%x): %x\n", 123)
	fmt.Printf("char 33 (%%c): %c\n", 33)
	fmt.Printf("pointer &p (%%p): %p\n", &p)

	flt := 1230000.45
	fmt.Printf("1230000.45 (%%f): %f\n", flt)
	fmt.Printf("1230000.45 (%%e): %e\n", flt)
	fmt.Printf("1230000.45 (%%E): %E\n", flt)
	fmt.Println()

	s := "string"
	fmt.Printf("\"string\" (%%s): %s\n", s)
	fmt.Printf("\"string\" (%%q): %q\n", s)
	fmt.Printf("base 16 \"string\" (%%x): %x\n", s)
	fmt.Println()

	fmt.Printf("fixed-width 123 (%%6d): %6d\n", 123)
	fmt.Printf("fixed-width 1.2 (%%6.2f): %6.2f\n", 1.2)
	fmt.Printf("left-justify 1.2, 1.3 (%%-6.2f): %-6.2f, %-6.2f\n", 1.2, 1.3)
	fmt.Printf("fixed-width \"hi\" (%%6s): %6s\n", "hi")
	fmt.Printf("left-justify \"ho\", \"ho\" (%%-6s): %-6s, %-6s\n", "ho", "ho")
	fmt.Println()

	fmt.Printf("s2 := fmt.Sprintf(\"a %%s\", \"string\")\n")
	s2 := fmt.Sprintf("a %s", "string")
	fmt.Printf("s2 (%%s): %s\n", s2)
	fmt.Println()

	fmt.Fprintf(os.Stderr, "print to stderr and other files: fmt.Fprintf(os.Stderr, \"...\", arg)")
}

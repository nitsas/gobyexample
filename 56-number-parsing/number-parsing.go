package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	str := "1234"
	numI, err := strconv.Atoi(str)
	fmt.Printf("strconv.Atoi(\"%s\"): %d, %#v\n", str, numI, err)

	str = "non-int-string"
	numI, err = strconv.Atoi(str)
	fmt.Printf("strconv.Atoi(\"%s\"): %d, %#v\n", str, numI, err)
	fmt.Println()

	str = "1.234"
	bits := 64
	f, err := strconv.ParseFloat(str, bits) // 64: bits of precision
	fmt.Printf("strconv.ParseFloat(\"%s\", %d): %f, %#v", str, bits, f, err)
	fmt.Println()

	str = "1234"
	base := 0 // 0 means "discover the base from the string"
	bits = 64
	numI64, err := strconv.ParseInt(str, base, bits) // 64: bits of precision
	fmt.Printf("strconv.ParseInt(\"%s\", %d, %d): %d, %#v", str, base, bits, numI64, err)
	fmt.Println()

	str = "1234"
	base = 8 // 0 means "discover the base from the string"
	bits = 64
	numI64, err = strconv.ParseInt(str, base, bits) // 64: bits of precision
	fmt.Printf("strconv.ParseInt(\"%s\", %d, %d): %d, %#v", str, base, bits, numI64, err)
	fmt.Println()

	str = fmt.Sprint(uint64(math.Pow(2, 64)))
	base = 0 // 0 means "discover the base from the string"
	bits = 64
	numI64, err = strconv.ParseInt(str, base, bits) // 64: bits of precision
	fmt.Printf("strconv.ParseInt(\"%s\", %d, %d): %d, %#v", str, base, bits, numI64, err)
	fmt.Println()

	str = fmt.Sprint(uint64(math.Pow(2, 63)))
	base = 0 // 0 means "discover the base from the string"
	bits = 64
	numU64, err := strconv.ParseUint(str, base, bits) // 64: bits of precision
	fmt.Printf("strconv.ParseUint(\"%s\", %d, %d): %d, %#v", str, base, bits, numU64, err)
	fmt.Println()
}

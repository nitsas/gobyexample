package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	fmt.Println("s:", s)

	// We can compute sha256 immediately:
	fmt.Printf("sha256.Sum256([]byte(s)) (as %%x): %x\n", sha256.Sum256([]byte(s)))

	// Or build it gradually:
	hasher := sha256.New()
	hasher.Write([]byte(s))
	fmt.Printf("sha256.New() -> .Write([]byte(s)) -> .Sum(nil) (as %%x): %x\n", hasher.Sum(nil))
	// In hasher.Sum(nil), instead of nil we can pass another []byte to append
	// to the previous data and then compute the hash.
}

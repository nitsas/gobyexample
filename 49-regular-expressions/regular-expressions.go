package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	pattern := "p([a-z]+)ch"
	fmt.Println("pattern:", pattern)
	match, err := regexp.MatchString(pattern, "peach")
	if err != nil {
		panic(err)
	}
	fmt.Println("regexp.MatchString(pattern, \"peach\"):", match)
	fmt.Println()

	r, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}
	fmt.Println("r, err := regexp.Compile(pattern); r:", r)
	r = regexp.MustCompile(pattern)
	fmt.Println("r := regexp.MustCompile(pattern):", r)
	str := "Have a lovely peach punch pinch!"
	fmt.Println("str:", str)
	fmt.Println()

	fmt.Println("r.MatchString(str):", r.MatchString(str))
	fmt.Println("r.FindString(str):", r.FindString(str))
	fmt.Println("r.FindStringIndex(str):", r.FindStringIndex(str))
	fmt.Println("r.FindStringSubmatch(str):", r.FindStringSubmatch(str))
	fmt.Println("r.FindStringSubmatchIndex(str):", r.FindStringSubmatchIndex(str))
	fmt.Println()

	fmt.Println("r.FindAllString(str, 0):", r.FindAllString(str, 0))
	fmt.Println("r.FindAllString(str, 1):", r.FindAllString(str, 1))
	fmt.Println("r.FindAllString(str, 2):", r.FindAllString(str, 2))
	fmt.Println("r.FindAllString(str, -1):", r.FindAllString(str, -1))
	fmt.Println("r.FindAllStringSubmatchIndex(str, -1):", r.FindAllStringSubmatchIndex(str, -1))
	fmt.Println()

	bytes := []byte("a peach of bytes")
	fmt.Println("bytes: []byte(\"a peach of bytes\"):", bytes)
	fmt.Println("r.Match(bytes):", r.Match(bytes))
	fmt.Println("r.Find(bytes):", r.Find(bytes))
	fmt.Println("r.FindIndex(bytes):", r.FindIndex(bytes))
	fmt.Println("r.FindSubmatch(bytes):", r.FindSubmatch(bytes))
	fmt.Println("r.FindSubmatchIndex(bytes):", r.FindSubmatchIndex(bytes))
	fmt.Println()

	fmt.Println("r.ReplaceAllString(str, \"kitch\"):", r.ReplaceAllString(str, "kitch"))
	fmt.Println("r.ReplaceAllStringFunc(str, strings.ToUpper):", r.ReplaceAllStringFunc(str, strings.ToUpper))
	fmt.Println()

	fmt.Printf("regexp.MustCompile(\" \").Split(str, 0): %#v\n", regexp.MustCompile(" ").Split(str, 0))
	fmt.Printf("regexp.MustCompile(\" \").Split(str, 1): %#v\n", regexp.MustCompile(" ").Split(str, 1))
	fmt.Printf("regexp.MustCompile(\" \").Split(str, 2): %#v\n", regexp.MustCompile(" ").Split(str, 2))
	fmt.Printf("regexp.MustCompile(\" \").Split(str, -1): %#v\n", regexp.MustCompile(" ").Split(str, -1))
}

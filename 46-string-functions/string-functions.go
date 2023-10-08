package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func aToO(r rune) rune {
	if r == 'a' {
		return 'o'
	} else {
		return r
	}
}

func main() {
	p("Contains:    ", s.Contains("test", "es"))
	p("ContainsAny: ", s.ContainsAny("abcdefg", "aeiouy"))
	p("Count:       ", s.Count("test", "t"))
	p("HasPrefix:   ", s.HasPrefix("test", "te"))
	p("HasSuffix:   ", s.HasSuffix("test", "st"))
	p("Index:       ", s.Index("test", "e"))
	p("Join:        ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:      ", s.Repeat("a", 8))
	p("Replace:     ", s.Replace("Hello world", "l", "_", -1))
	p("Replace:     ", s.Replace("Hello world", "l", "_", 2))
	p("Split:       ", s.Split("a-b-c-d-e", "-"))
	p("ToUpper:     ", s.ToUpper("what's up"))
	p("ToLower:     ", s.ToLower("HERE'S WHAT'S GOING DOWN"))
	p("Map:         ", s.Map(aToO, "lalala"))
	p("TrimSpace:   ", s.TrimSpace("	\n   surrounded by whitespace\n		  "))
}

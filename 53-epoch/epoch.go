package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)
	fmt.Println()

	// time since epoch in seconds
	fmt.Println("now.Unix():", now.Unix())
	// time since epoch in milliseconds
	fmt.Println("now.UnixMilli():", now.UnixMilli())
	// time since epoch in nanoseconds
	fmt.Println("now.UnixNano():", now.UnixNano())
	fmt.Println()

	// Parse epoch time into a time.Date.
	// The second argument is the nanoseconds.
	fmt.Println("time.Unix(now.Unix(), 0):", time.Unix(now.Unix(), 0))
	fmt.Println("time.Unix(now.Unix(), 123):", time.Unix(now.Unix(), 123))
	fmt.Println("time.Unix(0, now.UnixNano()):", time.Unix(0, now.UnixNano()))
	fmt.Println("time.UnixMilli(now.UnixMilli()):", time.UnixMilli(now.UnixMilli()))
	fmt.Println("time.UnixMicro(now.UnixMicro()):", time.UnixMicro(now.UnixMicro()))
}

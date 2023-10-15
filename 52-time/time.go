package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	then := time.Date(2023, 11, 7, 17, 28, 10, 651387237, time.UTC)
	// Location time.Local knows about summer time!
	// It chooses EEST for October 7th, and EET for November 7th.
	thenLocalEEST := time.Date(2023, 10, 7, 17, 28, 10, 651387237, time.Local)
	thenLocalEET := time.Date(2023, 11, 7, 17, 28, 10, 651387237, time.Local)

	fmt.Println("now:", now)
	fmt.Println("then:", then)
	fmt.Println("thenLocalEEST:", thenLocalEEST)
	fmt.Println("thenLocalEET:", thenLocalEET)
	fmt.Println()

	fmt.Println("now.Year():", now.Year())
	fmt.Println("now.Month():", now.Month())
	fmt.Println("now.Day():", now.Day())
	fmt.Println("now.Weekday():", now.Weekday())
	fmt.Println("now.Hour():", now.Hour())
	fmt.Println("now.Minute():", now.Minute())
	fmt.Println("now.Second():", now.Second())
	fmt.Println("now.Nanosecond():", now.Nanosecond())
	fmt.Println("now.Location():", now.Location())
	fmt.Println()

	fmt.Println("then.Before(now):", then.Before(now))
	fmt.Println("then.After(now):", then.After(now))
	fmt.Println("then.Equal(now):", then.Equal(now))
	fmt.Println()

	diff := then.Sub(now)
	fmt.Println("diff := then.Sub(now):", diff)
	fmt.Println("now.Sub(then):", now.Sub(then))
	fmt.Printf("type-of diff: %T\n", diff)
	fmt.Println("diff.Hours():", diff.Hours())
	fmt.Println("diff.Minutes():", diff.Minutes())
	fmt.Println("diff.Seconds():", diff.Seconds())
	fmt.Println("diff.Nanoseconds():", diff.Nanoseconds())
	fmt.Println()

	fmt.Println("now.Add(diff):", now.Add(diff))
	fmt.Println("then.Add(-diff):", then.Add(-diff))
	fmt.Println("then.Add(-diff).Local():", then.Add(-diff).Local())
}

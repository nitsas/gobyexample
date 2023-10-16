package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now.Format(time.RFC3339))
	fmt.Println()

	fmt.Println("Custom formats:")
	fmt.Println(now.Format("3:04 PM"))
	fmt.Println(now.Format("Mon Jan _2 15:04:05 2006"))
	fmt.Println(now.Format("2006-01-02T15:04:05.999999-07:00"))
	zone, zoneOffset := now.Zone()
	fmt.Printf("with printf: date%d_%02d_%02dtime%02dx%02dx%02dzone%soffset%d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(),
		zone, zoneOffset,
	)
	fmt.Println()

	fmt.Println("Parsing:")
	christmasMorning, err := time.Parse(
		time.RFC3339,
		"2023-12-25T09:00:00+02:00",
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("RFC3339 Christmas morning:", christmasMorning)

	form := "3 04 PM"
	hourOnly, err := time.Parse(form, "9 17 AM")
	if err != nil {
		panic(err)
	}
	fmt.Println("hourOnly:", hourOnly)
	fmt.Println()

	formLong := "Mon Jan _2 15:04:05 2006"
	_, err = time.Parse(formLong, "8:41PM")
	fmt.Println("parsing error:", err)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("something")
	}

	fmt.Print("Today is ")
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("the weekend!")
	default:
		fmt.Println("a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatami := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println(i, "is a bool")
		case int:
			fmt.Println(i, "is an int")
		case string:
			fmt.Println(i, "is a string")
		default:
			fmt.Println(i, "is of a type that I cannot recognize")
		}
	}
	whatami(true)
	whatami(15)
	whatami("hello")
	whatami(14.2)
}

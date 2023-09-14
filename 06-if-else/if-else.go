package main

import "fmt"

func main() {
	fmt.Println("if else:")
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	fmt.Println("\nif without else:")
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	fmt.Println("\nif else if else; declared variables are available later")
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has a single digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

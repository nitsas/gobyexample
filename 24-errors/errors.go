package main

import (
	"errors"
	"fmt"
)

func isEven(num int) bool {
	return num%2 == 0
}

func halveEven(num int) (int, error) {
	if !isEven(num) {
		return -1, errors.New("can't work with odd numbers")
	}

	return num / 2, nil
}

type intArgError struct {
	arg int
	msg string
}

func (ae *intArgError) Error() string {
	return fmt.Sprintf("argument %d %s", ae.arg, ae.msg)
}

func doubleOdd(num int) (int, error) {
	if isEven(num) {
		return -1, &intArgError{num, "is even - cannot work with it"}
	}

	return num * 2, nil
}

func main() {
	result, err := halveEven(4)
	fmt.Println("halveEven(4): result:", result, "| err:", err)
	result, err = halveEven(2)
	fmt.Println("halveEven(2): result:", result, "| err:", err)
	result, err = halveEven(1)
	fmt.Println("halveEven(1): result:", result, "| err:", err)
	fmt.Println()

	ae := intArgError{5, "cannot work with it"}
	fmt.Println("ae:", ae)
	fmt.Println("ae.Error():", ae.Error())
	fmt.Println()

	result, err = doubleOdd(1)
	fmt.Println("doubleOdd(1): result:", result, "| err:", err)
	result, err = doubleOdd(2)
	fmt.Println("doubleOdd(2): result:", result, "| err:", err)
	fmt.Println()

	if ae, ok := err.(*intArgError); ok {
		fmt.Println("Typecast err to *intArgError to examine its contents, ae:", ae)
		fmt.Println("ae.arg:", ae.arg)
		fmt.Println("ae.msg:", ae.msg)
	}
}

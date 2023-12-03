package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	url := "https://gobyexample.com"
	fmt.Println("GET", url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("...")
	}
}

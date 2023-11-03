/*
Convert lowercase input to uppercase output.
To try this program, pipe lowercase lines into it.
Example:
echo "hello world\nsecond line" | go run line-filters.go
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		uppercaseLine := strings.ToUpper(line)
		fmt.Println(uppercaseLine)
	}

	err := scanner.Err()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

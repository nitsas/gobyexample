package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("word", "default", "a string argument")
	numberPtr := flag.Int("number", 42, "an integer arugment")
	flagPtr := flag.Bool("flag", false, "a boolean argument (flag)")

	var word2 string
	flag.StringVar(&word2, "word2", "default2", "another string argument")

	flag.Parse()

	fmt.Println("Command-line flags:")
	fmt.Println("word:", *wordPtr)
	fmt.Println("number:", *numberPtr)
	fmt.Println("flag:", *flagPtr)
	fmt.Println("word2:", word2)
	fmt.Printf("flag.Args(): %#v\n", flag.Args())

	// Run this with `-h` or `--help` to get an auto-generated help message.
	// If you provide an unknown command-line flag, you'll get an error and the help.
	// If you provide positional arguments between flags, all the flags after
	// the positional arguments will be interpreted as plain positional arguments.
}

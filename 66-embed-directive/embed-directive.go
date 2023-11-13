package main

import (
	"embed"
	"fmt"
)

//go:embed folder/single_file.txt
var fileString string

//go:embed folder/single_file.txt
var fileBytes []byte

//go:embed folder
var folder embed.FS

func main() {
	fmt.Println("Contents of embedded folder/single_file.txt:")
	print(fileString)
	fmt.Println("Same contents read as a []byte:")
	print(string(fileBytes))

	file1Content, _ := folder.ReadFile("folder/file1.hash")
	fmt.Println("Contents of embedded folder/file1.hash:")
	print(string(file1Content))

	file2Content, _ := folder.ReadFile("folder/file2.hash")
	fmt.Println("Contents of embedded folder/file2.hash:")
	print(string(file2Content))
}

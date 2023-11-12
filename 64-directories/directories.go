package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func must(e error) {
	if e != nil {
		panic(e)
	}
}

func createEmptyFile(filename string) {
	content := []byte("")
	must(os.WriteFile(filename, content, 0644))
}

func dirOrFile(dirEntry fs.DirEntry) string {
	if dirEntry.IsDir() {
		return "(directory)"
	}
	return "(file)"
}

func ls(dirName string) {
	fmt.Printf("Listing %s:\n", dirName)
	parentContents, err := os.ReadDir(dirName)
	must(err)
	for _, dirEntry := range parentContents {
		fmt.Println("  ", dirEntry.Name(), dirOrFile(dirEntry))
	}
}

func visitDirEntry(path string, dirEntry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println("  ", path, dirOrFile(dirEntry))
	return nil
}

func main() {
	os.RemoveAll("subdir")

	fmt.Println("Creating a few temporary directories under subdir/ ...")
	must(os.Mkdir("subdir", 0755))
	defer os.RemoveAll("subdir")
	must(os.MkdirAll("subdir/parent/child", 0755))

	fmt.Println("... and a few empty files in them")
	createEmptyFile("subdir/file1")
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	ls("subdir/parent")
	fmt.Println("Chdir subdir/parent/child")
	must(os.Chdir("subdir/parent/child"))
	ls(".")
	fmt.Println("Chdir ../../..")
	must(os.Chdir("../../.."))

	fmt.Println("Walk directory subdir:")
	must(filepath.WalkDir("subdir", visitDirEntry))
}

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// var lsPath string
	lsPath, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
		// e.g.: panic: exec: "...": executable file not found in $PATH
	}
	fmt.Printf("Found command `ls` at: %#v\n", lsPath)
	fmt.Printf("The rest of this program will execute `%s -a -l -h`\n", lsPath)

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	err = syscall.Exec(lsPath, args, env)
	fmt.Println("Back in Go after exec-ing `ls ...`") // never runs, ls took over
	if err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"io"
	"os/exec"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dateCmd := exec.Command("date")
	fmt.Println("$ date")
	dateOutBytes, err := dateCmd.Output()
	panicIf(err)
	fmt.Println(string(dateOutBytes))

	fmt.Println("$ date -x")
	dateOutBytes, err = exec.Command("date", "-x").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.ExitError:
			fmt.Println(e)
			// exit status 1
			fmt.Println("e.String():", e.String(), ", e.ExitCode():", e.ExitCode())
			// e.ExitCode() holds the exit status
		default:
			panic(err)
		}
	}
	fmt.Println()

	fmt.Println("$ adsfadsbasdf")
	_, err = exec.Command("adsfadsbasdf").Output()
	if err != nil {
		switch err.(type) {
		case *exec.Error:
			fmt.Println(err)
			// exec: "adsfadsbasdf": executable file not found in $PATH
		default:
			panic(err)
		}
	}
	fmt.Println()

	grepInput := []byte("hello grep\ngoodbye grep")
	fmt.Printf("$ echo %#v | grep hello\n", string(grepInput))
	grepCmd := exec.Command("grep", "hello")
	grepStdin, err := grepCmd.StdinPipe()
	panicIf(err)
	grepStdout, err := grepCmd.StdoutPipe()
	panicIf(err)
	// We could also collect grepCmd.StderrPipe().
	grepCmd.Start()
	grepStdin.Write(grepInput)
	grepStdin.Close()
	grepOutBytes, err := io.ReadAll(grepStdout)
	panicIf(err)
	grepCmd.Wait()
	fmt.Println(string(grepOutBytes))

	fmt.Println("$ ls -a -l -h")
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOutBytes, err := lsCmd.Output()
	panicIf(err)
	fmt.Println(string(lsOutBytes))
}

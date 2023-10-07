package main

import "fmt"

func justPanic(message string) {
	panic(message)
}

func recoverAndPrint() {
	err := recover()
	if err != nil {
		fmt.Printf("recoverAndPrint: Recovered error %v\n", err)
		// Let's continue the execution of our program
		mainProgramLogic()
	}
}

func mainProgramLogic() {
	fmt.Println("mainProgramLogic: Running")
	// But what if this panics too?
	// We can recover again.
	// But won't that nest function calls too much over time?
}

func main() {
	defer recoverAndPrint()

	justPanic("Sorry, I panicked!")

	// Will the rest of main run? No
	fmt.Println("main: Rest of the code after justPanic")

	// If we want the program to keep working after the panic, then we have to
	// continue execution in the deferred function where recover happened.
}

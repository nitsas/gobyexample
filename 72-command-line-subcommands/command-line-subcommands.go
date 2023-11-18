package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	// The 2nd argument to NewFlagSet is about error-handling:
	// - flag.ContinueOnError returns a descriptive error
	// - flag.ExitOnError calls os.Exit(2) or for -h/-help Exit(0)
	// - flag.PanicOnError calls panic with a descriptive error
	feedCmd := flag.NewFlagSet("feed", flag.ExitOnError)
	petNameP := feedCmd.String("pet-name", "cat", "which pet to feed")
	gramsP := feedCmd.Uint("grams", 0, "how many grams of food")
	treatP := feedCmd.Bool("treat", false, "do you want to give a treat too?")

	// More methods of FlagSet:
	// Args, ErrorHandling, Init, Lookup, NArg, NFlag, Func, Name, Parse, Parsed, Set,
	// BoolFunc, BoolVar, Duration, Float64, Int, Int64, TextVar, Uint64,
	// Var, Visit, VisitAll, UnquoteUsage

	petCmd := flag.NewFlagSet("pet", flag.ExitOnError)
	howLongP := petCmd.Duration("how-long", 10*time.Second, "pet it for how long?")

	if len(os.Args) < 2 {
		fmt.Println("expected subcommand 'feed' or 'pet'")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "feed":
		feedCmd.Parse(os.Args[2:])
		fmt.Println("Called subcommand 'feed' with args:")
		fmt.Println("  pet-name:", *petNameP)
		fmt.Println("  grams:", *gramsP)
		fmt.Println("  treat:", *treatP)
		fmt.Println("  trailing args:", feedCmd.Args())
	case "pet":
		petCmd.Parse(os.Args[2:])
		fmt.Println("Called subcommand 'pet' with args:")
		fmt.Printf("  how-long: %#v\n", *howLongP)
		fmt.Println("  trailing args:", petCmd.Args())
	default:
		fmt.Println("expected subcommand 'feed' or 'pet'")
		os.Exit(1)
	}
}

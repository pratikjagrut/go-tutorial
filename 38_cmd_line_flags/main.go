package main

import (
	"flag"
	"fmt"
)

func main() {

	// Basic flag declarations are available for string, integer, and boolean options
	// (flagName, defaultValue, usage)
	// This flag function retruns pointer to the value not the value itself
	stringPtr := flag.String("word", "foo", "a string")
	intPtr := flag.Int("number", 42, "an int")
	boolPtr := flag.Bool("ok", false, "a bool")

	// Also we can decalare option that uses an existing variable
	var extVar string
	flag.StringVar(&extVar, "extvar", "bar", "a string var")

	// After all flags are declared, flag.Parse() is called to execute the command-line parsing.
	flag.Parse()

	fmt.Println("Word:", *stringPtr)
	fmt.Println("Number:", *intPtr)
	fmt.Println("Ok:", *boolPtr)
	fmt.Println("ExtVar:", extVar)

	// This will collect all tailing var or positional arg
	// Note that the flag package requires all flags to appear before positional arguments
	// (otherwise the flags will be interpreted as positional arguments).
	fmt.Println("Tail:", flag.Args())
}

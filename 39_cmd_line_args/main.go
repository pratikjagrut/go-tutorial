//Command-line arguments are a common way to parameterize execution of programs.
package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args[0]) //First argument is always path of program

	fmt.Println()
	for i, arg := range args {
		fmt.Printf("arg index: %d, arg: %v\n", i, arg)
	}

	//try running like below
	//go run main.go a b c d
}

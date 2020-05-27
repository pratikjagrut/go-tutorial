package main

import (
	"fmt"
	"os"
)

func main() {
	// Set env var
	os.Setenv("VERSION", "1")
	// Get env var
	fmt.Println("Version: ", os.Getenv("VERSION"))

	fmt.Println()
	fmt.Println("*****Env from System*****")
	fmt.Println()
	// get all the env var from system in [kay]value form
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

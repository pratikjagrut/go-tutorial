package main

import (
	"fmt"
)

func main() {
	// ********* if, example *********
	if true {
		fmt.Println("True")
	}

	// if else, example

	if 5%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd nubmer")
	}

	// if else ladder

	i := 10
	if i == 0 {
		fmt.Println("It's 0")
	} else if i < 0 {
		fmt.Println("Negative number")
	} else {
		fmt.Println("Positive number")
	}

	//if with short statement
	if j := 10; j%2 == 0 {
		fmt.Println(j, "is even number")
	}
}

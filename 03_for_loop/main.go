package main

import "fmt"

func main() {

	//Classic for loop pattern
	for b := 5; b >= 1; b-- {
		fmt.Printf("%d ", b)
	}
	fmt.Println()

	// while() loop in Go
	a := 1
	for a <= 5 {
		fmt.Printf("%d ", a)
		a++
	}
	fmt.Println()

	//Infinite for loop, will stop only after break or return
	for {
		fmt.Println("hi")
		break //If you remove break, it will go in infinite looping
	}

}

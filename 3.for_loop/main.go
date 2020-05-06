//This program demonstrates for loop
// Go has only only for loop, no while loop

package main

import "fmt"

func main() {

	// Basic for implimentation
	a := 1
	for a <= 5 {
		fmt.Printf("%d ", a)
		a++
	}

	fmt.Println()

	//Classic for loop pattern
	for b := 5; b >= 1; b-- {
		fmt.Printf("%d ", b)
	}

	fmt.Println()

	//Infinite for loop, will stop only after break or return
	for {
		fmt.Println("hi")
		break
	}

}

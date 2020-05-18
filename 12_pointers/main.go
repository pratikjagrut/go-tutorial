//Go has pointers. A pointer holds the memory address of a value.
package main

import "fmt"

func main() {
	a, b := 12, 1200

	ptr := &a                          // point to i
	fmt.Println("Add of a: ", ptr)     // Address of `a`
	fmt.Println("Value of a: ", *ptr)  // read `a` through the pointer
	*ptr = 67                          // set `a` through the pointer
	fmt.Println("New value of a: ", a) // see the new value of `a`

	fmt.Println()

	ptr = &b                           // point to b
	fmt.Println("Add of b: ", ptr)     // Address of `b`
	fmt.Println("Value of b: ", *ptr)  // read `b` through the pointer
	*ptr = *ptr * 40                   // divide b through the pointer
	fmt.Println("New value of b: ", b) // see the new value of b

}

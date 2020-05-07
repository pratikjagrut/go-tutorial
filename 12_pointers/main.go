//Go has pointers. A pointer holds the memory address of a value.
package main

import "fmt"

func main() {
	i := 10
	ptr := &i //The & operator generates a pointer to its operand.
	fmt.Printf("Initial i:\t%d\n", i)
	fmt.Printf("Address of i:\t%v\n", ptr)
	*ptr++ //The * operator denotes the pointer's underlying value.
	fmt.Printf("Final i:\t%d\n", i)

}

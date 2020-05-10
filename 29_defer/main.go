//Defer keyword is used to execute a function later, especially for cleanup
package main

import "fmt"

func firstFunc() {
	fmt.Println("First function called")
}

func cleanup() {
	fmt.Println("Cleanup....")
}

func main() {
	defer cleanup()
	firstFunc()

	write := make(chan int)
	defer close(write)

	go func(arg int) {
		write <- arg
	}(10)

	fmt.Println(<-write)
	// <-write now using this channel will give deadlock error, as it is closed
	//defer works in inside out fashion,
	//first defer close will execute and then cleanup function
}

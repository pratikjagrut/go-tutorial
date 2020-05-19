//The select statement lets a goroutine wait on multiple communication operations.
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go func(arg string) {
		ch1 <- arg
	}("First channel response")

	go func(arg string) {
		ch2 <- arg
	}("Second channel respose")

	go func(arg string) {
		ch3 <- arg
	}("Third channel respose")

	//Go’s select lets you wait on multiple channel operations.
	for i := 0; i < 4; i++ {
		select {
		case m := <-ch1:
			fmt.Println(m)
		case n := <-ch2:
			fmt.Println(n)
		case o := <-ch3:
			fmt.Println(o)
		default:
			fmt.Println("No other statement is ready yet")
			time.Sleep(1 * time.Second) // waits for other statements to get ready
		}
	}
}

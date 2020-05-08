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
		time.Sleep(2 * time.Second)
		ch1 <- arg
	}("First channel response")

	go func(arg string) {
		time.Sleep(1 * time.Second)
		ch2 <- arg
	}("Second channel respose")

	go func(arg string) {
		time.Sleep(1 * time.Second)
		ch3 <- arg
	}("Third channel respose")

	//Go’s select lets you wait on multiple channel operations.
	for i := 0; i < 3; i++ {
		select {
		case m := <-ch1:
			fmt.Println(m)
		case n := <-ch2:
			fmt.Println(n)
		case o := <-ch3:
			fmt.Println(o)
		}
	}
}

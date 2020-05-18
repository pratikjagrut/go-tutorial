//A goroutine is a lightweight thread managed by the Go runtime.
package main

import (
	"fmt"
	"time"
)

func show(arg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(arg)
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	show("go go go") //Normal function call, it will block execution of next statement till it returns the cotrol to main
	fmt.Println()
	go show("gopher") //Goroutine, won't block execution of next statement, it will run in background.

	//Anonymous goroutine, it will run asynchronously with show() goroutine
	go func(msg string) {
		for i := 0; i < 3; i++ {
			fmt.Println(msg)
			time.Sleep(1 * time.Second)
		}
	}("go is fun")

	time.Sleep(3 * time.Second) //This is to block main goroutine from terminating before other goroutine finish theire execution
	fmt.Println()
	fmt.Println("Main goroutine terminated")
}

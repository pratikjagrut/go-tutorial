//A goroutine is a lightweight thread managed by the Go runtime.
package main

import (
	"fmt"
	"time"
)

func show(arg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(arg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	fmt.Println("**Start**")
	go show("go")
	show("pher")
	fmt.Println("**Stop**")

	//Anonymous goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("go is fun")

	time.Sleep(1 * time.Second)
	fmt.Println("Done!")
}

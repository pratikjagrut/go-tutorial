//A WaitGroup is used to wait for a collection of Goroutines to finish executing.
package main

import (
	"fmt"
	"sync"
	"time"
)

func goroutine(id int, wg *sync.WaitGroup) {
	fmt.Printf("Goroutine number %d started\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine number %d stoped\n", id)
	wg.Done() //Decreases the counter
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i < 4; i++ {
		wg.Add(1) //Increments the counter
		go goroutine(i, &wg)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Wait() //makes the main goroutine to wait until the counter becomes zero
	fmt.Println("All goroutines finished their execution")
}

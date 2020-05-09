//To synchronize access to shared state across multiple goroutines
//we can also use the built-in synchronization features of
//goroutines and channels to achieve the same result as mutex.
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//This channel-based approach aligns with Go’s ideas of sharing memory by communicating
//and having each piece of data owned by exactly 1 goroutine.
func main() {
	var readOps uint64  //Read count
	var writeOps uint64 //Write count

	read := make(chan bool)
	write := make(chan int)
	done := make(chan bool)
	finalSum := make(chan int)

	//In this example our state (in this case sum) will be owned by a single goroutine.
	//This will guarantee that the data is never corrupted with concurrent access.
	go func() {
		sum := 0
		for {
			select {
			case <-read: //Read operation on shared operation
				fmt.Println(sum)
			case i := <-write: //Write operation on shared resource
				sum += i
			case <-done: //Send final sum via channel
				finalSum <- sum
			}
		}
	}()

	//Writer goroutines, will create 5 workers
	for w := 0; w < 5; w++ {
		go func() {
			for {
				number := rand.Intn(10000)
				write <- number
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	//Reader goroutine
	go func() {
		for {
			read <- true
			atomic.AddUint64(&readOps, 1)
			time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(time.Second) //Reader and writer goroutines will run for 1 second, before main goroutine is continued
	done <- true            //To get final sum
	fmt.Println("Read count: ", readOps)
	fmt.Println("Writer count: ", writeOps)
	fmt.Println("Sum: ", <-finalSum)
}

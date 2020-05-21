package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var readOps uint64  //Read count
	var writeOps uint64 //Write count

	sum := 0 // Shared mem state

	var mutex = &sync.Mutex{}

	//Writer goroutines, will create 5 workers
	for w := 0; w < 5; w++ {
		go func() {
			for {
				number := rand.Intn(10000)
				mutex.Lock()
				sum += number //Write operation on shared resource
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(10 * time.Millisecond)
			}
		}()
	}

	//Reader goroutine
	go func() {
		for {
			mutex.Lock()
			fmt.Println(sum) //Read operation on shared operation
			mutex.Unlock()
			atomic.AddUint64(&readOps, 1)
			time.Sleep(time.Millisecond)
		}
	}()

	time.Sleep(time.Second) //Reader and writer goroutines will run for 1 second, before main goroutine is continued
	fmt.Println("Read count: ", readOps)
	fmt.Println("Writer count: ", writeOps)
	fmt.Println("Sum: ", sum)

}

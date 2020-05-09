//We saw managing state in go using channels
//now we will look at using the `sync/atomic` package
//for atomic counters accessed by multiple goroutines.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main(){
	//In this example we will perform counter which will be shared by all goroutines 
	var counter int64 //this is global counter variable 

	var wg sync.WaitGroup

	//We are creating 100 groutines
	for i := 0; i < 100; i++ { 
		wg.Add(1)
		go func ()  {
			//each goroutine will increament counter 1000 times
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	//This suppose to print 100000
	//as 100 goroutines * 1000 times counter = 100000
	fmt.Println("Counter: ", counter)
	
}
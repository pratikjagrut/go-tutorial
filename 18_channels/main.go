//a channel is a technique which allows to let one goroutine to send data to another goroutine
package main

import (
	"fmt"
	"time"
)

func main() {
	//simple send and receive msg through channel
	c := make(chan string)
	go func(ch chan string) {
		time.Sleep(2 * time.Second)
		ch <- "msg" // receiving values
	}(c)
	// it will block excution untill channel receives values
	msg := <-c // sending values
	fmt.Println(msg)

	fmt.Println()

	//Close channel after use
	c1 := make(chan int)

	go func(n int, ch chan int) {
		sum := 0
		for i := 0; i <= n; i++ {
			sum += i
		}
		ch <- sum
		close(ch)
	}(5, c1)

	ans, ok := <-c1
	fmt.Println("Sum: ", ans, ", Is ch open: ", ok)
	_, ok = <-c1
	fmt.Println("Is channel c1 open: ", ok)

	fmt.Println()

	//Buffered channel
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	// c2 <- 3, this will throw runtime error
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	c2 <- 3 //Once the data is sent we can use buffer again
	fmt.Println(<-c2)

	fmt.Println()

	//Range iterates over channel
	c3 := make(chan int, 12)

	//This goroutine will calculate fibonacci series and will send each element of the series via channel
	go func(n int, ch chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			time.Sleep(500 * time.Millisecond) //this will help to simulate range over channel in slow speed
			ch <- x
			x, y = y, x+y
		}
		close(ch)
	}(cap(c3), c3)

	for i := range c3 {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	fmt.Println()

	//Channel Directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping("Ping", pings)
	pong(pings, pongs)
	fmt.Println(<-pongs)

	fmt.Println()

	//Channel synchronization
	done := make(chan bool)
	go func(done chan bool) {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond) // To simulate the hold
			fmt.Printf(" %d", i)
		}
		done <- true
	}(done)

	// It will hold the main goroutine until it receives value
	<-done // If you comment this line, you won't see any output from above goroutine
}

//Channel Directions
func ping(msg string, pings chan<- string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

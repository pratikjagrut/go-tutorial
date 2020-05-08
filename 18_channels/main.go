//a channel is a technique which allows to let one goroutine to send data to another goroutine
package main

import (
	"fmt"
	"time"
)

func send(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "ping" //send msg
}

func sum(n int, ch chan int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	ch <- sum // send sum
	close(ch)
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

//Channel Directions
//This ping function only accepts a channel for sending values.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

//The pong function accepts one channel for receives (pings) and a second for sends (pongs).
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	//simple send and receive msg through channel
	c := make(chan string)
	go send(c)
	msg := <-c // receive msg
	fmt.Println(msg)

	//Close channel after use
	c1 := make(chan int)
	go sum(5, c1)
	ans, ok := <-c1
	fmt.Println("Sum: ", ans, ", Is ch open: ", ok)
	_, ok = <-c1
	fmt.Println("Is channel c1 open: ", ok)

	//Buffered channel
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	//c2 <- 3, this will give an error, as buffer overflows
	fmt.Println(<-c2)
	fmt.Println(<-c2)

	//Range iterates over channel
	c3 := make(chan int, 12)
	go fibonacci(cap(c3), c3)
	for i := range c3 {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	//Channel Directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Ping")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

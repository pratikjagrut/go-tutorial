//tickers are for when you want to do something repeatedly at regular intervals
package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	stoped := make(chan bool)

	go func() {
		for {
			select {
			case <-stoped:
				fmt.Println("Ticker stoped!")
				return //exit for loop
			case t := <-ticker.C:
				fmt.Println("Tick at ", t)
			}
		}
	}()

	time.Sleep(5000 * time.Millisecond) //sleeps for ticker execution
	ticker.Stop()                       //stops ticker
	stoped <- true
}

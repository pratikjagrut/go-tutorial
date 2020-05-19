//Timers are for when you want to do something once in the future
package main

import (
	"fmt"
	"time"
)

func main() {
	//this is simple illustration, much similar to time.Sleep()
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer1 expired")

	//But unlike time.Sleep()
	//we can stop this timer before expiring
	timer2 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer2.C //fire timer
		fmt.Println("Timer2 expired")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer2 stoped")
	}

	//reset timer
	timer3 := time.NewTimer(1 * time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer3 expired")
	}()

	reset := timer3.Reset(2 * time.Second)
	if reset {
		fmt.Println("Timer3 reseted to 2 seconds")
	}

	time.Sleep(3 * time.Second) // Needed to wait for goroutine to return control before exiting program
}

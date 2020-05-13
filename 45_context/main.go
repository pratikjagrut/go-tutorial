package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func call(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("called")

	}
}

func main() {
	// context.WithCancel
	ctx1 := context.Background()              //It is root context
	ctx1, cancel1 := context.WithCancel(ctx1) //child context
	go func() {
		time.Sleep(3 * time.Second)
		cancel1() //all context will be canceled once cancel is invoked
	}()
	call(ctx1)

	// context.WithTimeout
	ctx2 := context.Background()
	ctx2, cancel2 := context.WithTimeout(ctx2, 3*time.Second)
	//It is necessary to call cancel as context will allocate some memory
	//to time so it is necessary to free that memory
	defer cancel2()
	call(ctx2)

	// Context withDeadLine is same as context.WithTimeout
	ctx3 := context.Background()
	ctx3, cancle3 := context.WithDeadline(ctx3, time.Now())
	defer cancle3()
	call(ctx3)
}

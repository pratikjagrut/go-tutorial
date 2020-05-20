# WaitGroup

A `WaitGroup` is used to wait for a collection of Goroutines to finish executing.
The control is blocked until all Goroutines finish executing.
Suppose we've 2 concurrently executing goroutines spawned from the `main` goroutine.
Therefore `main` goroutine is supposed to wait until other 2 goroutines finish their execution.
This can be done by `Waitgroup`.

## Declaring WaitGroup object

`import "sync"`

`WaitGroup` is a struct type, and we usually create zero value variable of it.

```go
var wg sync.WaitGroup
```

We mainly use 3 methods defined on WaitGroup: `Add()`, `Done()` and `Wait()`.

The way `WaitGroup` works is by using a counter. 
Whenever we call goroutine we call `Add(value)` passed value is `int`.
The counter is incremented by `value` passed in Add method. 
Then wherever we finish execution of goroutine we end it with `Done()` method.
`Done()` method decrements the counter by `one`. In the main goroutine, we call `Wait()` method which will block the execution until the counter becomes `zero`.

We parameterize goroutine with the pointer of WaitGroup `func goroutine(id int, wg *sync.WaitGroup) {}`

```go
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
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
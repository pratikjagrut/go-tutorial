# Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

Any function called with `go` keyword starts new goroutine `go f(x, y, z)`. All the goroutines run simultaneously or concurrently in the same address space. Every program contains `main Goroutine`. All the goroutines work under the main goroutine, if the main goroutine is terminated all the other goroutines, working under the main goroutine will also be terminated. Goroutines work in the background.

*Example:*

This is just a normal function, which will print given arg 3 times after a certain interval of sleep.

```go
func show(arg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(arg)
        time.Sleep(700 * time.Millisecond)
    }
}
```

Sleep is just added to simulate concurrency if you remove it function execution will be very fast and we won't be able to notice concurrency very well.

*We can call this function normally or create its goroutine.*

**Normal function call**

The normal function call will block the execution of the next statement until it returns the control to the main.

```go
show("go go go")
```

**Goroutine**

Goroutine won't block execution of the next statement, it will run in the background asynchronously.

```go
go show("gopher")
```

**Anonymous goroutine**

We can also create anonymous goroutine using an anonymous function

```go
go func(msg string) {
    for i := 0; i < 3; i++ {
        fmt.Println(msg)
        time.Sleep(1 * time.Second)
    }
}("go is fun")
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
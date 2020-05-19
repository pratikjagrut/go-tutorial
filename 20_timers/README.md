# Timers

Sometimes we want to execute code after sometimes in future. GO has a built-in feature called `timers`.This allows us to execute specific code after some time. This may collide with `time.Sleep()` or `time.After()`. This functions blocks execution until the specific time is over, and the timer does the similar operation but in timer, we can cancel the timer or reset it before it is expired.

## Declaring timer

First, we need to import the `time` package
```go
import "time"
```

Then set time using `time.NewTimer(value time.duration)`

```go
timer := time.NewTimer(2 * time.Second)
```

The `timeName.C` blocks on timer's channel `C` of type `time.Time`. It unblocks when the timer has expired.

```go
<-timer.C
```

## Cancel timer

In some scenarios, you want to be able to cancel a process.

`timer.Stop()` cancels the timer.

Here we fire timer in goroutine so it won't block execution of next statement.

```go
timer2 := time.NewTimer(1 * time.Second)
go func() {
    <-timer2.C
    fmt.Println("Timer2 expired")
}()

stop := timer2.Stop()
if stop {
    fmt.Println("Timer2 stoped")
}
```

In some scenarios, you want to be able to reset the timer.

`timer.reset(duration)` will reset the timer to the specified channel.

Here we fire timer in goroutine so it won't block execution of next statement.

```go
timer3 := time.NewTimer(1 * time.Second)
go func() {
    <-timer3.C
    fmt.Println("Timer3 expired")
}()

reset := timer3.Reset(2 * time.Second)
if reset {
    fmt.Println("Timer3 reseted to 2 seconds")
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
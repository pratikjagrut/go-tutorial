# Tickers

Tickers are for when you want to do something repeatedly at regular intervals. This is unlike timers which are used for timeouts. A ticker can repeat a block of code.

## Declaring tickers

First, we need to import the `time` package

```go
import "time"
```

Then set time using `time.NewTicker(value time.duration)`. The ticker will tick after a duration specified.

```go
ticker := time.NewTicker(500 * time.Millisecond)
```

Tickers use a similar mechanism to timers: a channel that is sent values.

Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.

In below example, we create ticker which ticks after each `500 Millisecond`. We create boolean channel `stoped` which will send and receive `true/false` values. We are using `select` to control channel cases.

```go
ticker := time.NewTicker(500 * time.Millisecond)
stoped := make(chan bool)

go func() {
    for {
        select {
        case <-stoped:
            fmt.Println("Ticker stoped!")
            break //breaks for loop
        case t := <-ticker.C:
            fmt.Println("Tick at ", t)
        }
    }
}()

time.Sleep(5 * time.Second) //sleeps for ticker execution
ticker.Stop()               //stops ticker
stoped <- true
```

In the above example, `select` runs in `forever for loop` which means it will look for its cases of send or receive operation of channel infinitely. Ticker channel will send values over each `500 Millisecond` till `time.Sleep()` is expired. 
Then `ticker.Stop()` will turn off ticker and `stoped` channel will send `true` value and 
select will receive value from the stoped channel and `breaks` the for a loop.

***You can refer main.go file for examples***

To run:
```
go run main.go
```
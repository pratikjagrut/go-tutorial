# Channels

Channels can be thought of as pipes using which Goroutines communicate.
Similar to how water flows from one end to another in a pipe, 
data can be sent from one end and received from another end.

## Declaring channels

Each channel is declared with a type and only declared type of value is allowed to transport through it. 
The zero value of a channel is nil. Channel should be declared with `make` like maps and slices.

```
channelName := make(chan valueType)
```

```go
ch := make(chan int)
```
Data flow in the direction of this `<-` operator.

By default, sends and receives block until the other side is ready. 
This allows goroutines to synchronize without explicit locks or condition variables.

Below is a simple example of send and receive data through a channel where it will block execution until the channel receives values(here 2 sec). 

```go
c := make(chan string)
go func(ch chan string) {
    time.Sleep(2 * time.Second)
    ch <- "msg" // receiving values
}(c)
// it will block excution untill channel receives values
msg := <-c // sending values
fmt.Println(msg)
```

## Buffered channel

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel.

`ch := make(chan int, 100)`

If the buffer is full it will block channel until buffered is freed. If we try to send value through the blocked channel it will throw runtime error `all goroutines are asleep - deadlock!`.

```go
c2 := make(chan int, 2)
c2 <- 1
c2 <- 2
// c2 <- 3, this will throw runtime error
fmt.Println(<-c2)
fmt.Println(<-c2)
c2 <- 3 //Once the data is sent we can use buffer again
fmt.Println(<-c2)
```

## Closing channel

Closing channels after use are the secure practice of using them. A sender can close a channel to indicate that no more values will be sent. Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.

`close(ch)` this way we can close the channel.

`value, ok := <-ch`, ch will return two values first one is data and another is bool, true if the channel is open, false if channel close 

```go
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
```

**Range over channel**

Like slices or maps, channels can also be iterated over using the range.

```
for i:= range ch{}
```

Below is anonymous goroutine which will calculate Fibonacci series and will send each element to the main goroutine through a channel. Below that there is a for loop which will iterate over each element using range on a channel.

```go
c3 := make(chan int, 12)

go func(n int, ch chan int) {
    x, y := 0, 1
    for i := 0; i < n; i++ {
        time.Sleep(500 * time.Millisecond)
        ch <- x
        x, y = y, x+y
    }
    close(ch)
}(cap(c3), c3)

for i := range c3 {
    fmt.Printf("%d ", i)
}
```
*Sleep is just to simulate iteration in slow speed, it is not compulsory.*

## Unidirectional channel

When using channels as function parameters, you can specify if a channel is meant to only send or receive values.

While declaring channel we can declare them with:

`var sendOnlyChan chan<- valueType` : send-only channel, we cannot read from, but can write to and close(), in simple word we can only put stuff in this channel we can not receive anything from it.

`var receiveOnlyChan <-chan valueType` : receive-only channel, we can read from, but cannot write to or close(), in simple words we can receive stuff from this channel can not put anything into it.

The point of the unidirectional channel is, sometimes we need the receiver to only receive data not send any but sender still has receive and send channel.

In below snippet, function ping has two parameter: 

1<sup>st</sup> is `msg` which will store string value.

2<sup>nd</sup> is `pings` a send-only channel, we are writing `msg` which we will send through it.

```go
func ping(msg string, pings chan<- string) {
    pings <- msg
}
```

In below snippet, function pong has two parameters:

1<sup>st</sup> is `pings` receive-only channel, through which we'll read or receive `msg`.

2<sup>nd</sup> is `pongs` a send-only channel, through which we'll send `msg`. 

```go
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
```

Below snippet is calling or implementation of a channel in the main function.

```go
pings := make(chan string, 1)
pongs := make(chan string, 1)
ping("Ping", pings)
pong(pings, pongs)
fmt.Println(<-pongs)
```

## Goroutine synchronization using channels

We know if the main goroutine is terminated then all the other goroutine associated with will be terminated too. So to let goroutines finish their execution we added sleep in the main goroutine in previous goroutines example. But channels allow us to do this neatly.

```go
done := make(chan bool)
go func(done chan bool) {
    for i := 0; i < 10; i++ {
        fmt.Printf(" %d", i)
    }
    done <- true
}(done)

// It will hold the main goroutine until it receives value
<-done 
```

If you comment last line `<-done` main will terminate and hence you won't see any output from above goroutine.

***You can refer main.go file for examples***

To run:
```
go run main.go
```
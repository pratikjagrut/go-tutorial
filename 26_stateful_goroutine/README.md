# Stateful Goroutines

To synchronize access to shared state across multiple goroutines we can also use the built-in synchronization features of
goroutines and channels to achieve the same result as the mutex.

This channel-based approach aligns with Go’s ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.

*Example:*

In this example, we are calculating the sum of random integers.

**Shared variable**

`sum` will be shared variable by two goroutines, one will be the calculating sum and writing it in sum variable, and the other one will be reading the value from sum variable concurrently.

**Write Operation**

Goroutines will be performing some write operations on the `sum` variable. We'll be keeping count of write operations using 
`var writeOps uint64`.

Here we are creating 5 worker thread to perform write operations simultaneously.

`number` is a random integer between 0 to 10000.

`write := make(chan int)`: This channel will be used send `number` variable to calculate and write sum to shared variable `sum`

`atomic.AddUint64(&writeOps, 1)`: It will be incrementing write operation counter.

```go
for w := 0; w < 5; w++ {
    go func() {
        for {
            number := rand.Intn(10000)
            write <- number
            atomic.AddUint64(&writeOps, 1)
            time.Sleep(10 * time.Millisecond)
        }
    }()
}
```

**Read Operation**

This goroutine is much simpler than earlier one. We'll be keeping count of reading operations using `var writeOps uint64`.

`read bool` channel: Will notify `true` to read the value from the shared variable `sum`.

`atomic.AddUint64(&readOps, 1)`: It will be incrementing read operation counter.

```go
go func() {
    for {
        read <- true
        atomic.AddUint64(&readOps, 1)
        time.Sleep(time.Millisecond)
    }
}()
```

**Synchronizing access using select**

The select statement lets a goroutine wait on multiple communication operations.

Select blocks until one of its cases can run, then it executes that case. 
It chooses one at random if multiple are ready.

Here, select will receive value from:

`read` channel: `true` value will be received to read the value from `sum`.

`write` channel: `number` will be received to calculate and write sum to `sum`.

`done` channel: `true` value will be received at the end, to send `sum` to `finalSum` channel.

```go
select {
case <-read: //Read operation on shared operation
    fmt.Println(sum)
case i := <-write: //Write operation on shared resource
    sum += i
case <-done: //Send final sum via channel
    finalSum <- sum
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
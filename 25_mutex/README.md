We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a time to avoid conflicts?

This concept is called `mutual exclusion`, and the conventional name for the data structure that provides it is `mutex`.

Go's standard library provides mutual exclusion with `sync.Mutex` and its two methods:

`Lock`
`Unlock`

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to Lock and Unlock.

## Declaring mutex

Mutex is a struct present in sync package so we need to import it first.

```go
import "sync"
```

We declare a zero value variable of mutex struct.

```go
var mutex = &sync.Mutex{}
```

Then we call `mutex.Lock()` and `mutex.Unlock()` to sync the access of shared variable or block code.

## Example

In this example, we are calculating the sum of random integers.

**Shared variable**

`sum` will be shared variable by two goroutines, one will be the calculating sum and writing it in sum variable, and the other one will be reading the value from sum variable concurrently.

**Read/Write Operation count**

Goroutines will be performing some read as well as write operations on the `sum` variable. We'll be keeping count of this operations.

We will be using `sync.atomic` package to increment the counter.

`var readOps uint64` will count read operations. 

`var writeOps uint64` will count write operations.

**Defining write operations goroutines**

Here we are creating 5 worker thread to perform write operations simultaneously. 

`number` is a random integer between 0 to 10000.

We are locking the access to `sum` at one goroutine at a time using `mutex.Lock()` and once we finish the use of shared variable we are unlocking it using `mutex.Unlock()`.

We are passing address of `writeOps` and ` int value` to `atomic.AddUint64()` function to increament write operation counter.

```go
for w := 0; w < 5; w++ {
    go func() {
        for {
            number := rand.Intn(10000)
            mutex.Lock()
            sum += number //Write operation on shared resource
            mutex.Unlock()
            atomic.AddUint64(&writeOps, 1)
            time.Sleep(10 * time.Millisecond)
        }
    }()
}
```

**Defining read operations goroutines**

This goroutine is much simpler than earlier one. 
We are reading shared variable `sum` surrounded by `mutex.Lock()` and `mutex.Unlock()` functions. Similar to write operation counter we are increment read operation counter.

```go
go func() {
    for {
        mutex.Lock()
        fmt.Println(sum) //Read operation on shared operation
        mutex.Unlock()
        atomic.AddUint64(&readOps, 1)
        time.Sleep(time.Millisecond)
    }
}()
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
# Context

In Go, we can create many goroutines to handle HTTP requests or spawned goroutines to make any complex calculation concurrently. But if any of this request is cancelled due to any reason then all the goroutines working on that request should exit quickly so the system can reclaim any resources they are using.
This can be performed using `context package`.

The context package makes it easy to pass request-scoped values, cancellation signals, and deadlines across API boundaries to all the goroutines involved in handling a request.

## Creating context

**Creating root context**

`context.Background()` is the root context mainly consumed by main, all other contexts are children of this context.
Background returns a non-nil, empty Context. It is never cancelled, has no values, and has no deadline. It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

Signature: `func context.Background() context.Context`

```go
ctx := context.Background()
```

**context.WithCancel**

Signature: `func context.WithCancel(parent context.Context) (ctx context.Context, cancel context.CancelFunc)`

This function creates a new context derived from the parent context that is passed in. The parent can be a background context or a context that was passed into the function.

`This returns a derived context and the cancel function`.

Cancelling this context releases resources associated with it, so code should call cancel as soon as the operations running in this Context complete.

```go
ctx, cancel := context.WithCancel(ctx)
```

**context.WithTimeout**

Signature: `func context.WithTimeout(parent context.Context, timeout time.Duration) (context.Context, context.CancelFunc)`

This function returns a derived context that gets cancelled if the cancel function is called or the timeout duration is exceeded.

calling cancel function is important here regardless we care or not because the context package will allocate some memory to 

```go
ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
defer cancel()
```

**context.WithDeadline**

Signature: `func context.WithDeadline(parent context.Context, d time.Time) (context.Context, context.CancelFunc)`

This function is similar to `context.WithCancel`. The difference is that it takes in time object as an input instead of the time duration.

```go
ctx, cancle := context.WithDeadline(ctx, time.Now())
defer cancle()
```

**Using this context in our function**

We can use `select` construct to wait and synchronize the channel's communication.

Wichever case `<-ctx.Done()` or `<-time.After(5 * time.Second)` recives select runs executes accordingly.

```go
func callMe(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("called")
	}
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
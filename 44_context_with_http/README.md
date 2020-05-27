# Context with HTTP

A Context carries deadlines, cancellation signals, and other request-scoped values across API boundaries and goroutines.
HTTP servers are useful for demonstrating the usage of context.Context for controlling cancellation.

## Context handling on the server

We know how to create a simple server as we've seen [here](https://github.com/pratikjagrut/go-tutorial/tree/master/42_HTTP_Server).

Now exactly what context we should use, should we use context.Background or something else? But here already a context is hidden inside an `http.Request`.

The context is useful here in the case where the client cancels the request before server finishes serving then context associated with that request will be cancelled and any goroutine serving that request will be terminated.

Following is a simple example of a context.

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    log.Println("Server listening......")
    defer log.Println("Server served")
    select {
    case <-time.After(3 * time.Second):
        fmt.Fprint(w, "Welcome to go world")
    case <-ctx.Done():
        log.Println(ctx.Err())
        http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
    }
}
```

`ctx := r.Context()` is a context from `r *http.Request`. 

We are using `select` construct to wait and synchronize the channel's communication.

**Output**

*Without cancelling the request*

Client-side:

Wait 5 Seconds for response

```
curl http:/localhost:8081/

Welcome to go world
```

Server-side:

```
go run main.go

2020/05/25 13:34:34 Server listening......                                                                 
2020/05/25 13:34:37 Server served  
```

*Cancel the request*

Client-Side:

Cancel curl request with `Ctrl+C` before `5 Seconds`.

```
curl http:/localhost:8081/
^C
```

Server Side:

Serve will receive and `ctx.err` saying `context cancelled`.

```
go run main.go

2020/05/25 13:39:41 Server listening
2020/05/25 13:39:43 context cancelled 
2020/05/25 13:39:43 Server served  
```

## Context handling on the client-side 

We know how to create simple client as we've seen [here](https://github.com/pratikjagrut/go-tutorial/tree/master/41_HTTP_client)


First, we create a root context

```go
ctx := context.Background()
```

Crete `context.WithTimeout` which will cancel HTTP request after 3 seconds.

```
ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
defer cancel()
```

Create HTTP request with the context

```go
req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8081", nil)
if err != nil {
    log.Fatal(err)
}
```

`Do` sends an HTTP request and returns an HTTP response, following policy (such as redirects, cookies, auth) as configured on the client.

```go
res, err := http.DefaultClient.Do(req)
if err != nil {
    log.Fatal(err)
}
defer res.Body.Close()
if res.StatusCode != http.StatusOK {
    log.Fatal(res.Status)
}

io.Copy(os.Stdout, res.Body)
```

**Output**

The client will cancel the request after the timeout.

```
 go run main.go

2020/05/25 14:13:12 Get http://localhost:8081: context deadline exceeded
exit status 1
```
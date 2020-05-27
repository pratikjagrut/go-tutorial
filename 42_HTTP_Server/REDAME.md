# HTTP Server

Writing a basic HTTP server is easy using the `net/http` package.

A fundamental concept in net/http servers is `handlers`. A handler is an object implementing the `http.Handler interface`. A common way to write a handler is by using the `http.HandlerFunc` adapter on functions with the appropriate signature.

Functions serving as handlers take a `http.ResponseWriter` and a `http.Request` as `arguments`. The response writer is used to fill in the HTTP response.

Ping handler function responds with pong upon request

```go
func Ping(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}
```

Header handler gives headers information upon hit

```go
func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
```

We register our handlers on server routes using the http.HandleFunc

```go
http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Please hit endpoint /ping or /header")
})
http.HandleFunc("/ping", ping)
http.HandleFunc("/header", headers)
```

Finally, we call the `ListenAndServe` with the port and a handler. nil tells it to use the default router we’ve just set up.

```go
http.ListenAndServe(":8081", nil) 
```

## Running our server

***You can refer main.go file for examples***

To run:
```
go run main.go
```

In other terminal you can hit endpoints using curl.

*Or you can hit this URLs in browser*

```
curl http:/localhost:8081/

Please hit endpoint /ping or /header
```

```
curl http:/localhost:8081/ping

pong
```

```
curl http:/localhost:8081/header

User-Agent: curl/7.65.3
Accept: */*
```
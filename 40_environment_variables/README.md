# Environment variables

`Environment variables` are a universal mechanism for conveying configuration information to Unix programs. 

Let’s look at how to `set`, `get`, and `list` environment variables using go.

***Set env var***

```go
os.Setenv("VERSION", "1")
```

***Get env var***

```go
fmt.Println("Version: ", os.Getenv("VERSION"))
```

***List all env from System***
```go
for _, e := range os.Environ() {
    fmt.Println(e)
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
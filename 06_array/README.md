# Arrays

An array is a numbered sequence of elements of a specific length.

In Go, the type `[n]T` is an array of length `n` and type `T`.

## Array Declarations 

```go
// i is an array of length 3, 
// it means it can hold upto 3 int values
var i [3]int
i[0] = 1
i[1] = 10
i[2] = 100
```

**Short-hand declaration**
```go
j := [3]string{"Go", "is", "fun"}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
# Errors

GO doesn't have exceptions so it doesn't have a try-catch construct like JAVA, RUBY or other language has. In GO error does not have conventional meaning. An error is just a value which function can return if something unexpected happens. `Go programs express error state with error values`

`error` is a built-in type in Go and it's zero value is `nil`.

A general way to handle an error is to return it as the last return value of a function call and check for the nil condition.

```go
func isEven(i int) (bool, error) {
    if i%2 != 0 {
        return false, errors.New("Not an even number")
    }
    return true, nil
}
```

Code should handle errors by testing whether the error equals nil

```go
ans, err := isEven(9)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(ans)
}
```

The error type is a built-in interface.

```go
type error interface {
    Error() string
}
```

We can create custom error types using `Error()` method declared in the `error` interface.

Create a derived type to implement Error method.

Here we are creating a struct
```go
type myError struct {
    number int
    msg    string
}
```

Implement Error() method

```go
func (e *myError) Error() string {
    return fmt.Sprintf("%d %s", e.number, e.msg)
}
``` 

```go
func isOdd(i int) (bool, error) {
    if i%2 == 0 {
        return false, &myError{
            number: i,
            msg:    "is not an odd number",
        }
    }
    return true, nil
}
```

In above function we are return type is `bool` and `interface error`, and return statement is returning `true/false`, `struct object/nil` this way we are enforcing `myError` to implement `Error()` method.

***You can refer main.go file for examples***

To run:
```
go run main.go
```
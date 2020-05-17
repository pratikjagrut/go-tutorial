# Closures

Go supports anonymous functions which can act as function closures. 
Anonymous functions are used when we want to define a function inline without passing any name to it.

The purpose of this function is to close over a variable of upper function to form a closure.

The function may access and assign to the referenced variables; in this sense, the function is "bound" to the variables.

**Define closure function**

```go
func getLimit() func() int {
    limit := 10
    return func() int {
        limit--
        return limit
    }
}
```
This function `getLimit` returns another function, which we define anonymously in the body of `getLimit`. That returned function closes over the variable `limit` to form a closure.

**Call closure function**

```go
showLimit := getLimit()
fmt.Println(showLimit())
fmt.Println(showLimit())
fmt.Println(showLimit())
```

We call `getLimit`, assigning the result (a function) to `showLimit`. This function value captures its own `limit` value, which will be updated each time we call `showLimit`.

***You can refer main.go file for examples***

To run:
```
go run main.go
```
# Looping Construct

Go has only one looping construct, the `for` loop.

## For loop

**Basic for loop**

The basic for loop has three components separated by semicolons:

- init statement: `i := 0` *exec before 1<sup>st</sup> iteration*
- condition expression: `i < n` *eval on every* interation
- post statement: `i++` *exec after each iteration*

*The expression need not be surrounded by parentheses ( ) but the braces { } are required.*
```go
for i := 0; i < n; i++ {
    //business logic
}
```
Init and post statement are optional.

```go
for ; i < n; {
    //business logic
}
```

**For is C's while() loop**

```go
for i < n {
    //business logic
}
```

**Forever or Infinite loop**
```go
for {
    //business logic
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```


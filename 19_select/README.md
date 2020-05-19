# Select

The select statement lets a goroutine wait on multiple communication operations.

Select blocks until one of its cases can run, then it executes that case. 
It chooses one at random if multiple are ready.

Select can be seen as a switch for channels where case statements refer to send or receive operation on channels.

Syntax:

```go
select {
    case SendOrReceiveOp1:
        // Business logic
    case SendOrReceiveOp2:
        // Business logic
    default:
        // Default business logic
}
```

The default case in a select is run if no other case is ready.
Default case can be used to send or receive without blocking.

***You can refer main.go file for examples***

To run:
```
go run main.go
```
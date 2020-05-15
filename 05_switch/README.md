# Switch statement

A switch statement is a shorter way to write a sequence of if - else statements.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that `Go only runs the selected case, not all the cases that follow`. So we don't need break statement here. Go's switch cases need not be constants, and the values involved need not be integers.

```go
switch expression {
    case exp1:
        //Executes if expression matches exp1
    case exp2:
        //Executes if expression matches exp2
    default:
        //Executes if expression does not matches with any case
}
```

**Switch evaluation order**

Switch cases evaluate cases from top to bottom and stops when a case succeeds. In below example it checks `case "Linux":` if os matches with Linux then it stops at that case else goes to next case.

```go
//Prints which OS you're using
switch os := runtime.GOOS; os {
case "linux":
    fmt.Println("Linux.")
case "darwin":
    fmt.Println("OS X.")
default:
    fmt.Printf("%s.\n", os)
}
```

**Switch with NO condition**

Switch with no condition is like `switch true`. It is useful for writing log if-else-if ladder.
```go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```







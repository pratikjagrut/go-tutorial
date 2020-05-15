# If and Else

## If

Go's if statements are like its for loops, the expression `need not be surrounded by parentheses ( ) but the braces { } are required`.

```go
if condition/expression {
    //business logic
}
```

## if else

```go
if i % 2 == 0 {
    fmt.Println("even")
} else {
    fmt.Println("Odd")
}   
```

## if-else-if ladder

```go
if i == 0 {
    fmt.Println("It's zero")
} else if i < 0 {
    fmt.Println("Negative number")
} else {
    fmt.Println("Positive number")
}
```

## If with a short statement

```go
if j := 10; j%2 == 0 {
	fmt.Println("Even number")
} 
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```

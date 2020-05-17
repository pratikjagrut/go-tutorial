# Functions

A function is a group of statements that together perform a task.

## Defining a function

```
func function_name( [parameter list] ) [return_types]
{
   body of the function
}
```

```go
func addition(i int, j int) int {
    return i + j
}
```

When two or more consecutive parameters have the same type we can omit the type from all but last.

```go
func addition(i, j, k int) int {
    return i + j + k
}
```

**Multiple results**

A function can return any number of results.

```go
func multiply(i, j int) (int, int, int) {
    return i, j, i * j
}
```
```go 
i, j, k := multiply(4, 5)
```

**Named return values**

Go's return values may be named. If so, they are treated as variables defined at the top of the function.

```go
func calc(i int) (x, y int) {
    x = i / 10
    y = i % 10
    return
}
```

**Variadic functions**

Variadic functions can be called with any number of trailing arguments. 
For example, `fmt.Println` is a common variadic function.

```go
func total(nums ...int) (total int) {
    for _, i := range nums {
        total += i
    }
    return total
}
```
```go
total := total(1, 2, 3, 4, 5)
total = total(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
```

**Recursion**

The function which calls itself directly or indirectly is called recursion.

```go
func fact(i int) int {
    if i == 0 {
        return 1
    }
    return i * fact(i-1)
}
```

```go
fact := fact(5)
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
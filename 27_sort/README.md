# Sorting

Go’s sort package implements sorting for builtins and user-defined types. 

*Sorting is in-place, so it changes the given slice and doesn’t return a new one.*
```go
strs := []string{"z", "x", "y"}
sort.Strings(strs)
fmt.Println("Strings:", strs)
```

```go
ints := []int{3, 8, 1}
sort.Ints(ints)
fmt.Println("Ints:   ", ints)
```

This is to check if given integer slice is sorted or not.

```go
s := sort.IntsAreSorted(ints) //Returns bool
fmt.Println("Sorted: ", s)
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
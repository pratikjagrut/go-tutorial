# Range

The range form of the for loop iterates over a slice or map.
When ranging over a slice, two values are returned for each iteration. 
The first is the index, and the second is a copy of the element at that index.

```go
var s = []int{2, 1}
for index, value := range s{
    fmt.Println("index: ", index, "value: ", value)
}
```
You can skip the index or value by assigning to _.
We can omit the value 

```go
for i, _ := range s
for _, value := range s
for i := range s
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```

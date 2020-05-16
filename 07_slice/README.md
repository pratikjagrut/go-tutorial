# Slices

Slices are a pivotal data type in Go, giving a more powerful interface to sequences than arrays. An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array.

The type `[]T` is a slice with elements of type `T`.

## Slice formation

Slice can be formed by specifying 2 indices, high bound and low bound, `s[low : high]`. This selects a half-open range which includes the first element but excludes the last one.
 
```go
odds := [5]int{1, 3, 5, 7, 9}

var o []int = odds[1:4]
```

We can skip low or high bound to use their defaults. For low bound default is `0` and for high bound default is `length of an array`.

`a[:high]`, `a[low:]` 

**They are like references to arrays**

A slice does not store any data, it just describes a section of an underlying array.

Changing element of the slice will make those changes to an underlying array and all other slices which share the same underlying array.

```go
alphabet := [5]string{"A", "B", "C", "D"}

i := alphabet[0:3]
j := alphabet[1:4]

i[2] = "P" // This changes will affect alphabet, i ,j slices
```

**Slice literals**

A slice literal is like an array literal without the length.

```go
i := []int{1, 2, 3, 4, 5,}
```

**Length and capacity**

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

Length: `len(s)`

Capacity: `cap(s)`

**Creating a slice with make**

Slices can be created with the built-in make function; this is how you create dynamically-sized arrays. `make([]T, len, cap(optional))`

```go
a := make([]int, 5)    // len(a) = 5
b := make([]int, 0, 5) // len(b) = 0, cap(b) = 5
```

**Appending**

In addition to these basic operations, slices support several more that make them richer than arrays. One is the builtin append, which returns a slice containing one or more new values. 

```go
z := []int{1, 2, 3, 4, 5}
z = append(z, 6, 7, 8) //1st arg is slice, and then can take any number of arg to append
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
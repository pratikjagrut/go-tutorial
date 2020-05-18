# Pointers

Go has pointers. 

`A pointer holds the memory address of a value`.

The type `*T` is a pointer to a `T` value. It's zero value is `nil`.

```go
var p *int
```

The `&` operator generates a pointer to its operand.

```go
i := 42
p = &i
```

The `*` operator denotes the pointer's underlying value.

```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as `dereferencing` or `indirecting`.

***You can refer main.go file for examples***

To run:
```
go run main.go
```
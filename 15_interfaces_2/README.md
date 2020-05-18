# Interfaces type assertion

A type assertion provides access to an interface value's underlying concrete value.

## Defination

```go
v := i.(T)
```

The above statement states that `interface value i` holds `type T`, and assigns the concrete `T` value to `v`

E.g.

In this example, interface i holds the `string type` value and assigns the `string type` value to `s`.

```go
var i interface{} = "golang"
s := i.(string)
```

A type assertion can return two values: the `underlying value` and a `boolean value` that reports whether the assertion succeeded.

```go
s, ok := i.(string)
```

This boolean value can be used to avoid panic error if that type's value is not in the interface.

## Type switches

Type switch permits us to do several types of assertion in series.

A type switch is like a regular switch statement,
but the cases in a type switch specify `types (not values)`, 
and those values are compared against the type of the value held by the given interface value.

```go
var i interface{} = "golang"

switch t := i.(type) {
case int:
    fmt.Println("Integer")
case string:
    fmt.Println("String")
default:
    fmt.Printf("Unknown value type: %T", t)
}
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
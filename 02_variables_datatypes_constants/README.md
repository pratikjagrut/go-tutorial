# Variables, data types and constants

## Variable

Variable is a symbolic name associated with a value and whose associated value may be changed anytime in the program.
It is the storage unit of a particular data type. A variable must be defined with the type of data it is holding.

## Data types
There are several data types in Go.

```go
bool        int         uint        float32     complex64
string      int8        uint8       float64     complex128
byte        int16       uint16
rune        int32       uint32
error       int64       uint64
```

## Zero value
In some programming languages variable holds a `null` or `undefined` value when not initialized, Go gives it a zero-value of its data type. A `boolean` variable if not initialized, gets the `false` value and an `integer` variable gets `0` value, `string` variable will get `nil` value.

## Declaring a variable

The `var` statement declares a list of variables. Type is at the last of the statement.

```go
var varName dataType
```
E.g.:

```go
var i, j, k int
var f float32
var u uint
```

**Variables with initializers**

We can initialize variable while declaring.

```go
var varName dataType = value
```

```go
var i int = 10
var s string = "hello"
var ok bool = true
```

*We can omit the type of variable, it will take the type of initializers*

```go
var i = 42           // int
var f = 3.142        // float64
var g = "hello"      // string
```

**Short-hand declaration**

This is the most commonly used variable declaration inside functions.

```go
varName := value
``` 
E.g.:
```go
i := 10
s := "hello"
ok := true
```

*When declaring a variable without specifying it's type the variable's type is inferred from the value on the right-hand side.*

```go
var i int
j := i // j is an int

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
```

**Assign or Re-assign**

You can assign or re-assign value to declared variable as:
```go
varName = value
```
E.g.:
```go
i = 12
s = "world"
ok = false
```

**Multiple variable declarations**

We can declare multiple var in one statement

```go
var i, j, k int
var x, y, z = 0.867 + 0.5i, 4.65, true
```

## Constants

Go supports constants of character, string, boolean, and numeric values. Constants are declared using `const` keyword

E.g:
```go
const str string = "constants"
```

**Untyped and Typed**

Constants can be declared with or without a type in Go. If we are declaring literal constant, then we are declaring constants that are untyped and unnamed.

E.g.:
```go
const str string = "constants" //Typed constant
const i = 10 //Untyped constant, literal declaration
const f = 3.14
```
The constants on the LHS of the declaration are named constants and the literal values on the RHS are unnamed constants.

Typed constants don’t use the same type system as variables, they've their own implementation for representing the values that we associate with them.

In the case of typed constant, the declared type is used to associate the type’s precision limitations or kind of constant.

In the case of untyped constant, the literal value will determine what kind/type the constant takes

Every GO compiler has the flexibility to implement constant as they wish, within the mandatory set of [requirements](http://golang.org/ref/spec#Constants).

***You can refer main.go file for examples***

To run:
```
go run main.go
```
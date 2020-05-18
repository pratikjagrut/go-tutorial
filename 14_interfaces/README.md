# Interfaces

An interface is a collection of method signatures that a Type can implement. The interface helps us to define the behaviour of an object. 

E.g. A geometrical figure can have area and perimeter. So `figure` interface will define the signature of `area` method and `perimeter` method, and if an object such as `circle` or `rectangle` defines those methods then `circle or rectangle implements figure interface`.

## Interface declaration

Like struct, we need to create a derived type of interface.

```go
type figure interface {
    area() float32
    perimeter() float32
}
```

## Implementing an interface

There is no explicit declaration for implementing, a type implements an interface by implementing its methods.

```go
type circle struct{
    radius float32
}

//Implementing area() method
func (c *circle) area() float32 {
    return math.Pi * c.radius * c.radius
}
```

```go
func main(){
    c := circle{10}
    fmt.Println("Area of circle: ", c.area())
}
```

Here we've implemented only area method but we can enforce the type to implement all the methods.

Create a type

```go
type rect struct {
    length float32
    width  float32
}
```

Here we are associating object of `rect` struct with `figure` interface

```go
func newRect(length, width float32) figure {
    return &rect{
        length: length,
        width:  width,
    }
}
```

Now if GO does not see an implementation of any or all the methods it will throw a compile-time error:

```
# .../main.go:33:3: cannot use &rect literal (type *rect) as type figure in return argument:
    *rect does not implement figure (missing area method)
Process exiting with code: 2
```
This way we can enforce type to implement all the methods.

## Interface values

Interface values can be inferred as a tuple of `(value, type)`

```go
func main(){
    c := circle{10}
    fmt.Printf("(%v, %T)\n", c, c)
}
```

output: 
```

{10}, main.circle)
```

## Empty interface

*An empty interface may hold values of any type*

```go
var q interface{}
fmt.Printf("(%v, %T)\n", q, q)
q = 10
fmt.Printf("(%v, %T)\n", q, q)
q = "hello"
fmt.Printf("(%v, %T)\n", q, q)
```

output:
```
(<nil>, <nil>)
(10, int)
(hello, string)
```

***You can refer main.go file for examples***

To run:
```
go run main.go
```
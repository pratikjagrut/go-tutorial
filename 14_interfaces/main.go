//Interfaces are collections of method signatures.
package main

import (
	"fmt"
	"math"
)

type figure interface {
	area() float32
	perimeter() float32
}

type circle struct {
	radius float32
}

//Type circle implements interface figure
func (c *circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

//enforcing to implement interface
type rect struct {
	length float32
	width  float32
}

//Type rect is enfocred to use interface figure
func newRect(length, width float32) figure {
	return &rect{
		length: length,
		width:  width,
	}
}

func (r *rect) area() float32 {
	return r.length * r.width
}

func (r *rect) perimeter() float32 {
	return 2 * (r.length * r.width)
}

//Interface values
//Under the hood, interface values can be thought of
//as a tuple of a value and a concrete type:(value, type)
type I interface {
	M()
}

type F float32

func (f F) M() {
	fmt.Printf("(%v, %T)\n", f, f)
}

func main() {
	//Type circle implements interface figure
	c := circle{10}
	fmt.Println("Area of circle: ", c.area())

	//enforcing to implement interface
	fmt.Println("\nEnforcing to implement interface")
	r := newRect(22.6, 12.3)
	fmt.Println("Area of rect: ", r.area())
	fmt.Println("Perimeter of rect: ", r.perimeter())

	fmt.Println("\nInterface values")
	//Interface values
	i := F(math.Pi)
	i.M()
	var p I //Nil Interface values
	fmt.Printf("(%v, %T)\n", p, p)

	//The empty interface
	fmt.Println("\nThe empty interface")
	var q interface{}
	fmt.Printf("(%v, %T)\n", q, q)
	//An empty interface may hold values of any type
	q = 10
	fmt.Printf("(%v, %T)\n", q, q)
	q = "hello"
	fmt.Printf("(%v, %T)\n", q, q)

}

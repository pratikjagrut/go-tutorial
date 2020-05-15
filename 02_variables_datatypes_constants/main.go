package main

import "fmt"

func main() {
	var i int // Varaible decalration
	i = 10    // initialization
	j := 20   // Shorthand for declaring and initializing
	fmt.Println("i + j = ", i+j)

	var k = "golang" // String
	fmt.Println(k)

	var a, b float32 = 1.45, 2.98 // Float32
	fmt.Printf("a = %v, b = %v -> a + b = %v\n", a, b, a+b)

	var c = true // Boolean
	fmt.Println(!c)

	var x, y, z = 0.867 + 0.5i, 4.65, true
	fmt.Println(x) // complex128
	fmt.Println(y)
	fmt.Println(z)

	//constant values
	//we can not reassign value to declared constant
	const str string = "constant" //Typed constant
	fmt.Println(str)

	//const on LHS will take value and type of const on RHS
	const number = 10.00 //Untyped constant
	fmt.Printf("%v, %T\n", number, number)

}

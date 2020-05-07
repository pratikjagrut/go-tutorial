//Interfaces continues
package main

import "fmt"

func main() {
	fmt.Println("**Type assertion**")
	var i interface{} = "golang"
	//A type assertion provides access to an interface value's underlying concrete value.
	s := i.(string)
	fmt.Println(s)

	// a type assertion can return two values:
	//the underlying value and a boolean value that reports whether the assertion succeeded.
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float32)
	fmt.Println(f, ok)

	// //panic error, if interface does not hold type, can be prevented as shown above
	// f = i.(float32)
	// fmt.Println(f)

	//Type switches
	fmt.Println("\n**Type switch**")
	// /A type switch is like a regular switch statement,
	//but the cases in a type switch specify types (not values),
	//and those values are compared against the type of the value held by the given interface value.
	valueType(10)
	valueType("hello")
	valueType(10.987)

}

func valueType(i interface{}) {
	switch t := i.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Printf("Unknown value type: %T", t)
	}
}

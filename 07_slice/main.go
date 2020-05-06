//An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
//flexible view into the elements of an array.
package main

import "fmt"

func main() {
	//create a slice of non zero length
	a := make([]int, 3)
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println("a: ", a)
	a = append(a, 4, 5, 6)
	fmt.Println("a: ", a)

	var b []int = a[2:5] // it will copy elemets from index 2-5 of slice 'a'
	fmt.Println("b: ", b)

	//Slices are like references to arrays
	// A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.

	alphabet := [5]string{"A", "B", "C", "D"}
	fmt.Println(alphabet)

	i := alphabet[0:3]
	j := alphabet[1:4]
	fmt.Println(i, j)

	i[2] = "P"
	fmt.Println(i, j)
	fmt.Println(alphabet)
}

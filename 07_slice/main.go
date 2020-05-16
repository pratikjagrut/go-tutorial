//An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
//flexible view into the elements of an array.
package main

import "fmt"

func main() {
	odds := [5]int{1, 3, 5, 7, 9} //Array of odd numbers

	var o []int = odds[1:4] // Create a slice of odds from 1-4(exculded) index
	fmt.Println(o)

	//Slices are like references to arrays
	alphabet := [5]string{"A", "B", "C", "D"}
	fmt.Println(alphabet)

	i := alphabet[0:3]
	j := alphabet[1:4]
	fmt.Println(i, j)

	i[2] = "P"
	fmt.Println(i, j)
	fmt.Println(alphabet)

	//Length and capacity
	fmt.Printf("Len of i: %d, cap of i: %d", len(i), cap(i))

	//create a slice of non zero length
	a := make([]int, 3)
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println("a: ", a)
	a = append(a, 4, 5, 6) //Append
	fmt.Println("a: ", a)

	var b []int = a[2:5] // it will copy elemets from index 2-5 of slice 'a'
	fmt.Println("b: ", b)
}

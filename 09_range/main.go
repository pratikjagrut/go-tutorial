//The range form of the for loop iterates over a slice or map.
//When ranging over a slice, two values are returned for each iteration.
//The first is the index, and the second is a copy of the element at that index.
package main

import "fmt"

func main() {
	var i = []int{2, 1}

	for j, k := range i {
		fmt.Printf("Index: %d, Value: %d\n", j, k)
	}

	fmt.Println()

	alphabet := map[string]string{"A": "Apple", "B": "Ball", "C": "Cat"}

	for k, v := range alphabet {
		fmt.Printf("Key: %s, Value: %s\n", k, v)
	}

	fmt.Println()

	//range on strings can iterates over unicode points.
	for j, k := range "git" {
		fmt.Printf("Index: %v, rune code: %v\n", j, k)
	}
}

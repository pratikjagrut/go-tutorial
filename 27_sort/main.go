//Sorting
package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"z", "x", "y"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{3, 8, 1}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints) //Returns bool
	fmt.Println("Sorted: ", s)
}

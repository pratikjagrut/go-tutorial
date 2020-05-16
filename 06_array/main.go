//In Go, an array is a numbered sequence of elements of a specific length.
package main

import "fmt"

func main() {
	var i [3]int
	i[0], i[1], i[2] = 1, 2, 3
	fmt.Println(i)

	j := [3]int{4, 5, 6}
	fmt.Println(j)
	fmt.Println("Len of J: ", len(j))

	// 2D array
	var k [2][2]int
	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			k[p][q] = p + q
		}
	}

	fmt.Println(k)
}

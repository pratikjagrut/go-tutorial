package main

import "fmt"

// func func_name(parameter type) return_type{}

func addition(i int, j int) int {
	return i + j
}

//When two or more consecutive named function parameters share a type,
//you can omit the type from all but the last.
func add(i, j, k int) int {
	return i + j + k
}

func main() {
	fmt.Println(addition(4, 5))
	fmt.Println(add(4, 5, 6))
}

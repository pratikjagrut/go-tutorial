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

//multiple return value
func multiply(i, j int) (int, int, int) {
	return i, j, i * j
}

//Named return values
func calc(i int) (x, y int) {
	x = i / 10
	y = i % 10
	return
}

//Variadic Functions, can be called with any number of trailing arguments.
func total(nums ...int) (total int) {
	for _, i := range nums {
		total += i
	}
	return total
}

//Recursion
func fact(i int) int {
	if i == 0 {
		return 1
	}
	return i * fact(i-1)
}

func main() {
	fmt.Println("Add 4,5: ", addition(4, 5))

	fmt.Println("Add 4, 5, 6: ", add(4, 5, 6))

	//multiple return value
	i, j, k := multiply(4, 5)
	fmt.Printf("Mul %d, %d: %d\n", i, j, k)

	//Named return values
	x, y := calc(17)
	fmt.Printf("Quotient: %d, Remainder: %d\n", x, y)

	//Variadic Functions
	fmt.Println("Total 1-5: ", total(1, 2, 3, 4, 5))
	fmt.Println("Total 1-10: ", total(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	//Recursion
	fmt.Println("Factorial of 5: ", fact(5))
}
